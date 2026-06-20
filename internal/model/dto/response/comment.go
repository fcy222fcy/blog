package response

import "time"

// CommentResponse 评论响应
type CommentResponse struct {
	ID         uint              `json:"id"`
	Content    string            `json:"content"`
	Nickname   string            `json:"nickname"`
	Email      string            `json:"email"`
	Website    string            `json:"website"`
	Avatar     string            `json:"avatar"`
	Status     string            `json:"status"`
	Article    ArticleBriefResponse `json:"article"`
	ParentID   uint              `json:"parent_id"`
	Replies    []CommentResponse `json:"replies,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
}

// ArticleBriefResponse 文章简要响应
type ArticleBriefResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
