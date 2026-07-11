# Gin 博客系统

一个基于 Go + Gin + GORM 开发的现代化博客系统，采用前后端分离架构，包含前台展示和后台管理两个独立前端应用。

## 项目特性

- **现代化技术栈**：Go 1.25 + Gin 1.12 + GORM 1.31 + Vue 3 + Vite 5
- **前后端分离**：独立的前台展示 (blog-web) 和后台管理 (blog-admin) 前端
- **完善的内容管理**：文章、分类、标签、评论、每日一问等功能模块
- **用户系统**：JWT 认证、博主标识、游客/登录用户双模式评论
- **富文本编辑**：Markdown 编辑器支持，代码高亮
- **SEO 友好**：自动生成 RSS 订阅和 Sitemap 站点地图
- **审计日志**：后台操作全记录
- **头像智能获取**：QQ 头像 > Gravatar > 首字母生成三级兜底
- **Docker 一键部署**：MySQL + Redis + 后端 + 双前端完整编排

## 技术架构

### 后端

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.25 | 编程语言 |
| Gin | 1.12 | Web 框架 |
| GORM | 1.31 | ORM 框架 |
| MySQL | 8.0+ | 主数据库 |
| SQLite | - | 嵌入式数据库支持 |
| Redis | 7 | 缓存/会话存储 |
| JWT | v5 | 身份认证 |
| Viper | 1.19 | 配置管理 |
| Zap | 1.27 | 结构化日志 |
| bcrypt | - | 密码加密 |

### 前台前端 (blog-web)

| 技术 | 版本 | 说明 |
|------|------|------|
| Vue | 3.4 | UI 框架 |
| Vite | 5.2 | 构建工具 |
| Vue Router | 4.3 | 路由管理 |
| Pinia | 2.1 | 状态管理 |
| Tailwind CSS | 3.4 | CSS 框架 |
| Axios | 1.7 | HTTP 客户端 |
| Marked | 12.0 | Markdown 渲染 |
| highlight.js | 11.9 | 代码高亮 |

### 后台前端 (blog-admin)

| 技术 | 版本 | 说明 |
|------|------|------|
| Vue | 3.4 | UI 框架 |
| Vite | 5.2 | 构建工具 |
| Vue Router | 4.3 | 路由管理 |
| Pinia | 2.1 | 状态管理 |
| Element Plus | 2.7 | UI 组件库 |
| md-editor-v3 | 4.21 | Markdown 编辑器 |
| Axios | 1.7 | HTTP 客户端 |

## 目录结构

```
gin博客/
├── blog-web/                 # 前台展示前端
│   ├── src/
│   │   ├── api/           # API 接口封装
│   │   ├── assets/        # 静态资源
│   │   ├── components/    # 公共组件
│   │   ├── router/      # 路由配置
│   │   ├── stores/      # Pinia 状态
│   │   ├── utils/       # 工具函数
│   │   ├── views/       # 页面视图
│   │   ├── App.vue
│   │   └── main.js
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── Dockerfile
├── blog-admin/             # 后台管理前端
│   ├── src/
│   │   ├── api/           # API 接口封装
│   │   ├── components/    # 公共组件
│   │   ├── router/      # 路由配置
│   │   ├── views/       # 页面视图
│   │   ├── App.vue
│   │   └── main.js
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── Dockerfile
├── internal/              # Go 后端核心代码
│   ├── api/v1/          # API 控制器与路由
│   │   ├── article/      # 文章模块
│   │   ├── auth/         # 认证模块
│   │   ├── category/    # 分类模块
│   │   ├── comment/  # 评论模块
│   │   ├── tag/          # 标签模块
│   │   ├── daily_question/  # 每日一问
│   │   ├── user/       # 用户模块
│   │   ├── audit_log/   # 审计日志
│   │   ├── about_page/  # 关于页
│   │   ├── rss/        # RSS订阅
│   │   └── sitemap/  # 站点地图
│   ├── app/             # 应用启动入口
│   ├── middleware/      # 中间件（认证、审计、日志、恢复）
│   ├── model/
│   │   ├── dto/        # 数据传输对象（请求/响应）
│   │   └── entity/   # 数据库实体模型
│   ├── repository/  # 数据访问层
│   └── service/     # 业务逻辑层
├── pkg/                 # 公共工具包
│   ├── bcrypt/         # 密码加密
│   ├── config/       # 配置加载
│   ├── database/   # 数据库初始化
│   ├── email/      # 邮件发送
│   ├── errors/     # 错误定义
│   ├── gravatar/  # 头像获取
│   ├── jwt/        # JWT 工具
│   ├── logger/   # 日志封装
│   ├── response/  # 统一响应
│   └── ua/         # UA 解析
├── configs/             # 配置与初始化 SQL
├── docs/              # 设计文档与 API 文档
├── scripts/           # 数据库脚本
├── test/              # 测试代码（单元/集成/E2E）
├── prototypes/      # 静态原型页面
├── public/          # 生成的静态文件（RSS/Sitemap）
├── docker-compose.yml    # Docker 编排文件
├── Dockerfile          # 后端镜像构建
├── go.mod
├── .env.example        # 环境变量示例
└── nginx.conf
```

