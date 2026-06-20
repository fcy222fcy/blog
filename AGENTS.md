# 博客前端项目

## 项目是什么

博客系统的前端页面，包含：

- 前台展示页面（prototypes/web/）
- 后台管理界面（admin/）

## 技术栈

- 前端：HTML/CSS/JavaScript（纯静态）
- 后端：Go + Gin + GORM（计划中）
- 未来接入 Vue 3

## 目录说明

- `admin/` - 后台管理界面，可直接在浏览器打开
- `prototypes/` - 前台页面原型
- `pkg/` - Go 公共工具包

## 开发流程

### 修改后台管理

直接编辑 `admin/index.html`、`admin/css/style.css`、`admin/js/app.js`

### 修改前台原型

编辑 `prototypes/web/` 下的文件

### 添加新页面

1. 在对应目录创建 HTML 文件
2. 添加对应的 CSS/JS
3. 如果是后台，在 `app.js` 中注册路由

## 注意事项

- 后台管理使用 localStorage 存储数据
- 前台页面是静态原型，不需要构建
- 修改代码前先了解现有结构
