package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// CommentService 评论服务接口
type CommentService interface {
	// GetCommentList 获取文章评论列表
	GetCommentList(articleID uint, req *request.PageRequest) (*response.PageResponse, error)

	// CreateComment 创建评论
	CreateComment(articleID uint, req *request.CreateCommentRequest) (uint, error)

	// GetAdminCommentList 获取评论列表（后台）
	GetAdminCommentList(req *request.CommentListRequest) (*response.PageResponse, error)

	// UpdateCommentStatus 审核评论
	UpdateCommentStatus(id uint, req *request.UpdateCommentStatusRequest) error

	// DeleteComment 删除评论
	DeleteComment(id uint) error
}
