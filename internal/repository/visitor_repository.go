package repository

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/entity"

	"gorm.io/gorm"
)

// visitorRepository 访客数据访问实现
type visitorRepository struct {
	db *gorm.DB
}

// NewVisitorRepository 创建访客仓库
func NewVisitorRepository(db *gorm.DB) VisitorRepository {
	return &visitorRepository{db: db}
}

func (r *visitorRepository) FindByID(id uint) (*entity.Visitor, error) {
	var visitor entity.Visitor
	err := r.db.First(&visitor, id).Error
	if err != nil {
		return nil, err
	}
	return &visitor, nil
}

func (r *visitorRepository) FindByEmail(email string) (*entity.Visitor, error) {
	var visitor entity.Visitor
	err := r.db.Where("email = ?", email).First(&visitor).Error
	if err != nil {
		return nil, err
	}
	return &visitor, nil
}

func (r *visitorRepository) Create(visitor *entity.Visitor) error {
	return r.db.Create(visitor).Error
}

func (r *visitorRepository) Update(visitor *entity.Visitor) error {
	return r.db.Save(visitor).Error
}

func (r *visitorRepository) GetVisitorList(req *request.VisitorListRequest) ([]*entity.Visitor, int64, error) {
	var visitors []*entity.Visitor
	var total int64

	query := r.db.Model(&entity.Visitor{})
	if req.Keyword != "" {
		query = query.Where("nickname LIKE ? OR email LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	err := query.Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&visitors).Error
	return visitors, total, err
}
