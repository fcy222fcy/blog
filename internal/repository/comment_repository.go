package repository

import (
	"blog/internal/model/entity"

	"gorm.io/gorm"
)

// commentRepository 评论数据访问实现
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository 创建评论数据访问
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

// FindByID 根据 ID 查找评论
func (r *commentRepository) FindByID(id uint) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.db.First(&comment, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

// Create 创建评论
func (r *commentRepository) Create(comment *entity.Comment) error {
	return r.db.Create(comment).Error
}

// Update 更新评论
func (r *commentRepository) Update(comment *entity.Comment) error {
	return r.db.Save(comment).Error
}

// Delete 删除评论（软删除）
func (r *commentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Comment{}, id).Error
}

// ListByArticleID 根据文章ID获取评论列表
func (r *commentRepository) ListByArticleID(articleID uint, offset, limit int) ([]*entity.Comment, int64, error) {
	var comments []*entity.Comment
	var total int64

	query := r.db.Model(&entity.Comment{}).Where("article_id = ? AND status = ?", articleID, "approved")
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error
	return comments, total, err
}

// AdminList 评论列表（后台）
func (r *commentRepository) AdminList(offset, limit int, status string) ([]*entity.Comment, int64, error) {
	var comments []*entity.Comment
	var total int64

	query := r.db.Model(&entity.Comment{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error
	return comments, total, err
}

// Count 统计评论数量
func (r *commentRepository) Count(status string) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Comment{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&total).Error
	return total, err
}

// CountByArticleID 统计文章评论数量
func (r *commentRepository) CountByArticleID(articleID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comment{}).
		Where("article_id = ? AND status = ?", articleID, "approved").
		Count(&count).Error
	return count, err
}

// BatchUpdateStatus 批量更新状态
func (r *commentRepository) BatchUpdateStatus(ids []uint, status string) error {
	return r.db.Model(&entity.Comment{}).Where("id IN ?", ids).
		Update("status", status).Error
}

// BatchDelete 批量删除
func (r *commentRepository) BatchDelete(ids []uint) error {
	return r.db.Delete(&entity.Comment{}, ids).Error
}
