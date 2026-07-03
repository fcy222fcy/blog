-- MySQL 初始化脚本
-- 设置字符集和排序规则
ALTER DATABASE IF EXISTS blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 注意：表结构由后端 GORM AutoMigrate 自动创建
-- 此脚本仅用于确保数据库字符集配置正确
