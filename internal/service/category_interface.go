package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// CategoryService 分类服务接口
type CategoryService interface {
	// GetCategoryList 获取分类列表
	GetCategoryList() ([]response.CategoryResponse, error)

	// CreateCategory 创建分类
	CreateCategory(req *request.CreateCategoryRequest) (uint, error)

	// UpdateCategory 更新分类
	UpdateCategory(id uint, req *request.UpdateCategoryRequest) error

	// DeleteCategory 删除分类
	DeleteCategory(id uint) error
}
