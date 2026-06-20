package entity

// User 用户
type User struct {
	BaseEntity
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Nickname string `gorm:"type:varchar(50)" json:"nickname"`
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Avatar   string `gorm:"type:varchar(500)" json:"avatar"`
	Bio      string `gorm:"type:text" json:"bio"`
	Status   int    `gorm:"type:tinyint;default:1" json:"status"` // 1: 启用 0: 禁用
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
