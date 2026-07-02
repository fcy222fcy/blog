package about_page

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/logger"
	"blog/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller 关于页面控制器
type Controller struct {
	aboutPageSvc service.AboutPageService
}

// NewController 创建控制器
func NewController(aboutPageSvc service.AboutPageService) *Controller {
	return &Controller{aboutPageSvc: aboutPageSvc}
}

// GetAboutPage 获取关于页面（公开）
func (c *Controller) GetAboutPage(ctx *gin.Context) {
	result, err := c.aboutPageSvc.GetAboutPage()
	if err != nil {
		logger.Error("获取关于页面失败", zap.Error(err))
		response.ServerError(ctx, "获取关于页面失败")
		return
	}
	response.Success(ctx, result)
}

// UpdateAboutPage 更新关于页面（需要登录）
func (c *Controller) UpdateAboutPage(ctx *gin.Context) {
	var req request.UpdateAboutPageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	if err := c.aboutPageSvc.UpdateAboutPage(&req); err != nil {
		logger.Error("更新关于页面失败", zap.Error(err))
		response.ServerError(ctx, "更新关于页面失败")
		return
	}

	response.Success(ctx, nil)
}
