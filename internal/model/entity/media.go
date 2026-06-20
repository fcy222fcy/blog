package entity

// Media 媒体文件
type Media struct {
	BaseEntity
	Filename string `gorm:"type:varchar(255);not null" json:"filename"`
	URL      string `gorm:"type:varchar(500);not null" json:"url"`
	Size     int64  `gorm:"not null" json:"size"`
	MimeType string `gorm:"type:varchar(100)" json:"mime_type"`
	Type     string `gorm:"type:varchar(20)" json:"type"` // image: 图片 document: 文档
}

// TableName 表名
func (Media) TableName() string {
	return "media"
}
