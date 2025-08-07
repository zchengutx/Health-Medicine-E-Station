/**
 * 认证相关API服务
 * 包含发送验证码和登录验证的API调用函数
 */

import { request } from './httpClient'
import { 
  SendCodeRequest, 
  SendCodeResponse, 
  LoginRequest, 
  LoginResponse 
} from '../types'
import { isValidPhone, isValidVerificationCode } from '../utils/validation'
import { ERROR_CODES, ERROR_MESSAGES } from '../constants'

/**
 * 发送验证码API
 * @param params 发送验证码请求参数
 * @returns Promise<SendCodeResponse>
 */
export const sendVerificationCode = async (params: SendCodeRequest): Promise<SendCodeResponse> => {
  // 参数验证
  if (!params.phone) {
    throw new Error(ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS])
  }

  if (!isValidPhone(params.phone)) {
    throw new Error(ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE])
  }

  try {
    const response = await request.post<SendCodeResponse>('/auth/send-code', {
      phone: params.phone,
      type: 'login', // 登录类型验证码
    })

    return response
  } catch (error) {
    // 处理特定的业务错误
    if (error instanceof Error) {
      if (error.message.includes('频繁')) {
        throw new Error(ERROR_MESSAGES[ERROR_CODES.TOO_MANY_REQUESTS])
      }
      if (error.message.includes('手机号')) {
        throw new Error(ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE])
      }
    }
    
    throw error
  }
}

/**
 * 登录验证API
 * @param params 登录请求参数
 * @returns Promise<LoginResponse>
 */
export const loginWithVerificationCode = async (params: LoginRequest): Promise<LoginResponse> => {
  // 参数验证
  if (!params.phone || !params.verificationCode) {
    throw new Error(ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS])
  }

  if (!isValidPhone(params.phone)) {
    throw new Error(ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE])
  }

  if (!isValidVerificationCode(params.verificationCode)) {
    throw new Error(ERROR_MESSAGES[ERROR_CODES.INVALID_CODE])
  }

  try {
    const response = await request.post<LoginResponse>('/auth/login', {
      phone: params.phone,
      verificationCode: params.verificationCode,
    })

    // 登录成功后保存token
    if (response.token) {
      localStorage.setItem('auth_token', response.token)
      localStorage.setItem('refresh_token', response.refreshToken)
      localStorage.setItem('user_info', JSON.stringify(response.user))
    }

    return response
  } catch (error) {
    // 处理特定的业务错误
    if (error instanceof Error) {
      if (error.message.includes('验证码')) {
        throw new Error(ERROR_MESSAGES[ERROR_CODES.CODE_INVALID])
      }
      if (error.message.includes('过期')) {
        throw new Error(ERROR_MESSAGES[ERROR_CODES.CODE_EXPIRED])
      }
      if (error.message.includes('手机号')) {
        throw new Error(ERROR_MESSAGES[ERROR_CODES.PHONE_NOT_REGISTERED])
      }
    }
    
    throw error
  }
}

/**
 * 刷新token API
 * @param refreshToken 刷新token
 * @returns Promise<{token: string, refreshToken: string}>
 */
export const refreshAuthToken = async (refreshToken: string): Promise<{
  token: string
  refreshToken: string
  expiresIn: number
}> => {
  if (!refreshToken) {
    throw new Error('刷新token不能为空')
  }

  try {
    const response = await request.post<{
      token: string
      refreshToken: string
      expiresIn: number
    }>('/auth/refresh', {
      refreshToken,
    })

    // 更新本地存储的token
    localStorage.setItem('auth_token', response.token)
    localStorage.setItem('refresh_token', response.refreshToken)

    return response
  } catch (error) {
    // 刷新失败，清除本地存储
    localStorage.removeItem('auth_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user_info')
    
    throw error
  }
}

/**
 * 登出API
 * @returns Promise<void>
 */
export const logout = async (): Promise<void> => {
  try {
    await request.post('/auth/logout')
  } catch (error) {
    // 即使登出API失败，也要清除本地存储
    console.warn('登出API调用失败:', error)
  } finally {
    // 清除本地存储
    localStorage.removeItem('auth_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user_info')
  }
}

/**
 * 检查token是否有效
 * @returns Promise<boolean>
 */
export const validateToken = async (): Promise<boolean> => {
  const token = localStorage.getItem('auth_token')
  
  if (!token) {
    return false
  }

  try {
    await request.get('/auth/validate')
    return true
  } catch (error) {
    // token无效，清除本地存储
    localStorage.removeItem('auth_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user_info')
    return false
  }
}