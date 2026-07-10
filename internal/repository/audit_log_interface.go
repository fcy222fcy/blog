package repository

import (
	"blog/internal/model/entity"
	"time"
)

// AuditLogRepository 审计日志数据访问接口
type AuditLogRepository interface {
	// Create 创建审计日志
	Create(log *entity.AuditLog) error

	// List 分页查询审计日志
	List(offset, limit int, action, targetType string, startTime, endTime *time.Time) ([]entity.AuditLog, int64, error)
}
