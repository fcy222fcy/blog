package service

import (
	"fmt"
	"strings"
	// "time"

	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"go.uber.org/zap"
)

type VisitorService interface {
	// 获取或创建访客
	GetOrCreateVisitor(nickname, email, website, ip string) (*entity.Visitor, error)
	// 更新访客信息
	UpdateVisitor(visitor *entity.Visitor) error
	// 获取访客管理列表
	GetVisitorList(req *request.VisitorListRequest) (*response.PageResponse, error)
	// 拉黑/解禁访客
	UpdateVisitorStatus(id uint, isBlocked bool, notes string) error
	// 验证访客邮箱
	VerifyVisitorEmail(id uint) error
}

type visitorService struct {
	visitorRepo repository.VisitorRepository
}

func NewVisitorService(visitorRepo repository.VisitorRepository) VisitorService {
	return &visitorService{
		visitorRepo: visitorRepo,
	}
}

// GetOrCreateVisitor 获取或创建访客
func (s *visitorService) GetOrCreateVisitor(nickname, email, website, ip string) (*entity.Visitor, error) {
	// 验证邮箱格式
	if !isValidEmail(email) {
		return nil, bizerrors.New(bizerrors.CodeInvalidParams, "邮箱格式不正确")
	}

	// 先尝试通过邮箱查找现有访客
	visitor, err := s.visitorRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("查询访客失败: %w", err)
	}

	if visitor != nil {
		// 更新现有访客信息
		needUpdate := false
		if visitor.Nickname != nickname {
			visitor.Nickname = nickname
			needUpdate = true
		}
		if visitor.Website != website {
			visitor.Website = website
			needUpdate = true
		}
		if visitor.LastIP != ip {
			visitor.LastIP = ip
			needUpdate = true
		}

		if needUpdate {
			if err := s.visitorRepo.Update(visitor); err != nil {
				logger.Warn("更新访客信息失败", zap.Error(err))
			}
		}
		return visitor, nil
	}

	// 创建新访客
	visitor = &entity.Visitor{
		Nickname:  nickname,
		Email:     email,
		Website:   website,
		LastIP:    ip,
		IsVerified: false,
		IsBlocked:  false,
		CommentCount: 0,
	}

	if err := s.visitorRepo.Create(visitor); err != nil {
		return nil, fmt.Errorf("创建访客失败: %w", err)
	}

	logger.Info("创建新访客", zap.String("email", email), zap.String("nickname", nickname))
	return visitor, nil
}

// UpdateVisitor 更新访客信息
func (s *visitorService) UpdateVisitor(visitor *entity.Visitor) error {
	return s.visitorRepo.Update(visitor)
}

// GetVisitorList 获取访客管理列表
func (s *visitorService) GetVisitorList(req *request.VisitorListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	visitors, total, err := s.visitorRepo.GetVisitorList(req)
	if err != nil {
		return nil, fmt.Errorf("获取访客列表失败: %w", err)
	}

	return &response.PageResponse{
		List:      visitors,
		Total:     total,
		Page:      req.Page,
		Size:      req.PageSize,
		TotalPage: 0,
	}, nil
}

// UpdateVisitorStatus 拉黑/解禁访客
func (s *visitorService) UpdateVisitorStatus(id uint, isBlocked bool, notes string) error {
	visitor, err := s.visitorRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询访客失败: %w", err)
	}
	if visitor == nil {
		return bizerrors.New(bizerrors.CodeRecordNotFound, "访客不存在")
	}

	visitor.IsBlocked = isBlocked
	if notes != "" {
		visitor.Notes = notes
	}

	if err := s.visitorRepo.Update(visitor); err != nil {
		return fmt.Errorf("更新访客状态失败: %w", err)
	}

	action := "解禁"
	if isBlocked {
		action = "拉黑"
	}
	logger.Info("访客状态更新", zap.Uint("id", id), zap.String("action", action))
	return nil
}

// VerifyVisitorEmail 验证访客邮箱
func (s *visitorService) VerifyVisitorEmail(id uint) error {
	visitor, err := s.visitorRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询访客失败: %w", err)
	}
	if visitor == nil {
		return bizerrors.New(bizerrors.CodeRecordNotFound, "访客不存在")
	}

	visitor.IsVerified = true
	return s.visitorRepo.Update(visitor)
}

// isValidEmail 简单邮箱格式验证
func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}