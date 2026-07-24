package scheduler

import (
	"blog/internal/model/entity"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type fakeArticleRepo struct {
	mu        sync.Mutex
	articles  map[uint]*entity.Article
	published []uint
}

func newFakeArticleRepo() *fakeArticleRepo {
	return &fakeArticleRepo{articles: make(map[uint]*entity.Article)}
}

func (r *fakeArticleRepo) FindByID(id uint) (*entity.Article, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	article := r.articles[id]
	if article == nil {
		return nil, nil
	}
	copied := *article
	return &copied, nil
}

func (r *fakeArticleRepo) ListScheduledAfter(now time.Time) ([]*entity.Article, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*entity.Article
	for _, article := range r.articles {
		if article.Status == entity.ArticleStatusScheduled && article.ScheduledAt != nil && article.ScheduledAt.After(now) {
			copied := *article
			result = append(result, &copied)
		}
	}
	return result, nil
}

func (r *fakeArticleRepo) PublishScheduledArticle(id uint, now time.Time) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	article := r.articles[id]
	if article == nil || article.Status != entity.ArticleStatusScheduled || article.ScheduledAt == nil || article.ScheduledAt.After(now) {
		return false, nil
	}
	article.Status = entity.ArticleStatusPublished
	article.ScheduledAt = nil
	r.published = append(r.published, id)
	return true, nil
}

func (r *fakeArticleRepo) PublishDueScheduledArticles(now time.Time) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var count int64
	for _, article := range r.articles {
		if article.Status == entity.ArticleStatusScheduled && article.ScheduledAt != nil && !article.ScheduledAt.After(now) {
			article.Status = entity.ArticleStatusPublished
			article.ScheduledAt = nil
			r.published = append(r.published, article.ID)
			count++
		}
	}
	return count, nil
}

func TestArticleSchedulerPublishesArticleAtScheduledTime(t *testing.T) {
	repo := newFakeArticleRepo()
	runAt := time.Now().Add(120 * time.Millisecond)
	repo.articles[1] = &entity.Article{
		BaseEntity:  entity.BaseEntity{ID: 1},
		Title:       "scheduled",
		Status:      entity.ArticleStatusScheduled,
		ScheduledAt: &runAt,
	}

	s := NewArticleScheduler(repo, nil)
	require.NoError(t, s.Start())
	defer s.Stop()

	require.Eventually(t, func() bool {
		article, _ := repo.FindByID(1)
		return article.Status == entity.ArticleStatusPublished
	}, time.Second, 20*time.Millisecond)
	require.Equal(t, 0, s.PendingCount())
}

func TestArticleSchedulerRefreshUnschedulesDraftArticle(t *testing.T) {
	repo := newFakeArticleRepo()
	runAt := time.Now().Add(time.Hour)
	article := &entity.Article{
		BaseEntity:  entity.BaseEntity{ID: 2},
		Status:      entity.ArticleStatusScheduled,
		ScheduledAt: &runAt,
	}
	repo.articles[2] = article

	s := NewArticleScheduler(repo, nil)
	require.NoError(t, s.RefreshArticle(article))
	require.Equal(t, 1, s.PendingCount())

	article.Status = entity.ArticleStatusDraft
	article.ScheduledAt = nil
	require.NoError(t, s.RefreshArticle(article))

	require.Equal(t, 0, s.PendingCount())
}

func TestArticleSchedulerStartPublishesOverdueAndLoadsFutureArticles(t *testing.T) {
	repo := newFakeArticleRepo()
	past := time.Now().Add(-time.Minute)
	future := time.Now().Add(time.Hour)
	repo.articles[1] = &entity.Article{BaseEntity: entity.BaseEntity{ID: 1}, Status: entity.ArticleStatusScheduled, ScheduledAt: &past}
	repo.articles[2] = &entity.Article{BaseEntity: entity.BaseEntity{ID: 2}, Status: entity.ArticleStatusScheduled, ScheduledAt: &future}

	s := NewArticleScheduler(repo, nil)
	require.NoError(t, s.Start())
	defer s.Stop()

	overdue, _ := repo.FindByID(1)
	require.Equal(t, entity.ArticleStatusPublished, overdue.Status)
	require.Equal(t, 1, s.PendingCount())
}
