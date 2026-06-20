package request

// CommentListRequest 评论列表请求
type CommentListRequest struct {
	PageRequest
	Status    string `json:"status" form:"status"`
	ArticleID uint   `json:"article_id" form:"article_id"`
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content   string `json:"content" binding:"required,min=1,max=1000"`
	Nickname  string `json:"nickname" binding:"required,min=1,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Website   string `json:"website" binding:"max=200"`
	ParentID  uint   `json:"parent_id"`
}

// UpdateCommentStatusRequest 更新评论状态请求
type UpdateCommentStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=approved rejected"`
}
