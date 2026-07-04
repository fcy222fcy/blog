-- 继续补充文章内容
USE blog;

-- 文章152: 2025年个人技术年度总结
UPDATE articles SET content = '## 引言

2025 年即将结束，回顾这一年的技术成长，收获颇丰。本文总结一下今年在技术方面的学习、实践和思考。

## 技术学习

### Go 语言深入
- 深入学习了 Go 的并发模式
- 掌握了 Context 的正确使用方式
- 学习了 Go 的性能优化技巧
- 阅读了多个开源项目的源码

### 云原生技术
- 学习了 Kubernetes 的基本原理
- 掌握了 Helm Chart 的编写
- 了解了 Service Mesh 的概念
- 实践了 GitOps 工作流

### AI 相关
- 学习了 LLM 的基本原理
- 尝试了 AI 编程助手
- 了解了 RAG 的实现方式
- 探索了 AI 在实际项目中的应用

## 项目成果

### 个人博客系统
- 完成了前后端分离的博客系统
- 实现了评论系统和邮件通知
- 添加了每日一问功能
- 优化了 SEO 和性能

### 开源贡献
- 提交了几个 PR 到开源项目
- 修复了一些文档错误
- 分享了个人项目和经验

## 技术思考

### 关于技术选型
- 不要盲目追新，选择适合的才是最好的
- 考虑团队的学习成本
- 权衡开发效率和运行效率

### 关于代码质量
- 代码可读性比技巧更重要
- 写好测试是保证质量的关键
- 重构要持续进行，不要积累技术债

### 关于学习方法
- 理论与实践相结合
- 输出倒逼输入
- 建立知识体系

## 2026 年规划

### 技术目标
- 深入学习分布式系统设计
- 提升系统设计能力
- 持续关注 AI 领域发展

### 个人目标
- 保持写技术博客的习惯
- 参与更多开源项目
- 提升英语阅读能力

## 结语

2025 年是充实的一年。在技术的道路上，永远没有终点。保持好奇心，持续学习，持续成长。

期待 2026 年能有更多的突破和收获！' WHERE id = 152;

-- 文章153: 使用 Docker Compose 部署个人博客
UPDATE articles SET content = '## 引言

本文记录使用 Docker Compose 部署个人博客的完整过程，包括环境准备、配置编写和部署上线。

## 环境准备

### 服务器要求
- 操作系统：Ubuntu 20.04+ 或 CentOS 7+
- 内存：2GB+
- 硬盘：20GB+
- 已安装 Docker 和 Docker Compose

### 域名和 SSL
- 准备一个域名
- 配置 DNS 解析
- 准备 SSL 证书（可使用 Let''s Encrypt）

## 项目结构

```
blog-deploy/
├── docker-compose.yml
├── nginx/
│   ├── nginx.conf
│   └── ssl/
├── mysql/
│   └── init.sql
├── uploads/
└── backups/
```

## Docker Compose 配置

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    container_name: blog-mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: ${DB_ROOT_PASSWORD}
      MYSQL_DATABASE: blog
      MYSQL_USER: blog_user
      MYSQL_PASSWORD: ${DB_PASSWORD}
    volumes:
      - mysql_data:/var/lib/mysql
      - ./mysql/init.sql:/docker-entrypoint-initdb.d/init.sql
    networks:
      - blog-network

  redis:
    image: redis:7-alpine
    container_name: blog-redis
    restart: always
    networks:
      - blog-network

  backend:
    build: .
    container_name: blog-backend
    restart: always
    environment:
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=blog_user
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=blog
      - REDIS_HOST=redis
    volumes:
      - ./uploads:/app/uploads
    depends_on:
      - mysql
      - redis
    networks:
      - blog-network

  frontend:
    build: ./blog-web
    container_name: blog-frontend
    restart: always
    networks:
      - blog-network

  admin:
    build: ./blog-admin
    container_name: blog-admin
    restart: always
    networks:
      - blog-network

  nginx:
    image: nginx:alpine
    container_name: blog-nginx
    restart: always
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
      - ./nginx/ssl:/etc/nginx/ssl
    depends_on:
      - backend
      - frontend
      - admin
    networks:
      - blog-network

volumes:
  mysql_data:

networks:
  blog-network:
    driver: bridge
```

## 环境变量配置

创建 `.env` 文件：

```env
DB_ROOT_PASSWORD=your_root_password
DB_PASSWORD=your_db_password
```

## Nginx 配置

```nginx
events {
    worker_connections 1024;
}

