# 个人信息页面统一设计文档

## 概述

本设计文档描述了如何将现有的两个个人信息相关页面（DoctorAuthView 和 ProfileView）统一成一个功能完整的个人信息页面。新的统一页面将支持查看、编辑和保存个人信息，同时移除冗余的路由和组件。

## 架构

### 当前架构分析

**现有组件：**
1. **ProfileView** (`/profile`) - 个人信息展示和编辑页面
   - 支持获取和更新个人信息
   - 使用 `getProfile` 和 `updateProfile` API
   - 包含完整的错误处理和加载状态

2. **DoctorAuthView** (`/doctor-auth`) - 医生认证页面
   - 用于首次填写个人信息
   - 使用 `authentication` API
   - 表单字段与ProfileView部分重叠但不完全相同

**API接口：**
- `getProfile` - 获取医生个人信息
- `updateProfile` - 更新医生个人信息  
- `authentication` - 医生认证（首次提交个人信息）

### 目标架构

**统一后的架构：**
1. **增强的ProfileView** - 作为唯一的个人信息页面
   - 自动检测是否为首次使用（无个人信息数据）
   - 根据数据状态显示不同的UI模式
   - 统一使用 `updateProfile` API（简化API调用逻辑）

2. **路由重定向** - `/doctor-auth` 重定向到 `/profile`

3. **MineView更新** - 移除医生头像点击跳转功能

## 组件和接口

### 增强的ProfileView组件

**状态管理：**
```typescript
interface ProfileState {
  loading: boolean
  authLoading: boolean
  hasError: boolean
  profileLoaded: boolean
  isFirstTime: boolean  // 新增：标识是否为首次使用
  mode: 'view' | 'edit'  // 新增：页面模式
}
```

**数据模型：**
```typescript
interface UnifiedProfileForm {
  // 基本信息
  DId: number
  Name: string
  Gender: string
  BirthDate: string
  Phone: string
  Email: string
  Avatar: string
  
  // 职业信息
  LicenseNumber: string
  DepartmentId: number
  HospitalId: number
  Title: string
  Speciality: string
  PracticeScope: string
  
  // 认证页面特有字段（需要整合）
  realName?: string      // 映射到 Name
  hospital?: string      // 映射到 HospitalId
  department?: string    // 映射到 DepartmentId
  specialty?: string     // 映射到 Speciality
  profile?: string       // 映射到 PracticeScope
  experience?: string    // 新增字段或映射到现有字段
}
```

**组件方法：**
```typescript
class EnhancedProfileView {
  // 初始化方法
  async initializeProfile(): Promise<void>
  
  // 数据获取
  async fetchProfile(): Promise<void>
  
  // 数据保存（统一处理首次和更新）
  async saveProfile(): Promise<void>
  
  // 模式切换
  toggleEditMode(): void
  
  // 表单验证
  validateForm(): boolean
  
  // 字段映射（处理认证页面字段到标准字段的转换）
  mapAuthFieldsToProfile(authData: any): UnifiedProfileForm
}
```

### 路由配置更新

**新的路由配置：**
```typescript
const routes = [
  // 保留原有的profile路由
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/ProfileView.vue')
  },
  
  // 重定向doctor-auth到profile
  {
    path: '/doctor-auth',
    redirect: '/profile'
  }
]
```

### MineView组件更新

**更新的交互逻辑：**
```typescript
// 移除医生头像点击跳转
const goDoctorAuth = () => {
  // 移除此方法或改为其他功能
}

// 保留编辑个人信息按钮
const goEditProfile = () => {
  router.push('/profile')
}
```

## 数据模型

### 统一的表单字段配置

