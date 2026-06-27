package user

import (
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/response"

	"github.com/gin-gonic/gin"
)

// Controller 用户控制器
type Controller struct {
	userSvc service.UserService
}

// NewController 创建用户控制器
func NewController(userSvc service.UserService) *Controller {
	return &Controller{userSvc: userSvc}
}

// GetUserInfo 获取用户信息
func (c *Controller) GetUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		response.Error(ctx, 401, "未登录")
		return
	}

	user, err := c.userSvc.GetUserByID(userID.(uint))
	if err != nil {
		response.Error(ctx, 500, "获取用户信息失败")
		return
	}

	response.Success(ctx, user)
}

// UpdateUserInfo 更新用户信息
func (c *Controller) UpdateUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		response.Error(ctx, 401, "未登录")
		return
	}

	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	err := c.userSvc.UpdateUser(userID.(uint), &req)
	if err != nil {
		response.Error(ctx, 500, "更新用户信息失败")
		return
	}

	response.Success(ctx, nil)
}
