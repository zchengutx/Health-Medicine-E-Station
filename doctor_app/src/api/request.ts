import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { errorHandler, ErrorType } from '@/utils/errorHandler'
import { log } from '@/utils/logger'

// API响应接口
export interface ApiResponse<T = any> {
  Message: string
  Code: number | string
  data?: T
}

// 创建axios实例
const request: AxiosInstance = axios.create({
  baseURL: '/api', // 使用代理
  timeout: 15000, // 增加超时时间
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
request.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    // 添加token到请求头
    const token = localStorage.getItem('doctor_token')
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    // 添加请求时间戳，防止缓存
    if (config.method === 'get') {
      config.params = {
        ...config.params,
        _t: Date.now()
      }
    }
    
    log.debug('API请求', {
      method: config.method?.toUpperCase(),
      url: config.url,
      hasData: !!(config.data || config.params)
    })
    return config
  },
  (error) => {
    log.error('请求拦截器错误', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse<ApiResponse>) => {
    const { data } = response
    
    log.debug('API响应', {
      url: response.config.url,
      status: response.status,
      code: data.Code,
      message: data.Message
    })
    
    // 检查业务状态码（支持字符串和数字类型）
    const code = typeof data.Code === 'string' ? parseInt(data.Code) : data.Code
    if (code === 200 || code === 0) {
      return data
    } else {
      // 业务错误
      const error = new Error(data.Message || '请求失败')
      error.name = 'BusinessError'
      return Promise.reject(error)
    }
  },
  (error) => {
    log.error('响应拦截器错误', error)
    
    // 使用全局错误处理器
    const errorMessage = errorHandler.handleApiError(error)
    
    // 特殊处理401未授权
    if (error.response?.status === 401) {
      // 清除本地存储
      localStorage.removeItem('doctor_token')
      localStorage.removeItem('doctor_info')
      
      // 延迟跳转，避免在某些情况下路由跳转失败
      setTimeout(() => {
        if (window.location.pathname !== '/login') {
          window.location.href = '/login'
        }
      }, 1500)
    }
    
    // 设置错误信息
    error.message = errorMessage
    
    return Promise.reject(error)
  }
)

export default request