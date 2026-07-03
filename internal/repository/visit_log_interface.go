package repository

import (
	"blog/internal/model/entity"
	"time"
)

// VisitRepository 访问日志数据访问接口
type VisitRepository interface {
	// Create 创建访问日志
	Create(log *entity.VisitLog) error

	// HasVisited 检查指定 IP 在给定时间范围内是否已访问过文章
	HasVisited(articleID uint, ip string, since time.Time) (bool, error)

	// CountTodayViews 统计今日浏览量
	CountTodayViews() (int64, error)

	// CountTodayViewsByArticle 统计文章今日浏览量
	CountTodayViewsByArticle(articleID uint) (int64, error)
}
