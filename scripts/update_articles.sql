-- 丰富文章内容
USE blog;

-- 文章8: Gin 框架快速上手
UPDATE articles SET content = '## 引言

Gin 是一个用 Go 编写的高性能 Web 框架，以其极简的 API 设计和出色的性能而闻名。本文将带你从零开始掌握 Gin 框架的核心用法。

## 为什么选择 Gin？

- **高性能**: 基于httprouter，路由匹配速度极快
- **极简设计**: API 简洁直观，学习成本低
- **中间件支持**: 强大的中间件机制，便于扩展
- **JSON 验证**: 内置请求数据验证功能
- **路由分组**: 支持路由分组和嵌套

## 安装

```bash
go get -u github.com/gin-gonic/gin
```

## 基础用法

### 创建第一个 API

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080")
}
```

### 获取请求参数

```go
// 查询参数: /user?name=tom
r.GET("/user", func(c *gin.Context) {
    name := c.DefaultQuery("name", "guest")
    c.JSON(200, gin.H{"name": name})
})

// 路径参数: /user/123
r.GET("/user/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
})

// 表单提交
r.POST("/form", func(c *gin.Context) {
    name := c.PostForm("name")
    age := c.DefaultPostForm("age", "0")
    c.JSON(200, gin.H{"name": name, "age": age})
})
```

### 绑定 JSON 请求体

```go
type LoginRequest struct {
    Username string `json:"username" binding:"required,min=3,max=20"`
    Password string `json:"password" binding:"required,min=6"`
}

r.POST("/login", func(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // 处理登录逻辑
    c.JSON(200, gin.H{"message": "登录成功"})
})
```

## 路由分组

路由分组可以将相关路由组织在一起，并共享中间件：

```go
func main() {
    r := gin.Default()

    // 公开 API
    v1 := r.Group("/api/v1")
    {
        v1.GET("/articles", listArticles)
        v1.GET("/articles/:id", getArticle)
    }

    // 需要认证的 API
    auth := r.Group("/api/v1")
    auth.Use(AuthMiddleware())
    {
        auth.POST("/articles", createArticle)
        auth.PUT("/articles/:id", updateArticle)
        auth.DELETE("/articles/:id", deleteArticle)
    }

    // 管理员 API
    admin := r.Group("/admin")
    admin.Use(AuthMiddleware(), AdminMiddleware())
    {
        admin.GET("/users", listUsers)
        admin.DELETE("/users/:id", deleteUser)
    }

    r.Run()
}
```

## 中间件

### 全局中间件

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        c.Next()

        latency := time.Since(start)
        status := c.Writer.Status()

        log.Printf("[GIN] %3d | %13v | %15s | %s",
            status, latency, c.ClientIP(), path)
    }
}

func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic recovered: %v", err)
                c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
            }
        }()
        c.Next()
    }
}

func main() {
    r := gin.New()
    r.Use(Logger(), Recovery())
    // ...
}
```

### 认证中间件

```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"error": "未登录"})
            c.Abort()
            return
        }

        // 验证 token
        userID, err := validateToken(token)
        if err != nil {
            c.JSON(401, gin.H{"error": "token 无效"})
            c.Abort()
            return
        }

        c.Set("userID", userID)
        c.Next()
    }
}
```

## 错误处理

```go
// 统一错误响应
func ErrorResponse(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{
        "code":    statusCode,
        "message": message,
    })
}

// 使用示例
r.GET("/article/:id", func(c *gin.Context) {
    id := c.Param("id")
    article, err := getArticleByID(id)
    if err != nil {
        ErrorResponse(c, 404, "文章不存在")
        return
    }
    c.JSON(200, article)
})
```

## 性能对比

Gin 在 TechEmpower 基准测试中表现优异：

| 框架 | 请求/秒 | 延迟 |
|-----|--------|------|
| Gin | ~100k | ~1ms |
| Echo | ~95k | ~1ms |
| net/http | ~80k | ~1.2ms |

## 最佳实践

