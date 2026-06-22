package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// LinkService 友链服务接口
type LinkService interface {
	// GetLinkList 获取友链列表
	GetLinkList() ([]response.LinkResponse, error)

	// GetAdminLinkList 获取友链列表（后台）
	GetAdminLinkList(req *request.LinkListRequest) (*response.PageResponse, error)

	// CreateLink 创建友链
	CreateLink(req *request.CreateLinkRequest) (uint, error)

	// UpdateLink 更新友链
	UpdateLink(id uint, req *request.UpdateLinkRequest) error

	// DeleteLink 删除友链
	DeleteLink(id uint) error

	// UpdateLinkStatus 更新友链状态
	UpdateLinkStatus(id uint, status string) error
}
