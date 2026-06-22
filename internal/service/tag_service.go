package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
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
		return nil, err
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
	existing, _ := s.tagRepo.FindByName(req.Name)
	if existing != nil {
		return 0, bizerrors.New(bizerrors.CodeTagNameExists, "标签名称已存在")
	}

	tag := &entity.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}

	err := s.tagRepo.Create(tag)
	if err != nil {
		return 0, err
	}

	return tag.ID, nil
}

// UpdateTag 更新标签
func (s *tagService) UpdateTag(id uint, req *request.UpdateTagRequest) error {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return err
	}
	if tag == nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, "标签不存在")
	}

	if req.Name != "" {
		existing, _ := s.tagRepo.FindByName(req.Name)
		if existing != nil && existing.ID != id {
			return bizerrors.New(bizerrors.CodeTagNameExists, "标签名称已存在")
		}
		tag.Name = req.Name
	}
	if req.Slug != "" {
		tag.Slug = req.Slug
	}

	return s.tagRepo.Update(tag)
}

// DeleteTag 删除标签
func (s *tagService) DeleteTag(id uint) error {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return err
	}
	if tag == nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, "标签不存在")
	}

	count, _ := s.tagRepo.GetTagArticleCount(id)
	if count > 0 {
		return bizerrors.New(bizerrors.CodeTagHasArticles, "该标签下还有文章，无法删除")
	}

	return s.tagRepo.Delete(id)
}
