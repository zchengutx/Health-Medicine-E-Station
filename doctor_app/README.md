# 优医医生端前端应用

基于Vue3 + TypeScript + Vite构建的医生端移动前端应用，提供完整的医生注册、登录和管理功能。

## 🌟 功能特性

- 🚀 **启动页面** - 精美的3秒倒计时启动页，包含医疗主题插图
- 🔐 **用户登录** - 手机验证码登录，支持记住手机号
- 📝 **用户注册** - 手机验证码注册，实时密码强度检测
- 📱 **移动端适配** - 完美的响应式设计，支持各种移动设备
- 🎨 **精美UI** - 基于Vant组件库，严格按照设计稿实现
- 🛡️ **错误处理** - 完善的全局错误处理和用户反馈系统
- ⚡ **性能优化** - 代码分割、懒加载、资源压缩
- 🧪 **测试覆盖** - 完整的单元测试、集成测试和E2E测试

## 🛠️ 技术栈

### 核心技术
- **前端框架**: Vue 3.4+ (Composition API)
- **构建工具**: Vite 5.0+
- **路由管理**: Vue Router 4
- **状态管理**: Pinia
- **类型检查**: TypeScript 5.0+

### UI和样式
- **UI组件库**: Vant 4 (移动端)
- **样式预处理**: SCSS
- **CSS框架**: 自定义响应式系统
- **图标**: Vant Icons

### 工具和测试
- **HTTP客户端**: Axios
- **测试框架**: Vitest + Cypress
- **代码质量**: ESLint + Prettier
- **包管理**: npm

## 🚀 快速开始

### 环境要求
- Node.js >= 16.0.0
- npm >= 8.0.0

### 安装和运行

```bash
# 克隆项目
git clone <repository-url>
cd doctor_app

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 在浏览器中打开 http://localhost:3000
```

### 构建和部署

```bash
# 构建生产版本
npm run build:prod

# 预览生产版本
npm run preview

# 构建分析
npm run build:analyze
```

## 📁 项目结构

```
doctor_app/
├── public/                 # 静态资源
├── src/
│   ├── api/               # API接口封装
│   │   ├── doctor.ts      # 医生相关API
│   │   └── request.ts     # HTTP请求配置
│   ├── assets/            # 资源文件
│   │   └── styles/        # 样式文件
│   ├── components/        # 公共组件
│   │   ├── CountdownTimer.vue
│   │   ├── LoadingSpinner.vue
│   │   ├── FeedbackButton.vue
│   │   └── ToastMessage.vue
│   ├── router/            # 路由配置
│   ├── stores/            # 状态管理
│   │   └── auth.ts        # 认证状态
│   ├── utils/             # 工具函数
│   │   ├── validation.ts  # 表单验证
│   │   ├── storage.ts     # 本地存储
│   │   ├── errorHandler.ts # 错误处理
│   │   └── toast.ts       # 消息提示
│   ├── views/             # 页面组件
│   │   ├── SplashView.vue # 启动页
│   │   ├── LoginView.vue  # 登录页
│   │   ├── RegisterView.vue # 注册页
│   │   └── HomeView.vue   # 首页
│   ├── test/              # 测试文件
│   ├── App.vue
│   └── main.ts
├── scripts/               # 构建脚本
├── cypress/               # E2E测试
├── DEPLOYMENT.md          # 部署文档
└── README.md
```

## 🔌 API接口

应用与后端医生服务进行集成，主要接口包括：

```typescript
// 发送短信验证码
POST /api/v1/doctor/SendSms
{
  "Phone": "13812345678",
  "SendSmsCode": "login"
}

// 医生登录
POST /api/v1/doctor/LoginDoctor
{
  "Phone": "13812345678",
  "Password": "",
  "SendSmsCode": "1234"
}

// 医生注册
POST /api/v1/doctor/RegisterDoctor
{
  "Phone": "13812345678",
  "Password": "Abc123456",
  "SendSmsCode": "1234"
}
```

## 🧪 测试

### 运行测试

```bash
# 单元测试
npm run test

# 单元测试（监听模式）
npm run test:ui

# 测试覆盖率
npm run test:coverage

# E2E测试
npm run test:e2e

# 运行所有测试
npm run test:all
```

### 测试覆盖

- ✅ 工具函数测试 (validation, storage)
- ✅ 组件单元测试 (CountdownTimer, LoadingSpinner, FeedbackButton)
- ✅ 状态管理测试 (auth store)
- ✅ 页面集成测试 (用户流程)
- ✅ E2E测试 (完整用户旅程)

## 🚀 部署

### 环境配置

项目支持多环境配置：

- `.env` - 通用配置
- `.env.development` - 开发环境
- `.env.production` - 生产环境

### 部署方式

#### 1. 静态文件部署
```bash
# 构建
npm run build:prod

# 部署到Nginx
cp -r dist/* /var/www/doctor-app/
```

#### 2. Docker部署
```bash
# 构建镜像
docker build -t doctor-app .

# 运行容器
docker run -p 80:80 doctor-app

# 或使用docker-compose
docker-compose up -d
```

#### 3. 自动化部署
```bash
# 使用部署脚本
./scripts/deploy.sh production v1.0.0
```

详细部署说明请参考 [DEPLOYMENT.md](./DEPLOYMENT.md)

## 🎨 设计系统

### 色彩系统
- **主色**: #007AFF (医疗蓝)
- **成功色**: #52C41A
- **警告色**: #FAAD14
- **错误色**: #FF4D4F

### 组件库
- 基于Vant 4构建
- 自定义医疗主题
- 完整的响应式支持
- 触摸优化和触觉反馈

## 📱 移动端优化

- **响应式设计**: 支持320px-768px屏幕
- **触摸优化**: 44px最小触摸目标
- **性能优化**: 代码分割、懒加载
- **离线支持**: Service Worker缓存
- **安全区域**: iPhone X系列适配

## 🛡️ 安全特性

- **数据加密**: 敏感信息本地加密存储
- **HTTPS**: 强制HTTPS通信
- **CSP**: 内容安全策略
- **XSS防护**: 输入验证和输出编码
- **CSRF防护**: Token验证

## 🔧 开发工具

### 推荐的VS Code扩展
- Vue Language Features (Volar)
- TypeScript Vue Plugin (Volar)
- ESLint
- Prettier
- SCSS IntelliSense

### 开发命令
```bash
# 类型检查
npm run type-check

# 代码格式化
npm run lint

# 依赖分析
npm run build:analyze
```

## 📊 性能指标

- **首屏加载**: < 2s
- **交互响应**: < 100ms
- **包大小**: < 500KB (gzipped)
- **Lighthouse评分**: > 90

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 📞 联系我们

- **项目维护者**: 优医技术团队
- **邮箱**: tech@youyi.com
- **文档**: [项目文档](https://docs.youyi.com/doctor-app)
- **问题反馈**: [GitHub Issues](https://github.com/youyi/doctor-app/issues)

## 🙏 致谢

感谢以下开源项目：
- [Vue.js](https://vuejs.org/)
- [Vite](https://vitejs.dev/)
- [Vant](https://vant-contrib.gitee.io/vant/)
- [TypeScript](https://www.typescriptlang.org/)

---

**优医医生端前端应用** - 让医疗服务更便捷 💙