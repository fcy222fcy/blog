package entity

import "time"

// Visitor 访客用户（匿名评论用户）
type Visitor struct {
	BaseEntity
	Nickname    string    `gorm:"type:varchar(50);not null" json:"nickname"`
	Email       string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Website     string    `gorm:"type:varchar(200)" json:"website"`
	Avatar      string    `gorm:"type:varchar(500)" json:"avatar"`
	IsVerified  bool      `gorm:"default:false" json:"is_verified"`     // 邮箱是否验证
	IsBlocked   bool      `gorm:"default:false" json:"is_blocked"`      // 是否被拉黑
	LastIP      string    `gorm:"type:varchar(50)" json:"last_ip"`
	CommentCount int      `gorm:"default:0" json:"comment_count"`       // 评论数量
	LastCommentAt *time.Time `json:"last_comment_at"`                  // 最后评论时间
	Notes       string    `gorm:"type:text" json:"notes"`              // 管理员备注
}

// TableName 表名
func (Visitor) TableName() string {
	return "visitors"
}