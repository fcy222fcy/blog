package tag

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	bizerrors "blog/pkg/errors"
	"blog/pkg/logger"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取标签列表失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

// CreateTag 创建标签
func (c *Controller) CreateTag(ctx *gin.Context) {
	var req request.CreateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	id, err := c.tagSvc.CreateTag(&req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("创建标签失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateTag 更新标签
func (c *Controller) UpdateTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	var req request.UpdateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	err = c.tagSvc.UpdateTag(uint(id), &req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("更新标签失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}

// DeleteTag 删除标签
func (c *Controller) DeleteTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	err = c.tagSvc.DeleteTag(uint(id))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("删除标签失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}
