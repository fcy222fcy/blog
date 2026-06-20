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

// GetLinkList 获取友链列表
func (s *linkService) GetLinkList() ([]response.LinkResponse, error) {
	links, err := s.linkRepo.ListApproved()
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
		})
	}
	return result, nil
}

// ApplyLink 申请友链
func (s *linkService) ApplyLink(req *request.ApplyLinkRequest) error {
	link := &entity.Link{
		Name:        req.Name,
		URL:         req.URL,
		Description: req.Description,
		Avatar:      req.Avatar,
		Status:      "pending",
	}
	return s.linkRepo.Create(link)
}

// GetAdminLinkList 获取友链列表（后台）
func (s *linkService) GetAdminLinkList(req *request.LinkListRequest) (*response.PageResponse, error) {
	list, total, err := s.linkRepo.ListAll(req.GetOffset(), req.GetPageSize(), req.Status)
	if err != nil {
		return nil, err
	}
	return response.NewPageResponse(list, total, req.Page, req.GetPageSize()), nil
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
		Status:      req.Status,
	}

	if link.Status == "" {
		link.Status = "approved"
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
	link.SortOrder = req.SortOrder
	if req.Status != "" {
		link.Status = req.Status
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

// UpdateLinkStatus 审核友链
func (s *linkService) UpdateLinkStatus(id uint, req *request.UpdateLinkStatusRequest) error {
	link, err := s.linkRepo.FindByID(id)
	if err != nil {
		return err
	}
	if link == nil {
		return bizerrors.New(bizerrors.CodeLinkNotFound, "友链不存在")
	}

	link.Status = req.Status
	return s.linkRepo.Update(link)
}
