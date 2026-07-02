# 博客前端项目

## 项目概述

博客系统的前端页面，包含前台展示和后台管理。

## 技术栈

- **框架**: Vue 3（待接入）
- **样式**: CSS 变量 + 深色/浅色主题
- **语言**: HTML/CSS/JavaScript

## 前端规范

### HTML 规范

- 使用语义化标签：`<header>`, `<nav>`, `<main>`, `<article>`, `<aside>`, `<footer>`
- 一个页面只有一个 `<h1>`
- 图片必须有 `alt` 属性
- 表单元素必须有关联的 `<label>`

### CSS 规范

- 使用 CSS 变量管理主题色
- 命名使用 BEM 或语义化命名
- 响应式断点：768px（平板）、1024px（桌面）
- 深色模式使用 `[data-scheme="dark"]` 选择器

### JavaScript 规范

- 使用 ES6+ 语法
- 避免使用 `var`，使用 `const` 和 `let`
- 函数命名使用驼峰式
- 事件处理使用 `addEventListener`

### 图标规范

- **少用 emoji，多用 SVG 图标**
- 按钮、导航、状态提示等场景优先使用 SVG 图标
- SVG 图标保持统一风格（线条粗细、颜色、尺寸）
- 仅在非正式场景（如聊天、评论）可使用 emoji

## 主题规范

### 颜色变量

```css
:root {
    --bg-color: #ffffff;
    --card-bg: #f8f9fa;
    --text-primary: #333333;
    --text-secondary: #666666;
    --accent-color: #007bff;
    --success-color: #28a745;
    --warning-color: #ffc107;
    --error-color: #dc3545;
}

[data-scheme="dark"] {
    --bg-color: #1a1a1a;
    --card-bg: #2d2d2d;
    --text-primary: #ffffff;
    --text-secondary: #b0b0b0;
}
```

## 响应式规范

- 移动端优先设计
- 断点：
  - `< 768px`: 移动端
  - `768px - 1024px`: 平板
  - `> 1024px`: 桌面端

## 组件规范

### 按钮

- 主要按钮：`.btn-primary`
- 次要按钮：`.btn-secondary`
- 危险按钮：`.btn-danger`
- 尺寸：`.btn-sm`, `.btn-md`, `.btn-lg`

### 卡片

- 统一使用 `.card` 类
- 圆角：8px
- 阴影：`0 2px 8px rgba(0,0,0,0.1)`

### 表单

- 输入框：`.form-input`
- 选择框：`.form-select`
- 文本域：`.form-textarea`
- 错误状态：`.form-error`

## 状态提示

- 成功：绿色提示
- 警告：黄色提示
- 错误：红色提示
- 使用 Toast 组件显示

## Git 规范

- 提交信息清晰描述改动
- 一次提交只包含一个逻辑改动
- 提交前格式化代码