1. **使用路由分组**: 按功能模块组织路由
2. **合理使用中间件**: 日志、认证、CORS 等
3. **参数验证**: 使用 binding tag 验证输入
4. **错误处理**: 统一错误响应格式
5. **项目结构**: 按功能模块组织代码

## 总结

Gin 框架以其出色的性能和简洁的 API，成为 Go Web 开发的首选框架。掌握这些核心特性，可以帮助你快速构建高质量的 Web 应用。' WHERE id = 8;

-- 文章9: MySQL 索引优化实战
UPDATE articles SET content = '## 引言

索引是数据库性能优化的核心手段。本文将深入讲解 MySQL 索引的原理、类型和优化实战技巧。

## 为什么需要索引？

想象一下在一本 500 页的书中查找特定内容：

- **没有索引**: 逐页翻阅，可能需要几分钟
- **有索引**: 查看索引目录，直接定位到目标页

数据库索引的原理类似，通过 B+ 树结构加速数据检索。

## 索引类型

### 1. B+ 树索引（最常用）

```sql
-- 创建普通索引
CREATE INDEX idx_username ON users(username);

-- 创建唯一索引
CREATE UNIQUE INDEX idx_email ON users(email);

-- 创建复合索引
CREATE INDEX idx_name_age ON users(name, age);
```

### 2. 哈希索引

- 仅支持等值查询
- 不支持范围查询
- 用于 Memory 引擎

### 3. 全文索引

```sql
CREATE FULLTEXT INDEX idx_content ON articles(content);

-- 使用全文搜索
SELECT * FROM articles
WHERE MATCH(content) AGAINST('MySQL 优化' IN BOOLEAN MODE);
```

## 索引优化实战

### 场景1: 慢查询优化

**问题**: 用户列表查询很慢

```sql
-- 原始查询 (全表扫描)
SELECT * FROM users WHERE status = 1 ORDER BY created_at DESC;
```

**解决方案**: 创建复合索引

```sql
-- 分析执行计划
EXPLAIN SELECT * FROM users WHERE status = 1 ORDER BY created_at DESC;

-- 创建索引
CREATE INDEX idx_status_created ON users(status, created_at);
```

### 场景2: 多条件查询

**问题**: 文章筛选查询慢

```sql
-- 原始查询
SELECT * FROM articles
WHERE category_id = 1 AND status = 'published' AND created_at > '2024-01-01';
```

**解决方案**: 遵循最左前缀原则

```sql
-- 创建复合索引 (注意字段顺序)
CREATE INDEX idx_category_status_date
ON articles(category_id, status, created_at);
```

### 场景3: 排序优化

**问题**: 分页查询很慢

```sql
-- 原始查询 (深分页问题)
SELECT * FROM articles
WHERE status = 'published'
ORDER BY created_at DESC
LIMIT 100000, 20;
```

**解决方案**: 使用书签方式

```sql
-- 优化后的查询
SELECT * FROM articles
WHERE status = 'published'
AND created_at < '2024-06-01 12:00:00'
ORDER BY created_at DESC
LIMIT 20;
```

## 索引设计原则

### 最左前缀原则

复合索引 `(a, b, c)` 可以支持：

```sql
WHERE a = 1                          -- 使用索引
WHERE a = 1 AND b = 2                -- 使用索引
WHERE a = 1 AND b = 2 AND c = 3      -- 使用索引
WHERE a = 1 AND c = 3                -- 部分使用索引
WHERE b = 2 AND c = 3                -- 不使用索引
```

### 避免索引失效

```sql
-- 不使用索引的情况
SELECT * FROM users WHERE name LIKE '%tom%';     -- 前导通配符
SELECT * FROM users WHERE age + 1 = 18;          -- 函数操作
SELECT * FROM users WHERE CAST(age AS CHAR) = '18'; -- 类型转换
SELECT * FROM users WHERE status != 1;           -- 不等于操作
```

### 覆盖索引

```sql
-- 如果只需要 id 和 username，可以创建覆盖索引
CREATE INDEX idx_covering ON users(id, username);

-- 这样查询可以完全使用索引，不需要回表
SELECT id, username FROM users WHERE username = 'tom';
```

## 索引监控

### 查看索引使用情况

