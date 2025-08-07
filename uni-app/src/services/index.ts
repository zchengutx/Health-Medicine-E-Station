/**
 * API服务层统一导出
 */

// HTTP客户端
export { httpClient, request } from './httpClient'

// 认证相关API
export {
  sendVerificationCode,
  loginWithVerificationCode,
  refreshAuthToken,
  logout,
  validateToken,
} from './authApi'

// 使用示例
export {
  exampleLoginFlow,
  exampleErrorHandling,
  exampleLogoutFlow,
  exampleCheckLoginStatus,
} from './example'

// 导出类型
export type { HttpError } from '../types'