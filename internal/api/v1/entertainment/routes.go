package entertainment

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	registerPublicRoutes(rg, controller)
	registerProtectedRoutes(rg, controller)
}

func registerPublicRoutes(rg *gin.RouterGroup, controller *Controller) {
	entertainment := rg.Group("/entertainment")
	{
		entertainment.GET("", controller.GetPublicList)
	}
}

func registerProtectedRoutes(rg *gin.RouterGroup, controller *Controller) {
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		registerAdminRoutes(protected, controller)
	}
}

func registerAdminRoutes(rg *gin.RouterGroup, controller *Controller) {
	admin := rg.Group("/admin/entertainment")
	{
		admin.GET("", controller.GetAdminList)
		admin.GET("/:id", controller.GetByID)
		admin.POST("", controller.Create)
		admin.PUT("/:id", controller.Update)
		admin.DELETE("/:id", controller.Delete)
	}
}
