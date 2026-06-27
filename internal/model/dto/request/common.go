package request

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page" form:"page" binding:"min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// GetOffset 获取偏移量
func (p *PageRequest) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.GetPageSize()
}

// GetPageSize 获取每页数量
func (p *PageRequest) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	return p.PageSize
}
