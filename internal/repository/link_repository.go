package repository

import (
	"blog/internal/model/entity"

	"gorm.io/gorm"
)

// linkRepository 友链数据访问实现
type linkRepository struct {
	db *gorm.DB
}

// NewLinkRepository 创建友链数据访问
func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{db: db}
}

// FindByID 根据 ID 查找友链
func (r *linkRepository) FindByID(id uint) (*entity.Link, error) {
	var link entity.Link
	err := r.db.First(&link, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &link, nil
}

// Create 创建友链
func (r *linkRepository) Create(link *entity.Link) error {
	return r.db.Create(link).Error
}

// Update 更新友链
func (r *linkRepository) Update(link *entity.Link) error {
	return r.db.Save(link).Error
}

// Delete 删除友链（软删除）
func (r *linkRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Link{}, id).Error
}

// List 友链列表（前台）
func (r *linkRepository) List() ([]*entity.Link, error) {
	var links []*entity.Link
	err := r.db.Where("status = ?", "approved").
		Order("sort_order ASC").
		Find(&links).Error
	return links, err
}

// AdminList 友链列表（后台）
func (r *linkRepository) AdminList(offset, limit int, status string) ([]*entity.Link, int64, error) {
	var links []*entity.Link
	var total int64

	query := r.db.Model(&entity.Link{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).
		Order("sort_order ASC").
		Find(&links).Error
	return links, total, err
}

// Count 统计友链数量
func (r *linkRepository) Count(status string) (int64, error) {
	var total int64
	query := r.db.Model(&entity.Link{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&total).Error
	return total, err
}

// BatchUpdateStatus 批量更新状态
func (r *linkRepository) BatchUpdateStatus(ids []uint, status string) error {
	return r.db.Model(&entity.Link{}).Where("id IN ?", ids).
		Update("status", status).Error
}

// BatchDelete 批量删除
func (r *linkRepository) BatchDelete(ids []uint) error {
	return r.db.Delete(&entity.Link{}, ids).Error
}
