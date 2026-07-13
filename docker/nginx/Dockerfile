# ============================================
# Stage 1: Build blog-web
# ============================================
FROM node:20-alpine AS web-builder

WORKDIR /app

COPY blog-web/package*.json ./
RUN npm ci

COPY blog-web/ ./
RUN npm run build

# ============================================
# Stage 2: Build blog-admin
# ============================================
FROM node:20-alpine AS admin-builder

WORKDIR /app

COPY blog-admin/package*.json ./
RUN npm ci

COPY blog-admin/ ./
RUN npm run build

# ============================================
# Stage 3: Nginx gateway
# ============================================
FROM nginx:alpine

# Install timezone data
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

# Copy nginx config
COPY docker/nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Copy web build output -> /web/
COPY --from=web-builder /app/dist /usr/share/nginx/html/web

# Copy admin build output -> /admin/ (admin is built with base="/admin/")
COPY --from=admin-builder /app/dist /usr/share/nginx/html/admin

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
