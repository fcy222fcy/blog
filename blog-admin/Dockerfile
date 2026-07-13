# 后台管理 standalone 部署 Dockerfile
# 注：通常与 docker-compose 统一网关一起部署（使用 docker/nginx/Dockerfile 统一构建）
# 本文件用于：单独部署 blog-admin（如单独二级域名 admin.example.com），或本地快速验证

ARG NODE_VERSION=20-alpine
ARG NGINX_VERSION=alpine

# ============================================
# Stage 1: Build
# ============================================
FROM node:${NODE_VERSION} AS builder

ARG NPM_REGISTRY
WORKDIR /app

COPY package*.json ./
RUN if [ -n "$NPM_REGISTRY" ]; then npm config set registry "$NPM_REGISTRY"; fi \
    && npm ci

COPY . .
RUN npm run build

# ============================================
# Stage 2: Runtime (Nginx)
# ============================================
FROM nginx:${NGINX_VERSION}

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

COPY --from=builder /app/dist /usr/share/nginx/html/admin
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
