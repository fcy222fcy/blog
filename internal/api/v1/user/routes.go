package user

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册用户模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	// 需要登录的路由
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		users := protected.Group("/user")
		{
			users.GET("/info", controller.GetUserInfo)
			users.PUT("/info", controller.UpdateUserInfo)
		}
	}
}
