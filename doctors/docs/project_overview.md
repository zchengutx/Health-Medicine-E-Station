# 医生端项目概述

## 项目架构

这是一个基于 Kratos 框架的医生端后台服务，采用 Clean Architecture 设计模式。

### 核心功能
- **MySQL 数据库连接** - 连接你现有的数据库
- **Redis 缓存** - 用于验证码存储和缓存
- **医生管理** - 注册、登录、认证等功能
- **短信验证** - 模拟短信验证码发送

### 目录结构
```
doctors/
├── api/                    # API 定义 (protobuf)
│   └── doctor/v1/         # 医生模块 API
├── cmd/doctors/           # 应用入口
├── configs/               # 配置文件
├── internal/              # 内部代码
│   ├── biz/              # 业务逻辑层
│   ├── data/             # 数据访问层
│   ├── service/          # 服务层
│   └── server/           # 服务器配置
├── utils/                # 工具函数
└── docs/                 # 文档
```

## 数据库表

### 1. doctors 表 (医生用户表)
- 存储医生的基本信息、认证信息
- 包含登录凭证、个人资料、执业信息等

### 2. system_configs 表 (系统配置表)
- 存储系统的各种配置信息
- 支持分组管理和类型区分

## 核心组件

### 数据访问层 (Data Layer)
- **data.go** - 数据库和Redis连接管理
- **model.go** - 数据模型定义
- **doctor.go** - 医生相关数据操作

### 业务逻辑层 (Business Layer)
- **doctor.go** - 医生业务逻辑处理
- 包含注册、登录、认证等核心业务

### 服务层 (Service Layer)
- **doctor.go** - HTTP/gRPC 接口实现
- 参数验证和响应处理

## API 接口

### 医生模块接口 (全部 POST 请求)
1. **POST /api/v1/doctor/sms/send** - 发送短信验证码
2. **POST /api/v1/doctor/register** - 医生注册
3. **POST /api/v1/doctor/login** - 医生登录
4. **POST /api/v1/doctor/authentication** - 医生认证

## 配置说明

### configs/config.yaml
```yaml
server:
  http:
    addr: 0.0.0.0:8000    # HTTP 服务端口
  grpc:
    addr: 0.0.0.0:9000    # gRPC 服务端口

data:
  database:
    driver: mysql
    source: "你的MySQL连接字符串"
  redis:
    addr: "你的Redis地址"
    password: "你的Redis密码"
    db: 0
  aliYun:
    accessKeyID: "你的阿里云AccessKey"
    accessKeySecret: "你的阿里云Secret"
```

## 特性

### 1. 数据库连接
- 使用 GORM 作为 ORM
- 自动连接池管理
- 连接状态监控

### 2. Redis 缓存
- 验证码存储 (5分钟过期)
- 连接状态监控
- 错误处理

### 3. 安全特性
- 密码 bcrypt 加密
- 手机号格式验证
- 参数验证

### 4. 日志记录
- 结构化日志
- 操作追踪
- 错误记录

## 启动方式

1. **配置数据库** - 确保 MySQL 和 Redis 可访问
2. **更新配置** - 修改 `configs/config.yaml`
3. **启动服务** - 运行 `kratos run`
4. **查看接口** - 启动后会显示所有可用接口

## 测试方式

使用任何 HTTP 客户端工具测试接口：
- Postman
- curl
- Insomnia
- VS Code REST Client

详细的测试方法请参考 `docs/api_test_guide.md`

## 项目优势

1. **简洁架构** - 只包含必要的功能，易于维护
2. **现有数据库** - 直接使用你的数据库表结构
3. **清晰分层** - 业务逻辑、数据访问、服务层分离
4. **易于扩展** - 基于 Kratos 框架，支持微服务架构
5. **生产就绪** - 包含日志、错误处理、连接池等生产特性

## 下一步

1. 根据业务需求添加更多接口
2. 集成真实的短信服务
3. 添加更多的业务逻辑
4. 部署到生产环境