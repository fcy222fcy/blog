package repository

import (
	"blog/internal/model/entity"
	"time"

	"gorm.io/gorm"
)

// auditLogRepository 审计日志数据访问实现
type auditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository 创建审计日志数据访问
func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

// Create 创建审计日志
func (r *auditLogRepository) Create(log *entity.AuditLog) error {
	return r.db.Create(log).Error
}

// List 分页查询审计日志
func (r *auditLogRepository) List(offset, limit int, action, targetType string, startTime, endTime *time.Time) ([]entity.AuditLog, int64, error) {
	var logs []entity.AuditLog
	var total int64

	query := r.db.Model(&entity.AuditLog{})

	if action != "" {
		query = query.Where("action = ?", action)
	}
	if targetType != "" {
		query = query.Where("target_type = ?", targetType)
	}
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
