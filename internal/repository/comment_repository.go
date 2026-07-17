package repository

import (
	"blog/internal/model/entity"
	"sort"
	"time"

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

// calcReplyCount 计算评论的总回复数（包括所有嵌套层级的回复）
func calcReplyCount(comment *entity.Comment) int {
	count := len(comment.Replies)
	for i := range comment.Replies {
		count += calcReplyCount(&comment.Replies[i])
	}
	return count
}

// calcHotScore 计算评论热度分：点赞权重更高 + 所有层级回复数
func calcHotScore(comment *entity.Comment) int {
	return comment.LikeCount*2 + calcReplyCount(comment)
}

// buildReplyTree 基于全部评论，为传入的根评论切片构建回复树
func buildReplyTree(rootComments []*entity.Comment, allComments []*entity.Comment) {
	rootIDSet := make(map[uint]struct{}, len(rootComments))
	for _, rc := range rootComments {
		rootIDSet[rc.ID] = struct{}{}
	}

	rootCommentMap := make(map[uint]*entity.Comment)
	for _, c := range rootComments {
		rootCommentMap[c.ID] = c
	}

	allCommentMap := make(map[uint]*entity.Comment)
	for _, c := range allComments {
		allCommentMap[c.ID] = c
	}

	for _, c := range allComments {
		if c.ParentID == nil {
			continue
		}

		// 先找父评论是否是根评论（属于当前分页的根集合）
		if parent, ok := rootCommentMap[*c.ParentID]; ok {
			parent.Replies = append(parent.Replies, *c)
		} else if parent, ok := allCommentMap[*c.ParentID]; ok {
			// 父评论不是根评论，需要确认父评论的最终根是否在当前分页中
			ancestor := parent
			for ancestor.ParentID != nil {
				if p, ok := allCommentMap[*ancestor.ParentID]; ok {
					ancestor = p
				} else {
					break
				}
			}
			if _, ok := rootIDSet[ancestor.ID]; ok {
				parent.Replies = append(parent.Replies, *c)
			}
		}
	}

	for _, c := range allComments {
		if c.ReplyToID != nil {
			if target, ok := allCommentMap[*c.ReplyToID]; ok {
				c.ReplyToNickname = target.Nickname
			}
		} else if c.ParentID != nil {
			if parent, ok := allCommentMap[*c.ParentID]; ok {
				c.ReplyToNickname = parent.Nickname
			}
		}
	}
}

// ListByArticleID 根据文章ID获取评论列表（构建树结构，sortBy: asc/desc/hot）
// 先对全量根评论在内存中排序 → 再切片分页，保证分页正确性
func (r *commentRepository) ListByArticleID(articleID uint, offset, limit int, sortBy string) ([]*entity.Comment, int64, error) {
	var allRootComments []*entity.Comment
	var total int64

	query := r.db.Model(&entity.Comment{}).Where("article_id = ? AND status = ? AND parent_id IS NULL", articleID, "approved")
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&entity.Comment{}).
		Where("article_id = ? AND status = ? AND parent_id IS NULL", articleID, "approved").
		Order("created_at DESC").
		Find(&allRootComments).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取该文章所有已批准的评论（用于构建回复树和计算热度）
	var allComments []*entity.Comment
	err = r.db.Model(&entity.Comment{}).
		Where("article_id = ? AND status = ?", articleID, "approved").
		Order("created_at ASC").
		Find(&allComments).Error
	if err != nil {
		return nil, 0, err
	}

	// 先对全量根评论构建完整回复树，以便计算每个根评论的真实回复数
	preBuildRoot := make([]*entity.Comment, len(allRootComments))
	for i := range allRootComments {
		clone := *allRootComments[i]
		clone.Replies = nil
		preBuildRoot[i] = &clone
	}
	buildReplyTree(preBuildRoot, allComments)

	// 根据 sortBy 对全量根评论排序（用 preBuildRoot 中的热度分，但排序 allRootComments）
	hotScores := make([]int, len(preBuildRoot))
	for i, rc := range preBuildRoot {
		hotScores[i] = calcHotScore(rc)
	}

	switch sortBy {
	case "asc":
		sort.SliceStable(allRootComments, func(i, j int) bool {
			return allRootComments[i].CreatedAt.Before(allRootComments[j].CreatedAt)
		})
		for i := range hotScores {
			_ = hotScores[i]
		}
	case "hot":
		// 预计算每个 allRootComments[i] 对应的索引：用 ID 匹配
		idToHot := make(map[uint]int, len(preBuildRoot))
		idToTime := make(map[uint]time.Time, len(preBuildRoot))
		for i, rc := range preBuildRoot {
			idToHot[rc.ID] = hotScores[i]
			idToTime[rc.ID] = rc.CreatedAt
		}
		sort.SliceStable(allRootComments, func(i, j int) bool {
			hi := idToHot[allRootComments[i].ID]
			hj := idToHot[allRootComments[j].ID]
			if hi != hj {
				return hi > hj
			}
			ti := allRootComments[i].CreatedAt
			tj := allRootComments[j].CreatedAt
			return ti.After(tj)
		})
	default: // desc
		sort.SliceStable(allRootComments, func(i, j int) bool {
			return allRootComments[i].CreatedAt.After(allRootComments[j].CreatedAt)
		})
	}

	// 切片分页
	start := offset
	end := offset + limit
	if start > int(total) {
		start = int(total)
	}
	if end > int(total) {
		end = int(total)
	}
	pagedRoots := allRootComments[start:end]
	for _, rc := range pagedRoots {
		rc.Replies = nil
	}

	// 只为分页后的根评论构建回复树
	buildReplyTree(pagedRoots, allComments)

	return pagedRoots, total, nil
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
		Preload("Article").
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

// DecrementLikeCount 减少评论点赞数
func (r *commentRepository) DecrementLikeCount(commentID uint) error {
	return r.db.Model(&entity.Comment{}).
		Where("id = ? AND like_count > 0", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).
		Error
}

// DeleteLikeLog 删除点赞记录
func (r *commentRepository) DeleteLikeLog(commentID uint, visitorIP string) error {
	return r.db.Where("comment_id = ? AND visitor_ip = ?", commentID, visitorIP).
		Delete(&entity.CommentLikeLog{}).
		Error
}
