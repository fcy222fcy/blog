package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
)

// categoryService 分类服务实现
type categoryService struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

// GetCategoryList 获取分类列表
func (s *categoryService) GetCategoryList() ([]response.CategoryResponse, error) {
	categories, err := s.categoryRepo.List()
	if err != nil {
		return nil, err
	}

	var result []response.CategoryResponse
	for _, cat := range categories {
		count, _ := s.categoryRepo.GetCategoryArticleCount(cat.ID)
		result = append(result, response.CategoryResponse{
			ID:           cat.ID,
			Name:         cat.Name,
			Slug:         cat.Slug,
			Description:  cat.Description,
			Icon:         cat.Icon,
			SortOrder:    cat.SortOrder,
			ArticleCount: count,
		})
	}
	return result, nil
}

// GetCategoryByID 根据ID获取分类
func (s *categoryService) GetCategoryByID(id uint) (*response.CategoryResponse, error) {
	cat, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, bizerrors.New(bizerrors.CodeCategoryNotFound, "分类不存在")
	}

	count, _ := s.categoryRepo.GetCategoryArticleCount(cat.ID)
	return &response.CategoryResponse{
		ID:           cat.ID,
		Name:         cat.Name,
		Slug:         cat.Slug,
		Description:  cat.Description,
		Icon:         cat.Icon,
		SortOrder:    cat.SortOrder,
		ArticleCount: count,
	}, nil
}

// CreateCategory 创建分类
func (s *categoryService) CreateCategory(req *request.CreateCategoryRequest) (uint, error) {
	existing, _ := s.categoryRepo.FindByName(req.Name)
	if existing != nil {
		return 0, bizerrors.New(bizerrors.CodeCategoryNameExists, "分类名称已存在")
	}

	category := &entity.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Icon:        req.Icon,
		SortOrder:   req.SortOrder,
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return 0, err
	}

	return category.ID, nil
}

// UpdateCategory 更新分类
func (s *categoryService) UpdateCategory(id uint, req *request.UpdateCategoryRequest) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return bizerrors.New(bizerrors.CodeCategoryNotFound, "分类不存在")
	}

	if req.Name != "" {
		existing, _ := s.categoryRepo.FindByName(req.Name)
		if existing != nil && existing.ID != id {
			return bizerrors.New(bizerrors.CodeCategoryNameExists, "分类名称已存在")
		}
		category.Name = req.Name
	}
	if req.Slug != "" {
		category.Slug = req.Slug
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.Icon != "" {
		category.Icon = req.Icon
	}
	if req.SortOrder != 0 {
		category.SortOrder = req.SortOrder
	}

	return s.categoryRepo.Update(category)
}

// DeleteCategory 删除分类
func (s *categoryService) DeleteCategory(id uint) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return bizerrors.New(bizerrors.CodeCategoryNotFound, "分类不存在")
	}

	count, _ := s.categoryRepo.GetCategoryArticleCount(id)
	if count > 0 {
		return bizerrors.New(bizerrors.CodeCategoryHasArticles, "该分类下还有文章，无法删除")
	}

	return s.categoryRepo.Delete(id)
}
