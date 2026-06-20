package entity

// Entertainment 娱乐（影视/游戏）
type Entertainment struct {
	BaseEntity
	Title          string  `gorm:"type:varchar(200);not null" json:"title"`
	TitleEn        string  `gorm:"type:varchar(200)" json:"title_en"`
	Type           string  `gorm:"type:varchar(20);not null" json:"type"` // movie: 电影 tv: 剧集 game: 游戏
	Year           int     `gorm:"not null" json:"year"`
	Cover          string  `gorm:"type:varchar(500)" json:"cover"`
	Rating         float64 `gorm:"type:decimal(3,1)" json:"rating"`
	RatingExternal float64 `gorm:"type:decimal(3,1)" json:"rating_external"`
	Platform       string  `gorm:"type:varchar(50)" json:"platform"`
	Playtime       string  `gorm:"type:varchar(20)" json:"playtime"`
	Comment        string  `gorm:"type:text" json:"comment"`
	Status         string  `gorm:"type:varchar(20)" json:"status"` // watching: 在看 watched: 已看 playing: 在玩 completed: 已完成
	Link           string  `gorm:"type:varchar(500)" json:"link"`
}

// TableName 表名
func (Entertainment) TableName() string {
	return "entertainment"
}
