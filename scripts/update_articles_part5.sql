-- 继续补充文章内容
USE blog;

-- 文章172: 2024年年终总结：成长与收获
UPDATE articles SET content = '## 引言

2024 年即将结束，回顾这一年，在工作、学习、生活等方面都有不少收获和感悟。

## 工作回顾

### 项目成果
- 完成了博客系统的开发和部署
- 优化了多个项目的性能
- 学习并实践了新技术栈

### 技术成长
- 深入学习了 Go 语言
- 掌握了 Docker 和 Kubernetes
- 了解了云原生架构

### 工作方法
- 学会了更高效的时间管理
- 提升了代码审查能力
- 改进了文档编写习惯

## 学习收获

### 技术书籍
阅读了几本有影响力的技术书籍：
- 《深入理解计算机系统》
- 《设计模式》
- 《重构》

### 在线课程
完成了多个在线课程的学习：
- Kubernetes 实战
- 分布式系统设计
- 系统性能优化

### 开源项目
参与了几个开源项目，学到了很多。

## 生活感悟

### 健康
- 开始注重运动和饮食
- 学会了劳逸结合

### 家庭
- 多陪伴家人
- 经营好家庭关系

### 兴趣
- 保持阅读的习惯
- 发展新的兴趣爱好

## 2025 年展望

### 技术目标
- 深入学习系统设计
- 提升架构能力
- 持续关注 AI 发展

### 个人目标
- 保持健康的生活方式
- 多读书多思考
- 平衡工作和生活

## 结语

2024 年是充实的一年。感谢所有帮助过我的人，也感谢努力的自己。

2025 年，继续前行！' WHERE id = 172;

-- 文章173: Gin 框架实战：构建 RESTful API
UPDATE articles SET content = '## 引言

Gin 是 Go 语言中最流行的 Web 框架之一，非常适合构建 RESTful API。本文通过实战项目介绍 Gin 的核心用法。

## 项目结构

```
my-api/
├── main.go
├── config/
│   └── config.go
├── middleware/
│   └── auth.go
├── models/
│   └── user.go
├── handlers/
│   └── user.go
├── routes/
│   └── routes.go
└── utils/
    └── response.go
```

## 基础设置

### 初始化项目

```bash
mkdir my-api
cd my-api
go mod init my-api
go get -u github.com/gin-gonic/gin
```

### 基础代码

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run(":8080")
}
```

## 路由设计

### RESTful 规范

| 方法 | 路径 | 说明 |
|-----|------|------|
| GET | /api/users | 获取用户列表 |
| GET | /api/users/:id | 获取用户详情 |
| POST | /api/users | 创建用户 |
| PUT | /api/users/:id | 更新用户 |
| DELETE | /api/users/:id | 删除用户 |

### 路由定义

```go
func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.GET("", listUsers)
            users.GET("/:id", getUser)
            users.POST("", createUser)
            users.PUT("/:id", updateUser)
            users.DELETE("/:id", deleteUser)
        }
    }
}
```

## 中间件

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

### 日志中间件

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        c.Next()

        latency := time.Since(start)
        status := c.Writer.Status()

        log.Printf("[%s] %s %s %d %v",
            c.ClientIP(), c.Request.Method, path, status, latency)
    }
}
```

## 数据验证

### 使用 binding tag

```go
type CreateUserRequest struct {
    Username string `json:"username" binding:"required,min=3,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func createUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // 创建用户...
}
```

## 统一响应格式

```go
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(200, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

func Error(c *gin.Context, code int, message string) {
    c.JSON(code, Response{
        Code:    code,
        Message: message,
    })
}
```

## 错误处理

```go
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            Error(c, 500, err.Error())
        }
    }
}
```

## 单元测试

```go
func TestCreateUser(t *testing.T) {
    router := SetupRouter()

    body := `{"username":"test","email":"test@example.com","password":"123456"}`
    req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
}
```

## 总结

Gin 框架简洁高效，非常适合构建 RESTful API。掌握路由、中间件、数据验证等核心概念，就能快速开发出高质量的 API 服务。' WHERE id = 173;

-- 文章175: 周末烘焙：第一次做蛋糕
UPDATE articles SET content = '## 起因

周末在家闲来无事，突然想尝试做蛋糕。虽然平时都是写代码，但偶尔换换脑子也不错。

## 准备工作

### 材料清单
- 鸡蛋 3 个
- 低筋面粉 100g
- 细砂糖 60g
- 牛奶 40ml
- 植物油 30ml
- 柠檬汁 几滴

### 工具准备
- 电动打蛋器
- 打蛋盆
- 筛网
- 烤箱
- 6寸蛋糕模具

## 制作过程

### 1. 分离蛋清蛋黄

这是最关键的一步，蛋清里不能有一点蛋黄，否则打发会失败。

### 2. 制作蛋黄糊

```步骤
1. 蛋黄中加入 20g 糖，搅拌均匀
2. 加入牛奶和植物油，搅拌均匀
3. 筛入低筋面粉，用刮刀翻拌均匀
```

### 3. 打发蛋清

```步骤
1. 蛋清中加入柠檬汁
2. 分三次加入 40g 糖
3. 打发至硬性发泡（提起打蛋器有直立尖角）
```

### 4. 混合面糊

取 1/3 蛋白霜到蛋黄糊中，翻拌均匀，再倒回剩余蛋白霜中，继续翻拌。

### 5. 烘烤

```温度
上火：150°C
下火：150°C
时间：45 分钟
```

## 结果

第一次做蛋糕，虽然没有那么完美，但成就感满满！

### 成功的地方
- 蛋糕成功膨胀起来了
- 口感还算松软
- 没有塌陷

### 需要改进
- 表面有些开裂
- 颜色不够均匀
- 脱模时有点粘

## 经验总结

1. 蛋清打发很重要，要打到硬性发泡
2. 翻拌时要轻柔，避免消泡
3. 烤箱温度要准确
4. 耐心等待，不要频繁开烤箱门

## 下次计划

- 尝试戚风蛋糕
- 学习裱花
- 做一个完整的生日蛋糕

## 结语

烘焙和编程一样，都需要耐心和实践。虽然第一次不够完美，但我会继续尝试的！

生活不止代码，还有诗和远方（以及蛋糕）。' WHERE id = 175;

-- 文章177: 我的 2024 阅读清单
UPDATE articles SET content = '## 引言

2024 年读了不少书，涵盖技术、思维、生活等方面。分享一下今年的阅读清单和感悟。

## 技术类

### 《深入理解计算机系统》
经典中的经典，虽然有些难度，但收获巨大。对计算机系统有了更深入的理解。

### 《Go 语言设计与实现》
深入了解了 Go 语言的内部实现，对编写高性能 Go 代码很有帮助。

### 《凤凰架构》
关于分布式系统架构的书，开阔了视野。

## 思维类

### 《思考，快与慢》
丹尼尔·卡尼曼的经典著作，讲述了人类思维的两个系统。

### 《原则》
瑞·达利欧的人生和工作原则，很有启发。

### 《刻意练习》
关于如何高效学习和提升技能。

## 生活类

### 《断舍离》
学会了放下不必要的东西，简化生活。

### 《小王子》
经典重读，每次读都有新的感悟。

### 《人类简史》
从宏观角度理解人类社会的发展。

## 阅读方法

### 我的阅读习惯
- 每天阅读 30 分钟
- 使用 Kindle 阅读
- 记录读书笔记
- 定期回顾

### 读书笔记方法
```markdown
## 书籍信息
- 书名：
- 作者：
- 阅读时间：

