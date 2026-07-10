package repository

import "blog/internal/model/entity"

// EntertainmentRepository 娱乐数据访问接口
type EntertainmentRepository interface {
	FindByID(id uint) (*entity.Entertainment, error)
	List(offset, limit int, typeStr, status string, year *int, keyword string) ([]*entity.Entertainment, int64, error)
	ListPublic(typeStr string, year *int) ([]*entity.Entertainment, error)
	Create(item *entity.Entertainment) error
	Update(item *entity.Entertainment) error
	Delete(id uint) error
	ListYears() ([]int, error)
}
