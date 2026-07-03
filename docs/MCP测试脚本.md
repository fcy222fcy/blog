# 博客系统 MCP Playwright 测试脚本

> **工具**: Playwright MCP  
> **版本**: v1.0  
> **创建时间**: 2026年7月3日  
> **测试目标**: 全面覆盖前后端功能

---

## 🚀 快速开始

### 1. 启动服务

```bash
# 启动后端
cd e:\goCode\goFile\gin博客
.\blog.exe

# 启动前台 (新终端)
cd blog-web
npm run dev

# 启动后台 (新终端)
cd blog-admin
npm run dev
```

### 2. 访问地址

- **前台**: http://localhost:5173
- **后台**: http://localhost:5174
- **API**: http://localhost:8080/api/v1

---

## 📦 测试场景分类

### A. 前台测试 (blog-web)
### B. 后台测试 (blog-admin)
### C. API测试
### D. 端到端测试

---

## A. 前台测试场景

### A1. 首页加载测试

**测试目标**: 验证首页能正常加载文章列表

**测试步骤**:
```
1. 打开 http://localhost:5173
2. 等待页面加载完成
3. 检查文章列表是否显示
4. 检查分页组件是否存在
```

**预期结果**:
- ✅ 页面标题包含"博客"或"Blog"
- ✅ 可以看到至少1篇文章卡片
- ✅ 文章卡片包含标题、摘要、日期
- ✅ 页面底部有分页按钮（如果文章数>10）

**Playwright测试命令**:
```javascript
// 导航到首页
await page.goto('http://localhost:5173');

// 检查页面标题
await expect(page).toHaveTitle(/博客|Blog/i);

// 检查文章列表
const articles = await page.locator('.article-card, .article-item, article').count();
expect(articles).toBeGreaterThan(0);

// 截图保存
await page.screenshot({ path: 'test-results/homepage.png', fullPage: true });
```

---

### A2. 文章详情页测试

**测试目标**: 验证点击文章后能正常查看详情

**测试步骤**:
```
1. 在首页点击第一篇文章
2. 等待文章详情页加载
3. 检查文章内容是否完整显示
4. 检查评论区是否存在
```

**预期结果**:
- ✅ URL变更为 /post/{slug} 或 /article/{slug}
- ✅ 可以看到完整的文章内容
- ✅ Markdown渲染正常（标题、代码块、列表等）
- ✅ 侧边栏显示分类和标签
- ✅ 评论区显示

**Playwright测试命令**:
```javascript
// 点击第一篇文章
await page.goto('http://localhost:5173');
await page.locator('.article-card, .article-item').first().click();

// 等待导航
await page.waitForURL(/post|article/);

// 检查文章内容
const content = await page.locator('.article-content, .post-content, .markdown-body').textContent();
expect(content.length).toBeGreaterThan(100);

// 检查评论区
await expect(page.locator('.comment-section, #comments')).toBeVisible();

// 截图
await page.screenshot({ path: 'test-results/article-detail.png', fullPage: true });
```

---

### A3. 分类筛选测试

**测试目标**: 验证点击分类可以筛选文章

**测试步骤**:
```
1. 在首页找到分类导航
2. 点击某个分类
3. 检查筛选结果
```

**预期结果**:
- ✅ URL包含分类参数 (?category=xxx)
- ✅ 显示的文章都属于该分类
- ✅ 分类名称高亮显示

**Playwright测试命令**:
```javascript
await page.goto('http://localhost:5173');

// 点击第一个分类
await page.locator('.category-nav a, .category-link').first().click();

// 检查URL
await page.waitForURL(/category=/);

// 检查文章列表
const articles = await page.locator('.article-card').count();
expect(articles).toBeGreaterThan(0);

// 截图
await page.screenshot({ path: 'test-results/category-filter.png' });
```

---

### A4. 搜索功能测试

**测试目标**: 验证搜索功能正常工作

