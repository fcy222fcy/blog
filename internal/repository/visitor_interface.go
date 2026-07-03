package repository

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/entity"
)

// VisitorRepository 访客数据访问接口
type VisitorRepository interface {
	FindByID(id uint) (*entity.Visitor, error)
	FindByEmail(email string) (*entity.Visitor, error)
	Create(visitor *entity.Visitor) error
	Update(visitor *entity.Visitor) error
	GetVisitorList(req *request.VisitorListRequest) ([]*entity.Visitor, int64, error)
}
