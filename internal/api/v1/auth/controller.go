package auth

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"

	"github.com/gin-gonic/gin"
)

// Controller 认证控制器
type Controller struct {
	authSvc service.AuthService
}

// NewController 创建认证控制器
func NewController(authSvc service.AuthService) *Controller {
	return &Controller{authSvc: authSvc}
}

// Login 用户登录
func (c *Controller) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	result, err := c.authSvc.Login(&req)
	if err != nil {
		response.Error(ctx, 401, err.Error())
		return
	}

	response.Success(ctx, result)
}

// Register 用户注册
func (c *Controller) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err := c.authSvc.Register(&req)
	if err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	response.Success(ctx, nil)
}
