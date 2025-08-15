import { log } from './logger'

// 错误类型枚举
export enum ErrorType {
  NETWORK = 'network',
  AUTHENTICATION = 'authentication',
  VALIDATION = 'validation',
  SERVER = 'server',
  UNKNOWN = 'unknown'
}

// 错误严重程度
export enum ErrorSeverity {
  LOW = 'low',
  MEDIUM = 'medium',
  HIGH = 'high',
  CRITICAL = 'critical'
}

// 错误信息接口
export interface ErrorInfo {
  type: ErrorType
  severity: ErrorSeverity
  message: string
  userMessage: string
  code?: string | number
  details?: any
  timestamp: number
  retryable: boolean
  fallbackAvailable: boolean
}

// 错误处理配置
interface ErrorHandlerConfig {
  maxRetries: number
  retryDelay: number
  enableFallback: boolean
  enableDiagnostics: boolean
}

// 默认配置
const defaultConfig: ErrorHandlerConfig = {
  maxRetries: 3,
  retryDelay: 1000,
  enableFallback: true,
  enableDiagnostics: true
}

/**
 * 统一错误处理类
 */
export class ErrorHandler {
  private config: ErrorHandlerConfig
  private retryCount: Map<string, number> = new Map()

  constructor(config: Partial<ErrorHandlerConfig> = {}) {
    this.config = { ...defaultConfig, ...config }
  }

  /**
   * 分析错误并返回错误信息
   */
  analyzeError(error: any, context?: string): ErrorInfo {
    const timestamp = Date.now()
    let errorInfo: ErrorInfo

    // 网络错误
    if (this.isNetworkError(error)) {
      errorInfo = {
        type: ErrorType.NETWORK,
        severity: ErrorSeverity.MEDIUM,
        message: error.message || '网络连接失败',
        userMessage: '网络连接不稳定，请检查网络设置后重试',
        timestamp,
        retryable: true,
        fallbackAvailable: true
      }
    }
    // 认证错误
    else if (this.isAuthenticationError(error)) {
      errorInfo = {
        type: ErrorType.AUTHENTICATION,
        severity: ErrorSeverity.HIGH,
        message: error.message || '认证失败',
        userMessage: '登录状态已过期，请重新登录',
        timestamp,
        retryable: false,
        fallbackAvailable: false
      }
    }
    // 验证错误
    else if (this.isValidationError(error)) {
      errorInfo = {
        type: ErrorType.VALIDATION,
        severity: ErrorSeverity.LOW,
        message: error.message || '数据验证失败',
        userMessage: error.message || '请检查输入的信息是否正确',
        timestamp,
        retryable: false,
        fallbackAvailable: false
      }
    }
    // 服务器错误
    else if (this.isServerError(error)) {
      errorInfo = {
        type: ErrorType.SERVER,
        severity: ErrorSeverity.HIGH,
        message: error.message || '服务器内部错误',
        userMessage: '服务器暂时不可用，请稍后重试',
        code: error.status || error.code,
        timestamp,
        retryable: true,
        fallbackAvailable: true
      }
    }
    // 未知错误
    else {
      errorInfo = {
        type: ErrorType.UNKNOWN,
        severity: ErrorSeverity.MEDIUM,
        message: error.message || '未知错误',
        userMessage: '操作失败，请重试',
        timestamp,
        retryable: true,
        fallbackAvailable: true
      }
    }

    // 添加上下文信息
    if (context) {
      errorInfo.details = { context, originalError: error }
    }

    // 记录错误日志
    this.logError(errorInfo, context)

    return errorInfo
  }

