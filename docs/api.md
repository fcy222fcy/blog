# 博客系统 API 接口文档

## 通用说明

### 基础信息

| 项目 | 说明 |
|------|------|
| 基础 URL | `http://localhost:8080/api/v1` |
| 数据格式 | JSON |
| 字符编码 | UTF-8 |
| 认证方式 | JWT Token（管理员接口需要） |

### 统一响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 错误码说明

| 错误码 | HTTP 状态码 | 说明 |
|--------|-------------|------|
| 200 | 200 | 成功 |
| 400 | 400 | 请求参数错误 |
| 401 | 401 | 未授权（未登录） |
| 403 | 403 | 禁止访问（权限不足） |
| 404 | 404 | 资源不存在 |
| 500 | 500 | 服务器内部错误 |

### 分页参数

| 参数名 | 类型 | 必填 | 默认值 | 说明 |
|--------|------|------|--------|------|
| page | int | 否 | 1 | 当前页码 |
| page_size | int | 否 | 10 | 每页数量（最大 50） |

### 分页响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

---

## 一、认证模块

### 1.1 管理员登录

**请求**

```
POST /auth/login
```

**请求体**

```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": 1719000000
  }
}
```

### 1.2 获取当前用户信息

**请求**

```
GET /auth/profile
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "nickname": "Liu Houliang",
    "avatar": "https://example.com/avatar.jpg",
    "email": "admin@example.com",
    "description": "Go 开发者",
    "social_links": [
      {"name": "GitHub", "url": "https://github.com/xxx"},
      {"name": "Twitter", "url": "https://twitter.com/xxx"}
    ]
  }
}
```

### 1.3 修改密码

**请求**

```
PUT /auth/password
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "old_password": "123456",
  "new_password": "new_password"
}
```

**响应**

```json
{
  "code": 200,
  "message": "密码修改成功",
  "data": null
}
```

---

## 二、文章模块

### 2.1 获取文章列表（前台）

**请求**

```
GET /articles
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| category | string | 否 | 分类 slug |
| tag | string | 否 | 标签 slug |
| keyword | string | 否 | 搜索关键词 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "我的博客主题已开源",
        "slug": "open-source-theme",
        "summary": "基于 Hugo Theme Stack 打造的开箱即用博客模板...",
        "cover": "https://example.com/cover.jpg",
        "category": {
          "id": 1,
          "name": "搭建网站",
          "slug": "build"
        },
        "tags": [
          {"id": 1, "name": "Hugo", "slug": "hugo"},
          {"id": 2, "name": "开源", "slug": "open-source"}
        ],
        "view_count": 252,
        "like_count": 42,
        "comment_count": 8,
        "created_at": "2026-04-24T10:00:00Z",
        "updated_at": "2026-04-24T10:00:00Z"
      }
    ],
    "total": 6,
    "page": 1,
    "page_size": 10
  }
}
```

### 2.2 获取文章详情（前台）

**请求**

```
GET /articles/:slug
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "title": "我的博客主题已开源",
    "slug": "open-source-theme",
    "content": "# 标题\n\n这里是文章正文...",
    "summary": "基于 Hugo Theme Stack 打造的开箱即用博客模板...",
    "cover": "https://example.com/cover.jpg",
    "category": {
      "id": 1,
      "name": "搭建网站",
      "slug": "build"
    },
    "tags": [
      {"id": 1, "name": "Hugo", "slug": "hugo"},
      {"id": 2, "name": "开源", "slug": "open-source"}
    ],
    "view_count": 252,
    "like_count": 42,
    "comment_count": 8,
    "reading_time": 3,
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 2.3 文章点赞

**请求**

```
POST /articles/:id/like
```

**响应**

```json
{
  "code": 200,
  "message": "点赞成功",
  "data": {
    "like_count": 43
  }
}
```

### 2.4 获取文章归档（前台）

**请求**

```
GET /articles/archives
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "year": 2026,
      "articles": [
        {
          "id": 1,
          "title": "我的博客主题已开源",
          "slug": "open-source-theme",
          "created_at": "2026-04-24T10:00:00Z"
        }
      ]
    }
  ]
}
```

---

## 三、后台文章管理

### 3.1 获取文章列表（后台）

**请求**

```
GET /admin/articles
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| category_id | int | 否 | 分类 ID |
| status | string | 否 | 状态：published/draft |
| keyword | string | 否 | 搜索关键词 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "我的博客主题已开源",
        "slug": "open-source-theme",
        "status": "published",
        "category": {
          "id": 1,
          "name": "搭建网站"
        },
        "tags": [
          {"id": 1, "name": "Hugo"},
          {"id": 2, "name": "开源"}
        ],
        "view_count": 252,
        "like_count": 42,
        "comment_count": 8,
        "created_at": "2026-04-24T10:00:00Z",
        "updated_at": "2026-04-24T10:00:00Z"
      }
    ],
    "total": 6,
    "page": 1,
    "page_size": 10
  }
}
```

### 3.2 获取文章详情（后台）

**请求**

```
GET /admin/articles/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "title": "我的博客主题已开源",
    "slug": "open-source-theme",
    "content": "# 标题\n\n这里是文章正文...",
    "summary": "基于 Hugo Theme Stack 打造的开箱即用博客模板...",
    "cover": "https://example.com/cover.jpg",
    "category_id": 1,
    "tag_ids": [1, 2],
    "status": "published",
    "is_top": false,
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 3.3 创建文章