```typescript
const unifiedFormItems = [
  // 基本信息组
  { 
    group: 'basic',
    key: 'Name', 
    label: '真实姓名', 
    placeholder: '请填写您的真实姓名',
    required: true,
    type: 'text'
  },
  {
    group: 'basic',
    key: 'Gender',
    label: '性别',
    placeholder: '请选择性别',
    required: true,
    type: 'select',
    options: [
      { value: '男', label: '男' },
      { value: '女', label: '女' }
    ]
  },
  {
    group: 'basic',
    key: 'BirthDate',
    label: '出生日期',
    placeholder: '请选择出生日期',
    required: false,
    type: 'date'
  },
  {
    group: 'basic',
    key: 'Phone',
    label: '手机号',
    placeholder: '请输入手机号',
    required: true,
    type: 'tel',
    readonly: true  // 手机号通常不允许修改
  },
  {
    group: 'basic',
    key: 'Email',
    label: '邮箱',
    placeholder: '请输入邮箱',
    required: false,
    type: 'email'
  },
  
  // 职业信息组
  {
    group: 'professional',
    key: 'HospitalId',
    label: '就职医院',
    placeholder: '请选择您目前所执业的医院',
    required: true,
    type: 'select'
  },
  {
    group: 'professional',
    key: 'DepartmentId',
    label: '所属科室',
    placeholder: '请选择您所属的科室',
    required: true,
    type: 'select'
  },
  {
    group: 'professional',
    key: 'Title',
    label: '职称',
    placeholder: '请填写您的职称',
    required: true,
    type: 'text'
  },
  {
    group: 'professional',
    key: 'LicenseNumber',
    label: '执业证号',
    placeholder: '请输入执业证号',
    required: false,
    type: 'text'
  },
  {
    group: 'professional',
    key: 'Speciality',
    label: '擅长领域',
    placeholder: '请填写您的擅长领域',
    required: false,
    type: 'textarea'
  },
  {
    group: 'professional',
    key: 'PracticeScope',
    label: '个人简介',
    placeholder: '请填写个人简介',
    required: false,
    type: 'textarea'
  }
]
```

### API数据映射

**认证API到更新API的字段映射：**
```typescript
const fieldMapping = {
  // 认证页面字段 -> 标准字段
  'realName': 'Name',
  'hospital': 'HospitalId',  // 需要转换为ID
  'department': 'DepartmentId',  // 需要转换为ID
  'title': 'Title',
  'specialty': 'Speciality',
  'profile': 'PracticeScope',
  'experience': 'PracticeScope'  // 合并到个人简介
}
```

## 错误处理

### 统一的错误处理策略

**错误类型分类：**
1. **网络错误** - 显示网络连接失败提示，提供重试选项
2. **认证错误** - 重定向到登录页面
3. **数据验证错误** - 显示具体的字段验证错误
4. **服务器错误** - 显示通用错误信息，提供重试选项

**错误恢复机制：**
```typescript
class ErrorRecovery {
  // 自动重试机制
  async retryWithBackoff(operation: () => Promise<any>, maxRetries: number = 3): Promise<any>
  
  // 本地缓存恢复
  async recoverFromCache(): Promise<boolean>
  
  // 降级处理
  async fallbackToReadOnlyMode(): Promise<void>
}
```

## 测试策略

### 单元测试

**测试覆盖范围：**
1. **组件渲染测试**
   - 首次使用时显示完整表单
   - 有数据时显示预填充表单
   - 加载状态显示
   - 错误状态显示

2. **数据处理测试**
   - 字段映射功能
   - 表单验证逻辑
   - API调用处理

3. **用户交互测试**
   - 表单提交流程
   - 编辑模式切换
   - 错误处理流程

### 集成测试

**测试场景：**
1. **首次用户流程** - 从空白表单到成功保存
2. **现有用户流程** - 加载现有数据并更新
3. **错误恢复流程** - 网络错误后的恢复
4. **路由重定向** - `/doctor-auth` 正确重定向到 `/profile`

### 用户验收测试

**测试用例：**
1. 用户点击"编辑个人信息"按钮能正确跳转
2. 首次使用用户能完整填写并保存信息
3. 现有用户能查看和编辑已有信息
4. 表单验证能正确提示错误信息
5. 保存成功后能正确跳转回"我的"页面

## 实现注意事项

### 数据兼容性

1. **API兼容性** - 确保新的统一页面能正确调用现有的API接口
2. **数据格式兼容** - 处理日期格式、选择框值等数据格式差异
3. **字段映射** - 正确处理认证页面字段到标准字段的映射

### 用户体验

1. **渐进式加载** - 优先显示基本信息，然后加载详细信息
2. **智能表单** - 根据用户输入自动填充相关字段
3. **保存提示** - 明确的保存状态反馈和成功提示

### 性能优化

1. **懒加载** - 按需加载下拉选项数据
2. **缓存策略** - 合理使用本地缓存减少API调用
3. **防抖处理** - 避免频繁的表单验证和保存操作

### 安全考虑

1. **输入验证** - 前端和后端双重验证
2. **敏感信息保护** - 避免在日志中记录敏感个人信息
3. **权限控制** - 确保用户只能修改自己的信息