```sql
-- 查看索引统计
SELECT
    database_name,
    table_name,
    index_name,
    rows_selected,
    rows_inserted,
    rows_updated
FROM performance_schema.table_io_waits_summary_by_index_usage
WHERE database_name = 'blog';
```

### 查看未使用的索引

```sql
SELECT
    object_schema,
    object_name,
    index_name
FROM performance_schema.table_io_waits_summary_by_index_usage
WHERE index_name IS NOT NULL
AND count_star = 0
AND object_schema = 'blog';
```

### 分析索引效率

```sql
-- 查看索引选择性
SELECT
    COUNT(DISTINCT username) / COUNT(*) AS username_selectivity,
    COUNT(DISTINCT status) / COUNT(*) AS status_selectivity
FROM users;
-- 选择性越高，索引效果越好
```

## 索引维护

### 定期优化

```sql
-- 优化表 (重建索引)
OPTIMIZE TABLE users;

-- 分析表 (更新统计信息)
ANALYZE TABLE users;
```

### 监控索引碎片

```sql
SELECT
    table_name,
    data_free,
    ROUND(data_free / 1024 / 1024, 2) AS碎片大小MB
FROM information_schema.tables
WHERE table_schema = 'blog'
AND data_free > 0;
```

## 实战案例

### 电商系统索引设计

```sql
-- 用户表
CREATE INDEX idx_user_phone ON users(phone);           -- 手机号登录
CREATE INDEX idx_user_status ON users(status, created_at); -- 用户列表

-- 商品表
CREATE INDEX idx_product_category ON products(category_id, status); -- 分类筛选
CREATE INDEX idx_product_price ON products(price);     -- 价格排序
CREATE INDEX idx_product_name ON products(name);       -- 商品搜索

-- 订单表
CREATE INDEX idx_order_user ON orders(user_id, created_at); -- 用户订单
CREATE INDEX idx_order_status ON orders(status, created_at); -- 订单列表
```

## 总结

1. **索引不是越多越多**: 每个索引都会影响写入性能
2. **遵循最左前缀**: 复合索引字段顺序很重要
3. **避免索引失效**: 注意查询条件的写法
4. **定期监控**: 发现并清理无用索引
5. **覆盖索引**: 减少回表，提升查询性能

合理使用索引是数据库性能优化的关键，但过度索引也会带来问题。需要在查询性能和写入性能之间找到平衡点。' WHERE id = 9;

-- 文章3: Cloudflare Workers AI 助手
UPDATE articles SET content = '## 引言

想给博客添加一个智能 AI 助手，但又不想花钱买服务器？Cloudflare Workers 提供了一个完美的免费解决方案。本文将手把手教你实现。

## 为什么选择 Cloudflare Workers？

- **免费额度充足**: 每天 10 万次请求
- **全球边缘节点**: 响应速度快
- **无需服务器**: 纯 Serverless 架构
- **AI 能力集成**: 支持多种 AI 模型

## 准备工作

1. 注册 Cloudflare 账号
2. 安装 Node.js
3. 安装 Wrangler CLI

```bash
npm install -g wrangler
wrangler login
```

## 创建 Worker

### 1. 初始化项目

```bash
wrangler init blog-ai-assistant
cd blog-ai-assistant
```

### 2. 配置 wrangler.toml

```toml
name = "blog-ai-assistant"
main = "src/index.js"
compatibility_date = "2024-01-01"

[ai]
binding = "AI"
```

### 3. 编写 Worker 代码

```javascript
export default {
  async fetch(request, env) {
    // CORS 处理
    if (request.method === "OPTIONS") {
      return new Response(null, {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "POST, OPTIONS",
          "Access-Control-Allow-Headers": "Content-Type",
        },
      });
    }

    if (request.method !== "POST") {
      return new Response("Method not allowed", { status: 405 });
    }

    try {
      const { question, context } = await request.json();

      // 使用 Cloudflare AI
      const response = await env.AI.run("@cf/meta/llama-3-8b-instruct", {
        messages: [
          {
            role: "system",
            content: `你是一个专业的博客助手。根据以下博客内容回答问题：
            ${context || "这是一个技术博客，分享编程、AI 和生活相关的内容。"}
            请用简洁友好的方式回答。`
          },
          {
            role: "user",
            content: question
          }
        ],
        max_tokens: 500,
      });

      return new Response(
        JSON.stringify({ answer: response.response }),
        {
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
        }
      );
    } catch (error) {
      return new Response(
        JSON.stringify({ error: "服务暂时不可用" }),
        { status: 500 }
      );
    }
  },
};
```