**请求**

```
POST /admin/articles
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "title": "新文章标题",
  "content": "# 标题\n\n文章正文内容...",
  "summary": "文章摘要（可选，留空自动截取）",
  "cover": "https://example.com/cover.jpg",
  "category_id": 1,
  "tag_ids": [1, 2],
  "status": "draft",
  "is_top": false
}
```

**响应**

```json
{
  "code": 200,
  "message": "文章创建成功",
  "data": {
    "id": 7
  }
}
```

### 3.4 更新文章

**请求**

```
PUT /admin/articles/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "title": "更新后的标题",
  "content": "更新后的内容...",
  "summary": "更新后的摘要",
  "cover": "https://example.com/new-cover.jpg",
  "category_id": 1,
  "tag_ids": [1, 2, 3],
  "status": "published",
  "is_top": true
}
```

**响应**

```json
{
  "code": 200,
  "message": "文章更新成功",
  "data": null
}
```

### 3.5 删除文章

**请求**

```
DELETE /admin/articles/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "文章删除成功",
  "data": null
}
```

### 3.6 批量删除文章

**请求**

```
POST /admin/articles/batch-delete
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "ids": [1, 2, 3]
}
```

**响应**

```json
{
  "code": 200,
  "message": "批量删除成功",
  "data": null
}
```

---

## 四、分类模块

### 4.1 获取分类列表（前台）

**请求**

```
GET /categories
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "搭建网站",
      "slug": "build",
      "description": "网站搭建、博客部署相关",
      "article_count": 3,
      "sort_order": 1
    }
  ]
}
```

### 4.2 创建分类（后台）

**请求**

```
POST /admin/categories
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "新分类",
  "slug": "new-category",
  "description": "分类描述",
  "icon": "📁",
  "sort_order": 0
}
```

**响应**

```json
{
  "code": 200,
  "message": "分类创建成功",
  "data": {
    "id": 4
  }
}
```

### 4.3 更新分类（后台）

**请求**

```
PUT /admin/categories/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "更新后的分类名",
  "slug": "updated-slug",
  "description": "更新后的描述",
  "icon": "📂",
  "sort_order": 1
}
```

**响应**

```json
{
  "code": 200,
  "message": "分类更新成功",
  "data": null
}
```

### 4.4 删除分类（后台）

**请求**

```
DELETE /admin/categories/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "分类删除成功",
  "data": null
}
```

---

## 五、标签模块

### 5.1 获取标签列表（前台）

**请求**

```
GET /tags
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "Hugo",
      "slug": "hugo",
      "article_count": 3
    }
  ]
}
```

### 5.2 创建标签（后台）

**请求**

```
POST /admin/tags
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "新标签",
  "slug": "new-tag"
}
```

**响应**

```json
{
  "code": 200,
  "message": "标签创建成功",
  "data": {
    "id": 13
  }
}
```

### 5.3 更新标签（后台）

**请求**

```
PUT /admin/tags/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "更新后的标签名",
  "slug": "updated-tag"
}
```

**响应**

```json
{
  "code": 200,
  "message": "标签更新成功",
  "data": null
}
```

### 5.4 删除标签（后台）

**请求**

```
DELETE /admin/tags/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "标签删除成功",
  "data": null
}
```

---

## 六、评论模块

### 6.1 获取文章评论列表（前台）

**请求**

