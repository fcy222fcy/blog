package audit_log

import (
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册审计日志模块路由（仅后台）
func RegisterRoutes(rg *gin.RouterGroup, controller *Controller) {
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		admin := protected.Group("/admin/audit-logs")
		{
			admin.GET("", controller.GetList)
		}
	}
}
