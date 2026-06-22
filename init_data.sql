USE blog;

-- 分类
INSERT INTO categories (name, slug, description, icon, sort_order, created_at, updated_at) VALUES
('搭建网站', 'build', '网站搭建、博客部署相关', '🌐', 1, NOW(), NOW()),
('软件开发', 'dev', '编程、开发工具、技术分享', '💻', 2, NOW(), NOW()),
('生活记录', 'life', '日常生活、随笔感悟', '🌱', 3, NOW(), NOW());

-- 标签
INSERT INTO tags (name, slug, created_at, updated_at) VALUES
('Hugo', 'hugo', NOW(), NOW()),
('Go', 'go', NOW(), NOW()),
('AI', 'ai', NOW(), NOW()),
('开源', 'open-source', NOW(), NOW()),
('Cloudflare', 'cloudflare', NOW(), NOW());

-- 文章
INSERT INTO articles (title, slug, content, summary, category_id, view_count, like_count, comment_count, status, is_top, reading_time, created_at, updated_at) VALUES
('我的博客主题已开源，欢迎使用', 'my-blog-theme-open-source', '## 引言\n\n基于 Hugo Theme Stack 打造的开箱即用博客模板，包含多项美化和功能增强。\n\n## 特性\n\n- 响应式设计\n- 深色/浅色主题\n- 代码高亮\n- 搜索功能\n\n如果你喜欢我的博客样式，不妨试试它！', '基于 Hugo Theme Stack 打造的开箱即用博客模板，包含多项美化和功能增强。', 1, 252, 12, 3, 'published', true, 3, NOW(), NOW()),
('DesktopSnap：开源轻量桌面图标恢复工具', 'desktopsnap', '## 项目介绍\n\nDesktopSnap 是一款基于 WinUI3 开发的轻量级桌面工具，支持多显示器布局一键备份与恢复。\n\n## 功能特点\n\n- 多显示器支持\n- 一键备份桌面图标布局\n- 已上架微软商店', '基于 WinUI3 开发的轻量级桌面工具。支持多显示器布局一键备份与恢复，已上架微软商店。', 2, 327, 8, 2, 'published', false, 3, NOW(), NOW()),
('用 Cloudflare Workers 免费给博客增加 AI 助手', 'cloudflare-ai-assistant', '## 背景\n\n利用 Cloudflare Workers 和 AI 能力，为博客添加智能问答功能。\n\n## 实现方案\n\n1. 创建 Cloudflare Worker\n2. 配置 AI 模型\n3. 接入博客\n\n完全免费，快来试试吧！', '利用 Cloudflare Workers 和 AI 能力，为博客添加智能问答功能。', 1, 160, 5, 1, 'published', false, 6, NOW(), NOW()),
('清明假期：收拾家务、整理心情', 'qingming-holiday', '## 假期日常\n\n清明假期三天，主要在家收拾家务。\n\n## 收拾成果\n\n- 整理了书房\n- 清理了衣柜\n- 断舍离了一些物品\n\n整理物品的过程也是整理心情的过程。', '清明假期的日常记录，整理家务，调整心态。', 3, 100, 3, 0, 'published', false, 9, NOW(), NOW()),
('逐渐难以逃离对于 AI 的焦虑', 'ai-anxiety', '## 引言\n\n最近看着各种 AI 工具的发展，感到越来越焦虑。\n\n## 焦虑来源\n\n1. 技术更新太快\n2. 担心被替代\n3. 学习跟不上\n\n与其焦虑，不如行动。', '养龙虾热潮让人们看到了 AI 的威力，也让人们更加因为没有用上 AI 技术而感到焦虑。', 3, 114, 6, 1, 'published', false, 10, NOW(), NOW()),
('Go 组合与方法覆盖的坑', 'go-interface-pitfall', '## 背景\n\n在使用 Go 语言开发时，遇到了一个关于接口组合和方法覆盖的问题。\n\n## 问题描述\n\n当一个结构体嵌入另一个实现了接口的结构体时，方法覆盖的行为可能不符合预期。\n\n## 解决方案\n\n需要显式定义方法来覆盖嵌入结构体的方法。', 'Go 语言开发中接口组合与方法覆盖的常见问题和解决方案。', 2, 89, 4, 0, 'draft', false, 5, NOW(), NOW());

-- 文章标签关联
INSERT INTO article_tags (article_id, tag_id) VALUES
(1, 1), (1, 4),
(2, 4),
(3, 3), (3, 5),
(5, 3),
(6, 2);

-- 友链
INSERT INTO links (name, url, description, avatar, logo, sort_order, status, created_at, updated_at) VALUES
('Jimmy Cai', 'https://jimmycai.com/', 'Author of Hugo Stack theme', 'https://avatars.githubusercontent.com/u/5889006?v=4', 'JC', 1, 'approved', NOW(), NOW()),
('Nan0inPsyLog', 'https://nan0in27.cn/', 'We love CS, thus we can pwn the world', 'https://nan0in27.cn/img/avatar_hu_85437ba0800ccf1d.png', 'N', 2, 'approved', NOW(), NOW()),
('Nanhu Shijie', 'https://jolla.pp.ua/', '一个生活爱好者!', '/avatar.jpg', 'N', 3, 'approved', NOW(), NOW()),
('Liu Zijian Blog', 'https://blog.liuzijian.com/', 'Keep on going never give up', '/avatar.jpg', 'L', 4, 'approved', NOW(), NOW());

-- 每日一问
INSERT INTO daily_questions (question, answer, date, like_count, comment_count, view_count, status, created_at, updated_at) VALUES
('什么是最好的编程语言？', '没有最好的编程语言，只有最适合的。Python 适合数据科学，JavaScript 适合 Web 开发，Go 适合后端服务。', '2026-06-20', 42, 8, 156, 1, NOW(), NOW()),
('如何成为一名优秀的软件工程师？', '成为优秀软件工程师的关键：1. 扎实的基础知识；2. 持续学习新技术；3. 多写代码多实践；4. 阅读优秀源码；5. 参与开源项目。', '2026-06-19', 128, 15, 234, 1, NOW(), NOW()),
('远程办公的优缺点是什么？', '优点：灵活的工作时间、节省通勤时间。缺点：缺乏面对面交流、容易分心。', '2026-06-18', 89, 12, 189, 1, NOW(), NOW()),
('如何看待人工智能的发展？', '人工智能正在快速发展，我们应该拥抱技术变革，学习相关知识，思考如何利用 AI 提高效率。', '2026-06-17', 234, 28, 456, 1, NOW(), NOW());

-- 用户
INSERT INTO users (username, password, nickname, email, avatar, bio, status, created_at, updated_at) VALUES
('admin', '123456', 'Liu Houliang', 'admin@liuhouliang.com', '/avatar.jpg', '日常落灰的个人博客，分享 Golang、AI 和 NAS 折腾经验', 1, NOW(), NOW());
