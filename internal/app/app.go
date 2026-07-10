package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"blog/internal/api"
	"blog/internal/repository"
	"blog/internal/service"
	"blog/pkg/config"
	"blog/pkg/database"
	"blog/pkg/email"
	"blog/pkg/jwt"
	"blog/pkg/logger"
)

// App 应用结构
type App struct {
	config *config.Config
	engine *http.Server
	router *api.Router
	db     *database.Database
}

// NewApp 创建应用实例
func NewApp() *App {
	return &App{}
}

// Run 启动应用
func (a *App) Run() {
	// 1. 加载配置
	a.loadConfig()

	// 2. 初始化日志
	a.initLogger()

	// 3. 初始化数据库
	a.initDatabase()

	// 4. 初始化依赖
	a.initDependencies()

	// 5. 启动 HTTP 服务
	a.startHTTPServer()
}

// loadConfig 加载配置
func (a *App) loadConfig() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	a.config = cfg
}

// initLogger 初始化日志
func (a *App) initLogger() {
	if err := logger.Init(a.config.Log); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
}

// initDatabase 初始化数据库
func (a *App) initDatabase() {
	db, err := database.NewDatabase(a.config.MySQL)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	a.db = db

	// 自动迁移
	if err := a.db.AutoMigrate(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
}

// initDependencies 初始化依赖注入
func (a *App) initDependencies() {
	// Repository
	userRepo := repository.NewUserRepository(a.db.DB)
	articleRepo := repository.NewArticleRepository(a.db.DB)
	categoryRepo := repository.NewCategoryRepository(a.db.DB)
	tagRepo := repository.NewTagRepository(a.db.DB)
	commentRepo := repository.NewCommentRepository(a.db.DB)
	linkRepo := repository.NewLinkRepository(a.db.DB)
	dailyQuestionRepo := repository.NewDailyQuestionRepository(a.db.DB)
	entertainmentRepo := repository.NewEntertainmentRepository(a.db.DB)
	aboutPageRepo := repository.NewAboutPageRepository(a.db.DB)
	visitRepo := repository.NewVisitRepository(a.db.DB)
	auditLogRepo := repository.NewAuditLogRepository(a.db.DB)

	// Service
	jwtInstance := jwt.NewJWT(a.config.JWT)
	userSvc := service.NewUserService(userRepo)
	authSvc := service.NewAuthService(userRepo, jwtInstance, a.config)
	articleSvc := service.NewArticleService(articleRepo, categoryRepo, tagRepo, visitRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	tagSvc := service.NewTagService(tagRepo)
	emailSvc := email.NewEmailService(a.config.Email)
	commentSvc := service.NewCommentService(commentRepo, articleRepo, userRepo, emailSvc, a.config)
	linkSvc := service.NewLinkService(linkRepo)
	dailyQuestionSvc := service.NewDailyQuestionService(dailyQuestionRepo)
	entertainmentSvc := service.NewEntertainmentService(entertainmentRepo)
	aboutPageSvc := service.NewAboutPageService(aboutPageRepo)
	auditLogSvc := service.NewAuditLogService(auditLogRepo)

	// Router
	a.router = api.NewRouter(
		authSvc,
		userSvc,
		articleSvc,
		categorySvc,
		tagSvc,
		commentSvc,
		linkSvc,
		dailyQuestionSvc,
		entertainmentSvc,
		aboutPageSvc,
		auditLogSvc,
		articleRepo,
		linkRepo,
		commentRepo,
		visitRepo,
		a.config,
	)
}

// startHTTPServer 启动 HTTP 服务
func (a *App) startHTTPServer() {
	addr := fmt.Sprintf(":%d", a.config.Server.Port)

	a.engine = &http.Server{
		Addr:         addr,
		Handler:      a.router.Setup(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 优雅退出
	go func() {
		fmt.Printf("服务启动成功，监听地址: %s\n", addr)
		if err := a.engine.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("正在关闭服务...")
	a.shutdown()
}

// shutdown 优雅关闭
func (a *App) shutdown() {
	// 关闭 HTTP 服务
	if err := a.engine.Close(); err != nil {
		log.Printf("关闭 HTTP 服务失败: %v", err)
	}

	// 关闭数据库连接
	if err := a.db.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}

	// 关闭日志
	logger.Sync()

	fmt.Println("服务已关闭")
}