  /**
   * 自动重试机制
   */
  async retryWithBackoff<T>(
    operation: () => Promise<T>,
    operationId: string,
    maxRetries?: number
  ): Promise<T> {
    const retries = maxRetries || this.config.maxRetries
    let lastError: any

    for (let attempt = 0; attempt <= retries; attempt++) {
      try {
        const result = await operation()
        // 成功后清除重试计数
        this.retryCount.delete(operationId)
        return result
      } catch (error) {
        lastError = error
        
        if (attempt === retries) {
          // 达到最大重试次数
          this.retryCount.delete(operationId)
          throw error
        }

        const errorInfo = this.analyzeError(error)
        if (!errorInfo.retryable) {
          // 不可重试的错误
          throw error
        }

        // 记录重试次数
        this.retryCount.set(operationId, attempt + 1)
        
        // 指数退避延迟
        const delay = this.config.retryDelay * Math.pow(2, attempt)
        log.warn(`操作失败，${delay}ms后进行第${attempt + 1}次重试`, {
          operationId,
          attempt: attempt + 1,
          maxRetries: retries,
          error: error.message
        })
        
        await this.delay(delay)
      }
    }

    throw lastError
  }

  /**
   * 获取用户友好的错误消息
   */
  getUserMessage(error: any, context?: string): string {
    const errorInfo = this.analyzeError(error, context)
    return errorInfo.userMessage
  }

  /**
   * 检查是否可以重试
   */
  canRetry(error: any): boolean {
    const errorInfo = this.analyzeError(error)
    return errorInfo.retryable
  }

  /**
   * 检查是否有备用方案
   */
  hasFallback(error: any): boolean {
    const errorInfo = this.analyzeError(error)
    return errorInfo.fallbackAvailable
  }

  /**
   * 获取重试次数
   */
  getRetryCount(operationId: string): number {
    return this.retryCount.get(operationId) || 0
  }

  // 私有方法

  private isNetworkError(error: any): boolean {
    return (
      error.code === 'NETWORK_ERROR' ||
      error.message?.includes('网络') ||
      error.message?.includes('network') ||
      error.message?.includes('timeout') ||
      error.name === 'NetworkError'
    )
  }

  private isAuthenticationError(error: any): boolean {
    return (
      error.status === 401 ||
      error.code === 401 ||
      error.message?.includes('unauthorized') ||
      error.message?.includes('认证') ||
      error.message?.includes('登录')
    )
  }

  private isValidationError(error: any): boolean {
    return (
      error.status === 400 ||
      error.code === 400 ||
      error.message?.includes('validation') ||
      error.message?.includes('验证') ||
      error.name === 'ValidationError'
    )
  }

  private isServerError(error: any): boolean {
    return (
      (error.status >= 500 && error.status < 600) ||
      (error.code >= 500 && error.code < 600) ||
      error.message?.includes('服务器') ||
      error.message?.includes('server')
    )
  }

  private logError(errorInfo: ErrorInfo, context?: string): void {
    const logData = {
      type: errorInfo.type,
      severity: errorInfo.severity,
      message: errorInfo.message,
      code: errorInfo.code,
      context,
      timestamp: errorInfo.timestamp,
      retryable: errorInfo.retryable,
      fallbackAvailable: errorInfo.fallbackAvailable
    }

    switch (errorInfo.severity) {
      case ErrorSeverity.CRITICAL:
        log.error('严重错误', logData)
        break
      case ErrorSeverity.HIGH:
        log.error('高级错误', logData)
        break
      case ErrorSeverity.MEDIUM:
        log.warn('中级错误', logData)
        break
      case ErrorSeverity.LOW:
        log.info('低级错误', logData)
        break
    }
  }

  private delay(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms))
  }
}

// 导出默认实例
export const errorHandler = new ErrorHandler()

// 导出便捷函数
export const analyzeError = (error: any, context?: string) => 
  errorHandler.analyzeError(error, context)

export const getUserMessage = (error: any, context?: string) => 
  errorHandler.getUserMessage(error, context)

export const retryWithBackoff = <T>(
  operation: () => Promise<T>,
  operationId: string,
  maxRetries?: number
) => errorHandler.retryWithBackoff(operation, operationId, maxRetries)