-- 博客数据库 Schema
-- 基于 GORM AutoMigrate 生成

CREATE DATABASE IF NOT EXISTS blog DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE blog;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    nickname VARCHAR(50) NULL,
    email VARCHAR(100) UNIQUE,
    avatar VARCHAR(500) NULL,
    bio TEXT NULL,
    status TINYINT DEFAULT 1,
    INDEX idx_users_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 分类表
CREATE TABLE IF NOT EXISTS categories (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(200) NULL,
    icon VARCHAR(10) NULL,
    sort_order INT DEFAULT 0,
    INDEX idx_categories_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 标签表
CREATE TABLE IF NOT EXISTS tags (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL UNIQUE,
    INDEX idx_tags_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 文章表
CREATE TABLE IF NOT EXISTS articles (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    title VARCHAR(200) NOT NULL,
    slug VARCHAR(200) NOT NULL UNIQUE,
    content LONGTEXT NULL,
    summary TEXT NULL,
    cover VARCHAR(500) NULL,
    category_id BIGINT UNSIGNED,
    view_count BIGINT DEFAULT 0,
    comment_count BIGINT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'published',
    is_top BOOLEAN DEFAULT FALSE,
    reading_time INT DEFAULT 0,
    INDEX idx_articles_deleted_at (deleted_at),
    INDEX idx_articles_category_id (category_id),
    CONSTRAINT fk_articles_category FOREIGN KEY (category_id) REFERENCES categories(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 文章标签关联表
CREATE TABLE IF NOT EXISTS article_tags (
    article_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (article_id, tag_id),
    CONSTRAINT fk_article_tags_article FOREIGN KEY (article_id) REFERENCES articles(id),
    CONSTRAINT fk_article_tags_tag FOREIGN KEY (tag_id) REFERENCES tags(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 评论表
CREATE TABLE IF NOT EXISTS comments (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    content TEXT NOT NULL,
    nickname VARCHAR(50) NULL,
    email VARCHAR(100) NULL,
    website VARCHAR(200) NULL,
    avatar VARCHAR(500) NULL,
    article_id BIGINT UNSIGNED,
    parent_id BIGINT UNSIGNED NULL,
    status VARCHAR(20) DEFAULT 'pending',
    ip VARCHAR(50) NULL,
    user_agent VARCHAR(500) NULL,
    INDEX idx_comments_deleted_at (deleted_at),
    INDEX idx_comments_article_id (article_id),
    INDEX idx_comments_parent_id (parent_id),
    CONSTRAINT fk_comments_article FOREIGN KEY (article_id) REFERENCES articles(id),
    CONSTRAINT fk_comments_parent FOREIGN KEY (parent_id) REFERENCES comments(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 每日一问表
CREATE TABLE IF NOT EXISTS daily_questions (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    question TEXT NOT NULL,
    answer TEXT NULL,
    date VARCHAR(10) NOT NULL UNIQUE,
    like_count BIGINT DEFAULT 0,
    comment_count BIGINT DEFAULT 0,
    view_count BIGINT DEFAULT 0,
    status TINYINT DEFAULT 1,
    INDEX idx_daily_questions_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 媒体文件表
CREATE TABLE IF NOT EXISTS media (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    filename VARCHAR(255) NOT NULL,
    url VARCHAR(500) NOT NULL,
    size BIGINT NOT NULL,
    mime_type VARCHAR(100) NULL,
    type VARCHAR(20) NULL,
    INDEX idx_media_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 关于页面表
CREATE TABLE IF NOT EXISTS about_pages (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    title VARCHAR(200) NULL,
    subtitle VARCHAR(500) NULL,
    bio TEXT NULL,
    skills TEXT NULL,
    about_me TEXT NULL,
    about_site TEXT NULL,
    projects TEXT NULL,
    contact_info TEXT NULL,
    INDEX idx_about_pages_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
