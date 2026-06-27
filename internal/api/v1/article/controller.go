package article

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller 文章控制器
type Controller struct {
	articleSvc service.ArticleService
}

// NewController 创建文章控制器
func NewController(articleSvc service.ArticleService) *Controller {
	return &Controller{articleSvc: articleSvc}
}

// GetArticleList 获取文章列表（前台）
func (c *Controller) GetArticleList(ctx *gin.Context) {
	var req request.ArticleListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.articleSvc.GetArticleList(&req)
	if err != nil {
		response.Error(ctx, 500, "获取文章列表失败")
		return
	}

	response.Success(ctx, result)
}

// GetArticleDetail 获取文章详情（前台）
func (c *Controller) GetArticleDetail(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		response.Error(ctx, 400, "文章slug不能为空")
		return
	}

	result, err := c.articleSvc.GetArticleDetail(slug)
	if err != nil {
		response.Error(ctx, 404, "文章不存在")
		return
	}

	response.Success(ctx, result)
}

// GetArchives 获取文章归档
func (c *Controller) GetArchives(ctx *gin.Context) {
	result, err := c.articleSvc.GetArticleArchives()
	if err != nil {
		response.Error(ctx, 500, "获取文章归档失败")
		return
	}

	response.Success(ctx, result)
}

// LikeArticle 文章点赞
func (c *Controller) LikeArticle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "文章ID无效")
		return
	}

	likeCount, err := c.articleSvc.LikeArticle(uint(id))
	if err != nil {
		response.Error(ctx, 500, "点赞失败")
		return
	}

	response.Success(ctx, gin.H{"like_count": likeCount})
}

// GetAdminArticleList 获取文章列表（后台）
func (c *Controller) GetAdminArticleList(ctx *gin.Context) {
	var req request.ArticleListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.articleSvc.GetAdminArticleList(&req)
	if err != nil {
		response.Error(ctx, 500, "获取文章列表失败")
		return
	}

	response.Success(ctx, result)
}

// GetAdminArticleDetail 获取文章详情（后台）
func (c *Controller) GetAdminArticleDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "文章ID无效")
		return
	}

	result, err := c.articleSvc.GetAdminArticleDetail(uint(id))
	if err != nil {
		response.Error(ctx, 404, "文章不存在")
		return
	}

	response.Success(ctx, result)
}

// CreateArticle 创建文章
func (c *Controller) CreateArticle(ctx *gin.Context) {
	var req request.CreateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.articleSvc.CreateArticle(&req)
	if err != nil {
		response.Error(ctx, 500, "创建文章失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateArticle 更新文章
func (c *Controller) UpdateArticle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "文章ID无效")
		return
	}

	var req request.UpdateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err = c.articleSvc.UpdateArticle(uint(id), &req)
	if err != nil {
		response.Error(ctx, 500, "更新文章失败: "+err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteArticle 删除文章
func (c *Controller) DeleteArticle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "文章ID无效")
		return
	}

	err = c.articleSvc.DeleteArticle(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除文章失败")
		return
	}

	response.Success(ctx, nil)
}

// BatchDeleteArticles 批量删除文章
func (c *Controller) BatchDeleteArticles(ctx *gin.Context) {
	var req struct {
		IDs []uint `json:"ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	err := c.articleSvc.BatchDeleteArticles(req.IDs)
	if err != nil {
		response.Error(ctx, 500, "批量删除文章失败")
		return
	}

	response.Success(ctx, nil)
}