**测试步骤**:
```
1. 找到搜索框
2. 输入关键词"Vue"
3. 提交搜索
4. 检查搜索结果
```

**预期结果**:
- ✅ 跳转到搜索结果页
- ✅ 显示包含关键词的文章
- ✅ 高亮显示关键词（可选）

**Playwright测试命令**:
```javascript
await page.goto('http://localhost:5173');

// 输入搜索关键词
await page.fill('input[type="search"], input[placeholder*="搜索"]', 'Vue');
await page.press('input[type="search"]', 'Enter');

// 等待搜索结果
await page.waitForURL(/search|q=/);

// 检查结果
const results = await page.locator('.search-result, .article-card').count();
console.log(`找到 ${results} 个搜索结果`);

// 截图
await page.screenshot({ path: 'test-results/search-results.png' });
```

---

### A5. 评论提交测试

**测试目标**: 验证游客可以提交评论

**测试步骤**:
```
1. 进入任意文章详情页
2. 滚动到评论区
3. 填写评论表单
4. 提交评论
```

**预期结果**:
- ✅ 显示成功提示"评论提交成功"
- ✅ 提示"等待审核"
- ✅ 表单清空

**Playwright测试命令**:
```javascript
// 进入文章详情页
await page.goto('http://localhost:5173');
await page.locator('.article-card').first().click();

// 滚动到评论区
await page.locator('#comment-form, .comment-form').scrollIntoViewIfNeeded();

// 填写表单
await page.fill('input[name="nickname"], input[placeholder*="昵称"]', '测试用户');
await page.fill('input[name="email"], input[type="email"]', 'test@example.com');
await page.fill('textarea[name="content"], textarea[placeholder*="评论"]', '这是一条测试评论');

// 提交
await page.click('button[type="submit"], .submit-btn');

// 等待提示
await page.waitForSelector('.success-message, .toast, .notification');

// 截图
await page.screenshot({ path: 'test-results/comment-submit.png' });
```

---

### A6. 归档页面测试

**测试目标**: 验证归档页面按年份展示文章

**测试步骤**:
```
1. 点击导航栏的"归档"链接
2. 检查归档列表
```

**预期结果**:
- ✅ URL为 /archives
- ✅ 按年份分组显示
- ✅ 每年显示文章标题和日期

**Playwright测试命令**:
```javascript
await page.goto('http://localhost:5173');

// 点击归档链接
await page.click('a[href="/archives"], a:has-text("归档")');

// 检查URL
await page.waitForURL(/archives/);

// 检查年份标题
const years = await page.locator('.year-title, h2').count();
expect(years).toBeGreaterThan(0);

// 截图
await page.screenshot({ path: 'test-results/archives.png', fullPage: true });
```

---

### A7. 友链页面测试

**测试目标**: 验证友链页面显示正常

**测试步骤**:
```
1. 点击导航栏的"友链"链接
2. 检查友链列表
```

**预期结果**:
- ✅ URL为 /links
- ✅ 显示友链卡片
- ✅ 每个友链包含头像、名称、描述

**Playwright测试命令**:
```javascript
await page.goto('http://localhost:5173');

// 点击友链
await page.click('a[href="/links"], a:has-text("友链")');

// 检查友链卡片
const links = await page.locator('.link-card, .friend-link').count();
console.log(`显示 ${links} 个友链`);

// 截图
await page.screenshot({ path: 'test-results/links.png' });
```

---

## B. 后台测试场景

### B1. 后台登录测试

**测试目标**: 验证管理员可以登录后台

**测试步骤**:
```
1. 打开后台登录页
2. 输入用户名和密码
3. 点击登录
4. 检查是否跳转到仪表盘
```

**预期结果**:
- ✅ 登录成功跳转到 /dashboard
- ✅ 显示欢迎信息
- ✅ 侧边栏显示菜单
- ✅ Token保存到localStorage

