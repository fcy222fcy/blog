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

// ListByArticleID 根据文章ID获取评论列表（构建树结构）
func (r *commentRepository) ListByArticleID(articleID uint, offset, limit int) ([]*entity.Comment, int64, error) {
	var rootComments []*entity.Comment
	var total int64

	// 先获取总数（只统计根评论）
	query := r.db.Model(&entity.Comment{}).Where("article_id = ? AND status = ? AND parent_id IS NULL", articleID, "approved")
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取根评论
	err = query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&rootComments).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取该文章所有已批准的评论（用于构建回复树）
	var allComments []*entity.Comment
	err = r.db.Model(&entity.Comment{}).
		Where("article_id = ? AND status = ?", articleID, "approved").
		Order("created_at ASC").
		Find(&allComments).Error
	if err != nil {
		return nil, 0, err
	}

	// 构建评论 ID 到评论的映射（使用根评论列表中的指针）
	rootCommentMap := make(map[uint]*entity.Comment)
	for _, c := range rootComments {
		rootCommentMap[c.ID] = c
	}

	// 构建所有评论的映射
	allCommentMap := make(map[uint]*entity.Comment)
	for _, c := range allComments {
		allCommentMap[c.ID] = c
	}

	// 构建树结构：将子评论添加到父评论的 Replies 中
	for _, c := range allComments {
		if c.ParentID != nil {
			// 先在根评论中查找
			if parent, ok := rootCommentMap[*c.ParentID]; ok {
				parent.Replies = append(parent.Replies, *c)
			} else if parent, ok := allCommentMap[*c.ParentID]; ok {
				// 如果父评论不是根评论，也添加到父评论的 Replies 中
				parent.Replies = append(parent.Replies, *c)
			}
		}
	}

	// 为回复评论设置 ReplyToNickname
	for _, c := range allComments {
		if c.ParentID != nil {
			if parent, ok := allCommentMap[*c.ParentID]; ok {
				c.ReplyToNickname = parent.Nickname
			}
		}
	}

	return rootComments, total, err
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

// IncrementLikeCount 增加评论点赞数
func (r *commentRepository) IncrementLikeCount(commentID uint) error {
	return r.db.Model(&entity.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).
		Error
}

// CreateLikeLog 创建点赞记录
func (r *commentRepository) CreateLikeLog(log *entity.CommentLikeLog) error {
	return r.db.Create(log).Error
}

// HasLiked 检查是否已点赞
func (r *commentRepository) HasLiked(commentID uint, visitorIP string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.CommentLikeLog{}).
		Where("comment_id = ? AND visitor_ip = ?", commentID, visitorIP).
		Count(&count).Error
	return count > 0, err
}