## 接入博客

### 前端代码

```html
<!-- AI 助手组件 -->
<div id="ai-assistant" class="ai-chat">
  <div class="ai-header">
    <span>AI 助手</span>
    <button onclick="toggleAI()">×</button>
  </div>
  <div id="chat-messages" class="ai-messages">
    <div class="message ai">
      你好！我是 AI 助手，有什么问题可以问我~
    </div>
  </div>
  <div class="ai-input">
    <input type="text" id="ai-input" placeholder="输入你的问题..." />
    <button onclick="sendQuestion()">发送</button>
  </div>
</div>

<script>
const WORKER_URL = "https://your-worker.workers.dev";

async function sendQuestion() {
  const input = document.getElementById("ai-input");
  const question = input.value.trim();
  if (!question) return;

  // 显示用户消息
  addMessage(question, "user");
  input.value = "";

  try {
    const response = await fetch(WORKER_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        question: question,
        context: document.querySelector("article")?.innerText || ""
      }),
    });

    const data = await response.json();
    addMessage(data.answer, "ai");
  } catch (error) {
    addMessage("抱歉，服务暂时不可用", "ai");
  }
}

function addMessage(text, type) {
  const container = document.getElementById("chat-messages");
  const div = document.createElement("div");
  div.className = `message ${type}`;
  div.textContent = text;
  container.appendChild(div);
  container.scrollTop = container.scrollHeight;
}
</script>
```

### CSS 样式

```css
.ai-chat {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 350px;
  background: var(--card-bg);
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.15);
  overflow: hidden;
}

.ai-header {
  padding: 12px 16px;
  background: var(--accent-color);
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.ai-messages {
  height: 300px;
  overflow-y: auto;
  padding: 16px;
}

.message {
  margin-bottom: 12px;
  padding: 8px 12px;
  border-radius: 8px;
  max-width: 80%;
}

.message.user {
  background: var(--accent-color);
  color: white;
  margin-left: auto;
}

.message.ai {
  background: var(--bg-color);
  border: 1px solid #eee;
}

.ai-input {
  display: flex;
  padding: 12px;
  border-top: 1px solid #eee;
}

.ai-input input {
  flex: 1;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.ai-input button {
  margin-left: 8px;
  padding: 8px 16px;
  background: var(--accent-color);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
```

## 部署

```bash
wrangler deploy
```

## 成本对比

| 方案 | 月费用 | 并发能力 | 维护成本 |
|-----|-------|---------|---------|
| Cloudflare Workers | $0 | 高 | 低 |
| VPS + API | $5-20 | 中 | 高 |
| 云函数 | $0-5 | 中 | 中 |

## 进阶优化

### 1. 缓存常见问题

```javascript
const cache = new Map();

async function getAnswer(question, env) {
  const key = question.toLowerCase().trim();
  if (cache.has(key)) {
    return cache.get(key);
  }

  const answer = await env.AI.run(/* ... */);
  cache.set(key, answer);

  // 5分钟后过期
  setTimeout(() => cache.delete(key), 5 * 60 * 1000);

  return answer;
}
```

### 2. 流式响应

```javascript
// 使用 Streaming 提升用户体验
const stream = await env.AI.run("@cf/meta/llama-3-8b-instruct", {
  messages: [{ role: "user", content: question }],
  stream: true,
});

return new Response(stream, {
  headers: { "Content-Type": "text/event-stream" },
});
```

## 总结

Cloudflare Workers 为博客 AI 助手提供了完美的免费解决方案。零成本、高性能、易部署，非常适合个人博客使用。快来试试吧！' WHERE id = 3;
