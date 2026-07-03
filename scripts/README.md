# SQL Scripts

本目录包含博客系统的数据库脚本文件。

## 文件说明

| 文件名 | 说明 |
|--------|------|
| `schema.sql` | 数据库表结构定义，包含所有表的创建语句 |
| `init_data.sql` | 初始数据，包含分类、标签、文章、友链等基础数据 |
| `test_articles.sql` | 测试文章数据（2024-2025年） |
| `test_2024_articles.sql` | 2024年文章测试数据 |

## 使用方法

### 1. 初始化数据库

```bash
# 创建数据库和表结构
mysql -u root -p < scripts/schema.sql

# 导入初始数据
mysql -u root -p blog < scripts/init_data.sql
```

### 2. 导入测试数据

```bash
# 导入测试文章
mysql -u root -p blog < scripts/test_articles.sql

# 或导入2024年文章
mysql -u root -p blog < scripts/test_2024_articles.sql
```

### 3. 使用 Go 项目自动迁移

项目使用 GORM AutoMigrate，启动时会自动创建表结构：

```bash
go run main.go
```

## 数据库表结构

- `users` - 用户表
- `categories` - 分类表
- `tags` - 标签表
- `articles` - 文章表
- `article_tags` - 文章标签关联表
- `comments` - 评论表
- `links` - 友链表
- `daily_questions` - 每日一问表
- `media` - 媒体文件表
- `about_pages` - 关于页面表

## 注意事项

1. 所有表都包含 `created_at`、`updated_at`、`deleted_at` 字段（软删除）
2. 字符集使用 `utf8mb4`，支持 emoji 等特殊字符
3. 外键约束确保数据完整性
