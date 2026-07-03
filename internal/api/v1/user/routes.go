package user

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册用户模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	// 公开路由（无需登录）- 获取博客主人信息
	users := rg.Group("/user")
	{
		users.GET("/info", controller.GetPublicUserInfo)
	}

	// 需要登录的路由
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		users := protected.Group("/user")
		{
			users.GET("/profile", controller.GetUserInfo)
			users.PUT("/profile", controller.UpdateUserInfo)
		}
	}
}
