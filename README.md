<p align="center">
  <h1 align="center">✍️ Gin 博客系统</h1>
</p>

<p align="center">
  <a href="https://github.com/fcy222fcy/blog/stargazers"><img src="https://img.shields.io/github/stars/fcy222fcy/blog?style=flat-square" alt="Stars"></a>
  <a href="https://github.com/fcy222fcy/blog/forks"><img src="https://img.shields.io/github/forks/fcy222fcy/blog?style=flat-square" alt="Forks"></a>
  <a href="https://github.com/fcy222fcy/blog/issues"><img src="https://img.shields.io/github/issues/fcy222fcy/blog?style=flat-square" alt="Issues"></a>
  <a href="https://github.com/fcy222fcy/blog/blob/main/LICENSE"><img src="https://img.shields.io/github/license/fcy222fcy/blog?style=flat-square" alt="License"></a>
  <img src="https://img.shields.io/github/languages/top/fcy222fcy/blog?style=flat-square" alt="Language">
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome">
</p>

<p align="center">
  基于 <strong>Go + Gin + GORM + Vue 3</strong> 开发的现代化博客系统
  <br>
  前后端分离 · 双前端架构 · Docker 一键部署
</p>

---

## ✨ 功能特性

- 🚀 **现代化技术栈** — Go 1.25 / Gin 1.12 / GORM 1.31 / Vue 3 / Vite 5
- 🎨 **双前端架构** — 独立的前台展示 `blog-web` + 后台管理 `blog-admin`
- 📝 **内容管理** — 文章、分类、标签、草稿、定时发布
- 💬 **评论系统** — 游客/登录双模式、多级回复、热度排序、点赞、博主标识防冒充
- 🔐 **安全认证** — JWT 令牌 + bcrypt 密码加密 + 审计日志
- 🖼️ **智能头像** — QQ 头像 → Gravatar → 首字母生成，三级兜底
- 🤔 **每日一问** — 问答卡片，答案渐变遮罩，边看边想
- 🔍 **全文搜索** + 📂 时间归档 + 🏷️ 分类/标签浏览
- 📡 **SEO 友好** — 自动生成 RSS 订阅 + Sitemap 站点地图
- 🐳 **Docker 部署** — 一条命令启动 MySQL + Redis + 后端 + 双前端
- 📱 **响应式设计** — 完美适配桌面 / 平板 / 手机

---

## 🏗️ 技术栈

### 后端

<p>
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Gin-1.12-00ADD8?style=flat-square" alt="Gin">
  <img src="https://img.shields.io/badge/GORM-1.31-20B2AA?style=flat-square" alt="GORM">
  <img src="https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat-square&logo=mysql&logoColor=white" alt="MySQL">
  <img src="https://img.shields.io/badge/Redis-7-DC382D?style=flat-square&logo=redis&logoColor=white" alt="Redis">
  <img src="https://img.shields.io/badge/JWT-v5-000000?style=flat-square&logo=json-web-tokens&logoColor=white" alt="JWT">
  <img src="https://img.shields.io/badge/Zap-1.27-00ADD8?style=flat-square" alt="Zap">
  <img src="https://img.shields.io/badge/Viper-1.19-00ADD8?style=flat-square" alt="Viper">
</p>

### 前台前端 (blog-web)

<p>
  <img src="https://img.shields.io/badge/Vue-3.4-4FC08D?style=flat-square&logo=vue.js&logoColor=white" alt="Vue">
  <img src="https://img.shields.io/badge/Vite-5.2-646CFF?style=flat-square&logo=vite&logoColor=white" alt="Vite">
  <img src="https://img.shields.io/badge/Pinia-2.1-F7D336?style=flat-square" alt="Pinia">
  <img src="https://img.shields.io/badge/Vue_Router-4.3-4FC08D?style=flat-square&logo=vue.js&logoColor=white" alt="Vue Router">
  <img src="https://img.shields.io/badge/TailwindCSS-3.4-38B2AC?style=flat-square&logo=tailwind-css&logoColor=white" alt="TailwindCSS">
  <img src="https://img.shields.io/badge/Axios-1.7-5A29E4?style=flat-square&logo=axios&logoColor=white" alt="Axios">
</p>

### 后台前端 (blog-admin)

<p>
  <img src="https://img.shields.io/badge/Vue-3.4-4FC08D?style=flat-square&logo=vue.js&logoColor=white" alt="Vue">
  <img src="https://img.shields.io/badge/Element_Plus-2.7-409EFF?style=flat-square&logo=element&logoColor=white" alt="Element Plus">
  <img src="https://img.shields.io/badge/Markdown-Editor-519ABA?style=flat-square&logo=markdown&logoColor=white" alt="Markdown Editor">
  <img src="https://img.shields.io/badge/Pinia-2.1-F7D336?style=flat-square" alt="Pinia">
</p>

---

## 📁 项目结构

