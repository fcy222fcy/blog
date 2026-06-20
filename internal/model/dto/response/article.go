package response

import "time"

// ArticleResponse 文章响应
type ArticleResponse struct {
	ID           uint              `json:"id"`
	Title        string            `json:"title"`
	Slug         string            `json:"slug"`
	Content      string            `json:"content,omitempty"`
	Summary      string            `json:"summary"`
	Cover        string            `json:"cover"`
	Category     CategoryResponse  `json:"category"`
	Tags         []TagResponse     `json:"tags"`
	ViewCount    int64             `json:"view_count"`
	LikeCount    int64             `json:"like_count"`
	CommentCount int64             `json:"comment_count"`
	Status       string            `json:"status"`
	IsTop        bool              `json:"is_top"`
	ReadingTime  int               `json:"reading_time"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

// ArticleListResponse 文章列表响应（简化版）
type ArticleListResponse struct {
	ID           uint             `json:"id"`
	Title        string           `json:"title"`
	Slug         string           `json:"slug"`
	Summary      string           `json:"summary"`
	Cover        string           `json:"cover"`
	Category     CategoryResponse `json:"category"`
	Tags         []TagResponse    `json:"tags"`
	ViewCount    int64            `json:"view_count"`
	LikeCount    int64            `json:"like_count"`
	CommentCount int64            `json:"comment_count"`
	ReadingTime  int              `json:"reading_time"`
	CreatedAt    time.Time        `json:"created_at"`
}

// ArchiveResponse 归档响应
type ArchiveResponse struct {
	Year     int                    `json:"year"`
	Articles []ArchiveArticleResponse `json:"articles"`
}

// ArchiveArticleResponse 归档文章响应
type ArchiveArticleResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}
