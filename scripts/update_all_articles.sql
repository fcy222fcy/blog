-- 补充所有文章内容
USE blog;

-- 文章1: 我的博客主题已开源，欢迎使用
UPDATE articles SET content = '## 引言

基于 Hugo Theme Stack 打造的开箱即用博客模板，包含多项美化和功能增强。这个主题已经开源，欢迎使用和贡献！

## 特性

### 响应式设计
- 完美适配手机、平板、桌面等各种设备
- 移动端友好的导航菜单
- 图片自适应缩放

### 主题系统
- 支持深色/浅色主题切换
- 主题颜色跟随系统设置
- 平滑的过渡动画

### 代码高亮
- 支持 100+ 编程语言
- 一键复制代码
- 行号显示
- 代码折叠

### 搜索功能
- 全文搜索
- 搜索结果高亮
- 快捷键支持 (Ctrl+K)

### 其他特性
- 目录自动生成
- 图片懒加载
- 代码块样式美化
- 评论系统集成
- RSS 订阅支持

## 安装方法

```bash
# 克隆仓库
git clone https://github.com/your-username/hugo-blog-theme.git

# 进入目录
cd hugo-blog-theme

# 安装依赖
hugo mod get -u

# 启动开发服务器
hugo server -D
```

## 配置说明

在 `config.toml` 中配置你的博客信息：

```toml
title = "我的博客"
 baseURL = "https://your-domain.com"
theme = "hugo-blog-theme"

[params]
  author = "Your Name"
  description = "个人技术博客"
  dateFormat = "2006-01-02"
```

## 自定义

### 添加文章

```bash
hugo new content posts/my-first-post.md
```

### 修改颜色

编辑 `assets/css/variables.css` 文件中的 CSS 变量。

### 添加页面

在 `content` 目录下创建新的 `.md` 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

如果你喜欢我的博客样式，不妨试试它！有任何问题欢迎交流。' WHERE id = 1;

-- 文章2: DesktopSnap：开源轻量桌面图标恢复工具
UPDATE articles SET content = '## 项目介绍

DesktopSnap 是一款基于 WinUI3 开发的轻量级桌面工具，支持多显示器布局一键备份与恢复。作为一名经常需要在多个显示器之间切换的开发者，我深受桌面图标混乱的困扰，于是开发了这款工具。

## 为什么要做这个工具？

你是否遇到过以下情况：

- 重装系统后桌面图标全部打乱
- 连接不同显示器后图标位置变化
- 多显示器切换后图标重叠

DesktopSnap 就是为了解决这些问题而诞生的。

## 功能特点

### 多显示器支持
- 自动识别所有连接的显示器
- 分别保存每个显示器的图标布局
- 支持不同分辨率和缩放比例

### 一键备份
- 记录所有图标的位置信息
- 保存图标名称和路径
- 支持多个备份版本

### 一键恢复
- 快速恢复之前的图标布局
- 智能匹配图标
- 恢复失败时的回滚机制

### 已上架微软商店
- 安全可靠
- 自动更新
- 免费使用

## 技术实现

### 技术栈
- **UI 框架**: WinUI 3
- **开发语言**: C#
- **构建工具**: MSBuild
- **打包方式**: MSIX

### 核心原理
1. 通过 Windows API 获取桌面图标位置
2. 将位置信息序列化为 JSON
3. 存储到本地文件或云同步
4. 恢复时反序列化并设置图标位置

## 下载地址

