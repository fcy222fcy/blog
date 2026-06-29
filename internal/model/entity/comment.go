package entity

// Comment 评论
type Comment struct {
	BaseEntity
	Content      string    `gorm:"type:text;not null" json:"content"`
	Nickname     string    `gorm:"type:varchar(50)" json:"nickname"`
	Email        string    `gorm:"type:varchar(100)" json:"email"`
	Website      string    `gorm:"type:varchar(200)" json:"website"`
	Avatar       string    `gorm:"type:varchar(500)" json:"avatar"`
	ArticleID    uint      `gorm:"index" json:"article_id"`
	Article      Article   `gorm:"foreignKey:ArticleID" json:"-"`
	ParentID     *uint     `gorm:"index" json:"parent_id"`
	ReplyTo      *Comment  `gorm:"-" json:"-"`
	ReplyToID    *uint     `gorm:"-" json:"-"`
	ReplyToNickname string `gorm:"-" json:"reply_to_nickname"`
	Replies      []Comment `gorm:"-" json:"replies,omitempty"`
	Status       string    `gorm:"type:varchar(20);default:pending" json:"status"` // pending: 待审核 approved: 已通过 rejected: 已拒绝
	IP           string    `gorm:"type:varchar(50)" json:"ip"`
	UserAgent    string    `gorm:"type:varchar(500)" json:"user_agent"`
}

// TableName 表名
func (Comment) TableName() string {
	return "comments"
}
