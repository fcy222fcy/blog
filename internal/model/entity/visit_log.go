package entity

import "time"

// VisitLog 访问日志
type VisitLog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	ArticleID uint      `gorm:"index" json:"article_id"`
	IP        string    `gorm:"type:varchar(50)" json:"ip"`
	UserAgent string    `gorm:"type:varchar(255)" json:"user_agent"`
	Referer   string    `gorm:"type:varchar(255)" json:"referer"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 表名
func (VisitLog) TableName() string {
	return "visit_logs"
}
