package response

import "time"

// LinkResponse 友链响应
type LinkResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar"`
	Logo        string    `json:"logo"`
	SortOrder   int       `json:"sort_order"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
