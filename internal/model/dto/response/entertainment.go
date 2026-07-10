package response

import "time"

// EntertainmentResponse 娱乐响应
type EntertainmentResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Type           string    `json:"type"`
	Year           int       `json:"year"`
	Cover          string    `json:"cover"`
	Rating         float64   `json:"rating"`
	RatingExternal float64   `json:"rating_external"`
	Platform       string    `json:"platform"`
	Playtime       string    `json:"playtime"`
	Comment        string    `json:"comment"`
	Status         string    `json:"status"`
	Link           string    `json:"link"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
