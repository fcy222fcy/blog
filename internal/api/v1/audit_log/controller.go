package audit_log

import (
	"blog/internal/service"
	"blog/pkg/logger"
	"blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller 审计日志控制器
type Controller struct {
	auditLogSvc service.AuditLogService
}

// NewController 创建审计日志控制器
func NewController(auditLogSvc service.AuditLogService) *Controller {
	return &Controller{auditLogSvc: auditLogSvc}
}

// GetList 获取审计日志列表
func (c *Controller) GetList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	action := ctx.Query("action")
	targetType := ctx.Query("target_type")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	result, err := c.auditLogSvc.GetList(page, pageSize, action, targetType, startDate, endDate)
	if err != nil {
		logger.Error("获取审计日志列表失败", zap.Error(err))
		response.ServerError(ctx, "服务器内部错误")
		return
	}

	response.Success(ctx, result)
}
