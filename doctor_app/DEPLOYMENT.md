# 医生端前端应用部署指南

## 概述

本文档描述了医生端前端应用的部署流程和配置要求。

## 系统要求

### 开发环境
- Node.js >= 16.0.0
- npm >= 8.0.0
- Git

### 生产环境
- Nginx >= 1.18
- SSL证书（HTTPS）
- 域名解析

## 环境配置

### 环境变量

应用支持以下环境变量配置：

```bash
# 应用配置
VITE_APP_TITLE=优医医生版
VITE_APP_VERSION=1.0.0

# API配置
VITE_API_BASE_URL=https://api.youyi.com
VITE_API_TIMEOUT=10000

# 功能开关
VITE_ENABLE_MOCK=false
VITE_ENABLE_DEVTOOLS=false
VITE_ENABLE_ERROR_REPORT=true
```

### 配置文件

- `.env` - 通用配置
- `.env.development` - 开发环境配置
- `.env.production` - 生产环境配置

## 构建流程

### 本地构建

```bash
# 安装依赖
npm install

# 开发环境构建
npm run build:dev

# 生产环境构建
npm run build:prod

# 构建分析
npm run build:analyze
```

### 自动化构建

使用提供的构建脚本：

```bash
# 使用Node.js脚本
node scripts/build.js production

# 使用Shell脚本（Linux/macOS）
chmod +x scripts/deploy.sh
./scripts/deploy.sh production v1.0.0
```

## 部署方式

### 1. 静态文件部署

#### Nginx配置

创建 `/etc/nginx/sites-available/doctor-app` 配置文件：

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name doctor.youyi.com;
    
    # 重定向到HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    server_name doctor.youyi.com;
    
    # SSL配置
    ssl_certificate /path/to/ssl/cert.pem;
    ssl_certificate_key /path/to/ssl/key.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384;
    
    # 网站根目录
    root /var/www/doctor-app;
    index index.html;
    
    # Gzip压缩
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css text/xml text/javascript application/javascript application/xml+rss application/json;
    
    # 静态资源缓存
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
        add_header Vary Accept-Encoding;
    }
    
    # HTML文件不缓存
    location ~* \.html$ {
        expires -1;
        add_header Cache-Control "no-cache, no-store, must-revalidate";
        add_header Pragma "no-cache";
    }
    
    # API代理
    location /api/ {
        proxy_pass http://localhost:8000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    
    # SPA路由支持
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    # 安全头
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
    add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self' https:; media-src 'self'; object-src 'none'; child-src 'none'; worker-src 'self'; frame-ancestors 'self'; form-action 'self'; base-uri 'self';" always;
}
```

启用配置：

```bash
sudo ln -s /etc/nginx/sites-available/doctor-app /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 2. Docker部署

#### Dockerfile

```dockerfile
# 构建阶段
FROM node:18-alpine as build-stage

WORKDIR /app

# 复制package文件
COPY package*.json ./

# 安装依赖
RUN npm ci --only=production=false

# 复制源代码
COPY . .

# 构建应用
RUN npm run build:prod

# 生产阶段
FROM nginx:alpine as production-stage

# 复制构建结果
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 复制Nginx配置
COPY nginx.conf /etc/nginx/nginx.conf

# 暴露端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]
```

#### docker-compose.yml

```yaml
version: '3.8'

services:
  doctor-app:
    build: .
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./ssl:/etc/nginx/ssl:ro
    environment:
      - NODE_ENV=production
    restart: unless-stopped
    
  # 可选：添加后端服务
  doctor-api:
    image: doctor-api:latest
    ports:
      - "8000:8000"
    environment:
      - NODE_ENV=production
    restart: unless-stopped
```

### 3. CDN部署

#### 阿里云OSS + CDN

```bash
# 安装阿里云CLI
npm install -g @alicloud/cli

# 配置访问密钥
aliyun configure

# 上传到OSS
aliyun oss cp -r dist/ oss://your-bucket/doctor-app/ --update

# 刷新CDN缓存
aliyun cdn RefreshObjectCaches --ObjectPath https://cdn.youyi.com/doctor-app/
```

## 监控和维护

### 健康检查

创建健康检查端点：

```bash
# 检查应用是否正常运行
curl -f http://localhost/health

# 检查API连接
curl -f http://localhost/api/health
```

### 日志监控

```bash
# Nginx访问日志
tail -f /var/log/nginx/access.log

# Nginx错误日志
tail -f /var/log/nginx/error.log

# 系统资源监控
htop
```

### 备份策略

```bash
# 创建备份脚本
#!/bin/bash
DATE=$(date +%Y%m%d-%H%M%S)
tar -czf /backup/doctor-app-$DATE.tar.gz -C /var/www doctor-app

# 保留最近7天的备份
find /backup -name "doctor-app-*.tar.gz" -mtime +7 -delete
```

## 性能优化

### 1. 构建优化

- 启用代码分割
- 压缩静态资源
- 移除未使用的代码
- 优化图片资源

### 2. 服务器优化

- 启用Gzip压缩
- 配置静态资源缓存
- 使用HTTP/2
- 启用Keep-Alive

### 3. CDN优化

- 静态资源CDN加速
- 图片CDN优化
- 全球节点分发

## 安全配置

### 1. HTTPS配置

```bash
# 使用Let's Encrypt免费证书
sudo certbot --nginx -d doctor.youyi.com
```

### 2. 安全头配置

在Nginx配置中添加安全头（见上面的配置示例）。

### 3. 访问控制

```nginx
# 限制访问频率
limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;

location /api/ {
    limit_req zone=api burst=20 nodelay;
    # ... 其他配置
}
```

## 故障排除

### 常见问题

1. **白屏问题**
   - 检查构建是否成功
   - 检查静态资源路径
   - 查看浏览器控制台错误

2. **API请求失败**
   - 检查代理配置
   - 验证后端服务状态
   - 检查CORS配置

3. **路由404错误**
   - 确认Nginx配置了SPA路由支持
   - 检查`try_files`配置

### 调试命令

```bash
# 检查Nginx配置
sudo nginx -t

# 重新加载Nginx
sudo systemctl reload nginx

# 查看Nginx状态
sudo systemctl status nginx

# 查看端口占用
netstat -tlnp | grep :80
```

## 版本管理

### 发布流程

1. 代码合并到主分支
2. 创建版本标签
3. 自动化构建和测试
4. 部署到测试环境
5. 验证功能
6. 部署到生产环境
7. 监控和回滚准备

### 回滚策略

```bash
# 快速回滚到上一个版本
sudo cp -r /backup/doctor-app-previous /var/www/doctor-app
sudo systemctl reload nginx
```

## 联系信息

如有部署问题，请联系：
- 技术支持：tech@youyi.com
- 运维团队：ops@youyi.com