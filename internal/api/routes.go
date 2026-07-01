package api

import (
	"blog/internal/api/v1/article"
	"blog/internal/api/v1/auth"
	"blog/internal/api/v1/category"
	"blog/internal/api/v1/comment"
	"blog/internal/api/v1/daily_question"
	"blog/internal/api/v1/link"
	"blog/internal/api/v1/tag"
	"blog/internal/api/v1/user"
	"blog/internal/middleware"
	"blog/internal/model/dto/request"
	"blog/internal/service"
	"blog/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Router 路由器
type Router struct {
	engine *gin.Engine
	config *config.Config

	// 各模块 controller
	authController         *auth.Controller
	userController         *user.Controller
	articleController      *article.Controller
	categoryController     *category.Controller
	tagController          *tag.Controller
	commentController      *comment.Controller
	linkController         *link.Controller
	dailyQuestionController *daily_question.Controller
}

// NewRouter 创建路由器
func NewRouter(
	authSvc service.AuthService,
	userSvc service.UserService,
	articleSvc service.ArticleService,
	categorySvc service.CategoryService,
	tagSvc service.TagService,
	commentSvc service.CommentService,
	linkSvc service.LinkService,
	dailyQuestionSvc service.DailyQuestionService,
	config *config.Config,
) *Router {
	engine := gin.Default()

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("datetime", request.ValidateDate)
	}

	return &Router{
		engine:                 engine,
		config:                 config,
		authController:         auth.NewController(authSvc),
		userController:         user.NewController(userSvc),
		articleController:      article.NewController(articleSvc),
		categoryController:     category.NewController(categorySvc),
		tagController:          tag.NewController(tagSvc),
		commentController:      comment.NewController(commentSvc),
		linkController:         link.NewController(linkSvc),
		dailyQuestionController: daily_question.NewController(dailyQuestionSvc),
	}
}

// Setup 设置路由
func (r *Router) Setup() *gin.Engine {
	// 全局中间件
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.Logger())
	// 使用配置的 CORS 来源，如果没有配置则限制为同源
	r.engine.Use(middleware.CORS("http://localhost:3000", "http://localhost:8888"))

	// API v1 路由组
	apiV1 := r.engine.Group("/api/v1")

	// 注入 JWT 配置到上下文，供各模块 Auth 中间件使用
	apiV1.Use(func(c *gin.Context) {
		c.Set("jwt_config", r.config.JWT)
		c.Next()
	})

	// 注册各模块路由
	auth.RegisterRoutes(apiV1, r.authController)
	article.RegisterRoutes(apiV1, r.articleController)
	category.RegisterRoutes(apiV1, r.categoryController)
	tag.RegisterRoutes(apiV1, r.tagController)
	comment.RegisterRoutes(apiV1, r.commentController)
	link.RegisterRoutes(apiV1, r.linkController)
	daily_question.RegisterRoutes(apiV1, r.dailyQuestionController)
	user.RegisterRoutes(apiV1, r.userController)

	// 注册仪表盘路由（需要登录）
	r.registerDashboardRoutes(apiV1)

	return r.engine
}

// registerDashboardRoutes 注册仪表盘路由
func (r *Router) registerDashboardRoutes(rg *gin.RouterGroup) {
	// 需要登录的路由
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		dashboard := protected.Group("/admin/dashboard")
		{
			dashboard.GET("/stats", r.getDashboardStats)
			dashboard.GET("/recent-articles", r.getRecentArticles)
		}
	}
}

// getDashboardStats 获取仪表盘统计
func (r *Router) getDashboardStats(c *gin.Context) {
	// TODO: 实现仪表盘统计
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    gin.H{},
	})
}

// getRecentArticles 获取最近文章
func (r *Router) getRecentArticles(c *gin.Context) {
	// TODO: 实现最近文章
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    []interface{}{},
	})
}