http {
    upstream frontend {
        server frontend:80;
    }

    upstream backend {
        server backend:8080;
    }

    upstream admin {
        server admin:80;
    }

    server {
        listen 80;
        server_name your-domain.com;
        return 301 https://$server_name$request_uri;
    }

    server {
        listen 443 ssl;
        server_name your-domain.com;

        ssl_certificate /etc/nginx/ssl/cert.pem;
        ssl_certificate_key /etc/nginx/ssl/key.pem;

        location / {
            proxy_pass http://frontend;
        }

        location /api {
            proxy_pass http://backend;
        }

        location /admin {
            proxy_pass http://admin;
        }

        location /uploads {
            alias /app/uploads;
        }
    }
}
```

## 部署步骤

```bash
# 1. 克隆代码
git clone your-repo-url
cd blog-deploy

# 2. 构建镜像
docker-compose build

# 3. 启动服务
docker-compose up -d

# 4. 查看日志
docker-compose logs -f

# 5. 初始化数据库
docker-compose exec mysql mysql -u root -p blog < ./mysql/init.sql
```

## 常用维护命令

```bash
# 查看容器状态
docker-compose ps

# 重启服务
docker-compose restart

# 更新代码并重新部署
git pull
docker-compose build
docker-compose up -d

# 备份数据库
docker-compose exec mysql mysqldump -u root -p blog > backup.sql

# 查看日志
docker-compose logs -f backend
```

## 总结

使用 Docker Compose 部署博客，可以实现一键部署、轻松迁移、方便维护。配合 CI/CD，还可以实现自动化部署。' WHERE id = 153;

-- 文章154: Go 语言高性能 Web 服务实践
UPDATE articles SET content = '## 引言

Go 语言天生适合构建高性能的 Web 服务。本文分享一些在实际项目中积累的性能优化经验。

## 连接池配置

```go
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### 配置建议
- MaxIdleConns: 根据并发量设置，一般 10-50
- MaxOpenConns: 根据数据库承受能力设置
- ConnMaxLifetime: 建议不超过 1 小时

## 并发控制

### 使用 Worker Pool

```go
type WorkerPool struct {
    tasks chan func()
    wg    sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
    p := &WorkerPool{
        tasks: make(chan func()),
    }
    for i := 0; i < size; i++ {
        p.wg.Add(1)
        go p.worker()
    }
    return p
}

func (p *WorkerPool) worker() {
    defer p.wg.Done()
    for task := range p.tasks {
        task()
    }
}

func (p *WorkerPool) Submit(task func()) {
    p.tasks <- task
}

func (p *WorkerPool) Close() {
    close(p.tasks)
    p.wg.Wait()
}
```

### 限制并发数

```go
sem := make(chan struct{}, 100)

for _, item := range items {
    sem <- struct{}{}
    go func(item Item) {
        defer func() { <-sem }()
        process(item)
    }(item)
}
```

## 缓存策略

### 本地缓存

```go
type Cache struct {
    data map[string]interface{}
    mu   sync.RWMutex
    ttl  time.Duration
}

func NewCache(ttl time.Duration) *Cache {
    return &Cache{
        data: make(map[string]interface{}),
        ttl:  ttl,
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```

### Redis 缓存

```go
func GetWithCache(key string, fn func() (interface{}, error)) (interface{}, error) {
    // 先查缓存
    val, err := redis.Get(key).Result()
    if err == nil {
        return val, nil
    }

    // 缓存未命中，查数据库
    val, err = fn()
    if err != nil {
        return nil, err
    }

    // 写入缓存
    redis.Set(key, val, time.Hour)
    return val, nil
}
```

## JSON 优化

### 使用 sonic

```go
import "github.com/bytedance/sonic"

// 序列化
data, _ := sonic.Marshal(obj)

// 反序列化
var obj Type
sonic.Unmarshal(data, &obj)
```

### 复用 Buffer

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func Marshal(v interface{}) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()

    enc := json.NewEncoder(buf)
    if err := enc.Encode(v); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
```

## HTTP 优化

### 使用 fasthttp

```go
import "github.com/valyala/fasthttp"

func requestHandler(ctx *fasthttp.RequestCtx) {
    ctx.SetContentType("application/json")
    ctx.Write([]byte(`{"message":"hello"}`))
}

func main() {
    fasthttp.ListenAndServe(":8080", requestHandler)
}
```

### 连接复用

```go
var client = &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

## 性能分析

### 使用 pprof

```go
import _ "net/http/pprof"

go func() {
    http.ListenAndServe("localhost:6060", nil)
}()
```

### 查看 CPU Profile

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

## 监控指标

### 关键指标
- QPS (每秒请求数)
- 响应时间 (P50, P95, P99)
- 错误率
- 内存使用
- Goroutine 数量

### 使用 Prometheus

```go
import "github.com/prometheus/client_golang/prometheus"

var requestCount = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total HTTP requests",
    },
)

func init() {
    prometheus.MustRegister(requestCount)
}
```

## 总结

高性能不是靠单点优化，而是全方位的持续改进。从架构设计到代码实现，从数据库到缓存，每个环节都有优化空间。' WHERE id = 154;
