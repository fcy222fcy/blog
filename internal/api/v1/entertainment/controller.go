package entertainment

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

type Controller struct {
	entertainmentSvc service.EntertainmentService
}

func NewController(entertainmentSvc service.EntertainmentService) *Controller {
	return &Controller{entertainmentSvc: entertainmentSvc}
}

func (c *Controller) GetPublicList(ctx *gin.Context) {
	typeStr := ctx.Query("type")
	yearStr := ctx.Query("year")

	var year *int
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = &y
		}
	}

	result, err := c.entertainmentSvc.GetPublicList(typeStr, year)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取公开娱乐列表业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取公开娱乐列表失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

func (c *Controller) GetAdminList(ctx *gin.Context) {
	var req request.EntertainmentListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	result, err := c.entertainmentSvc.GetAdminList(&req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取后台娱乐列表业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取后台娱乐列表失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

func (c *Controller) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		response.BadRequest(ctx, "无效的ID")
		return
	}

	result, err := c.entertainmentSvc.GetByID(uint(id))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取娱乐条目业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("获取娱乐条目失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, result)
}

func (c *Controller) Create(ctx *gin.Context) {
	var req request.CreateEntertainmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	id, err := c.entertainmentSvc.Create(&req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("创建娱乐条目业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("创建娱乐条目失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, gin.H{"id": id})
	ctx.Set("audit_created_id", uint(id))
}

func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		response.BadRequest(ctx, "无效的ID")
		return
	}

	var req request.UpdateEntertainmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	err = c.entertainmentSvc.Update(uint(id), &req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("更新娱乐条目业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("更新娱乐条目失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}

func (c *Controller) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		response.BadRequest(ctx, "无效的ID")
		return
	}

	err = c.entertainmentSvc.Delete(uint(id))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("删除娱乐条目业务错误", zap.Error(err))
			response.BizError(ctx, err)
		} else {
			logger.Error("删除娱乐条目失败", zap.Error(err))
			response.ServerError(ctx, "服务器内部错误")
		}
		return
	}

	response.Success(ctx, nil)
}
