package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"time"
)

// dailyQuestionService 每日一问服务实现
type dailyQuestionService struct {
	dailyQuestionRepo repository.DailyQuestionRepository
}

// NewDailyQuestionService 创建每日一问服务
func NewDailyQuestionService(dailyQuestionRepo repository.DailyQuestionRepository) DailyQuestionService {
	return &dailyQuestionService{dailyQuestionRepo: dailyQuestionRepo}
}

// GetTodayQuestion 获取今日问题
func (s *dailyQuestionService) GetTodayQuestion() (*response.DailyQuestionResponse, error) {
	today := time.Now().Format("2006-01-02")
	return s.GetQuestionByDate(today)
}

// GetQuestionByDate 获取指定日期问题
func (s *dailyQuestionService) GetQuestionByDate(date string) (*response.DailyQuestionResponse, error) {
	question, err := s.dailyQuestionRepo.FindByDate(date)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "该日期暂无问题")
	}

	return &response.DailyQuestionResponse{
		ID:           question.ID,
		Question:     question.Question,
		Answer:       question.Answer,
		Date:         question.Date,
		LikeCount:    question.LikeCount,
		CommentCount: question.CommentCount,
		ViewCount:    question.ViewCount,
		Status:       question.Status,
		CreatedAt:    question.CreatedAt,
	}, nil
}

// LikeDailyQuestion 每日一问点赞
func (s *dailyQuestionService) LikeDailyQuestion(id uint) (int64, error) {
	err := s.dailyQuestionRepo.IncrementLikeCount(id)
	if err != nil {
		return 0, err
	}
	question, err := s.dailyQuestionRepo.FindByID(id)
	if err != nil {
		return 0, err
	}
	return question.LikeCount, nil
}

// GetAdminQuestionList 获取问题列表（后台）
func (s *dailyQuestionService) GetAdminQuestionList(req *request.DailyQuestionListRequest) (*response.PageResponse, error) {
	list, total, err := s.dailyQuestionRepo.List(req.GetOffset(), req.GetPageSize(), req.Status, req.Keyword)
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
}

// CreateDailyQuestion 创建问题
func (s *dailyQuestionService) CreateDailyQuestion(req *request.CreateDailyQuestionRequest) (uint, error) {
	// 检查日期是否已存在
	existing, _ := s.dailyQuestionRepo.FindByDate(req.Date)
	if existing != nil {
		return 0, bizerrors.New(bizerrors.CodeInvalidParams, "该日期已有问题")
	}

	question := &entity.DailyQuestion{
		Question: req.Question,
		Answer:   req.Answer,
		Date:     req.Date,
		Status:   req.Status,
	}

	if question.Status == 0 {
		question.Status = 1
	}

	err := s.dailyQuestionRepo.Create(question)
	if err != nil {
		return 0, err
	}
	return question.ID, nil
}

// UpdateDailyQuestion 更新问题
func (s *dailyQuestionService) UpdateDailyQuestion(id uint, req *request.UpdateDailyQuestionRequest) error {
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
		question.Date = req.Date
	}
	if req.Status == 0 || req.Status == 1 {
		question.Status = req.Status
	}

	return s.dailyQuestionRepo.Update(question)
}

// DeleteDailyQuestion 删除问题
func (s *dailyQuestionService) DeleteDailyQuestion(id uint) error {
	question, err := s.dailyQuestionRepo.FindByID(id)
	if err != nil {
		return err
	}
	if question == nil {
		return bizerrors.New(bizerrors.CodeDailyQuestionNotFound, "问题不存在")
	}
	return s.dailyQuestionRepo.Delete(id)
}
