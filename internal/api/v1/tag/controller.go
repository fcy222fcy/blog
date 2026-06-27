package tag

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller 标签控制器
type Controller struct {
	tagSvc service.TagService
}

// NewController 创建标签控制器
func NewController(tagSvc service.TagService) *Controller {
	return &Controller{tagSvc: tagSvc}
}

// GetTagList 获取标签列表
func (c *Controller) GetTagList(ctx *gin.Context) {
	result, err := c.tagSvc.GetTagList()
	if err != nil {
		response.Error(ctx, 500, "获取标签列表失败")
		return
	}

	response.Success(ctx, result)
}

// CreateTag 创建标签
func (c *Controller) CreateTag(ctx *gin.Context) {
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
func (c *Controller) UpdateTag(ctx *gin.Context) {
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
func (c *Controller) DeleteTag(ctx *gin.Context) {
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
