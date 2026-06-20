package response

// PageResponse 分页响应
type PageResponse struct {
	List      interface{} `json:"list"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	TotalPage int         `json:"total_page"`
}

// NewPageResponse 创建分页响应
func NewPageResponse(list interface{}, total int64, page, size int) *PageResponse {
	totalPage := int(total) / size
	if int(total)%size > 0 {
		totalPage++
	}
	return &PageResponse{
		List:      list,
		Total:     total,
		Page:      page,
		Size:      size,
		TotalPage: totalPage,
	}
}
