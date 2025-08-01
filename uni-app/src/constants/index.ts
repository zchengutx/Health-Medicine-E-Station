/**
 * 应用常量定义
 */

// API相关常量
export const API_CONFIG = {
  BASE_URL: 'http://localhost:3001/api', // 可以通过环境变量配置
  TIMEOUT: 10000, // 10秒超时
  RETRY_TIMES: 3, // 重试次数
} as const

// 验证码相关常量
export const VERIFICATION_CODE = {
  LENGTH: 6, // 验证码长度
  COUNTDOWN: 60, // 倒计时秒数
  RESEND_LIMIT: 5, // 每日重发限制
} as const

// 手机号相关常量
export const PHONE = {
  LENGTH: 11, // 手机号长度
  PREFIX_REGEX: /^1[3-9]/, // 手机号前缀正则
} as const

// 表单验证相关常量
export const VALIDATION = {
  DEBOUNCE_DELAY: 300, // 防抖延迟时间（毫秒）
  MIN_PHONE_LENGTH: 11,
  MAX_PHONE_LENGTH: 11,
  MIN_CODE_LENGTH: 6,
  MAX_CODE_LENGTH: 6,
} as const

// 错误码定义
export const ERROR_CODES = {
  // 网络错误
  NETWORK_ERROR: 'NETWORK_ERROR',
  TIMEOUT_ERROR: 'TIMEOUT_ERROR',

  // 参数错误
  INVALID_PHONE: 'INVALID_PHONE',
  INVALID_CODE: 'INVALID_CODE',
  MISSING_PARAMS: 'MISSING_PARAMS',

  // 业务错误
  CODE_EXPIRED: 'CODE_EXPIRED',
  CODE_INVALID: 'CODE_INVALID',
  PHONE_NOT_REGISTERED: 'PHONE_NOT_REGISTERED',
  TOO_MANY_REQUESTS: 'TOO_MANY_REQUESTS',

  // 系统错误
  SERVER_ERROR: 'SERVER_ERROR',
  UNKNOWN_ERROR: 'UNKNOWN_ERROR',
} as const

// 错误信息映射
export const ERROR_MESSAGES = {
  [ERROR_CODES.NETWORK_ERROR]: '网络连接失败，请检查网络设置',
  [ERROR_CODES.TIMEOUT_ERROR]: '请求超时，请稍后重试',
  [ERROR_CODES.INVALID_PHONE]: '请输入正确的手机号',
  [ERROR_CODES.INVALID_CODE]: '请输入正确的验证码',
  [ERROR_CODES.MISSING_PARAMS]: '参数不完整',
  [ERROR_CODES.CODE_EXPIRED]: '验证码已过期，请重新获取',
  [ERROR_CODES.CODE_INVALID]: '验证码错误，请重新输入',
  [ERROR_CODES.PHONE_NOT_REGISTERED]: '手机号未注册',
  [ERROR_CODES.TOO_MANY_REQUESTS]: '请求过于频繁，请稍后重试',
  [ERROR_CODES.SERVER_ERROR]: '服务器错误，请稍后重试',
  [ERROR_CODES.UNKNOWN_ERROR]: '未知错误，请稍后重试',
} as const

// 快捷登录选项
export const QUICK_LOGIN_OPTIONS = [
  {
    id: 'wechat',
    name: '微信',
    icon: '🔗',
    color: '#07C160',
  },
  {
    id: 'qq',
    name: 'QQ',
    icon: '🔗',
    color: '#12B7F5',
  },
  {
    id: 'alipay',
    name: '支付宝',
    icon: '🔗',
    color: '#1677FF',
  },
] as const

// 用户协议配置
export const USER_AGREEMENTS = [
  {
    id: 'terms',
    title: '用户服务协议',
    url: '/terms',
    required: true,
  },
  {
    id: 'privacy',
    title: '隐私政策',
    url: '/privacy',
    required: true,
  },
] as const

// 本地存储键名
export const STORAGE_KEYS = {
  TOKEN: 'auth_token',
  REFRESH_TOKEN: 'refresh_token',
  USER_INFO: 'user_info',
  PHONE_HISTORY: 'phone_history',
  AGREEMENT_ACCEPTED: 'agreement_accepted',
} as const

// 页面路由
export const ROUTES = {
  LOGIN: '/login',
  HOME: '/',
  TERMS: '/terms',
  PRIVACY: '/privacy',
} as const
