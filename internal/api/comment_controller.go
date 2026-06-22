package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentController 评论控制器
type CommentController struct {
	commentSvc service.CommentService
}

// NewCommentController 创建评论控制器
func NewCommentController(commentSvc service.CommentService) *CommentController {
	return &CommentController{commentSvc: commentSvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *CommentController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	comments := rg.Group("/comments")
	{
		comments.GET("/article/:articleId", c.GetCommentsByArticle)
		comments.POST("", c.CreateComment)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *CommentController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/comments")
	{
		admin.GET("", c.GetAdminCommentList)
		admin.PUT("/:id/status", c.UpdateCommentStatus)
		admin.DELETE("/:id", c.DeleteComment)
		admin.POST("/batch-delete", c.BatchDeleteComments)
	}
}

// GetCommentsByArticle 获取文章评论列表
func (c *CommentController) GetCommentsByArticle(ctx *gin.Context) {
	articleId, err := strconv.ParseUint(ctx.Param("articleId"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "文章ID无效")
		return
	}

	var req request.CommentListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.commentSvc.GetCommentsByArticle(uint(articleId), &req)
	if err != nil {
		response.Error(ctx, 500, "获取评论列表失败")
		return
	}

	response.Success(ctx, result)
}

// CreateComment 创建评论
func (c *CommentController) CreateComment(ctx *gin.Context) {
	var req request.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.commentSvc.CreateComment(&req)
	if err != nil {
		response.Error(ctx, 500, "创建评论失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// GetAdminCommentList 获取评论列表（后台）
func (c *CommentController) GetAdminCommentList(ctx *gin.Context) {
	var req request.CommentListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.commentSvc.GetAdminCommentList(&req)
	if err != nil {
		response.Error(ctx, 500, "获取评论列表失败")
		return
	}

	response.Success(ctx, result)
}

// UpdateCommentStatus 更新评论状态
func (c *CommentController) UpdateCommentStatus(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "评论ID无效")
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	err = c.commentSvc.UpdateCommentStatus(uint(id), req.Status)
	if err != nil {
		response.Error(ctx, 500, "更新评论状态失败")
		return
	}

	response.Success(ctx, nil)
}

// DeleteComment 删除评论
func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "评论ID无效")
		return
	}

	err = c.commentSvc.DeleteComment(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除评论失败")
		return
	}

	response.Success(ctx, nil)
}

// BatchDeleteComments 批量删除评论
func (c *CommentController) BatchDeleteComments(ctx *gin.Context) {
	var req struct {
		IDs []uint `json:"ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	err := c.commentSvc.BatchDeleteComments(req.IDs)
	if err != nil {
		response.Error(ctx, 500, "批量删除评论失败")
		return
	}

	response.Success(ctx, nil)
}
