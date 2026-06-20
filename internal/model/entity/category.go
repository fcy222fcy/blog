package entity

// Category 分类
type Category struct {
	BaseEntity
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Slug        string `gorm:"type:varchar(50);uniqueIndex;not null" json:"slug"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Icon        string `gorm:"type:varchar(10)" json:"icon"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`
}

// TableName 表名
func (Category) TableName() string {
	return "categories"
}
