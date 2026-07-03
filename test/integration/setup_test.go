package integration

import (
	"blog/internal/api"
	"blog/internal/model/entity"
	"blog/internal/repository"
	"blog/internal/service"
	"blog/pkg/config"
	"blog/pkg/jwt"
	"blog/pkg/logger"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// TestServer 测试服务器
type TestServer struct {
	Engine   *gin.Engine
	DB       *gorm.DB
	JWT      *jwt.JWT
	Token    string
	UserID   uint
	Username string
}

// 测试用 JWT 配置
var testJWTConfig = config.JWTConfig{
	Secret:     "test-secret-for-integration-testing",
	ExpireHour: 24,
}

// 测试用配置
var testConfig = &config.Config{
	Server: config.ServerConfig{Port: 8080},
	JWT:    testJWTConfig,
	Email:  config.EmailConfig{},
}

// NewTestServer 创建测试服务器
func NewTestServer(t *testing.T) *TestServer {
	t.Helper()

	// 初始化日志（静默模式）
	logger.Init(config.LogConfig{
		Level:      "error",
		Filename:   "",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
	})

	// 使用 SQLite 内存数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("连接测试数据库失败: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Article{},
		&entity.Category{},
		&entity.Tag{},
		&entity.Comment{},
		&entity.Link{},
		&entity.DailyQuestion{},
	)
	if err != nil {
		t.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化 Repository
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	tagRepo := repository.NewTagRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	linkRepo := repository.NewLinkRepository(db)
	dailyQuestionRepo := repository.NewDailyQuestionRepository(db)
	aboutPageRepo := repository.NewAboutPageRepository(db)

	// 初始化 JWT
	jwtInstance := jwt.NewJWT(testJWTConfig)

	// 初始化 Service
	userSvc := service.NewUserService(userRepo)
	authSvc := service.NewAuthService(userRepo, jwtInstance)
	articleSvc := service.NewArticleService(articleRepo, categoryRepo, tagRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	tagSvc := service.NewTagService(tagRepo)
	commentSvc := service.NewCommentService(commentRepo, articleRepo, userRepo, nil, testConfig)
	linkSvc := service.NewLinkService(linkRepo)
	dailyQuestionSvc := service.NewDailyQuestionService(dailyQuestionRepo)
	aboutPageSvc := service.NewAboutPageService(aboutPageRepo)

	// 初始化 Router
	router := api.NewRouter(
		authSvc, userSvc, articleSvc, categorySvc, tagSvc,
		commentSvc, linkSvc, dailyQuestionSvc, aboutPageSvc,
		articleRepo, linkRepo, commentRepo,
		testConfig,
	)

	// 设置路由
	engine := router.Setup()

	return &TestServer{
		Engine: engine,
		DB:     db,
		JWT:    jwtInstance,
	}
}

// Login 登录并设置 Token
func (ts *TestServer) Login(t *testing.T, username, password string) {
	t.Helper()

	// 先确保用户存在
	ts.DB.Create(&entity.User{
		Username: username,
		Password: password,
		Nickname: username,
		Email:    username + "@test.com",
		Status:   1,
	})

	// 查找用户
	var u entity.User
	ts.DB.Where("username = ?", username).First(&u)
	ts.UserID = u.ID
	ts.Username = u.Username

	// 生成 Token
	token, _, err := ts.JWT.GenerateToken(u.ID, u.Username)
	if err != nil {
		t.Fatalf("生成 Token 失败: %v", err)
	}
	ts.Token = token
}

// DoRequest 执行 HTTP 请求（无 body）
func (ts *TestServer) DoRequest(method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ts.Engine.ServeHTTP(w, req)
	return w
}

// DoAuthRequest 执行需要认证的 HTTP 请求（无 body）
func (ts *TestServer) DoAuthRequest(method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set("Content-Type", "application/json")
	if ts.Token != "" {
		req.Header.Set("Authorization", "Bearer "+ts.Token)
	}
	w := httptest.NewRecorder()
	ts.Engine.ServeHTTP(w, req)
	return w
}

// DoRequestWithBody 执行带 JSON body 的 HTTP 请求
func (ts *TestServer) DoRequestWithBody(method, path string, jsonBody string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ts.Engine.ServeHTTP(w, req)
	return w
}

// DoAuthRequestWithBody 执行带 JSON body 且需要认证的 HTTP 请求
func (ts *TestServer) DoAuthRequestWithBody(method, path string, jsonBody string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	if ts.Token != "" {
		req.Header.Set("Authorization", "Bearer "+ts.Token)
	}
	w := httptest.NewRecorder()
	ts.Engine.ServeHTTP(w, req)
	return w
}
