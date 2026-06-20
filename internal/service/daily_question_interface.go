package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// DailyQuestionService 每日一问服务接口
type DailyQuestionService interface {
	// GetTodayQuestion 获取今日问题
	GetTodayQuestion() (*response.DailyQuestionResponse, error)

	// GetQuestionByDate 获取指定日期问题
	GetQuestionByDate(date string) (*response.DailyQuestionResponse, error)

	// LikeDailyQuestion 每日一问点赞
	LikeDailyQuestion(id uint) (int64, error)

	// GetAdminQuestionList 获取问题列表（后台）
	GetAdminQuestionList(req *request.DailyQuestionListRequest) (*response.PageResponse, error)

	// CreateDailyQuestion 创建问题
	CreateDailyQuestion(req *request.CreateDailyQuestionRequest) (uint, error)

	// UpdateDailyQuestion 更新问题
	UpdateDailyQuestion(id uint, req *request.UpdateDailyQuestionRequest) error

	// DeleteDailyQuestion 删除问题
	DeleteDailyQuestion(id uint) error
}
