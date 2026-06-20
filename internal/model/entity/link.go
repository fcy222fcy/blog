package entity

// Link 友链
type Link struct {
	BaseEntity
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	URL         string `gorm:"type:varchar(500);not null" json:"url"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Avatar      string `gorm:"type:varchar(500)" json:"avatar"`
	Logo        string `gorm:"type:varchar(10)" json:"logo"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`
	Status      string `gorm:"type:varchar(20);default:approved" json:"status"` // pending: 待审核 approved: 已通过 rejected: 已拒绝
}

// TableName 表名
func (Link) TableName() string {
	return "links"
}