```
GET /articles/:article_id/comments
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "content": "很棒的文章，学到了很多！",
        "nickname": "张三",
        "email": "zhangsan@example.com",
        "website": "https://zhangsan.com",
        "avatar": "https://www.gravatar.com/avatar/xxx",
        "status": "approved",
        "created_at": "2026-04-24T12:00:00Z",
        "replies": [
          {
            "id": 2,
            "content": "谢谢支持！",
            "nickname": "博主",
            "created_at": "2026-04-24T13:00:00Z"
          }
        ]
      }
    ],
    "total": 8,
    "page": 1,
    "page_size": 10
  }
}
```

### 6.2 提交评论（前台）

**请求**

```
POST /articles/:article_id/comments
```

**请求体**

```json
{
  "content": "评论内容",
  "nickname": "张三",
  "email": "zhangsan@example.com",
  "website": "https://zhangsan.com",
  "parent_id": 0
}
```

**响应**

```json
{
  "code": 200,
  "message": "评论提交成功，等待审核",
  "data": {
    "id": 3
  }
}
```

### 6.3 获取评论列表（后台）

**请求**

```
GET /admin/comments
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| status | string | 否 | 状态：pending/approved/rejected |
| article_id | int | 否 | 文章 ID |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "content": "很棒的文章，学到了很多！",
        "nickname": "张三",
        "email": "zhangsan@example.com",
        "status": "pending",
        "article": {
          "id": 1,
          "title": "我的博客主题已开源"
        },
        "created_at": "2026-04-24T12:00:00Z"
      }
    ],
    "total": 12,
    "page": 1,
    "page_size": 10
  }
}
```

### 6.4 审核评论（后台）

**请求**

```
PUT /admin/comments/:id/status
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "status": "approved"
}
```

**响应**

```json
{
  "code": 200,
  "message": "评论状态更新成功",
  "data": null
}
```

### 6.5 删除评论（后台）

**请求**

```
DELETE /admin/comments/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "评论删除成功",
  "data": null
}
```

---

## 七、友链模块

### 7.1 获取友链列表（前台）

**请求**

```
GET /links
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "吴同学的笔记本",
      "url": "https://www.wutongqi.cn",
      "description": "全栈开发者，分享技术与生活",
      "avatar": "https://example.com/avatar.jpg",
      "logo": "🤖",
      "sort_order": 1,
      "status": "approved"
    }
  ]
}
```

### 7.2 申请友链（前台）

**请求**

```
POST /links/apply
```

**请求体**

```json
{
  "name": "我的博客",
  "url": "https://myblog.com",
  "description": "一个技术博客",
  "avatar": "https://myblog.com/avatar.jpg",
  "email": "me@myblog.com"
}
```

**响应**

```json
{
  "code": 200,
  "message": "申请已提交，等待审核",
  "data": null
}
```

### 7.3 获取友链列表（后台）

**请求**

```
GET /admin/links
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| status | string | 否 | 状态：pending/approved/rejected |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "吴同学的笔记本",
        "url": "https://www.wutongqi.cn",
        "description": "全栈开发者，分享技术与生活",
        "avatar": "https://example.com/avatar.jpg",
        "status": "approved",
        "created_at": "2026-01-01T00:00:00Z"
      }
    ],
    "total": 4,
    "page": 1,
    "page_size": 10
  }
}
```

### 7.4 创建友链（后台）

**请求**

```
POST /admin/links
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "新友链",
  "url": "https://newblog.com",
  "description": "友链描述",
  "avatar": "https://newblog.com/avatar.jpg",
  "logo": "🔗",
  "sort_order": 0,
  "status": "approved"
}
```

**响应**

```json
{
  "code": 200,
  "message": "友链创建成功",
  "data": {
    "id": 5
  }
}
```

### 7.5 更新友链（后台）

**请求**

```
PUT /admin/links/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "name": "更新后的名称",
  "url": "https://updated-url.com",
  "description": "更新后的描述",
  "avatar": "https://updated-avatar.jpg",
  "logo": "🔗",
  "sort_order": 1,
  "status": "approved"
}
```

**响应**

```json
{
  "code": 200,
  "message": "友链更新成功",
  "data": null
}
```

### 7.6 删除友链（后台）

**请求**

```
DELETE /admin/links/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "友链删除成功",
  "data": null
}
```

### 7.7 审核友链（后台）

**请求**

```
PUT /admin/links/:id/status
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "status": "approved"
}
```

**响应**

```json
{
  "code": 200,
  "message": "友链状态更新成功",
  "data": null
}
```

---

## 八、每日一问模块

### 8.1 获取今日问题（前台）

**请求**