**Playwright测试命令**:
```javascript
await page.goto('http://localhost:5174');

// 填写登录表单
await page.fill('input[name="username"], input[placeholder*="用户名"]', 'admin');
await page.fill('input[name="password"], input[type="password"]', '123456');

// 点击登录
await page.click('button[type="submit"], .login-btn');

// 等待跳转
await page.waitForURL(/dashboard/);

// 检查侧边栏
await expect(page.locator('.sidebar, .menu')).toBeVisible();

// 检查Token
const token = await page.evaluate(() => localStorage.getItem('token'));
expect(token).toBeTruthy();

// 截图
await page.screenshot({ path: 'test-results/admin-login.png' });
```

---

### B2. 仪表盘数据显示测试

**测试目标**: 验证仪表盘统计数据正常显示

**测试步骤**:
```
1. 登录后台
2. 查看仪表盘
3. 检查统计卡片
```

**预期结果**:
- ✅ 显示文章总数
- ✅ 显示评论总数
- ✅ 显示浏览总数
- ✅ 显示今日浏览数

**Playwright测试命令**:
```javascript
// 前提：已登录
await page.goto('http://localhost:5174');

// 等待统计卡片加载
await page.waitForSelector('.stat-card, .dashboard-card');

// 检查统计数据
const stats = await page.locator('.stat-card .number, .stat-value').count();
expect(stats).toBeGreaterThanOrEqual(4);

// 截图
await page.screenshot({ path: 'test-results/dashboard.png' });
```

---

### B3. 创建文章测试

**测试目标**: 验证可以创建新文章

**测试步骤**:
```
1. 点击侧边栏"文章管理"
2. 点击"新建文章"按钮
3. 填写文章信息
4. 点击保存
```

**预期结果**:
- ✅ 显示成功提示
- ✅ 文章列表中出现新文章
- ✅ 可以查看新文章详情

**Playwright测试命令**:
```javascript
// 前提：已登录
await page.goto('http://localhost:5174/dashboard');

// 进入文章管理
await page.click('a:has-text("文章")');

// 点击新建
await page.click('button:has-text("新建"), .create-btn');

// 填写表单
await page.fill('input[name="title"]', '测试文章标题');
await page.fill('textarea[name="summary"]', '这是测试文章的摘要');

// 填写Markdown编辑器（如果是CodeMirror或其他编辑器需要特殊处理）
await page.fill('.md-editor textarea, .markdown-editor textarea', '# 测试内容\n\n这是测试文章的正文。');

// 选择分类
await page.selectOption('select[name="category_id"]', '1');

// 保存
await page.click('button:has-text("保存"), .save-btn');

// 等待成功提示
await page.waitForSelector('.success-message, .el-message--success');

// 截图
await page.screenshot({ path: 'test-results/article-create.png' });
```

---

### B4. 评论审核测试

**测试目标**: 验证可以审核评论

**测试步骤**:
```
1. 进入评论管理
2. 找到待审核评论
3. 点击"通过"按钮
```

**预期结果**:
- ✅ 评论状态变为"已审核"
- ✅ 前台可以看到该评论

**Playwright测试命令**:
```javascript
// 前提：已登录
await page.goto('http://localhost:5174/dashboard');

// 进入评论管理
await page.click('a:has-text("评论")');

// 找到待审核评论
const pendingComment = page.locator('tr:has-text("待审核"), .comment-item:has-text("pending")').first();

// 点击通过按钮
await pendingComment.locator('button:has-text("通过"), .approve-btn').click();

// 等待确认
if (await page.locator('.el-message-box, .confirm-dialog').isVisible()) {
    await page.click('button:has-text("确定")');
}

// 等待成功提示
await page.waitForSelector('.success-message');

// 截图
await page.screenshot({ path: 'test-results/comment-approve.png' });
```

---

### B5. 媒体上传测试

**测试目标**: 验证可以上传图片到媒体库

**测试步骤**:
```
1. 进入媒体库
2. 点击上传按钮
3. 选择图片文件
4. 等待上传完成
```

