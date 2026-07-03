package service

import (
	"blog/internal/model/dto/request"
	"blog/internal/model/dto/response"
)

// CommentService 评论服务接口
type CommentService interface {
	// GetCommentsByArticle 获取文章评论列表（支持 slug 或数字 ID）
	GetCommentsByArticle(articleParam string, req *request.CommentListRequest) (*response.PageResponse, error)

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

	// LikeComment 点赞评论（防重复）
	LikeComment(commentID uint, visitorIP string) error
}
