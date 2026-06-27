package link

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller 友链控制器
type Controller struct {
	linkSvc service.LinkService
}

// NewController 创建友链控制器
func NewController(linkSvc service.LinkService) *Controller {
	return &Controller{linkSvc: linkSvc}
}

// GetLinkList 获取友链列表（前台）
func (c *Controller) GetLinkList(ctx *gin.Context) {
	result, err := c.linkSvc.GetLinkList()
	if err != nil {
		response.Error(ctx, 500, "获取友链列表失败")
		return
	}

	response.Success(ctx, result)
}

// GetAdminLinkList 获取友链列表（后台）
func (c *Controller) GetAdminLinkList(ctx *gin.Context) {
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
func (c *Controller) CreateLink(ctx *gin.Context) {
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
func (c *Controller) UpdateLink(ctx *gin.Context) {
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
func (c *Controller) DeleteLink(ctx *gin.Context) {
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
func (c *Controller) UpdateLinkStatus(ctx *gin.Context) {
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