**预期结果**:
- ✅ 显示上传进度
- ✅ 上传成功后显示在列表中
- ✅ 返回图片URL

**Playwright测试命令**:
```javascript
// 前提：已登录
await page.goto('http://localhost:5174/dashboard');

// 进入媒体库
await page.click('a:has-text("媒体")');

// 准备测试图片
const testImagePath = 'e:\\goCode\\goFile\\gin博客\\test\\test-image.jpg';

// 上传文件
await page.setInputFiles('input[type="file"]', testImagePath);

// 等待上传完成
await page.waitForSelector('.upload-success, .el-message--success', { timeout: 10000 });

// 检查图片列表
const images = await page.locator('.media-item, .image-card').count();
expect(images).toBeGreaterThan(0);

// 截图
await page.screenshot({ path: 'test-results/media-upload.png' });
```

---

## C. API测试场景

### C1. 登录API测试

**测试目标**: 验证登录接口返回正确

**测试步骤**:
```
POST /api/v1/auth/login
Body: {"username": "admin", "password": "123456"}
```

**预期结果**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOi...",
    "expires_at": 1783151517,
    "user": {
      "id": 1,
      "username": "admin",
      "nickname": "管理员"
    }
  }
}
```

**Playwright测试命令**:
```javascript
const response = await page.request.post('http://localhost:8080/api/v1/auth/login', {
  data: {
    username: 'admin',
    password: '123456'
  }
});

const data = await response.json();
console.log('Login Response:', data);

expect(response.ok()).toBeTruthy();
expect(data.code).toBe(0);
expect(data.data.token).toBeTruthy();
expect(data.data.expires_at).toBeGreaterThan(Date.now() / 1000);

// 保存token供后续使用
const token = data.data.token;
```

---

### C2. 获取文章列表API测试

**测试目标**: 验证文章列表接口

**测试步骤**:
```
GET /api/v1/articles?page=1&page_size=10
```

**预期结果**:
```json
{
  "code": 0,
  "data": {
    "list": [...],
    "total": 6,
    "page": 1,
    "page_size": 10
  }
}
```

**Playwright测试命令**:
```javascript
const response = await page.request.get('http://localhost:8080/api/v1/articles?page=1&page_size=10');

const data = await response.json();
console.log('Articles:', data);

expect(response.ok()).toBeTruthy();
expect(data.code).toBe(0);
expect(Array.isArray(data.data.list)).toBeTruthy();
expect(data.data.total).toBeGreaterThanOrEqual(0);
```

---

### C3. 创建文章API测试

**测试目标**: 验证创建文章接口

**Playwright测试命令**:
```javascript
// 前提：已有token
const token = 'eyJhbGciOi...'; // 从登录获取

const response = await page.request.post('http://localhost:8080/api/v1/admin/articles', {
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  data: {
    title: 'API测试文章',
    content: '# 测试\n\n这是通过API创建的文章',
    summary: '测试摘要',
    category_id: 1,
    tag_ids: [1, 2],
    status: 'draft'
  }
});

const data = await response.json();
console.log('Create Article:', data);

expect(response.ok()).toBeTruthy();
expect(data.code).toBe(0);
expect(data.data.id).toBeTruthy();

// 保存文章ID供后续使用
const articleId = data.data.id;
```

---

## D. 端到端测试场景

### D1. 完整发布流程测试

**测试目标**: 从创建到发布到前台查看的完整流程

**测试步骤**:
```
1. 后台登录
2. 创建新文章（草稿）
3. 编辑文章内容
4. 发布文章
5. 前台查看文章
6. 提交评论
7. 后台审核评论
8. 前台看到评论
```

**Playwright测试命令**:
```javascript
// 1. 后台登录
await page.goto('http://localhost:5174');
await page.fill('input[name="username"]', 'admin');
await page.fill('input[type="password"]', '123456');
await page.click('button[type="submit"]');
await page.waitForURL(/dashboard/);

