package request

// EntertainmentListRequest 娱乐列表请求
type EntertainmentListRequest struct {
	PageRequest
	Type    string `json:"type" form:"type"`
	Status  string `json:"status" form:"status"`
	Keyword string `json:"keyword" form:"keyword"`
	Year    *int   `json:"year" form:"year"`
}

// CreateEntertainmentRequest 创建娱乐请求
type CreateEntertainmentRequest struct {
	Title          string  `json:"title" binding:"required,min=1,max=200"`
	TitleEn        string  `json:"title_en" binding:"max=200"`
	Type           string  `json:"type" binding:"required,oneof=movie tv game"`
	Year           int     `json:"year" binding:"required,min=1900,max=2100"`
	Cover          string  `json:"cover" binding:"max=500"`
	Rating         float64 `json:"rating" binding:"omitempty,min=0,max=10"`
	RatingExternal float64 `json:"rating_external" binding:"omitempty,min=0,max=10"`
	Platform       string  `json:"platform" binding:"max=50"`
	Playtime       string  `json:"playtime" binding:"max=20"`
	Comment        string  `json:"comment"`
	Status         string  `json:"status" binding:"omitempty,oneof=watching watched playing completed"`
	Link           string  `json:"link" binding:"max=500"`
}

// UpdateEntertainmentRequest 更新娱乐请求
type UpdateEntertainmentRequest struct {
	Title          string  `json:"title" binding:"min=1,max=200"`
	TitleEn        string  `json:"title_en" binding:"max=200"`
	Type           string  `json:"type" binding:"omitempty,oneof=movie tv game"`
	Year           int     `json:"year" binding:"omitempty,min=1900,max=2100"`
	Cover          string  `json:"cover" binding:"max=500"`
	Rating         float64 `json:"rating" binding:"omitempty,min=0,max=10"`
	RatingExternal float64 `json:"rating_external" binding:"omitempty,min=0,max=10"`
	Platform       string  `json:"platform" binding:"max=50"`
	Playtime       string  `json:"playtime" binding:"max=20"`
	Comment        string  `json:"comment"`
	Status         string  `json:"status" binding:"omitempty,oneof=watching watched playing completed"`
	Link           string  `json:"link" binding:"max=500"`
}
