package service

import (
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	"blog/pkg/logger"
	"fmt"
)

// AuditLogService 审计日志服务接口
type AuditLogService interface {
	// Create 创建审计日志
	Create(operatorID uint, operatorName, action, targetType string, targetID uint, targetTitle, detail, ip, userAgent string)

	// GetList 获取审计日志列表
	GetList(page, pageSize int, action, targetType, startDate, endDate string) (*response.PageResponse, error)
}

// auditLogService 审计日志服务实现
type auditLogService struct {
	auditLogRepo repository.AuditLogRepository
}

// NewAuditLogService 创建审计日志服务
func NewAuditLogService(auditLogRepo repository.AuditLogRepository) AuditLogService {
	return &auditLogService{
		auditLogRepo: auditLogRepo,
	}
}

// Create 创建审计日志（不阻塞主流程）
func (s *auditLogService) Create(operatorID uint, operatorName, action, targetType string, targetID uint, targetTitle, detail, ip, userAgent string) {
	log := &entity.AuditLog{
		OperatorID:   operatorID,
		OperatorName: operatorName,
		Action:       action,
		TargetType:   targetType,
		TargetID:     targetID,
		TargetTitle:  targetTitle,
		Detail:       detail,
		IP:           ip,
		UserAgent:    userAgent,
	}

	if err := s.auditLogRepo.Create(log); err != nil {
		logger.Errorf("创建审计日志失败: %v", err)
	}
}

// GetList 获取审计日志列表
func (s *auditLogService) GetList(page, pageSize int, action, targetType, startDate, endDate string) (*response.PageResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	list, total, err := s.auditLogRepo.List(offset, pageSize, action, targetType, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("查询审计日志失败, %w", err)
	}

	// 转换为 AuditLogResponse 切片
	var items []response.AuditLogResponse
	for _, log := range list {
		items = append(items, response.AuditLogResponse{
			ID:           log.ID,
			OperatorID:   log.OperatorID,
			OperatorName: log.OperatorName,
			Action:       log.Action,
			TargetType:   log.TargetType,
			TargetID:     log.TargetID,
			TargetTitle:  log.TargetTitle,
			Detail:       log.Detail,
			IP:           log.IP,
			CreatedAt:    log.CreatedAt,
		})
	}

	return response.NewPageResponse(items, total, page, pageSize), nil
}
