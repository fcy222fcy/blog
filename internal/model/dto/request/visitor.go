package request

// VisitorListRequest 访客列表请求
type VisitorListRequest struct {
	PageRequest
	Keyword  string `json:"keyword" form:"keyword"`
	IsBlocked *bool `json:"is_blocked" form:"is_blocked"`
}
