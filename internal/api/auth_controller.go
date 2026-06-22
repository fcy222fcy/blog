package api

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"

	"github.com/gin-gonic/gin"
)

// AuthController 认证控制器
type AuthController struct {
	authSvc service.AuthService
}

// NewAuthController 创建认证控制器
func NewAuthController(authSvc service.AuthService) *AuthController {
	return &AuthController{authSvc: authSvc}
}

// RegisterRoutes 注册路由
func (c *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", c.Login)
		auth.POST("/register", c.Register)
	}
}

// Login 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
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
func (c *AuthController) Register(ctx *gin.Context) {
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