// 2-4. 创建并发布文章
const testTitle = `E2E测试文章 ${Date.now()}`;
await page.click('a:has-text("文章")');
await page.click('button:has-text("新建")');
await page.fill('input[name="title"]', testTitle);
await page.fill('textarea[name="summary"]', 'E2E测试摘要');
await page.fill('.md-editor textarea', '# E2E测试\n\n这是端到端测试文章。');
await page.selectOption('select[name="category_id"]', '1');
await page.selectOption('select[name="status"]', 'published');
await page.click('button:has-text("保存")');
await page.waitForSelector('.success-message');

// 5. 前台查看
await page.goto('http://localhost:5173');
await page.fill('input[type="search"]', testTitle);
await page.press('input[type="search"]', 'Enter');
await page.click(`.article-card:has-text("${testTitle}")`);

// 6. 提交评论
await page.fill('input[name="nickname"]', 'E2E测试用户');
await page.fill('input[type="email"]', 'e2e@test.com');
await page.fill('textarea[name="content"]', 'E2E测试评论');
await page.click('button[type="submit"]');
await page.waitForSelector('.success-message');

// 7. 后台审核
await page.goto('http://localhost:5174/dashboard');
await page.click('a:has-text("评论")');
await page.locator('tr:has-text("E2E测试评论")').locator('.approve-btn').click();
await page.waitForSelector('.success-message');

// 8. 前台验证
await page.goto('http://localhost:5173');
await page.fill('input[type="search"]', testTitle);
await page.press('input[type="search"]', 'Enter');
await page.click(`.article-card:has-text("${testTitle}")`);
await expect(page.locator('.comment:has-text("E2E测试评论")')).toBeVisible();

// 截图
await page.screenshot({ path: 'test-results/e2e-complete.png', fullPage: true });

console.log('✅ 端到端测试完成');
```

---

## 📊 测试用例统计

| 测试类型 | 用例数 | 优先级 |
|---------|--------|--------|
| 前台测试 | 7 | P0 |
| 后台测试 | 5 | P0 |
| API测试 | 3 | P0 |
| 端到端测试 | 1 | P0 |
| **总计** | **16** | - |

---

## 🔧 使用MCP工具执行测试

### 方式1: 逐个执行

使用MCP Playwright工具，复制上述测试命令逐个执行。

### 方式2: 批量执行脚本

创建 `test-suite.js` 文件：

```javascript
// test-suite.js
const tests = [
  {
    name: '首页加载测试',
    run: async (page) => {
      await page.goto('http://localhost:5173');
      await expect(page).toHaveTitle(/博客|Blog/i);
      const articles = await page.locator('.article-card').count();
      expect(articles).toBeGreaterThan(0);
    }
  },
  // ... 添加更多测试
];

// 执行所有测试
for (const test of tests) {
  console.log(`执行: ${test.name}`);
  try {
    await test.run(page);
    console.log(`✅ ${test.name} 通过`);
  } catch (error) {
    console.error(`❌ ${test.name} 失败:`, error.message);
  }
}
```

---

## 📝 测试报告模板

执行完测试后，生成测试报告：

```markdown
# 测试报告

**测试时间**: 2026-07-03  
**测试人员**: XXX  
**测试环境**: 开发环境

## 测试结果

- 总用例数: 16
- 通过: 15
- 失败: 1
- 跳过: 0
- 通过率: 93.75%

## 失败用例

### TC-B3: 创建文章测试
- **失败原因**: Markdown编辑器未正确加载
- **截图**: test-results/article-create.png
- **建议**: 检查编辑器组件加载时机

## 测试总结

大部分功能正常，需要修复文章编辑器加载问题。
```

---

## 🎯 下一步

1. 执行所有测试用例
2. 修复发现的问题
3. 重新执行失败的用例
4. 生成完整测试报告
5. 部署到生产环境前再次执行

**祝测试顺利！** 🚀