[微软商店](https://apps.microsoft.com/detail/desktopsnap)

## 源码

项目已开源，欢迎查看和贡献：

```bash
git clone https://github.com/your-username/DesktopSnap.git
```

## 后续计划

- 支持云端同步
- 定时自动备份
- 支持更多桌面管理功能

感谢使用 DesktopSnap！' WHERE id = 2;

-- 文章4: 清明假期：收拾家务、整理心情
UPDATE articles SET content = '## 假期日常

清明假期三天，没有出去人挤人，选择在家收拾家务。平时工作忙，家里堆积了不少需要整理的东西，趁着假期好好打理一下。

## 收拾成果

### 书房整理
- 重新整理了书架，把看过的和没看过的分开
- 清理了桌上堆积的杂物
- 整理了各种数据线和充电器
- 把不用的电子设备收了起来

### 衣柜清理
- 断舍离了一批不再穿的衣服
- 按季节重新分类整理
- 清理了衣柜底部的"遗忘角落"

### 厨房整理
- 清理了过期的调味料
- 整理了餐具和厨具
- 给冰箱来了一次大扫除

## 断舍离心得

这次整理最大的感触是：很多东西留着并不是因为需要，而是因为"舍不得"。

- 三年没穿的衣服 → 该放手了
- 用不上的各种线材 → 找不到对应的设备就捐掉
- 过期的食品调料 → 健康第一

## 整理的意义

整理物品的过程也是整理心情的过程。

当你把一个杂乱的角落收拾得井井有条时，那种成就感是难以言喻的。看着整洁的房间，心情也会变得舒畅。

断舍离让生活更简单，也让我更清楚自己真正需要什么。

## 假期最后一天

假期最后一天，泡了一壶茶，坐在整理好的书房里，翻了几页书。这大概就是假期最好的打开方式了。

新的一个月，从整洁的环境开始。' WHERE id = 4;

-- 文章5: 逐渐难以逃离对于 AI 的焦虑
UPDATE articles SET content = '## 引言

最近看着各种 AI 工具的发展，感到越来越焦虑。从 ChatGPT 到各种 AI 编程助手，从 AI 绘画到 AI 视频，技术的迭代速度似乎越来越快。

## 焦虑来源

### 1. 技术更新太快

上周还在学的框架，这周可能就被新技术取代了。AI 领域更是如此，几乎每周都有新的模型发布。

### 2. 担心被替代

看到 AI 可以写代码、写文章、做设计，不禁会想：我的工作会不会被替代？

### 3. 学习跟不上

想学习 AI 相关技术，但发现需要的知识面太广：机器学习、深度学习、各种框架...时间根本不够用。

### 4. 同行压力

看到同行都在学习 AI、使用 AI，自己如果落后了怎么办？

## 如何应对

### 1. 保持学习，但不要焦虑

学习是必要的，但不必为此焦虑。技术是工具，关键是用它来解决实际问题。

### 2. 关注本质

AI 再强大，也需要人来定义问题、判断结果。提升自己的核心竞争力才是关键。

### 3. 做好当下

与其焦虑未来，不如专注当下。把手头的工作做好，自然不会被淘汰。

### 4. 拥抱变化

AI 不是敌人，而是工具。学会使用 AI，让它成为你的助力。

## 我的实践

- 开始使用 AI 编程助手提高效率
- 学习 AI 的基本原理，而不是所有细节
- 关注 AI 在自己领域的应用
- 保持好奇心，但不过度焦虑

## 结语

与其焦虑，不如行动。与其担心被替代，不如学会与 AI 共处。

技术在变，但解决问题的能力、创造力、沟通能力这些核心能力是不会被替代的。

保持学习，保持思考，做好当下。' WHERE id = 5;

-- 文章11: Docker 容器化部署实践
UPDATE articles SET content = '## 为什么要用 Docker？

Docker 解决了"在我机器上能运行"的经典问题，让应用部署变得标准化、可移植。

### 传统部署的痛点

- 环境配置复杂
- 不同机器环境不一致
- 部署过程繁琐
- 难以回滚

### Docker 的优势

- 环境一致性
- 快速部署
- 易于扩展
- 资源隔离

## 核心概念

### 镜像 (Image)
镜像是一个只读模板，包含运行应用所需的一切：代码、运行时、库、环境变量。

### 容器 (Container)
容器是镜像的运行实例，是真正运行应用的地方。

### 仓库 (Registry)
仓库用于存储和分发镜像，最常用的是 Docker Hub。

## Dockerfile 编写

```dockerfile
# 基础镜像
FROM golang:1.21-alpine

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN go build -o main .

# 运行
CMD ["./main"]
```

## 常用命令

```bash
# 构建镜像
docker build -t myapp:latest .

# 运行容器
docker run -d -p 8080:8080 myapp:latest

# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 停止容器
docker stop <container_id>

# 进入容器
docker exec -it <container_id> sh

# 查看日志
docker logs <container_id>

# 清理未使用的资源
docker system prune
```

## Docker Compose

对于多容器应用，使用 Docker Compose 更方便：

```yaml
version: '3.8'

services:
  web:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db
    environment:
      - DB_HOST=db

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: myapp
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
```

## 最佳实践

### 1. 使用多阶段构建

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# 运行阶段
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### 2. 优化镜像大小

- 使用 Alpine 基础镜像
- 合并 RUN 指令
- 使用 .dockerignore

### 3. 安全考虑

- 不要在镜像中存储敏感信息
- 使用非 root 用户运行
- 定期更新基础镜像

## 常见问题

### 容器无法启动
检查日志：`docker logs <container_id>`

### 端口冲突
修改端口映射或停止占用端口的进程

### 网络问题
使用 `docker network` 创建自定义网络

## 总结

Docker 是现代开发和运维的必备技能。掌握 Docker，可以让部署变得更加简单可靠。' WHERE id = 11;
