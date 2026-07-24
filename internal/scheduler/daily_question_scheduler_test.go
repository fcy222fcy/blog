package scheduler

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type fakeDailyQuestionRepo struct {
	mu             sync.Mutex
	publishedDates []string
}

func (r *fakeDailyQuestionRepo) PublishScheduledQuestions(today string) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.publishedDates = append(r.publishedDates, today)
	return 1, nil
}

func TestDailyQuestionSchedulerRunsPublishJob(t *testing.T) {
	repo := &fakeDailyQuestionRepo{}
	s := NewDailyQuestionScheduler(repo)

	require.NoError(t, s.PublishToday())

	require.Equal(t, []string{time.Now().Format("2006-01-02")}, repo.publishedDates)
}
