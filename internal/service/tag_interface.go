package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// TagService 标签服务接口
type TagService interface {
	// GetTagList 获取标签列表
	GetTagList() ([]response.TagResponse, error)

	// CreateTag 创建标签
	CreateTag(req *request.CreateTagRequest) (uint, error)

	// UpdateTag 更新标签
	UpdateTag(id uint, req *request.UpdateTagRequest) error

	// DeleteTag 删除标签
	DeleteTag(id uint) error
}