```
gin博客/
├── 🎨 blog-web/              # 前台展示前端 (Vue 3 + Tailwind)
├── 🛠️ blog-admin/            # 后台管理前端 (Vue 3 + Element Plus)
├── ⚙️ internal/              # Go 后端核心代码
│   ├── api/v1/            # API 控制器 & 路由
│   ├── app/               # 应用启动入口
│   ├── middleware/        # 中间件 (认证/审计/日志/恢复)
│   ├── model/             # 数据模型 (DTO / Entity)
│   ├── repository/        # 数据访问层
│   └── service/           # 业务逻辑层
├── 📦 pkg/                 # 公共工具包 (JWT / 日志 / 加密 / 头像 ...)
├── 🐳 docker-compose.yml   # Docker 编排
└── 🔧 .env.example         # 环境变量示例
```

> 💡 采用经典的 Controller → Service → Repository 三层架构

---

## 🚀 快速开始

### 方式一：Docker 一键部署 ⭐（推荐）

```bash
# 1. 克隆项目
git clone https://github.com/fcy222fcy/blog.git
cd blog

# 2. 配置环境变量
cp .env.example .env
# 编辑 .env，修改密码和 JWT 密钥

# 3. 一键启动全部服务
docker-compose up -d
```

启动完成后访问：
- 🌐 **前台**：http://localhost
- 🔧 **后台**：http://localhost:8081
- ⚙️ **API**：http://localhost:8080

### 方式二：本地开发

#### 前置要求

- [Go](https://go.dev/) ≥ 1.25
- [Node.js](https://nodejs.org/) ≥ 18
- [MySQL](https://www.mysql.com/) ≥ 8.0
- [Redis](https://redis.io/) ≥ 7（可选）

#### 启动后端

```bash
# 1. 配置 .env 中的数据库连接
cp .env.example .env

# 2. 安装依赖
go mod download

# 3. 启动服务（首次启动自动迁移数据库表）
go run internal/app/app.go
# 默认监听 :8080
```

#### 启动前台

```bash
cd blog-web
npm install
npm run dev    # http://localhost:5173
```

#### 启动后台

```bash
cd blog-admin
npm install
npm run dev    # http://localhost:5174
```

---

## ⚙️ 配置说明

核心环境变量参考 [.env.example](.env.example)：

| 变量 | 说明 |
|------|------|
| `DB_*` | MySQL 数据库连接信息 |
| `JWT_SECRET` | JWT 签名密钥（生产环境务必修改） |
| `JWT_EXPIRE_HOURS` | Token 过期时间，默认 7 天 |
| `REDIS_*` | Redis 连接配置 |
| `EMAIL_*` | SMTP 邮件配置（评论通知） |
| `CORS_ORIGINS` | 允许跨域的域名，逗号分隔 |

> 🔐 **博主身份**：通过配置中的 `blogger.user_id` 标识，防止前端伪造博主身份

---

## 💬 评论系统设计

### 双模式支持

| 模式 | 身份获取 | 可自定义 |
|------|----------|----------|
| 🧑‍💼 游客 | 手动填写昵称 + 邮箱 + 网站 | ✅ 昵称/邮箱/网站 |
| 👤 登录用户 | 自动关联用户信息 | ✅ 昵称/头像 |
| 👑 博主 | 强制使用配置值，不可覆盖 | ❌ 防冒充 |

### 热度排序（后端分页一致性保证）

```
热度分 = 点赞数 × 2 + 所有层级回复数
同分数按创建时间倒序
```

### 智能头像获取

```
QQ 邮箱 → 提取 QQ 号 → qlogo.cn
                ↓ 未命中
          Gravatar MD5
                ↓ 未命中
        首字母生成 SVG 兜底
```

---

## 🧪 测试

```bash
# 运行全部测试
go test ./test/...

# 分类测试
go test ./test/unit/...          # 单元测试
go test ./test/integration/...   # 集成测试
go test ./test/comprehensive/... # 综合测试
```

详细测试报告：[test/comprehensive/TEST_REPORT.md](test/comprehensive/TEST_REPORT.md)

---

## 📖 文档

- [API 接口文档](docs/README.md) · [OpenAPI 规范](docs/api.yaml)
- [后端代码开发规范](docs/后端代码开发规范.md)
- [项目设计文档](docs/设计文档.md)
- [搜索功能设计](docs/搜索功能设计文档.md)
- [用户头像获取设计](docs/用户头像获取功能设计.md)
- [完整测试用例](docs/完整测试用例.md)

---

## 🤝 参与贡献

任何形式的贡献都非常欢迎！👏

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/your-feature`
3. 提交改动：`git commit -m 'feat: add some feature'`
4. 推送到分支：`git push origin feature/your-feature`
5. 提交 Pull Request

> 💡 提交信息建议遵循 [Conventional Commits](https://www.conventionalcommits.org/) 规范

---

## 📄 许可证

MIT © [fcy222fcy](https://github.com/fcy222fcy)

---

<p align="center">
  <samp>
    如果这个项目对你有帮助的话，不妨点个 ⭐ Star 支持一下~
  </samp>
</p>
