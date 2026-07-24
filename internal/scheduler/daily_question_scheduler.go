package scheduler

import (
	"blog/pkg/logger"
	"time"

	"github.com/robfig/cron/v3"
)

type DailyQuestionRepository interface {
	PublishScheduledQuestions(today string) (int64, error)
}

type DailyQuestionScheduler struct {
	repo DailyQuestionRepository
	cron *cron.Cron
}

func NewDailyQuestionScheduler(repo DailyQuestionRepository) *DailyQuestionScheduler {
	return &DailyQuestionScheduler{
		repo: repo,
		cron: cron.New(),
	}
}

func (s *DailyQuestionScheduler) Start() error {
	if err := s.PublishToday(); err != nil {
		return err
	}
	_, err := s.cron.AddFunc("0 0 * * *", func() {
		if err := s.PublishToday(); err != nil {
			logger.Errorf("每日一问定时发布失败: %v", err)
		}
	})
	if err != nil {
		return err
	}
	s.cron.Start()
	return nil
}

func (s *DailyQuestionScheduler) Stop() {
	s.cron.Stop()
}

func (s *DailyQuestionScheduler) PublishToday() error {
	today := time.Now().Format("2006-01-02")
	count, err := s.repo.PublishScheduledQuestions(today)
	if err != nil {
		return err
	}
	if count > 0 {
		logger.Infof("每日一问定时发布完成, count: %d, date: %s", count, today)
	}
	return nil
}
