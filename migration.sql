-- ============================================
-- 博客系统数据库迁移脚本
-- 用于：GORM AutoMigrate 之外的手动 DDL
-- 注意：如果使用 GORM AutoMigrate，以下 ALTER/CREATE 会自动执行
-- ============================================

-- 1. 文章表新增字段：定时发布、SEO
ALTER TABLE `articles`
  ADD COLUMN `scheduled_at` DATETIME NULL COMMENT '定时发布时间' AFTER `reading_time`,
  ADD COLUMN `seo_title` VARCHAR(200) DEFAULT '' COMMENT 'SEO 标题' AFTER `scheduled_at`,
  ADD COLUMN `seo_description` VARCHAR(500) DEFAULT '' COMMENT 'SEO 描述' AFTER `seo_title`,
  ADD COLUMN `seo_keywords` VARCHAR(300) DEFAULT '' COMMENT 'SEO 关键词' AFTER `seo_description`;

-- 2. 创建审计日志表
CREATE TABLE IF NOT EXISTS `audit_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `operator_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '操作人ID',
  `operator_name` VARCHAR(50) DEFAULT '' COMMENT '操作人名称',
  `action` VARCHAR(50) NOT NULL COMMENT '操作类型：create/update/delete/approve/reject',
  `target_type` VARCHAR(50) DEFAULT '' COMMENT '目标类型：article/comment/category/tag/link',
  `target_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '目标ID',
  `target_title` VARCHAR(200) DEFAULT '' COMMENT '目标标题',
  `detail` TEXT COMMENT '操作详情（JSON 格式）',
  `ip` VARCHAR(50) DEFAULT '' COMMENT '操作IP',
  `user_agent` VARCHAR(500) DEFAULT '' COMMENT 'User-Agent',
  `created_at` DATETIME(3) NULL DEFAULT NULL,
  `updated_at` DATETIME(3) NULL DEFAULT NULL,
  `deleted_at` DATETIME(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_audit_logs_action` (`action`),
  INDEX `idx_audit_logs_target_type` (`target_type`),
  INDEX `idx_audit_logs_target_id` (`target_id`),
  INDEX `idx_audit_logs_operator_id` (`operator_id`),
  INDEX `idx_audit_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作审计日志表';
