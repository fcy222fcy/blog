package repository

import (
	"blog/internal/model/entity"
	"time"

	"gorm.io/gorm"
)

// visitLogRepository 访问日志数据访问实现
type visitLogRepository struct {
	db *gorm.DB
}

// NewVisitRepository 创建访问日志数据访问
func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitLogRepository{db: db}
}

// Create 创建访问日志
func (r *visitLogRepository) Create(log *entity.VisitLog) error {
	return r.db.Create(log).Error
}

// HasVisited 检查指定 IP 在给定时间范围内是否已访问过文章
func (r *visitLogRepository) HasVisited(articleID uint, ip string, since time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&entity.VisitLog{}).
		Where("article_id = ? AND ip = ? AND created_at >= ?", articleID, ip, since).
		Count(&count).Error
	return count > 0, err
}

// CountTodayViews 统计今日浏览量
func (r *visitLogRepository) CountTodayViews() (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := r.db.Model(&entity.VisitLog{}).
		Where("DATE(created_at) = ?", today).
		Count(&count).Error
	return count, err
}

// CountTodayViewsByArticle 统计文章今日浏览量
func (r *visitLogRepository) CountTodayViewsByArticle(articleID uint) (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := r.db.Model(&entity.VisitLog{}).
		Where("article_id = ? AND DATE(created_at) = ?", articleID, today).
		Count(&count).Error
	return count, err
}
