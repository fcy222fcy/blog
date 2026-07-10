package comment

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册评论模块路由
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	// 公开路由（无需登录）
	registerPublicRoutes(rg, controller)

	// 需要登录的路由
	registerProtectedRoutes(rg, controller)
}

// registerPublicRoutes 注册公开路由
func registerPublicRoutes(rg *gin.RouterGroup, controller *Controller) {
	comments := rg.Group("/comments")
	comments.Use(middleware.OptionalAuth())
	{
		comments.GET("/article/:articleId", controller.GetCommentsByArticle)
		comments.POST("", controller.CreateComment)
		comments.POST("/:id/like", controller.LikeComment)
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
	admin := rg.Group("/admin/comments")
	{
		admin.GET("", controller.GetAdminCommentList)
		admin.PUT("/:id/status", controller.UpdateCommentStatus)
		admin.DELETE("/:id", controller.DeleteComment)
		admin.POST("/batch-delete", controller.BatchDeleteComments)
	}
}