## 核心观点
1.
2.
3.

## 个人感悟

## 行动计划
```

## 2025 阅读计划

### 计划阅读
- 《代码整洁之道》
- 《重构》
- 《设计模式》
- 更多非技术类书籍

### 目标
- 每月至少 2 本书
- 输出更多读书笔记
- 保持阅读习惯

## 结语

阅读是最低成本的自我提升方式。2025 年，继续阅读，继续成长。' WHERE id = 177;

-- 文章179: 旅行：厦门三日游
UPDATE articles SET content = '## 出发

趁年假去了趟厦门，这座海滨城市一直是我向往的地方。

## Day 1：鼓浪屿

### 上岛
从东渡码头坐船上岛，船票需要提前预约。

### 游览路线
- 日光岩：鼓浪屿最高点，俯瞰全岛
- 菽庄花园：园林建筑很美
- 龙头路：各种小吃和纪念品

### 美食
- 鱼丸汤：正宗的厦门鱼丸
- 土笋冻：特色小吃
- 花生汤：甜甜的

## Day 2：厦门大学 & 南普陀寺

### 厦门大学
- 中国最美大学之一
- 芙蓉隧道的涂鸦很有特色
- 白城沙滩看日落

### 南普陀寺
- 闽南佛教圣地
- 素斋很好吃
- 登山可以俯瞰厦门

## Day 3：环岛路 & 曾厝垵

### 环岛路
- 骑行环岛路，欣赏海景
- 沙滩上玩水
- 看日落

### 曾厝垵
- 文艺小渔村
- 各种特色小店
- 海鲜大排档

## 住宿

住在中山路附近，交通方便，去各个景点都很近。

## 旅行感悟

### 关于旅行
- 旅行是放松和充电
- 不用赶行程，随心所欲
- 享受当地的生活节奏

### 关于厦门
- 城市干净整洁
- 人民热情友好
- 美食很多
- 适合慢游

## 实用建议

1. 提前订好船票和住宿
2. 鼓浪屿要早点去，避开人流
3. 防晒一定要做好
4. 带好相机

## 结语

厦门之行非常愉快，这座城市的慢节奏让人放松。有机会还会再来！' WHERE id = 179;

-- 文章181: 新手如何搭建个人博客
UPDATE articles SET content = '## 为什么要搭建博客？

- 记录学习笔记
- 分享技术经验
- 建立个人品牌
- 锻炼写作能力

## 博客类型

### 托管平台
- 掘金
- CSDN
- 知乎

### 自建博客
- Hugo + GitHub Pages
- Hexo + Vercel
- WordPress
- 自己开发

## 方案对比

| 方案 | 难度 | 成本 | 自定义 |
|-----|------|------|--------|
| 掘金等 | 低 | 免费 | 低 |
| Hugo | 中 | 免费 | 高 |
| WordPress | 中 | 低 | 高 |
| 自己开发 | 高 | 中 | 最高 |

## 推荐方案：Hugo

### 优点
- 免费
- 快速
- 简单
- 部署方便

### 搭建步骤

```bash
# 安装 Hugo
brew install hugo  # macOS
choco install hugo  # Windows

# 创建博客
hugo new site my-blog
cd my-blog

# 安装主题
git init
git submodule add https://github.com/xxx/xxx.git themes/xxx

# 创建文章
hugo new content posts/my-first-post.md

# 本地预览
hugo server -D
```

### 部署到 GitHub Pages

1. 创建 GitHub 仓库
2. 推送代码
3. 开启 GitHub Pages

## 写作建议

1. 坚持更新
2. 写有价值的内容
3. 保持简洁
4. 配图说明

## 总结

搭建博客不难，难的是坚持。开始行动，记录你的学习和成长！' WHERE id = 181;
