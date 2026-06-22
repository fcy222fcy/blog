package request

// DailyQuestionListRequest 每日一问列表请求
type DailyQuestionListRequest struct {
	PageRequest
	Status  *int   `json:"status" form:"status"`
	Keyword string `json:"keyword" form:"keyword"`
	Date    string `json:"date" form:"date"`
}

// CreateDailyQuestionRequest 创建每日一问请求
type CreateDailyQuestionRequest struct {
	Question string `json:"question" binding:"required,min=1"`
	Answer   string `json:"answer" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Status   int    `json:"status" binding:"omitempty,oneof=0 1"`
}

// UpdateDailyQuestionRequest 更新每日一问请求
type UpdateDailyQuestionRequest struct {
	Question string `json:"question" binding:"min=1"`
	Answer   string `json:"answer"`
	Date     string `json:"date"`
	Status   int    `json:"status" binding:"omitempty,oneof=0 1"`
}
