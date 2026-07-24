package entity

const (
	DailyQuestionStatusDisabled  = 0
	DailyQuestionStatusPublished = 1
	DailyQuestionStatusScheduled = 2
)

// DailyQuestion 每日一问
type DailyQuestion struct {
	BaseEntity
	Question     string `gorm:"type:text;not null" json:"question"`
	Answer       string `gorm:"type:text" json:"answer"`
	Date         string `gorm:"type:varchar(10);uniqueIndex;not null" json:"date"`
	LikeCount    int64  `gorm:"default:0" json:"like_count"`
	CommentCount int64  `gorm:"default:0" json:"comment_count"`
	ViewCount    int64  `gorm:"default:0" json:"view_count"`
	Status       int    `gorm:"type:tinyint;default:1" json:"status"` // 1: 启用 0: 禁用
}

// TableName 表名
func (DailyQuestion) TableName() string {
	return "daily_questions"
}
