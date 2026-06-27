package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"fmt"
)

// tagService 标签服务实现
type tagService struct {
	tagRepo repository.TagRepository
}

// NewTagService 创建标签服务
func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{tagRepo: tagRepo}
}

// GetTagList 获取标签列表
func (s *tagService) GetTagList() ([]response.TagResponse, error) {
	tags, err := s.tagRepo.List()
	if err != nil {
		return nil, fmt.Errorf("获取标签列表失败, %w", err)
	}

	var result []response.TagResponse
	for _, tag := range tags {
		count, _ := s.tagRepo.GetTagArticleCount(tag.ID)
		result = append(result, response.TagResponse{
			ID:           tag.ID,
			Name:         tag.Name,
			Slug:         tag.Slug,
			ArticleCount: count,
			CreatedAt:    tag.CreatedAt,
		})
	}
	return result, nil
}

// CreateTag 创建标签
func (s *tagService) CreateTag(req *request.CreateTagRequest) (uint, error) {
	existing, err := s.tagRepo.FindByName(req.Name)
	if err != nil {
		return 0, fmt.Errorf("查询标签失败, %w", err)
	}
	if existing != nil {
		return 0, bizerrors.New(bizerrors.CodeTagNameExists, bizerrors.GetMessage(bizerrors.CodeTagNameExists))
	}

	tag := &entity.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}

	if err := s.tagRepo.Create(tag); err != nil {
		return 0, fmt.Errorf("创建标签失败, %w", err)
	}

	logger.Infof("创建标签成功, id=%d, name=%s", tag.ID, tag.Name)
	return tag.ID, nil
}

// UpdateTag 更新标签
func (s *tagService) UpdateTag(id uint, req *request.UpdateTagRequest) error {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询标签失败, %w", err)
	}
	if tag == nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, bizerrors.GetMessage(bizerrors.CodeTagNotFound))
	}

	if req.Name != "" {
		existing, err := s.tagRepo.FindByName(req.Name)
		if err != nil {
			return fmt.Errorf("查询标签名称失败, %w", err)
		}
		if existing != nil && existing.ID != id {
			return bizerrors.New(bizerrors.CodeTagNameExists, bizerrors.GetMessage(bizerrors.CodeTagNameExists))
		}
		tag.Name = req.Name
	}
	if req.Slug != "" {
		tag.Slug = req.Slug
	}

	if err := s.tagRepo.Update(tag); err != nil {
		return fmt.Errorf("更新标签失败, %w", err)
	}

	logger.Infof("更新标签成功, id=%d", id)
	return nil
}

// DeleteTag 删除标签
func (s *tagService) DeleteTag(id uint) error {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("查询标签失败, %w", err)
	}
	if tag == nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, bizerrors.GetMessage(bizerrors.CodeTagNotFound))
	}

	count, err := s.tagRepo.GetTagArticleCount(id)
	if err != nil {
		return fmt.Errorf("获取标签文章数失败, %w", err)
	}
	if count > 0 {
		return bizerrors.New(bizerrors.CodeTagHasArticles, bizerrors.GetMessage(bizerrors.CodeTagHasArticles))
	}

	if err := s.tagRepo.Delete(id); err != nil {
		return fmt.Errorf("删除标签失败, %w", err)
	}

	logger.Infof("删除标签成功, id=%d", id)
	return nil
}
