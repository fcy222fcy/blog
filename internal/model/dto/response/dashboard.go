package response

// DashboardStatsResponse 仪表盘统计响应
type DashboardStatsResponse struct {
	ArticleCount       int64 `json:"article_count"`
	PublishedCount     int64 `json:"published_count"`
	DraftCount         int64 `json:"draft_count"`
	TotalViews         int64 `json:"total_views"`
	TotalLikes         int64 `json:"total_likes"`
	CommentCount       int64 `json:"comment_count"`
	PendingCommentCount int64 `json:"pending_comment_count"`
	LinkCount          int64 `json:"link_count"`
	CategoryCount      int64 `json:"category_count"`
	TagCount           int64 `json:"tag_count"`
	DailyQuestionCount int64 `json:"daily_question_count"`
}

// RecentArticleResponse 最近文章响应
type RecentArticleResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ViewCount int64  `json:"view_count"`
	CreatedAt string `json:"created_at"`
}
