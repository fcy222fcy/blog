package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
	"blog/internal/model/entity"
	"blog/internal/repository"
	bizerrors "blog/pkg/errors"
)

// linkService 友链服务实现
type linkService struct {
	linkRepo repository.LinkRepository
}

// NewLinkService 创建友链服务
func NewLinkService(linkRepo repository.LinkRepository) LinkService {
	return &linkService{linkRepo: linkRepo}
}

// GetLinkList 获取友链列表（前台）
func (s *linkService) GetLinkList() ([]response.LinkResponse, error) {
	links, err := s.linkRepo.List()
	if err != nil {
		return nil, err
	}

	var result []response.LinkResponse
	for _, link := range links {
		result = append(result, response.LinkResponse{
			ID:          link.ID,
			Name:        link.Name,
			URL:         link.URL,
			Description: link.Description,
			Avatar:      link.Avatar,
			Logo:        link.Logo,
			SortOrder:   link.SortOrder,
			Status:      link.Status,
			CreatedAt:   link.CreatedAt,
		})
	}
	return result, nil
}

// GetAdminLinkList 获取友链列表（后台）
func (s *linkService) GetAdminLinkList(req *request.LinkListRequest) (*response.PageResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := s.linkRepo.AdminList((req.Page-1)*req.PageSize, req.PageSize, req.Status)
	if err != nil {
		return nil, err
	}

	var result []response.LinkResponse
	for _, link := range list {
		result = append(result, response.LinkResponse{
			ID:          link.ID,
			Name:        link.Name,
			URL:         link.URL,
			Description: link.Description,
			Avatar:      link.Avatar,
			Logo:        link.Logo,
			SortOrder:   link.SortOrder,
			Status:      link.Status,
			CreatedAt:   link.CreatedAt,
		})
	}

	return response.NewPageResponse(result, total, req.Page, req.PageSize), nil
}

// CreateLink 创建友链
func (s *linkService) CreateLink(req *request.CreateLinkRequest) (uint, error) {
	link := &entity.Link{
		Name:        req.Name,
		URL:         req.URL,
		Description: req.Description,
		Avatar:      req.Avatar,
		Logo:        req.Logo,
		SortOrder:   req.SortOrder,
		Status:      "pending",
	}

	err := s.linkRepo.Create(link)
	if err != nil {
		return 0, err
	}

	return link.ID, nil
}

// UpdateLink 更新友链
func (s *linkService) UpdateLink(id uint, req *request.UpdateLinkRequest) error {
	link, err := s.linkRepo.FindByID(id)
	if err != nil {
		return err
	}
	if link == nil {
		return bizerrors.New(bizerrors.CodeLinkNotFound, "友链不存在")
	}

	if req.Name != "" {
		link.Name = req.Name
	}
	if req.URL != "" {
		link.URL = req.URL
	}
	if req.Description != "" {
		link.Description = req.Description
	}
	if req.Avatar != "" {
		link.Avatar = req.Avatar
	}
	if req.Logo != "" {
		link.Logo = req.Logo
	}
	if req.SortOrder != 0 {
		link.SortOrder = req.SortOrder
	}

	return s.linkRepo.Update(link)
}

// DeleteLink 删除友链
func (s *linkService) DeleteLink(id uint) error {
	link, err := s.linkRepo.FindByID(id)
	if err != nil {
		return err
	}
	if link == nil {
		return bizerrors.New(bizerrors.CodeLinkNotFound, "友链不存在")
	}

	return s.linkRepo.Delete(id)
}

// UpdateLinkStatus 更新友链状态
func (s *linkService) UpdateLinkStatus(id uint, status string) error {
	link, err := s.linkRepo.FindByID(id)
	if err != nil {
		return err
	}
	if link == nil {
		return bizerrors.New(bizerrors.CodeLinkNotFound, "友链不存在")
	}

	link.Status = status
	return s.linkRepo.Update(link)
}
