package request

// ArticleListRequest 文章列表请求
type ArticleListRequest struct {
	PageRequest
	Category uint   `json:"category" form:"category"`
	Tag      uint   `json:"tag" form:"tag"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   string `json:"status" form:"status"`
}

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required,min=1,max=200"`
	Content    string `json:"content" binding:"required"`
	Summary    string `json:"summary"`
	Cover      string `json:"cover"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids"`
	Status     string `json:"status" binding:"omitempty,oneof=published draft"`
	IsTop      bool   `json:"is_top"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title      string `json:"title" binding:"min=1,max=200"`
	Content    string `json:"content"`
	Summary    string `json:"summary"`
	Cover      string `json:"cover"`
	CategoryID uint   `json:"category_id"`
	TagIDs     []uint `json:"tag_ids"`
	Status     string `json:"status" binding:"omitempty,oneof=published draft"`
	IsTop      bool   `json:"is_top"`
}
