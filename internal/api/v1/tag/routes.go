package tag

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册标签模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	// 公开路由（无需登录）
	registerPublicRoutes(rg, controller)

	// 需要登录的路由
	registerProtectedRoutes(rg, controller)
}

// registerPublicRoutes 注册公开路由
func registerPublicRoutes(rg *gin.RouterGroup, controller *Controller) {
	tags := rg.Group("/tags")
	{
		tags.GET("", controller.GetTagList)
	}
}

// registerProtectedRoutes 注册需要登录的路由
func registerProtectedRoutes(rg *gin.RouterGroup, controller *Controller) {
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		// 后台管理路由
		registerAdminRoutes(protected, controller)
	}
}

// registerAdminRoutes 注册后台管理路由
func registerAdminRoutes(rg *gin.RouterGroup, controller *Controller) {
	admin := rg.Group("/admin/tags")
	{
		admin.POST("", controller.CreateTag)
		admin.PUT("/:id", controller.UpdateTag)
		admin.DELETE("/:id", controller.DeleteTag)
	}
}
