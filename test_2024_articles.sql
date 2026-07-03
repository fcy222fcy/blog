USE blog;

-- 添加2024年的文章
INSERT INTO articles (title, slug, content, summary, category_id, view_count, comment_count, status, is_top, reading_time, created_at, updated_at) VALUES
('2024年年终总结：成长与收获', '2024-annual-summary-review', '2024年是充满挑战和收获的一年。完成了多个重要项目，技术能力显著提升。', '2024 年个人成长总结。', 3, 678, 22, 'published', false, 16, '2024-12-30 09:00:00', '2024-12-30 09:00:00'),
('Gin 框架实战：构建 RESTful API', 'gin-restful-api-practice', 'Gin是一个高性能的Go Web框架。核心功能包括路由管理、中间件、参数绑定。', '使用 Gin 框架构建 RESTful API 的完整指南。', 2, 534, 14, 'published', false, 13, '2024-11-25 15:30:00', '2024-11-25 15:30:00'),
('Git 工作流：团队协作的最佳实践', 'git-team-collaboration-workflow', '常见工作流包括Git Flow、GitHub Flow、GitLab Flow。提交规范使用feat、fix、docs等前缀。', 'Git 团队协作工作流和最佳实践指南。', 2, 423, 11, 'published', false, 10, '2024-10-18 11:45:00', '2024-10-18 11:45:00'),
('周末烘焙：第一次做蛋糕', 'first-time-baking-cake', '看到朋友圈的蛋糕照片，决定自己尝试。准备了鸡蛋、面粉、糖、牛奶等材料。', '第一次尝试烘焙蛋糕的经历和心得。', 3, 198, 8, 'published', false, 5, '2024-09-12 17:20:00', '2024-09-12 17:20:00'),
('MySQL 性能优化实战', 'mysql-performance-optimization-guide', '线上数据库查询变慢，需要优化。通过慢查询日志、EXPLAIN分析找出问题。', 'MySQL 数据库性能优化的实战经验和技巧。', 2, 467, 13, 'published', false, 15, '2024-08-20 14:10:00', '2024-08-20 14:10:00'),
('我的 2024 阅读清单', 'my-2024-reading-list-summary', '2024年计划阅读24本书，实际阅读了28本。推荐书籍包括《深入理解计算机系统》《设计模式》等。', '2024 年度阅读清单和读书心得分享。', 3, 312, 9, 'published', false, 8, '2024-07-15 20:00:00', '2024-07-15 20:00:00'),
('Docker 入门到实践', 'docker-practice-guide', 'Docker的核心概念包括镜像、容器、仓库。常用命令包括pull、run、ps、logs等。', 'Docker 从入门到实践的完整学习指南。', 2, 589, 16, 'published', false, 11, '2024-06-08 10:30:00', '2024-06-08 10:30:00'),
('旅行：厦门三日游', 'xiamen-trip-three-days', 'Day1游览鼓浪屿和南普陀寺，Day2参观厦门大学和曾厝垵，Day3环岛路骑行。', '厦门三日游的旅行攻略和游记。', 3, 267, 7, 'published', false, 9, '2024-05-02 16:45:00', '2024-05-02 16:45:00'),
('Go 语言错误处理最佳实践', 'go-error-handling-practice', 'Go语言没有异常机制，错误处理需要特别注意。常见方式包括直接返回、包装错误、自定义错误。', 'Go 语言错误处理的常见方式和最佳实践。', 2, 398, 10, 'published', false, 7, '2024-04-10 09:20:00', '2024-04-10 09:20:00'),
('新手如何搭建个人博客', 'build-personal-blog-guide', '博客可以记录学习笔记、分享技术经验、建立个人品牌。技术选型包括Hugo、Hexo、WordPress等。', '从零开始搭建个人博客的完整指南。', 1, 623, 18, 'published', true, 14, '2024-03-15 11:00:00', '2024-03-15 11:00:00');

-- 查看新插入文章的ID范围
SELECT id, title, created_at FROM articles WHERE created_at >= '2024-01-01' AND created_at < '2025-01-01' ORDER BY created_at DESC;