```
GET /daily-questions/today
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "question": "什么是最好的编程语言？",
    "answer": "没有最好的编程语言，只有最适合的...",
    "date": "2026-06-15",
    "like_count": 42,
    "comment_count": 8
  }
}
```

### 8.2 获取指定日期问题（前台）

**请求**

```
GET /daily-questions/:date
```

**路径参数**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| date | string | 日期，格式：2026-06-15 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "question": "什么是最好的编程语言？",
    "answer": "没有最好的编程语言，只有最适合的...",
    "date": "2026-06-15",
    "like_count": 42,
    "comment_count": 8
  }
}
```

### 8.3 每日一问点赞（前台）

**请求**

```
POST /daily-questions/:id/like
```

**响应**

```json
{
  "code": 200,
  "message": "点赞成功",
  "data": {
    "like_count": 43
  }
}
```

### 8.4 获取问题列表（后台）

**请求**

```
GET /admin/daily-questions
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| status | int | 否 | 状态：0-禁用 1-启用 |
| keyword | string | 否 | 搜索关键词 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "question": "什么是最好的编程语言？",
        "answer": "没有最好的编程语言，只有最适合的...",
        "date": "2026-06-15",
        "status": 1,
        "view_count": 100,
        "like_count": 42,
        "comment_count": 8,
        "created_at": "2026-06-15T00:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "page_size": 10
  }
}
```

### 8.5 创建问题（后台）

**请求**

```
POST /admin/daily-questions
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "question": "新问题是什么？",
  "answer": "这是问题的答案...",
  "date": "2026-06-20",
  "status": 1
}
```

**响应**

```json
{
  "code": 200,
  "message": "问题创建成功",
  "data": {
    "id": 11
  }
}
```

### 8.6 更新问题（后台）

**请求**

```
PUT /admin/daily-questions/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "question": "更新后的问题",
  "answer": "更新后的答案...",
  "date": "2026-06-20",
  "status": 1
}
```

**响应**

```json
{
  "code": 200,
  "message": "问题更新成功",
  "data": null
}
```

### 8.7 删除问题（后台）

**请求**

```
DELETE /admin/daily-questions/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "问题删除成功",
  "data": null
}
```

---

## 九、媒体库模块

### 9.1 上传文件（后台）

**请求**

```
POST /media/upload
```

**请求头**

```
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| file | file | 是 | 文件（支持 JPG、PNG、GIF、WebP、PDF） |

**响应**

```json
{
  "code": 200,
  "message": "上传成功",
  "data": {
    "id": 1,
    "filename": "image.jpg",
    "url": "https://example.com/uploads/2026/06/image.jpg",
    "size": 102400,
    "mime_type": "image/jpeg",
    "created_at": "2026-06-15T10:00:00Z"
  }
}
```

### 9.2 获取媒体列表（后台）

**请求**

```
GET /media
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| type | string | 否 | 类型：image/document |
| keyword | string | 否 | 搜索关键词 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "filename": "image.jpg",
        "url": "https://example.com/uploads/2026/06/image.jpg",
        "size": 102400,
        "mime_type": "image/jpeg",
        "created_at": "2026-06-15T10:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "page_size": 20
  }
}
```

### 9.3 删除媒体文件（后台）

**请求**

```
DELETE /media/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "文件删除成功",
  "data": null
}
```

---

## 十、娱乐模块（影视/游戏）

### 10.1 获取娱乐列表（前台）

**请求**

```
GET /entertainment
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| year | int | 否 | 年份 |
| type | string | 否 | 类型：movie/tv/game |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "movies": [
      {
        "id": 1,
        "title": "挽救计划",
        "title_en": "Project Hail Mary",
        "type": "movie",
        "year": 2026,
        "cover": "https://example.com/cover.jpg",
        "rating": 8.8,
        "rating_external": 8.2,
        "comment": "男主演技很棒，非常好的科幻电影",
        "status": "watched",
        "link": "https://www.themoviedb.org/movie/687163"
      }
    ],
    "tv": [],
    "games": [
      {
        "id": 1,
        "title": "Slay the Spire 2",
        "type": "game",
        "year": 2026,
        "cover": "https://example.com/cover.jpg",
        "rating": 9.6,
        "rating_external": 9.6,
        "platform": "PC",
        "playtime": "35H",
        "comment": "看起来跟一代差不多，玩起来还是沉迷",
        "status": "completed"
      }
    ]
  }
}
```

### 10.2 获取娱乐列表（后台）

**请求**

```
GET /admin/entertainment
```

**请求头**

```
Authorization: Bearer <token>
```

