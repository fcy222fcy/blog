package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TagController 标签控制器
type TagController struct {
	tagSvc service.TagService
}

// NewTagController 创建标签控制器
func NewTagController(tagSvc service.TagService) *TagController {
	return &TagController{tagSvc: tagSvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *TagController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	tags := rg.Group("/tags")
	{
		tags.GET("", c.GetTagList)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *TagController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/tags")
	{
		admin.POST("", c.CreateTag)
		admin.PUT("/:id", c.UpdateTag)
		admin.DELETE("/:id", c.DeleteTag)
	}
}

// GetTagList 获取标签列表
func (c *TagController) GetTagList(ctx *gin.Context) {
	result, err := c.tagSvc.GetTagList()
	if err != nil {
		response.Error(ctx, 500, "获取标签列表失败")
		return
	}

	response.Success(ctx, result)
}

// CreateTag 创建标签
func (c *TagController) CreateTag(ctx *gin.Context) {
	var req request.CreateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.tagSvc.CreateTag(&req)
	if err != nil {
		response.Error(ctx, 500, "创建标签失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateTag 更新标签
func (c *TagController) UpdateTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "标签ID无效")
		return
	}

	var req request.UpdateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err = c.tagSvc.UpdateTag(uint(id), &req)
	if err != nil {
		response.Error(ctx, 500, "更新标签失败: "+err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteTag 删除标签
func (c *TagController) DeleteTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "标签ID无效")
		return
	}

	err = c.tagSvc.DeleteTag(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除标签失败")
		return
	}

	response.Success(ctx, nil)
}
