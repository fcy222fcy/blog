package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// CommentService 评论服务接口
type CommentService interface {
	// GetCommentsByArticle 获取文章评论列表
	GetCommentsByArticle(articleID uint, req *request.CommentListRequest) (*response.PageResponse, error)

	// CreateComment 创建评论
	CreateComment(req *request.CreateCommentRequest) (uint, error)

	// GetAdminCommentList 获取评论列表（后台）
	GetAdminCommentList(req *request.CommentListRequest) (*response.PageResponse, error)

	// UpdateCommentStatus 更新评论状态
	UpdateCommentStatus(id uint, status string) error

	// DeleteComment 删除评论
	DeleteComment(id uint) error

	// BatchDeleteComments 批量删除评论
	BatchDeleteComments(ids []uint) error
}
