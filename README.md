# 优医健康医疗电子站 - 医生端系统

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Vue Version](https://img.shields.io/badge/Vue-3.4+-green.svg)](https://vuejs.org)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)]()

一个基于微服务架构的现代化医疗健康管理系统，专为医生端设计，提供完整的医生注册、认证、个人信息管理和医疗服务功能。

## 🌟 项目概述

优医健康医疗电子站是一个全栈医疗管理平台，采用前后端分离架构：

- **后端服务**: 基于 Go + Kratos 微服务框架
- **前端应用**: 基于 Vue 3 + TypeScript + Vite 的移动端应用
- **数据存储**: MySQL + Redis
- **通信协议**: gRPC + HTTP RESTful API

## 🏗️ 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   移动端前端     │    │   Web管理后台    │    │   第三方集成     │
│   (Vue 3)      │    │   (Future)      │    │   (SMS/支付)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
         ┌─────────────────────────────────────────────────┐
         │              API Gateway / Load Balancer        │
         └─────────────────────────────────────────────────┘
                                 │
    ┌────────────────────────────┼────────────────────────────┐
    │                            │                            │
┌─────────┐              ┌─────────────┐              ┌─────────────┐
│ Doctor  │              │Consultation │              │Prescription │
│Service  │              │  Service    │              │  Service    │
│(医生服务)│              │ (咨询服务)   │              │ (处方服务)   │
└─────────┘              └─────────────┘              └─────────────┘
    │                            │                            │
    └────────────────────────────┼────────────────────────────┘
                                 │
         ┌─────────────────────────────────────────────────┐
         │              Data Layer                         │
         │  ┌─────────────┐    ┌─────────────┐            │
         │  │   MySQL     │    │   Redis     │            │
         │  │ (主数据库)   │    │  (缓存)     │            │
         │  └─────────────┘    └─────────────┘            │
         └─────────────────────────────────────────────────┘
```

## 📁 项目结构

```
Health-Medicine-E-Station/
├── doctors/                    # 后端微服务
│   ├── api/                   # API定义 (Protocol Buffers)
│   │   ├── doctor/v1/         # 医生服务API
│   │   ├── consultation/v1/   # 咨询服务API
│   │   ├── patient/v1/        # 患者服务API
│   │   └── prescription/v1/   # 处方服务API
│   ├── cmd/                   # 应用入口
│   ├── configs/               # 配置文件
│   ├── internal/              # 内部代码
│   │   ├── biz/              # 业务逻辑层
│   │   ├── data/             # 数据访问层
│   │   ├── service/          # 服务层
│   │   └── server/           # 服务器配置
│   ├── migrations/           # 数据库迁移
│   └── utils/                # 工具函数
├── doctor_app/               # 前端移动应用
│   ├── src/
│   │   ├── api/             # API接口封装
│   │   ├── components/      # 公共组件
│   │   ├── views/           # 页面组件
│   │   ├── stores/          # 状态管理
│   │   ├── router/          # 路由配置
│   │   └── utils/           # 工具函数
│   ├── public/              # 静态资源
│   └── tests/               # 测试文件
└── .kiro/                   # Kiro IDE配置
    └── specs/               # 功能规范文档
```

## 🚀 快速开始

### 环境要求

**后端服务:**
- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Protocol Buffers 编译器

**前端应用:**
- Node.js 16.0+
- npm 8.0+

### 安装和运行

#### 1. 克隆项目

```bash
git clone <repository-url>
cd Health-Medicine-E-Station
```

#### 2. 启动后端服务

```bash
cd doctors

# 安装依赖
go mod download

# 生成代码
make api

# 配置数据库连接 (编辑 configs/config.yaml)
# 运行数据库迁移
make migrate

# 启动服务
make run

# 或者使用 Docker
docker-compose up -d
```

#### 3. 启动前端应用

```bash
cd doctor_app

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 在浏览器中打开 http://localhost:3000
```

## 🔧 技术栈详解

### 后端技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| **Go** | 1.21+ | 主要编程语言 |
| **Kratos** | 2.8.0 | 微服务框架 |
| **gRPC** | 1.65.0 | 服务间通信 |
| **Protocol Buffers** | 3.x | 接口定义语言 |
| **GORM** | 1.30.1 | ORM框架 |
| **MySQL** | 8.0+ | 主数据库 |
| **Redis** | 6.0+ | 缓存和会话存储 |
| **JWT** | 5.1.0 | 身份认证 |
| **Wire** | 0.6.0 | 依赖注入 |
| **Docker** | - | 容器化部署 |

### 前端技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| **Vue 3** | 3.4+ | 前端框架 |
| **TypeScript** | 5.3+ | 类型系统 |
| **Vite** | 5.0+ | 构建工具 |
| **Vue Router** | 4.2+ | 路由管理 |
| **Pinia** | 2.1+ | 状态管理 |
| **Vant** | 4.8+ | UI组件库 |
| **Axios** | 1.6+ | HTTP客户端 |
| **SCSS** | 1.69+ | CSS预处理器 |
| **Vitest** | 1.1+ | 单元测试 |
| **Cypress** | 13.6+ | E2E测试 |

## 📊 核心功能模块

### 1. 医生管理模块 (Doctor Service)

**功能特性:**
- ✅ 医生注册和登录
- ✅ 短信验证码认证
- ✅ 个人信息管理
- ✅ 医生资质认证
- ✅ 密码管理
- ✅ 账号注销

**API接口:**
```protobuf
service Doctor {
  rpc SendSms(SendSmsReq) returns (SendSmsResp);
  rpc RegisterDoctor(RegisterDoctorReq) returns (RegisterDoctorResp);
  rpc LoginDoctor(LoginDoctorReq) returns (LoginDoctorResp);
  rpc Authentication(AuthenticationReq) returns (AuthenticationResp);
  rpc GetDoctorProfile(GetDoctorProfileReq) returns (GetDoctorProfileResp);
  rpc UpdateDoctorProfile(UpdateDoctorProfileReq) returns (UpdateDoctorProfileResp);
  rpc ChangePassword(ChangePasswordReq) returns (ChangePasswordResp);
  rpc DeleteAccount(DeleteAccountReq) returns (DeleteAccountResp);
}
```

### 2. 咨询管理模块 (Consultation Service)

**功能特性:**
- 🚧 在线咨询管理
- 🚧 咨询记录查询
- 🚧 实时消息通信
- 🚧 咨询状态跟踪

### 3. 处方管理模块 (Prescription Service)

**功能特性:**
- 🚧 电子处方开具
- 🚧 处方历史管理
- 🚧 药品信息查询
- 🚧 处方审核流程

### 4. 患者管理模块 (Patient Service)

**功能特性:**
- 🚧 患者信息管理
- 🚧 病历记录
- 🚧 随访计划
- 🚧 健康档案

## 🗄️ 数据库设计

### 核心数据表

#### 医生表 (doctors)
```sql
CREATE TABLE doctors (
  id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  doctor_code VARCHAR(32) NOT NULL COMMENT '医生编码',
  name VARCHAR(50) COMMENT '医生姓名',
  gender VARCHAR(10) NOT NULL DEFAULT '男' COMMENT '性别',
  birth_date DATE COMMENT '出生日期',
  phone CHAR(11) NOT NULL COMMENT '手机号码',
  password VARCHAR(255) NOT NULL COMMENT '密码',
  email VARCHAR(100) COMMENT '邮箱地址',
  avatar VARCHAR(255) COMMENT '头像URL',
  license_number VARCHAR(50) COMMENT '执业医师资格证号',
  department_id BIGINT UNSIGNED COMMENT '科室ID',
  hospital_id BIGINT UNSIGNED COMMENT '医院ID',
  title VARCHAR(50) COMMENT '职称',
  speciality TEXT COMMENT '专业特长',
  practice_scope TEXT COMMENT '执业范围',
  status VARCHAR(10) NOT NULL DEFAULT '启用' COMMENT '状态',
  last_login_time TIMESTAMP COMMENT '最后登录时间',
  last_login_ip VARCHAR(45) COMMENT '最后登录IP',
  created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  deleted_at DATETIME(6) COMMENT '删除时间',
  
  UNIQUE KEY uk_phone (phone),
  UNIQUE KEY uk_doctor_code (doctor_code),
  KEY idx_department_id (department_id),
  KEY idx_hospital_id (hospital_id),
  KEY idx_status (status)
);
```

## 🔐 安全特性

### 身份认证和授权
- **JWT Token**: 基于JWT的无状态身份认证
- **短信验证**: 阿里云短信服务集成
- **密码加密**: bcrypt哈希加密存储
- **会话管理**: Redis存储会话信息

### 数据安全
- **SQL注入防护**: GORM参数化查询
- **XSS防护**: 前端输入验证和输出编码
- **CSRF防护**: Token验证机制
- **HTTPS**: 强制HTTPS通信

### 接口安全
- **参数验证**: Protocol Buffers类型验证
- **错误处理**: 统一错误码和错误信息
- **日志审计**: 完整的操作日志记录
- **限流控制**: Redis实现的接口限流

## 🧪 测试策略

### 后端测试
```bash
# 单元测试
go test ./...

# 集成测试
go test -tags=integration ./...

# 基准测试
go test -bench=. ./...

# 测试覆盖率
go test -cover ./...
```

### 前端测试
```bash
# 单元测试
npm run test

# E2E测试
npm run test:e2e

# 测试覆盖率
npm run test:coverage

# 组件测试
npm run test:component
```

### 测试覆盖率目标
- **后端代码覆盖率**: > 80%
- **前端代码覆盖率**: > 85%
- **API接口测试**: 100%
- **关键业务流程**: 100%

## 🚀 部署指南

### 开发环境部署

```bash
# 使用 Docker Compose
docker-compose -f docker-compose.dev.yml up -d

# 或者分别启动
# 启动数据库和Redis
docker-compose up -d mysql redis

# 启动后端服务
cd doctors && make run

# 启动前端应用
cd doctor_app && npm run dev
```

### 生产环境部署

```bash
# 构建Docker镜像
docker build -t doctor-backend ./doctors
docker build -t doctor-frontend ./doctor_app

# 使用生产配置启动
docker-compose -f docker-compose.prod.yml up -d

# 或者使用Kubernetes
kubectl apply -f k8s/
```

### 环境配置

#### 后端配置 (configs/config.yaml)
```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9001
    timeout: 1s

data:
  database:
    driver: mysql
    source: user:password@tcp(host:port)/database?parseTime=True&loc=Local
  redis:
    addr: host:port
    password: password
    db: 0
```

#### 前端配置 (.env.production)
```env
VITE_API_BASE_URL=https://api.yourdomain.com
VITE_APP_TITLE=优医医生端
VITE_APP_VERSION=1.0.0
```

## 📈 性能指标

### 后端性能
- **响应时间**: < 100ms (P95)
- **并发处理**: > 1000 QPS
- **内存使用**: < 512MB
- **CPU使用**: < 50%

### 前端性能
- **首屏加载**: < 2s
- **交互响应**: < 100ms
- **包大小**: < 500KB (gzipped)
- **Lighthouse评分**: > 90

### 数据库性能
- **查询响应**: < 50ms
- **连接池**: 10-100 连接
- **缓存命中率**: > 90%

## 🔄 CI/CD 流程

### 持续集成
```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  backend-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: 1.21
      - run: make test
      
  frontend-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-node@v3
        with:
          node-version: 16
      - run: npm ci && npm run test:all
```

### 持续部署
- **开发环境**: 自动部署到开发服务器
- **测试环境**: PR合并后自动部署
- **生产环境**: 手动触发部署流程

## 📊 监控和日志

### 应用监控
- **健康检查**: HTTP健康检查端点
- **指标收集**: Prometheus + Grafana
- **链路追踪**: Jaeger分布式追踪
- **告警通知**: 钉钉/邮件告警

### 日志管理
- **结构化日志**: JSON格式日志输出
- **日志级别**: Debug/Info/Warn/Error
- **日志收集**: ELK Stack
- **日志轮转**: 按大小和时间轮转

## 🤝 开发规范

### 代码规范
- **Go**: 遵循Go官方代码规范
- **TypeScript**: 使用ESLint + Prettier
- **Git**: 使用Conventional Commits规范
- **API**: RESTful API设计原则

### 分支管理
- **main**: 生产环境分支
- **develop**: 开发环境分支
- **feature/***: 功能开发分支
- **hotfix/***: 紧急修复分支

### 代码审查
- **PR Review**: 至少一人审查
- **自动化检查**: CI/CD流水线检查
- **测试覆盖**: 新功能必须包含测试
- **文档更新**: API变更需更新文档

## 🛣️ 发展路线图

### v1.0.0 (当前版本)
- ✅ 医生注册和登录
- ✅ 个人信息管理
- ✅ 基础认证功能

### v1.1.0 (计划中)
- 🚧 在线咨询功能
- 🚧 患者管理
- 🚧 消息通知系统

### v1.2.0 (规划中)
- 📋 电子处方功能
- 📋 药品管理
- 📋 处方审核流程

### v2.0.0 (远期规划)
- 📋 AI辅助诊断
- 📋 数据分析报表
- 📋 多租户支持

## 🐛 问题反馈

### 已知问题
- [ ] 个人信息页面首次加载问题 (已修复)
- [ ] 日期字段更新500错误 (已修复)

### 反馈渠道
- **GitHub Issues**: [项目Issues页面]
- **邮箱**: tech@youyi.com
- **钉钉群**: 优医技术交流群

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

感谢以下开源项目和技术社区：

- [Kratos](https://go-kratos.dev/) - 微服务框架
- [Vue.js](https://vuejs.org/) - 前端框架
- [GORM](https://gorm.io/) - Go ORM库
- [Vant](https://vant-contrib.gitee.io/vant/) - 移动端UI组件库
- [Protocol Buffers](https://developers.google.com/protocol-buffers) - 接口定义语言

## 👥 贡献者

- **项目负责人**: 优医技术团队
- **后端开发**: Go开发团队
- **前端开发**: Vue开发团队
- **测试工程师**: QA团队
- **运维工程师**: DevOps团队

---

**优医健康医疗电子站** - 让医疗服务更智能、更便捷 💙

[![Star History Chart](https://api.star-history.com/svg?repos=youyi/Health-Medicine-E-Station&type=Date)](https://star-history.com/#youyi/Health-Medicine-E-Station&Date)