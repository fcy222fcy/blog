package entity

import "time"

// CommentLikeLog 评论点赞记录
type CommentLikeLog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CommentID uint      `gorm:"index;not null" json:"comment_id"`
	VisitorIP string    `gorm:"type:varchar(50);not null" json:"visitor_ip"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 表名
func (CommentLikeLog) TableName() string {
	return "comment_like_logs"
}
