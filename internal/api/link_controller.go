package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LinkController 友链控制器
type LinkController struct {
	linkSvc service.LinkService
}

// NewLinkController 创建友链控制器
func NewLinkController(linkSvc service.LinkService) *LinkController {
	return &LinkController{linkSvc: linkSvc}
}

// RegisterPublicRoutes 注册公开路由
func (c *LinkController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	links := rg.Group("/links")
	{
		links.GET("", c.GetLinkList)
	}
}

// RegisterAdminRoutes 注册后台路由
func (c *LinkController) RegisterAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/links")
	{
		admin.GET("", c.GetAdminLinkList)
		admin.POST("", c.CreateLink)
		admin.PUT("/:id", c.UpdateLink)
		admin.DELETE("/:id", c.DeleteLink)
		admin.PUT("/:id/status", c.UpdateLinkStatus)
	}
}

// GetLinkList 获取友链列表（前台）
func (c *LinkController) GetLinkList(ctx *gin.Context) {
	result, err := c.linkSvc.GetLinkList()
	if err != nil {
		response.Error(ctx, 500, "获取友链列表失败")
		return
	}

	response.Success(ctx, result)
}

// GetAdminLinkList 获取友链列表（后台）
func (c *LinkController) GetAdminLinkList(ctx *gin.Context) {
	var req request.LinkListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	result, err := c.linkSvc.GetAdminLinkList(&req)
	if err != nil {
		response.Error(ctx, 500, "获取友链列表失败")
		return
	}

	response.Success(ctx, result)
}

// CreateLink 创建友链
func (c *LinkController) CreateLink(ctx *gin.Context) {
	var req request.CreateLinkRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	id, err := c.linkSvc.CreateLink(&req)
	if err != nil {
		response.Error(ctx, 500, "创建友链失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateLink 更新友链
func (c *LinkController) UpdateLink(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "友链ID无效")
		return
	}

	var req request.UpdateLinkRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err = c.linkSvc.UpdateLink(uint(id), &req)
	if err != nil {
		response.Error(ctx, 500, "更新友链失败: "+err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteLink 删除友链
func (c *LinkController) DeleteLink(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "友链ID无效")
		return
	}

	err = c.linkSvc.DeleteLink(uint(id))
	if err != nil {
		response.Error(ctx, 500, "删除友链失败")
		return
	}

	response.Success(ctx, nil)
}

// UpdateLinkStatus 更新友链状态
func (c *LinkController) UpdateLinkStatus(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx, 400, "友链ID无效")
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	err = c.linkSvc.UpdateLinkStatus(uint(id), req.Status)
	if err != nil {
		response.Error(ctx, 500, "更新友链状态失败")
		return
	}

	response.Success(ctx, nil)
}