## 功能模块

### 前台页面

| 模块 | 说明 |
|------|------|
| 首页 | 文章列表、分页、搜索入口 |
| 文章详情 | Markdown 渲染、代码高亮、点赞、评论区 |
| 分类归档 | 按分类浏览文章 |
| 标签归档 | 按标签浏览文章 |
| 时间归档 | 按年月浏览文章 |
| 每日一问 | 每日问答卡片（答案渐变遮盖） |
| 关于页面 | 博主介绍、项目展示 |
| 搜索功能 | 全文搜索文章 |
| 评论系统 | 支持游客/登录用户、多级回复、点赞、热度排序 |
| RSS / Sitemap | SEO 支持 |

### 后台管理

| 模块 | 说明 |
|------|------|
| 仪表盘 | 数据统计概览 |
| 文章管理 | 新建/编辑/删除文章，支持草稿、定时发布 |
| 分类管理 | 分类增删改查 |
| 标签管理 | 标签增删改查 |
| 评论管理 | 评论审核、删除、回复 |
| 每日一问 | 问答管理 |
| 媒体库 | 文件上传管理 |
| 关于页 | 关于页内容编辑 |
| 审计日志 | 操作记录查询 |
| 修改密码 | 管理员密码修改 |

## 快速开始

### 环境要求

- Go 1.25+
- Node.js 18+
- MySQL 8.0+
- Redis 7+（可选）

### 方式一：Docker 一键部署（推荐）

```bash
# 1. 复制环境变量配置
cp .env.example .env

# 2. 根据需要修改 .env 中的配置（特别是密码和密钥）

# 3. 启动所有服务
docker-compose up -d

# 4. 访问
# 前台: http://localhost
# 后台: http://localhost:8081
# 后端 API: http://localhost:8080
```

### 方式二：本地开发

#### 1. 启动数据库

```bash
# 使用 Docker 启动 MySQL 和 Redis
docker run -d --name mysql-blog \
  -e MYSQL_ROOT_PASSWORD=root_password \
  -e MYSQL_DATABASE=blog \
  -e MYSQL_USER=blog_user \
  -e MYSQL_PASSWORD=your_password \
  -p 3306:3306 mysql:8.0

docker run -d --name redis-blog -p 6379:6379 redis:7-alpine
```

#### 2. 配置环境变量

```bash
cp .env.example .env
# 编辑 .env 填入数据库连接等配置
```

#### 3. 启动后端服务

```bash
# 安装依赖
go mod download

# 初始化数据库表结构（自动迁移）
# 首次启动会自动创建表

# 启动服务
go run internal/app/app.go
# 或编译后运行
go build -o blog-server internal/app/app.go
./blog-server
```

后端服务启动在 `http://localhost:8080`

#### 4. 启动前台前端

```bash
cd blog-web

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前台启动在 `http://localhost:5173`

#### 5. 启动后台前端

```bash
cd blog-admin

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

后台启动在 `http://localhost:5174`（或 Vite 自动分配的端口）

## 配置说明

