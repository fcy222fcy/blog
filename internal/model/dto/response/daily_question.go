package response

import "time"

// DailyQuestionResponse 每日一问响应
type DailyQuestionResponse struct {
	ID           uint      `json:"id"`
	Question     string    `json:"question"`
	Answer       string    `json:"answer"`
	Date         string    `json:"date"`
	LikeCount    int64     `json:"like_count"`
	CommentCount int64     `json:"comment_count"`
	ViewCount    int64     `json:"view_count"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
