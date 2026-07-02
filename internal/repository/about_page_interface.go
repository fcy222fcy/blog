package repository

import "blog/internal/model/entity"

// AboutPageRepository 关于页面数据访问接口
type AboutPageRepository interface {
	// Get 获取关于页面（只有一条记录）
	Get() (*entity.AboutPage, error)

	// Save 保存关于页面（更新或创建）
	Save(page *entity.AboutPage) error
}
