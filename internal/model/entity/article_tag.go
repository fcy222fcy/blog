package entity

// ArticleTag 文章标签关联表
type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey;autoIncrement:false"`
	TagID     uint `gorm:"primaryKey;autoIncrement:false"`
}

// TableName 表名
func (ArticleTag) TableName() string {
	return "article_tags"
}