### 环境变量 (.env)

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `DB_HOST` | MySQL 主机 | localhost |
| `DB_PORT` | MySQL 端口 | 3306 |
| `DB_USER` | MySQL 用户名 | blog_user |
| `DB_PASSWORD` | MySQL 密码 | your_password |
| `DB_NAME` | 数据库名 | blog |
| `JWT_SECRET` | JWT 签名密钥 | - |
| `JWT_EXPIRE_HOURS` | Token 过期时间（小时） | 168 |
| `REDIS_HOST` | Redis 主机 | localhost |
| `REDIS_PORT` | Redis 端口 | 6379 |
| `EMAIL_HOST` | SMTP 服务器 | - |
| `EMAIL_PORT` | SMTP 端口 | - |
| `EMAIL_USER` | 邮箱账号 | - |
| `EMAIL_PASSWORD` | 邮箱授权码 | - |
| `APP_ENV` | 运行环境 | production |
| `APP_PORT` | 服务端口 | 8080 |
| `APP_URL` | 站点 URL | - |
| `CORS_ORIGINS` | 允许的跨域来源 | - |

### 博主配置

博主账号信息通过配置文件管理，包含：
- `user_id`：博主用户 ID（防冒充标识）
- `username`：登录用户名
- `password`：登录密码
- `nickname`：显示昵称
- `avatar`：头像路径
- `email`：联系邮箱

## 评论系统设计

### 认证模式

- **游客模式**：填写昵称、邮箱、网站
- **登录用户模式**：自动关联用户信息，可修改昵称头像
- **博主标识**：仅当登录用户 ID 匹配配置 blogger.user_id 时显示博主标识（防冒充）

### 排序规则

评论在后端全量排序后再分页，确保分页一致性：

```
热度分 = 点赞数 × 2 + 所有层级回复数
分数相同 → 按创建时间倒序
```

### 头像获取优先级

1. QQ 邮箱 → 提取 QQ 号 → qlogo.cn 接口
2. 其他邮箱 → Gravatar MD5 哈希
3. 兜底 → 根据昵称首字母生成 SVG 头像

## API 接口

接口文档详见 [docs/README.md](docs/README.md) 及 [docs/api.yaml](docs/api.yaml)

### 主要接口分组

- `GET    /api/v1/articles`          # 文章列表
- `GET    /api/v1/articles/:id`  # 文章详情
- `POST   /api/v1/auth/login`      # 登录
- `GET    /api/v1/comments`         # 评论列表
- `POST   /api/v1/comments`        # 发表评论
- `GET    /api/v1/daily-questions/today`  # 今日一问
- `GET    /rss.xml`              # RSS 订阅
- `GET    /sitemap.xml`          # Sitemap

## 开发规范

### 后端代码规范

详见 [docs/后端代码开发规范.md](docs/后端代码开发规范.md)

### 分层架构

```
API Controller → Service → Repository → Database
     ↓            ↓            ↓
  请求校验    业务逻辑    数据访问
  参数绑定    事务控制    CRUD 操作
  响应封装    规则校验    持久化
```

### Git 提交流程

1. 新建分支：`feature/xxx` 或 `fix/xxx`
2. 提交信息格式：`<type>: <subject>`
   - `feat`: 新功能
   - `fix`: 修复 bug
   - `docs`: 文档更新
   - `style`: 代码格式调整
   - `refactor`: 重构
   - `test`: 测试相关
   - `chore`: 构建/工具变更

## 测试

项目包含完整的测试套件：

```bash
# 单元测试
go test ./test/unit/...

# 集成测试
go test ./test/integration/...

# 综合测试
go test ./test/comprehensive/...

# 全部测试
go test ./test/...
```

测试报告详见 [test/comprehensive/TEST_REPORT.md](test/comprehensive/TEST_REPORT.md)

## 部署

### Docker 生产部署

```bash
# 构建并启动
docker-compose up -d --build

# 查看日志
docker-compose logs -f backend frontend admin

# 停止服务
docker-compose down
```

### Nginx 反向代理

各服务已包含 `nginx.conf` 配置文件，可根据需要调整。

## 许可证

MIT License

## 贡献指南

欢迎提交 Issue 和 Pull Request！
