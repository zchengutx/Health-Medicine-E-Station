#!/bin/bash

# 医生端前端应用部署脚本
# 使用方法: ./scripts/deploy.sh [环境] [版本]
# 例如: ./scripts/deploy.sh production v1.0.0

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 默认参数
ENVIRONMENT=${1:-production}
VERSION=${2:-$(date +%Y%m%d-%H%M%S)}
BUILD_DIR="dist"
BACKUP_DIR="backup"

echo -e "${BLUE}🚀 开始部署医生端前端应用${NC}"
echo -e "${BLUE}环境: ${ENVIRONMENT}${NC}"
echo -e "${BLUE}版本: ${VERSION}${NC}"

# 检查Node.js版本
echo -e "${YELLOW}📋 检查环境...${NC}"
node_version=$(node -v)
echo "Node.js版本: $node_version"

if ! command -v node &> /dev/null; then
    echo -e "${RED}❌ Node.js 未安装${NC}"
    exit 1
fi

# 检查npm版本
if ! command -v npm &> /dev/null; then
    echo -e "${RED}❌ npm 未安装${NC}"
    exit 1
fi

# 安装依赖
echo -e "${YELLOW}📦 安装依赖...${NC}"
npm ci --production=false

# 运行测试
echo -e "${YELLOW}🧪 运行测试...${NC}"
npm run test:run

# 类型检查
echo -e "${YELLOW}🔍 类型检查...${NC}"
npm run type-check

# 构建应用
echo -e "${YELLOW}🏗️ 构建应用...${NC}"
if [ "$ENVIRONMENT" = "production" ]; then
    npm run build:prod
else
    npm run build:dev
fi

# 检查构建结果
if [ ! -d "$BUILD_DIR" ]; then
    echo -e "${RED}❌ 构建失败，未找到构建目录${NC}"
    exit 1
fi

echo -e "${GREEN}✅ 构建完成${NC}"

# 创建备份
if [ -d "/var/www/doctor-app" ]; then
    echo -e "${YELLOW}💾 创建备份...${NC}"
    mkdir -p "$BACKUP_DIR"
    tar -czf "$BACKUP_DIR/doctor-app-backup-$(date +%Y%m%d-%H%M%S).tar.gz" -C /var/www doctor-app
    echo -e "${GREEN}✅ 备份完成${NC}"
fi

# 部署文件
echo -e "${YELLOW}🚚 部署文件...${NC}"
if [ "$ENVIRONMENT" = "production" ]; then
    # 生产环境部署
    sudo rsync -av --delete "$BUILD_DIR/" /var/www/doctor-app/
    
    # 设置正确的权限
    sudo chown -R www-data:www-data /var/www/doctor-app
    sudo chmod -R 755 /var/www/doctor-app
    
    # 重启Nginx
    sudo systemctl reload nginx
    
elif [ "$ENVIRONMENT" = "staging" ]; then
    # 测试环境部署
    rsync -av --delete "$BUILD_DIR/" /var/www/staging/doctor-app/
    
else
    echo -e "${YELLOW}⚠️ 开发环境，跳过部署步骤${NC}"
fi

# 健康检查
echo -e "${YELLOW}🏥 健康检查...${NC}"
if [ "$ENVIRONMENT" = "production" ]; then
    sleep 5
    if curl -f -s http://localhost/health > /dev/null; then
        echo -e "${GREEN}✅ 应用运行正常${NC}"
    else
        echo -e "${RED}❌ 应用健康检查失败${NC}"
        exit 1
    fi
fi

# 清理
echo -e "${YELLOW}🧹 清理临时文件...${NC}"
rm -rf node_modules/.cache

echo -e "${GREEN}🎉 部署完成！${NC}"
echo -e "${GREEN}版本: ${VERSION}${NC}"
echo -e "${GREEN}环境: ${ENVIRONMENT}${NC}"
echo -e "${GREEN}时间: $(date)${NC}"

# 发送通知（可选）
if command -v curl &> /dev/null && [ -n "$WEBHOOK_URL" ]; then
    echo -e "${YELLOW}📢 发送部署通知...${NC}"
    curl -X POST "$WEBHOOK_URL" \
        -H "Content-Type: application/json" \
        -d "{\"text\":\"医生端前端应用部署完成\\n环境: $ENVIRONMENT\\n版本: $VERSION\\n时间: $(date)\"}"
fi