package api

import (
	"blog/internal/middleware"
	"blog/internal/service"
	"blog/pkg/config"

	"github.com/gin-gonic/gin"
)

// Router 路由器
type Router struct {
	engine             *gin.Engine
	config             *config.Config
	authController     *AuthController
	userController     *UserController
	articleController  *ArticleController
	categoryController *CategoryController
	tagController      *TagController
	commentController  *CommentController
	linkController     *LinkController
	dailyQController   *DailyQuestionController
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
	return &Router{
		engine:             gin.Default(),
		config:             config,
		authController:     NewAuthController(authSvc),
		userController:     NewUserController(userSvc),
		articleController:  NewArticleController(articleSvc),
		categoryController: NewCategoryController(categorySvc),
		tagController:      NewTagController(tagSvc),
		commentController:  NewCommentController(commentSvc),
		linkController:     NewLinkController(linkSvc),
		dailyQController:   NewDailyQuestionController(dailyQuestionSvc),
	}
}

// Setup 设置路由
func (r *Router) Setup() *gin.Engine {
	// 全局中间件
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.CORS())

	// API 路由组
	apiV1 := r.engine.Group("/api/v1")
	{
		// 认证相关（无需登录）
		r.authController.RegisterRoutes(apiV1)

		// 公开接口（无需登录）
		r.registerPublicRoutes(apiV1)

		// 需要登录的接口
		r.registerProtectedRoutes(apiV1)
	}

	return r.engine
}

// registerPublicRoutes 注册公开路由（无需登录）
func (r *Router) registerPublicRoutes(rg *gin.RouterGroup) {
	// 文章
	r.articleController.RegisterPublicRoutes(rg)

	// 分类
	r.categoryController.RegisterPublicRoutes(rg)

	// 标签
	r.tagController.RegisterPublicRoutes(rg)

	// 评论
	r.commentController.RegisterPublicRoutes(rg)

	// 友链
	r.linkController.RegisterPublicRoutes(rg)

	// 每日一问
	r.dailyQController.RegisterPublicRoutes(rg)
}

// registerProtectedRoutes 注册需要登录的路由
func (r *Router) registerProtectedRoutes(rg *gin.RouterGroup) {
	// 需要认证的路由组
	protected := rg.Group("")
	protected.Use(middleware.Auth())
	{
		// 用户
		r.userController.RegisterRoutes(protected)

		// 后台管理
		r.registerAdminRoutes(protected)
	}
}

// registerAdminRoutes 注册后台管理路由
func (r *Router) registerAdminRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/admin")
	{
		// 文章管理
		r.articleController.RegisterAdminRoutes(admin)

		// 分类管理
		r.categoryController.RegisterAdminRoutes(admin)

		// 标签管理
		r.tagController.RegisterAdminRoutes(admin)

		// 评论管理
		r.commentController.RegisterAdminRoutes(admin)

		// 友链管理
		r.linkController.RegisterAdminRoutes(admin)

		// 每日一问管理
		r.dailyQController.RegisterAdminRoutes(admin)

		// 仪表盘
		r.registerDashboardRoutes(admin)
	}
}

// registerDashboardRoutes 注册仪表盘路由
func (r *Router) registerDashboardRoutes(rg *gin.RouterGroup) {
	dashboard := rg.Group("/dashboard")
	{
		dashboard.GET("/stats", r.getDashboardStats)
		dashboard.GET("/recent-articles", r.getRecentArticles)
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
