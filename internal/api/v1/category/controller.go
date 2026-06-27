package category

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

// Controller 分类控制器
type Controller struct {
	categorySvc service.CategoryService
}

// NewController 创建分类控制器
func NewController(categorySvc service.CategoryService) *Controller {
	return &Controller{categorySvc: categorySvc}
}

// GetCategoryList 获取分类列表
func (c *Controller) GetCategoryList(ctx *gin.Context) {
	result, err := c.categorySvc.GetCategoryList()
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取分类列表失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

// GetCategoryDetail 获取分类详情
func (c *Controller) GetCategoryDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	result, err := c.categorySvc.GetCategoryByID(uint(id))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取分类详情失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

// CreateCategory 创建分类
func (c *Controller) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	id, err := c.categorySvc.CreateCategory(&req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("创建分类失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateCategory 更新分类
func (c *Controller) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	err = c.categorySvc.UpdateCategory(uint(id), &req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("更新分类失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}

// DeleteCategory 删除分类
func (c *Controller) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	err = c.categorySvc.DeleteCategory(uint(id))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("删除分类失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}
