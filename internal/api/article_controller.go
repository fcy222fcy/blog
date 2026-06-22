package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleController 文章控制器
type ArticleController struct {
	articleSvc service.ArticleService
}

// NewArticleController 创建文章控制器
func NewArticleController(articleSvc service.ArticleService) *ArticleController {
	return &ArticleController{articleSvc: articleSvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *ArticleController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	articles := rg.Group("/articles")
	{
		articles.GET("", c.GetArticleList)
		articles.GET("/archives", c.GetArchives)
		articles.GET("/:slug", c.GetArticleDetail)
		articles.POST("/:id/like", c.LikeArticle)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *ArticleController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/articles")
	{
		admin.GET("", c.GetAdminArticleList)
		admin.GET("/:id", c.GetAdminArticleDetail)
		admin.POST("", c.CreateArticle)
		admin.PUT("/:id", c.UpdateArticle)
		admin.DELETE("/:id", c.DeleteArticle)
		admin.POST("/batch-delete", c.BatchDeleteArticles)
	}
}

// GetArticleList 获取文章列表（前台）
func (c *ArticleController) GetArticleList(ctx *gin.Context) {
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
func (c *ArticleController) GetArticleDetail(ctx *gin.Context) {
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
func (c *ArticleController) GetArchives(ctx *gin.Context) {
	result, err := c.articleSvc.GetArticleArchives()
	if err != nil {
		response.Error(ctx, 500, "获取文章归档失败")
		return
	}

	response.Success(ctx, result)
}

// LikeArticle 文章点赞
func (c *ArticleController) LikeArticle(ctx *gin.Context) {
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
func (c *ArticleController) GetAdminArticleList(ctx *gin.Context) {
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
func (c *ArticleController) GetAdminArticleDetail(ctx *gin.Context) {
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
func (c *ArticleController) CreateArticle(ctx *gin.Context) {
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
func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
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
func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
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
func (c *ArticleController) BatchDeleteArticles(ctx *gin.Context) {
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
