package request

// LinkListRequest 友链列表请求
type LinkListRequest struct {
	PageRequest
	Status string `json:"status" form:"status"`
}

// CreateLinkRequest 创建友链请求
type CreateLinkRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	URL         string `json:"url" binding:"required,url"`
	Description string `json:"description" binding:"max=200"`
	Avatar      string `json:"avatar" binding:"omitempty,max=500"`
	Logo        string `json:"logo" binding:"max=10"`
	SortOrder   int    `json:"sort_order"`
	Status      string `json:"status" binding:"omitempty,oneof=pending approved rejected"`
}

// UpdateLinkRequest 更新友链请求
type UpdateLinkRequest struct {
	Name        string `json:"name" binding:"min=1,max=100"`
	URL         string `json:"url" binding:"url"`
	Description string `json:"description" binding:"max=200"`
	Avatar      string `json:"avatar" binding:"omitempty,max=500"`
	Logo        string `json:"logo" binding:"max=10"`
	SortOrder   int    `json:"sort_order"`
	Status      string `json:"status" binding:"omitempty,oneof=pending approved rejected"`
}

// UpdateLinkStatusRequest 更新友链状态请求
type UpdateLinkStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=approved rejected"`
}

// ApplyLinkRequest 申请友链请求
type ApplyLinkRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	URL         string `json:"url" binding:"required,url"`
	Description string `json:"description" binding:"max=200"`
	Avatar      string `json:"avatar" binding:"omitempty,max=500"`
	Email       string `json:"email" binding:"required,email"`
}
