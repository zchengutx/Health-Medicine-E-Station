# 设计文档

## 概述

移动端登录页面将采用现代化的响应式设计，使用React和TypeScript构建，提供流畅的用户体验。页面将实现手机号码验证、短信验证码登录、用户协议确认等核心功能，并支持多种屏幕尺寸的适配。

## 架构

### 技术栈
- **前端框架**: React 18+ with TypeScript
- **样式方案**: CSS Modules + PostCSS
- **状态管理**: React Hooks (useState, useEffect, useContext)
- **表单验证**: 自定义验证hooks
- **HTTP客户端**: Axios
- **构建工具**: Vite
- **移动端适配**: Viewport meta标签 + CSS媒体查询

### 架构模式
采用组件化架构，将登录页面拆分为多个可复用的组件：

```
LoginPage (容器组件)
├── Header (导航栏组件)
├── LoginForm (登录表单组件)
│   ├── PhoneInput (手机号输入组件)
│   ├── VerificationCodeInput (验证码输入组件)
│   └── LoginButton (登录按钮组件)
├── QuickLogin (快捷登录组件)
└── UserAgreement (用户协议组件)
```

## 组件和接口

### 1. LoginPage 容器组件

**职责**: 管理整个登录页面的状态和业务逻辑

**Props接口**:
```typescript
interface LoginPageProps {
  onLoginSuccess: (userInfo: UserInfo) => void;
  onNavigateBack: () => void;
}
```

**状态管理**:
```typescript
interface LoginState {
  phoneNumber: string;
  verificationCode: string;
  isPhoneValid: boolean;
  isCodeValid: boolean;
  isLoading: boolean;
  countdown: number;
  agreedToTerms: boolean;
  errorMessage: string;
}
```

### 2. Header 导航栏组件

**职责**: 显示页面标题和导航控制

**Props接口**:
```typescript
interface HeaderProps {
  title: string;
  onBack: () => void;
  showMenuButton?: boolean;
}
```

### 3. PhoneInput 手机号输入组件

**职责**: 处理手机号输入和验证

**Props接口**:
```typescript
interface PhoneInputProps {
  value: string;
  onChange: (value: string) => void;
  onValidationChange: (isValid: boolean) => void;
  disabled?: boolean;
}
```

**验证规则**:
- 中国大陆手机号格式: /^1[3-9]\d{9}$/
- 实时验证和错误提示

### 4. VerificationCodeInput 验证码输入组件

**职责**: 处理验证码输入和获取

**Props接口**:
```typescript
interface VerificationCodeInputProps {
  value: string;
  onChange: (value: string) => void;
  onValidationChange: (isValid: boolean) => void;
  phoneNumber: string;
  onSendCode: (phone: string) => Promise<void>;
  countdown: number;
  disabled?: boolean;
}
```

### 5. API接口设计

**发送验证码接口**:
```typescript
interface SendCodeRequest {
  phoneNumber: string;
  type: 'login';
}

interface SendCodeResponse {
  success: boolean;
  message: string;
  countdown: number;
}
```

**登录接口**:
```typescript
interface LoginRequest {
  phoneNumber: string;
  verificationCode: string;
}

interface LoginResponse {
  success: boolean;
  message: string;
  token?: string;
  userInfo?: UserInfo;
}
```

## 数据模型

### UserInfo 用户信息模型
```typescript
interface UserInfo {
  id: string;
  phoneNumber: string;
  nickname?: string;
  avatar?: string;
  isNewUser: boolean;
}
```

### ValidationResult 验证结果模型
```typescript
interface ValidationResult {
  isValid: boolean;
  errorMessage?: string;
}
```

### ApiResponse 通用API响应模型
```typescript
interface ApiResponse<T = any> {
  success: boolean;
  data?: T;
  message: string;
  code: number;
}
```

## 错误处理

### 错误类型分类
1. **输入验证错误**: 手机号格式错误、验证码格式错误
2. **网络请求错误**: 网络连接失败、服务器错误
3. **业务逻辑错误**: 验证码错误、手机号未注册
4. **系统错误**: 未知错误、服务不可用

### 错误处理策略
```typescript
enum ErrorType {
  VALIDATION_ERROR = 'VALIDATION_ERROR',
  NETWORK_ERROR = 'NETWORK_ERROR',
  BUSINESS_ERROR = 'BUSINESS_ERROR',
  SYSTEM_ERROR = 'SYSTEM_ERROR'
}

interface ErrorHandler {
  handleError(error: Error, type: ErrorType): void;
  showErrorMessage(message: string): void;
  clearError(): void;
}
```

### 用户友好的错误提示
- 手机号格式错误: "请输入正确的手机号码"
- 验证码错误: "验证码错误，请重新输入"
- 网络错误: "网络连接失败，请检查网络后重试"
- 服务器错误: "服务暂时不可用，请稍后重试"

## 测试策略

### 1. 单元测试
使用Jest + React Testing Library进行组件单元测试：

**测试覆盖范围**:
- 组件渲染测试
- 用户交互测试
- 表单验证逻辑测试
- API调用测试
- 错误处理测试

**关键测试用例**:
```typescript
describe('PhoneInput Component', () => {
  test('should validate phone number format', () => {});
  test('should show error message for invalid phone', () => {});
  test('should enable send code button when phone is valid', () => {});
});

describe('LoginForm Component', () => {
  test('should disable login button when form is invalid', () => {});
  test('should call login API when form is submitted', () => {});
  test('should show loading state during login', () => {});
});
```

### 2. 集成测试
测试组件间的交互和数据流：
- 表单提交流程测试
- 验证码发送和验证流程测试
- 错误状态传递测试

### 3. 端到端测试
使用Cypress进行完整用户流程测试：
- 完整登录流程测试
- 错误场景测试
- 不同设备尺寸适配测试

### 4. 可访问性测试
- 键盘导航测试
- 屏幕阅读器兼容性测试
- 颜色对比度测试
- ARIA标签测试

## 性能优化

### 1. 代码分割
- 使用React.lazy()进行组件懒加载
- 路由级别的代码分割

### 2. 资源优化
- 图片压缩和WebP格式支持
- CSS和JavaScript压缩
- 字体文件优化

### 3. 缓存策略
- 静态资源缓存
- API响应缓存
- 本地存储优化

### 4. 移动端优化
- Touch事件优化
- 防抖和节流处理
- 虚拟键盘适配

## 安全考虑

### 1. 输入安全
- XSS防护：输入内容转义
- 输入长度限制
- 特殊字符过滤

### 2. 网络安全
- HTTPS通信
- Token安全存储
- 请求签名验证

### 3. 验证码安全
- 验证码有效期限制
- 发送频率限制
- 验证次数限制

## 用户体验设计

### 1. 视觉设计
- 遵循Material Design或iOS Human Interface Guidelines
- 一致的颜色方案和字体规范
- 清晰的视觉层次

### 2. 交互设计
- 流畅的动画过渡
- 即时的用户反馈
- 直观的操作流程

### 3. 响应式设计
- 支持多种屏幕尺寸
- 横竖屏适配
- 高DPI屏幕支持