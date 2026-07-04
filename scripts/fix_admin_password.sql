-- 修复管理员密码 (密码: 123456)
-- 运行此脚本将 admin 密码更新为 bcrypt 加密格式

USE blog;

UPDATE users
SET password = '$2a$10$rr5uvsYJFrmZkGsSwDrG9.Q0aFQvu0uikcHU/0vefGiucCWtZGXYC'
WHERE username = 'admin';

-- 验证更新
SELECT id, username, nickname, email FROM users WHERE username = 'admin';