**查询参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页数量 |
| type | string | 否 | 类型：movie/tv/game |
| year | int | 否 | 年份 |
| status | string | 否 | 状态 |

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "挽救计划",
        "title_en": "Project Hail Mary",
        "type": "movie",
        "year": 2026,
        "cover": "https://example.com/cover.jpg",
        "rating": 8.8,
        "rating_external": 8.2,
        "comment": "男主演技很棒，非常好的科幻电影",
        "status": "watched",
        "link": "https://www.themoviedb.org/movie/687163",
        "created_at": "2026-06-15T00:00:00Z"
      }
    ],
    "total": 12,
    "page": 1,
    "page_size": 10
  }
}
```

### 10.3 创建娱乐条目（后台）

**请求**

```
POST /admin/entertainment
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "title": "新电影",
  "title_en": "New Movie",
  "type": "movie",
  "year": 2026,
  "cover": "https://example.com/cover.jpg",
  "rating": 8.5,
  "rating_external": 8.0,
  "comment": "电影评论",
  "status": "watched",
  "link": "https://www.themoviedb.org/movie/xxx"
}
```

**响应**

```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "id": 13
  }
}
```

### 10.4 更新娱乐条目（后台）

**请求**

```
PUT /admin/entertainment/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "title": "更新后的标题",
  "rating": 9.0,
  "status": "completed"
}
```

**响应**

```json
{
  "code": 200,
  "message": "更新成功",
  "data": null
}
```

### 10.5 删除娱乐条目（后台）

**请求**

```
DELETE /admin/entertainment/:id
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

---

## 十一、系统设置模块

### 11.1 获取系统设置（后台）

**请求**

```
GET /admin/settings
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "site_name": "Liu Houliang - 个人博客",
    "site_description": "日常落灰的个人博客，分享 Golang、AI 和 NAS 折腾经验",
    "site_url": "https://liuhouliang.com",
    "site_keywords": "博客,Golang,AI,NAS",
    "seo_title": "Liu Houliang - Go 开发者",
    "seo_description": "分享 Golang 开发、AI 和 NAS 折腾经验",
    "seo_keywords": "Golang,AI,NAS,编程,技术",
    "page_size": 10,
    "favicon": "https://example.com/favicon.ico",
    "logo": "https://example.com/logo.png"
  }
}
```

### 11.2 更新系统设置（后台）

**请求**

```
PUT /admin/settings
```

**请求头**

```
Authorization: Bearer <token>
```

**请求体**

```json
{
  "site_name": "新的博客名称",
  "site_description": "新的博客描述",
  "site_url": "https://new-url.com",
  "site_keywords": "新关键词",
  "seo_title": "新的 SEO 标题",
  "seo_description": "新的 SEO 描述",
  "seo_keywords": "新 SEO 关键词",
  "page_size": 15
}
```

**响应**

```json
{
  "code": 200,
  "message": "设置保存成功",
  "data": null
}
```

---

## 十二、仪表盘模块

### 12.1 获取仪表盘统计（后台）

**请求**

```
GET /admin/dashboard/stats
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "article_count": 6,
    "published_count": 5,
    "draft_count": 1,
    "total_views": 1050,
    "total_likes": 1200,
    "comment_count": 12,
    "pending_comment_count": 3,
    "link_count": 4,
    "category_count": 3,
    "tag_count": 12,
    "daily_question_count": 30
  }
}
```

### 12.2 获取最近文章（后台）

**请求**

```
GET /admin/dashboard/recent-articles
```

**请求头**

```
Authorization: Bearer <token>
```

**响应**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "我的博客主题已开源",
      "view_count": 252,
      "created_at": "2026-04-24T10:00:00Z"
    }
  ]
}
```

---

## 附录

### 1. 状态码枚举

**文章状态**
- `published` - 已发布
- `draft` - 草稿

**评论状态**
- `pending` - 待审核
- `approved` - 已通过
- `rejected` - 已拒绝

**友链状态**
- `pending` - 待审核
- `approved` - 已通过
- `rejected` - 已拒绝

**每日一问状态**
- `0` - 禁用
- `1` - 启用

**娱乐状态**
- `watching` - 在看
- `watched` - 已看
- `playing` - 在玩
- `completed` - 已完成

### 2. 数据类型说明

**娱乐类型**
- `movie` - 电影
- `tv` - 剧集
- `game` - 游戏

### 3. 媒体类型

- `image` - 图片
- `document` - 文档

### 4. 分页默认值

- 默认页码：1
- 默认每页数量：10
- 最大每页数量：50
