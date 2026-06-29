package entity

// Article 文章
type Article struct {
	BaseEntity
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Slug        string     `gorm:"type:varchar(200);uniqueIndex;not null" json:"slug"`
	Content     string     `gorm:"type:longtext" json:"content"`
	Summary     string     `gorm:"type:text" json:"summary"`
	Cover       string     `gorm:"type:varchar(500)" json:"cover"`
	CategoryID  uint       `gorm:"index" json:"category_id"`
	Category    Category   `gorm:"foreignKey:CategoryID" json:"category"`
	Tags        []Tag      `gorm:"many2many:article_tags;" json:"tags"`
	ViewCount   int64      `gorm:"default:0" json:"view_count"`
	CommentCount int64     `gorm:"default:0" json:"comment_count"`
	Status      string     `gorm:"type:varchar(20);default:published" json:"status"` // published: 已发布 draft: 草稿
	IsTop       bool       `gorm:"default:false" json:"is_top"`
	ReadingTime int        `gorm:"default:0" json:"reading_time"` // 阅读时长（分钟）
}

// TableName 表名
func (Article) TableName() string {
	return "articles"
}
