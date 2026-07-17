package repository

import "blog/internal/model/entity"

// CommentRepository 评论数据访问接口
type CommentRepository interface {
	// FindByID 根据 ID 查找评论
	FindByID(id uint) (*entity.Comment, error)

	// Create 创建评论
	Create(comment *entity.Comment) error

	// Update 更新评论
	Update(comment *entity.Comment) error

	// Delete 删除评论（软删除）
	Delete(id uint) error

	// ListByArticleID 根据文章ID获取评论列表，sortBy: asc/desc/hot
	ListByArticleID(articleID uint, offset, limit int, sortBy string) ([]*entity.Comment, int64, error)

	// AdminList 评论列表（后台）
	AdminList(offset, limit int, status string) ([]*entity.Comment, int64, error)

	// Count 统计评论数量
	Count(status string) (int64, error)

	// CountByArticleID 统计文章评论数量
	CountByArticleID(articleID uint) (int64, error)

	// BatchUpdateStatus 批量更新状态
	BatchUpdateStatus(ids []uint, status string) error

	// BatchDelete 批量删除
	BatchDelete(ids []uint) error

	// IncrementLikeCount 增加评论点赞数
	IncrementLikeCount(commentID uint) error

	// CreateLikeLog 创建点赞记录
	CreateLikeLog(log *entity.CommentLikeLog) error

	// HasLiked 检查是否已点赞
	HasLiked(commentID uint, visitorIP string) (bool, error)

	// DecrementLikeCount 减少评论点赞数
	DecrementLikeCount(commentID uint) error

	// DeleteLikeLog 删除点赞记录
	DeleteLikeLog(commentID uint, visitorIP string) error
}
