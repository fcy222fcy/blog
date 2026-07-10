package repository

import (
	"blog/internal/model/entity"
	"strings"

	"gorm.io/gorm"
)

type entertainmentRepository struct {
	db *gorm.DB
}

func NewEntertainmentRepository(db *gorm.DB) EntertainmentRepository {
	return &entertainmentRepository{db: db}
}

func (r *entertainmentRepository) FindByID(id uint) (*entity.Entertainment, error) {
	var item entity.Entertainment
	err := r.db.First(&item, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *entertainmentRepository) List(offset, limit int, typeStr, status string, year *int, keyword string) ([]*entity.Entertainment, int64, error) {
	var items []*entity.Entertainment
	var total int64

	query := r.db.Model(&entity.Entertainment{})
	if typeStr != "" {
		query = query.Where("type = ?", typeStr)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if year != nil && *year > 0 {
		query = query.Where("year = ?", *year)
	}
	if keyword != "" {
		query = query.Where("title LIKE ? OR title_en LIKE ? OR comment LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).
		Order("year DESC, id DESC").
		Find(&items).Error
	return items, total, err
}

func (r *entertainmentRepository) ListPublic(typeStr string, year *int) ([]*entity.Entertainment, error) {
	var items []*entity.Entertainment
	query := r.db.Model(&entity.Entertainment{})
	if typeStr != "" {
		query = query.Where("type = ?", typeStr)
	}
	if year != nil && *year > 0 {
		query = query.Where("year = ?", *year)
	}
	err := query.Order("year DESC, id DESC").Find(&items).Error
	return items, err
}

func (r *entertainmentRepository) Create(item *entity.Entertainment) error {
	return r.db.Create(item).Error
}

func (r *entertainmentRepository) Update(item *entity.Entertainment) error {
	return r.db.Save(item).Error
}

func (r *entertainmentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Entertainment{}, id).Error
}

func (r *entertainmentRepository) ListYears() ([]int, error) {
	var years []int
	rows, err := r.db.Model(&entity.Entertainment{}).
		Select("DISTINCT year").
		Order("year DESC").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var y int
		if err := rows.Scan(&y); err == nil {
			years = append(years, y)
		}
	}
	return years, nil
}

func (r *entertainmentRepository) CountByStatus(status string) int64 {
	var total int64
	query := r.db.Model(&entity.Entertainment{})
	if strings.TrimSpace(status) != "" {
		query = query.Where("status = ?", status)
	}
	query.Count(&total)
	return total
}
