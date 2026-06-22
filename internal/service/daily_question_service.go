package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
)

// dailyQuestionService 每日一问服务实现
type dailyQuestionService struct {
	dailyQuestionRepo repository.DailyQuestionRepository
}

// NewDailyQuestionService 创建每日一问服务
func NewDailyQuestionService(dailyQuestionRepo repository.DailyQuestionRepository) DailyQuestionService {
	return &dailyQuestionService{dailyQuestionRepo: dailyQuestionRepo}
}

// GetLatestQuestion 获取最新问题
func (s *dailyQuestionService) GetLatestQuestion() (*response.DailyQuestionResponse, error) {
	question, err := s.dailyQuestionRepo.GetLatest()
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "暂无问题")
	}

	_ = s.dailyQuestionRepo.IncrementViewCount(question.ID)
	return s.toResponse(question), nil
}

// GetQuestionByDate 根据日期获取问题
func (s *dailyQuestionService) GetQuestionByDate(date string) (*response.DailyQuestionResponse, error) {
	question, err := s.dailyQuestionRepo.FindByDate(date)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "该日期没有问题")
	}

	_ = s.dailyQuestionRepo.IncrementViewCount(question.ID)
	return s.toResponse(question), nil
}

// GetPreviousQuestion 获取前一天的问题
func (s *dailyQuestionService) GetPreviousQuestion(date string) (*response.DailyQuestionResponse, error) {
	question, err := s.dailyQuestionRepo.GetPrevious(date)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "没有前一天的问题")
	}

	return s.toResponse(question), nil
}

// GetNextQuestion 获取后一天的问题
func (s *dailyQuestionService) GetNextQuestion(date string) (*response.DailyQuestionResponse, error) {
	question, err := s.dailyQuestionRepo.GetNext(date)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "没有后一天的问题")
	}

	return s.toResponse(question), nil
}

// LikeQuestion 问题点赞
func (s *dailyQuestionService) LikeQuestion(id uint) (int64, error) {
	return s.dailyQuestionRepo.IncrementLikeCount(id)
}

// GetAdminQuestionList 获取问题列表（后台）
func (s *dailyQuestionService) GetAdminQuestionList(req *request.DailyQuestionListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	status := -1
	if req.Status != nil {
		status = *req.Status
	}

	list, total, err := s.dailyQuestionRepo.List((req.Page-1)*req.PageSize, req.PageSize, status)
	if err != nil {
		return nil, err
	}

	var result []response.DailyQuestionResponse
	for _, q := range list {
		result = append(result, *s.toResponse(q))
	}

	return response.NewPageResponse(result, total, req.Page, req.PageSize), nil
}

// CreateQuestion 创建问题
func (s *dailyQuestionService) CreateQuestion(req *request.CreateDailyQuestionRequest) (uint, error) {
	existing, _ := s.dailyQuestionRepo.FindByDate(req.Date)
	if existing != nil {
		return 0, bizerrors.New(bizerrors.CodeDailyQuestionDateExists, "该日期已有问题")
	}

	question := &entity.DailyQuestion{
		Question: req.Question,
		Answer:   req.Answer,
		Date:     req.Date,
		Status:   1,
	}

	err := s.dailyQuestionRepo.Create(question)
	if err != nil {
		return 0, err
	}

	return question.ID, nil
}

// UpdateQuestion 更新问题
func (s *dailyQuestionService) UpdateQuestion(id uint, req *request.UpdateDailyQuestionRequest) error {
	question, err := s.dailyQuestionRepo.FindByID(id)
	if err != nil {
		return err
	}
	if question == nil {
		return bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "问题不存在")
	}

	if req.Question != "" {
		question.Question = req.Question
	}
	if req.Answer != "" {
		question.Answer = req.Answer
	}
	if req.Date != "" {
		existing, _ := s.dailyQuestionRepo.FindByDate(req.Date)
		if existing != nil && existing.ID != id {
			return bizerrors.New(bizerrors.CodeDailyQuestionDateExists, "该日期已有问题")
		}
		question.Date = req.Date
	}
	if req.Status != 0 {
		question.Status = req.Status
	}

	return s.dailyQuestionRepo.Update(question)
}

// DeleteQuestion 删除问题
func (s *dailyQuestionService) DeleteQuestion(id uint) error {
	question, err := s.dailyQuestionRepo.FindByID(id)
	if err != nil {
		return err
	}
	if question == nil {
		return bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "问题不存在")
	}

	return s.dailyQuestionRepo.Delete(id)
}

// UpdateQuestionStatus 更新问题状态
func (s *dailyQuestionService) UpdateQuestionStatus(id uint, status int) error {
	question, err := s.dailyQuestionRepo.FindByID(id)
	if err != nil {
		return err
	}
	if question == nil {
		return bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "问题不存在")
	}

	question.Status = status
	return s.dailyQuestionRepo.Update(question)
}

// toResponse 转换为响应DTO
func (s *dailyQuestionService) toResponse(q *entity.DailyQuestion) *response.DailyQuestionResponse {
	return &response.DailyQuestionResponse{
		ID:           q.ID,
		Question:     q.Question,
		Answer:       q.Answer,
		Date:         q.Date,
		LikeCount:    q.LikeCount,
		CommentCount: q.CommentCount,
		ViewCount:    q.ViewCount,
		Status:       q.Status,
		CreatedAt:    q.CreatedAt,
	}
}
