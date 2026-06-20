package entity

// Tag 标签
type Tag struct {
	BaseEntity
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Slug string `gorm:"type:varchar(50);uniqueIndex;not null" json:"slug"`
}

// TableName 表名
func (Tag) TableName() string {
	return "tags"
}
