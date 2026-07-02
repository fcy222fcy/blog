package repository

import (
	"blog/internal/model/entity"

	"gorm.io/gorm"
)

type aboutPageRepository struct {
	db *gorm.DB
}

// NewAboutPageRepository 创建关于页面仓库
func NewAboutPageRepository(db *gorm.DB) AboutPageRepository {
	return &aboutPageRepository{db: db}
}

// Get 获取关于页面
func (r *aboutPageRepository) Get() (*entity.AboutPage, error) {
	var page entity.AboutPage
	err := r.db.First(&page).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &page, nil
}

// Save 保存关于页面
func (r *aboutPageRepository) Save(page *entity.AboutPage) error {
	// 检查是否已存在记录
	var count int64
	r.db.Model(&entity.AboutPage{}).Count(&count)

	if count == 0 {
		// 创建新记录
		return r.db.Create(page).Error
	}
	// 更新现有记录（使用 ID=1）
	return r.db.Model(&entity.AboutPage{}).Where("id = ?", page.ID).Updates(page).Error
}
