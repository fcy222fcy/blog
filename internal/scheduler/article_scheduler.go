package scheduler

import (
	"blog/internal/model/entity"
	"blog/pkg/logger"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type ArticleRepository interface {
	FindByID(id uint) (*entity.Article, error)
	ListScheduledAfter(now time.Time) ([]*entity.Article, error)
	PublishScheduledArticle(id uint, now time.Time) (bool, error)
	PublishDueScheduledArticles(now time.Time) (int64, error)
}

type ArticlePublishedHook func(article *entity.Article)

type ArticleScheduler struct {
	repo        ArticleRepository
	onPublished ArticlePublishedHook
	cron        *cron.Cron
	mu          sync.Mutex
	entries     map[uint]cron.EntryID
	started     bool
}

func NewArticleScheduler(repo ArticleRepository, onPublished ArticlePublishedHook) *ArticleScheduler {
	return &ArticleScheduler{
		repo:        repo,
		onPublished: onPublished,
		cron:        cron.New(),
		entries:     make(map[uint]cron.EntryID),
	}
}

func (s *ArticleScheduler) Start() error {
	now := time.Now()
	if count, err := s.repo.PublishDueScheduledArticles(now); err != nil {
		return err
	} else if count > 0 {
		logger.Infof("启动时发布过期定时文章: %d", count)
	}

	articles, err := s.repo.ListScheduledAfter(now)
	if err != nil {
		return err
	}
	for _, article := range articles {
		if err := s.RefreshArticle(article); err != nil {
			return err
		}
	}

	s.mu.Lock()
	if !s.started {
		s.cron.Start()
		s.started = true
	}
	s.mu.Unlock()
	return nil
}

func (s *ArticleScheduler) Stop() {
	s.cron.Stop()
}

func (s *ArticleScheduler) RefreshArticle(article *entity.Article) error {
	if article == nil {
		return nil
	}
	s.UnscheduleArticle(article.ID)
	if article.Status != entity.ArticleStatusScheduled || article.ScheduledAt == nil {
		return nil
	}
	if !article.ScheduledAt.After(time.Now()) {
		return s.publishArticle(article.ID)
	}

	entryID := s.cron.Schedule(oneTimeSchedule{at: *article.ScheduledAt}, cron.FuncJob(func() {
		if err := s.publishArticle(article.ID); err != nil {
			logger.Errorf("定时发布文章失败, id: %d, error: %v", article.ID, err)
		}
	}))

	s.mu.Lock()
	s.entries[article.ID] = entryID
	s.mu.Unlock()
	return nil
}

func (s *ArticleScheduler) UnscheduleArticle(id uint) {
	s.mu.Lock()
	entryID, ok := s.entries[id]
	if ok {
		delete(s.entries, id)
	}
	s.mu.Unlock()
	if ok {
		s.cron.Remove(entryID)
	}
}

func (s *ArticleScheduler) PendingCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.entries)
}

func (s *ArticleScheduler) publishArticle(id uint) error {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	published, err := s.repo.PublishScheduledArticle(id, time.Now())
	if err != nil {
		return err
	}
	s.UnscheduleArticle(id)
	if published && s.onPublished != nil {
		s.onPublished(article)
	}
	return nil
}

type oneTimeSchedule struct {
	at time.Time
}

func (s oneTimeSchedule) Next(after time.Time) time.Time {
	if s.at.After(after) {
		return s.at
	}
	return time.Time{}
}
