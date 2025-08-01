// 用户相关类型定义
export interface User {
  id: string
  phone: string
  nickname?: string
  avatar?: string
  createdAt: string
  updatedAt: string
}

// 登录表单数据类型
export interface LoginFormData {
  phone: string
  verificationCode: string
}

// API响应基础类型
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
  success: boolean
}

// 发送验证码请求参数
export interface SendCodeRequest {
  phone: string
}

// 发送验证码响应数据
export interface SendCodeResponse {
  success: boolean
  message: string
  countdown: number // 倒计时秒数
}

// 登录请求参数
export interface LoginRequest {
  phone: string
  verificationCode: string
}

// 登录响应数据
export interface LoginResponse {
  user: User
  token: string
  refreshToken: string
  expiresIn: number
}

// 表单验证错误类型
export interface FormErrors {
  phone?: string
  verificationCode?: string
}

// 验证码状态类型
export interface VerificationCodeState {
  countdown: number
  isCountingDown: boolean
  canSend: boolean
}

// 登录表单状态类型
export interface LoginFormState {
  formData: LoginFormData
  errors: FormErrors
  isLoading: boolean
  isValid: boolean
}

// 快捷登录选项类型
export interface QuickLoginOption {
  id: string
  name: string
  icon: string
  color: string
}

// 用户协议类型
export interface UserAgreement {
  id: string
  title: string
  url: string
  required: boolean
}

// 错误类型定义
export interface AppError {
  code: string
  message: string
  details?: unknown
}

// HTTP错误类型
export interface HttpError extends Error {
  status: number
  statusText: string
  response?: unknown
}
