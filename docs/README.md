# 接口文档说明

## 文档规范

本项目采用 **OpenAPI 3.0 + Swag** 两者结合的方式管理接口文档。

### 文档结构

```
docs/
├── api.yaml                    # OpenAPI 3.0 规范文件（完整接口定义）
├── swagger.yaml                # Swag 注释规范文件（代码注释参考）
└── README.md                   # 本说明文件
```

## 工作流程

### 1. 设计阶段（OpenAPI YAML）

在 `docs/api.yaml` 中定义接口规范，包括：
- 接口路径和方法
- 请求参数和响应格式
- 数据模型定义
- 认证方式

### 2. 开发阶段（Swag 注释）

在 Go 代码中添加 Swag 注释，格式参考 `docs/swagger.yaml`：

```go
// @Summary 获取文章列表
// @Description 分页获取已发布的文章列表
// @Tags 文章
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} response.Response{data=response.ArticleListResponse}
// @Router /articles [get]
func (ctrl *ArticleController) GetArticleList(c *gin.Context) {
    // ...
}
```

### 3. 生成文档

使用 `swag` 工具自动生成 Swagger 文档：

```bash
# 安装 swag
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init -g cmd/server/main.go -o docs/swagger

# 启动服务后访问
# http://localhost:8080/swagger/index.html
```

## 接口分组

| 分组 | 路径前缀 | 说明 |
|------|----------|------|
| 认证 | `/auth` | 登录、注册、修改密码 |
| 用户 | `/user` | 用户信息管理 |
| 文章 | `/articles` | 文章列表、详情、点赞 |
| 分类 | `/categories` | 分类管理 |
| 标签 | `/tags` | 标签管理 |
| 评论 | `/comments` | 评论管理 |
| 友链 | `/links` | 友情链接管理 |
| 每日一问 | `/daily-questions` | 每日问答管理 |
| 媒体库 | `/media` | 文件上传管理 |
| 系统设置 | `/settings` | 博客配置管理 |
| 仪表盘 | `/dashboard` | 后台统计数据 |

## 认证方式

需要登录的接口使用 JWT Token 认证，在请求头中添加：

```
Authorization: Bearer <token>
```

## 响应格式

### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误响应

```json
{
  "code": 1001,
  "message": "用户不存在"
}
```

### 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "size": 10,
    "total_page": 10
  }
}
```

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 请求错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 1000+ | 用户/认证业务错误 |
| 2000+ | 参数错误 |
| 3000+ | 资源错误 |

## 在线文档

启动服务后，访问以下地址查看在线文档：

- **Swagger UI**: http://localhost:8080/swagger/index.html
- **ReDoc**: http://localhost:8080/swagger/redoc.html

## 工具推荐

- **Swagger Editor**: https://editor.swagger.io/ - 在线编辑 OpenAPI 文档
- **Postman**: 导入 `docs/api.yaml` 进行 API 测试
- **Apifox**: 国产 API 工具，支持 OpenAPI 导入
