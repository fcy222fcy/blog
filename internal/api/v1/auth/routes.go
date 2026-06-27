package auth

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册认证模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}
}
