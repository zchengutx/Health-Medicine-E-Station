/**
 * HTTP客户端配置
 * 基于Axios的HTTP客户端，包含请求/响应拦截器、错误处理和重试机制
 */

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosError } from 'axios'
import { API_CONFIG, ERROR_CODES, ERROR_MESSAGES } from '../constants'
import { ApiResponse, HttpError } from '../types'

/**
 * 创建HTTP客户端实例
 */
const createHttpClient = (): AxiosInstance => {
  const client = axios.create({
    baseURL: API_CONFIG.BASE_URL,
    timeout: API_CONFIG.TIMEOUT,
    headers: {
      'Content-Type': 'application/json',
    },
  })

  // 请求拦截器
  client.interceptors.request.use(
    (config) => {
      // 添加认证token（如果存在）
      const token = localStorage.getItem('auth_token')
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }

      // 添加请求时间戳
      config.metadata = { startTime: Date.now() }

      return config
    },
    (error) => {
      return Promise.reject(error)
    }
  )

  // 响应拦截器
  client.interceptors.response.use(
    (response: AxiosResponse) => {
      // 记录响应时间
      const endTime = Date.now()
      const startTime = response.config.metadata?.startTime || endTime
      console.log(`API请求耗时: ${endTime - startTime}ms`)

      return response
    },
    (error: AxiosError) => {
      return Promise.reject(transformError(error))
    }
  )

  return client
}

/**
 * 转换Axios错误为应用错误
 */
const transformError = (error: AxiosError): HttpError => {
  const httpError = new Error() as HttpError
  httpError.name = 'HttpError'

  if (error.code === 'ECONNABORTED' || error.message.includes('timeout')) {
    httpError.message = ERROR_MESSAGES[ERROR_CODES.TIMEOUT_ERROR]
    httpError.status = 408
    httpError.statusText = 'Request Timeout'
  } else if (!error.response) {
    httpError.message = ERROR_MESSAGES[ERROR_CODES.NETWORK_ERROR]
    httpError.status = 0
    httpError.statusText = 'Network Error'
  } else {
    const { status, statusText, data } = error.response
    httpError.status = status
    httpError.statusText = statusText
    httpError.response = data

    // 根据状态码设置错误信息
    switch (status) {
      case 400:
        httpError.message = (data as ApiResponse)?.message || ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS]
        break
      case 401:
        httpError.message = '认证失败，请重新登录'
        break
      case 403:
        httpError.message = '权限不足'
        break
      case 404:
        httpError.message = '请求的资源不存在'
        break
      case 429:
        httpError.message = ERROR_MESSAGES[ERROR_CODES.TOO_MANY_REQUESTS]
        break
      case 500:
        httpError.message = ERROR_MESSAGES[ERROR_CODES.SERVER_ERROR]
        break
      default:
        httpError.message = (data as ApiResponse)?.message || ERROR_MESSAGES[ERROR_CODES.UNKNOWN_ERROR]
    }
  }

  return httpError
}

/**
 * 重试机制配置
 */
interface RetryConfig {
  retries: number
  retryDelay: number
  retryCondition?: (error: HttpError) => boolean
}

/**
 * 默认重试配置
 */
const defaultRetryConfig: RetryConfig = {
  retries: API_CONFIG.RETRY_TIMES,
  retryDelay: 1000,
  retryCondition: (error: HttpError) => {
    // 只对网络错误和5xx错误进行重试
    return error.status === 0 || (error.status >= 500 && error.status < 600)
  },
}

/**
 * 带重试机制的请求函数
 */
const requestWithRetry = async <T>(
  client: AxiosInstance,
  config: AxiosRequestConfig,
  retryConfig: RetryConfig = defaultRetryConfig
): Promise<T> => {
  let lastError: HttpError

  for (let attempt = 0; attempt <= retryConfig.retries; attempt++) {
    try {
      const response = await client.request<ApiResponse<T>>(config)
      
      // 检查业务层面的成功状态
      if (response.data.success) {
        return response.data.data
      } else {
        throw new Error(response.data.message) as HttpError
      }
    } catch (error) {
      lastError = error as HttpError

      // 如果是最后一次尝试，或者不满足重试条件，直接抛出错误
      if (attempt === retryConfig.retries || !retryConfig.retryCondition?.(lastError)) {
        throw lastError
      }

      // 等待后重试
      await new Promise(resolve => setTimeout(resolve, retryConfig.retryDelay * (attempt + 1)))
      console.log(`API请求重试 ${attempt + 1}/${retryConfig.retries}`)
    }
  }

  throw lastError!
}

// 创建全局HTTP客户端实例
export const httpClient = createHttpClient()

// 导出请求方法
export const request = {
  get: <T>(url: string, config?: AxiosRequestConfig) =>
    requestWithRetry<T>(httpClient, { ...config, method: 'GET', url }),

  post: <T>(url: string, data?: unknown, config?: AxiosRequestConfig) =>
    requestWithRetry<T>(httpClient, { ...config, method: 'POST', url, data }),

  put: <T>(url: string, data?: unknown, config?: AxiosRequestConfig) =>
    requestWithRetry<T>(httpClient, { ...config, method: 'PUT', url, data }),

  delete: <T>(url: string, config?: AxiosRequestConfig) =>
    requestWithRetry<T>(httpClient, { ...config, method: 'DELETE', url }),
}

// 扩展Axios配置类型以支持metadata
declare module 'axios' {
  interface AxiosRequestConfig {
    metadata?: {
      startTime: number
    }
  }
}