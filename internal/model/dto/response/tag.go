package response

// TagResponse 标签响应
type TagResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ArticleCount int64  `json:"article_count"`
}
