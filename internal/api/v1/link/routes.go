package link

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册友链模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	// 公开路由（无需登录）
	registerPublicRoutes(rg, controller)

	// 需要登录的路由
	registerProtectedRoutes(rg, controller)
}

// registerPublicRoutes 注册公开路由
func registerPublicRoutes(rg *gin.RouterGroup, controller *Controller) {
	links := rg.Group("/links")
	{
		links.GET("", controller.GetLinkList)
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
	admin := rg.Group("/admin/links")
	{
		admin.GET("", controller.GetAdminLinkList)
		admin.POST("", controller.CreateLink)
		admin.PUT("/:id", controller.UpdateLink)
		admin.DELETE("/:id", controller.DeleteLink)
		admin.PUT("/:id/status", controller.UpdateLinkStatus)
	}
}
