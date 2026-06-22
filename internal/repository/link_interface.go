package repository

import "blog/internal/model/entity"

// LinkRepository 友链数据访问接口
type LinkRepository interface {
	// FindByID 根据 ID 查找友链
	FindByID(id uint) (*entity.Link, error)

	// Create 创建友链
	Create(link *entity.Link) error

	// Update 更新友链
	Update(link *entity.Link) error

	// Delete 删除友链（软删除）
	Delete(id uint) error

	// List 友链列表（前台，只返回已审核的）
	List() ([]*entity.Link, error)

	// AdminList 友链列表（后台，所有状态）
	AdminList(offset, limit int, status string) ([]*entity.Link, int64, error)

	// Count 统计友链数量
	Count(status string) (int64, error)

	// BatchUpdateStatus 批量更新状态
	BatchUpdateStatus(ids []uint, status string) error

	// BatchDelete 批量删除
	BatchDelete(ids []uint) error
}
