import { showToast, showDialog } from 'vant'

// 错误类型枚举
export enum ErrorType {
  NETWORK = 'NETWORK',
  API = 'API',
  VALIDATION = 'VALIDATION',
  AUTH = 'AUTH',
  UNKNOWN = 'UNKNOWN'
}

// 错误级别枚举
export enum ErrorLevel {
  INFO = 'INFO',
  WARNING = 'WARNING',
  ERROR = 'ERROR',
  CRITICAL = 'CRITICAL'
}

// 错误信息接口
export interface ErrorInfo {
  type: ErrorType
  level: ErrorLevel
  message: string
  code?: string | number
  details?: any
  timestamp: number
  stack?: string
}

// 错误处理器类
export class GlobalErrorHandler {
  private static instance: GlobalErrorHandler
  private errorLog: ErrorInfo[] = []
  private maxLogSize = 100

  private constructor() {
    this.setupGlobalHandlers()
  }

  static getInstance(): GlobalErrorHandler {
    if (!GlobalErrorHandler.instance) {
      GlobalErrorHandler.instance = new GlobalErrorHandler()
    }
    return GlobalErrorHandler.instance
  }

  // 设置全局错误处理器
  private setupGlobalHandlers() {
    // 捕获未处理的Promise错误
    window.addEventListener('unhandledrejection', (event) => {
      console.error('Unhandled promise rejection:', event.reason)
      this.handleError({
        type: ErrorType.UNKNOWN,
        level: ErrorLevel.ERROR,
        message: '系统发生未知错误',
        details: event.reason,
        timestamp: Date.now(),
        stack: event.reason?.stack
      })
      event.preventDefault()
    })

    // 捕获JavaScript运行时错误
    window.addEventListener('error', (event) => {
      console.error('JavaScript error:', event.error)
      this.handleError({
        type: ErrorType.UNKNOWN,
        level: ErrorLevel.ERROR,
        message: event.message || '脚本执行错误',
        details: {
          filename: event.filename,
          lineno: event.lineno,
          colno: event.colno
        },
        timestamp: Date.now(),
        stack: event.error?.stack
      })
    })

    // 捕获资源加载错误
    window.addEventListener('error', (event) => {
      if (event.target !== window) {
        console.error('Resource loading error:', event.target)
        this.handleError({
          type: ErrorType.NETWORK,
          level: ErrorLevel.WARNING,
          message: '资源加载失败',
          details: {
            tagName: (event.target as any)?.tagName,
            src: (event.target as any)?.src || (event.target as any)?.href
          },
          timestamp: Date.now()
        })
      }
    }, true)
  }

  // 处理错误
  handleError(errorInfo: ErrorInfo) {
    // 添加到错误日志
    this.addToLog(errorInfo)

    // 根据错误级别和类型决定处理方式
    switch (errorInfo.level) {
      case ErrorLevel.INFO:
        this.showInfoMessage(errorInfo.message)
        break
      case ErrorLevel.WARNING:
        this.showWarningMessage(errorInfo.message)
        break
      case ErrorLevel.ERROR:
        this.showErrorMessage(errorInfo.message)
        break
      case ErrorLevel.CRITICAL:
        this.showCriticalError(errorInfo)
        break
    }

    // 上报错误（在生产环境中）
    if (process.env.NODE_ENV === 'production') {
      this.reportError(errorInfo)
    }
  }

  // 处理API错误
  handleApiError(error: any): string {
    let errorInfo: ErrorInfo

    if (error.response) {
      // HTTP错误响应
      const { status, data } = error.response
      errorInfo = {
        type: ErrorType.API,
        level: this.getErrorLevelByStatus(status),
        message: this.getApiErrorMessage(status, data?.Message),
        code: status,
        details: data,
        timestamp: Date.now()
      }
    } else if (error.request) {
      // 网络错误
      errorInfo = {
        type: ErrorType.NETWORK,
        level: ErrorLevel.ERROR,
        message: '网络连接失败，请检查网络设置',
        details: error.request,
        timestamp: Date.now()
      }
    } else {
      // 其他错误
      errorInfo = {
        type: ErrorType.UNKNOWN,
        level: ErrorLevel.ERROR,
        message: error.message || '请求失败，请稍后重试',
        details: error,
        timestamp: Date.now(),
        stack: error.stack
      }
    }

    this.handleError(errorInfo)
    return errorInfo.message
  }

  // 处理验证错误
  handleValidationError(field: string, message: string) {
    const errorInfo: ErrorInfo = {
      type: ErrorType.VALIDATION,
      level: ErrorLevel.WARNING,
      message: `${field}: ${message}`,
      timestamp: Date.now()
    }
    
    this.handleError(errorInfo)
  }

  // 处理认证错误
  handleAuthError(message: string = '登录已过期，请重新登录') {
    const errorInfo: ErrorInfo = {
      type: ErrorType.AUTH,
      level: ErrorLevel.ERROR,
      message,
      timestamp: Date.now()
    }
    
    this.handleError(errorInfo)
    
    // 清除登录状态并跳转到登录页
    setTimeout(() => {
      localStorage.removeItem('doctor_token')
      localStorage.removeItem('doctor_info')
      window.location.href = '/login'
    }, 1500)
  }

  // 根据HTTP状态码获取错误级别
  private getErrorLevelByStatus(status: number): ErrorLevel {
    if (status >= 500) return ErrorLevel.CRITICAL
    if (status >= 400) return ErrorLevel.ERROR
    if (status >= 300) return ErrorLevel.WARNING
    return ErrorLevel.INFO
  }

  // 获取API错误消息
  private getApiErrorMessage(status: number, message?: string): string {
    const defaultMessages: Record<number, string> = {
      400: '请求参数错误',
      401: '登录已过期，请重新登录',
      403: '没有权限访问',
      404: '请求的资源不存在',
      408: '请求超时',
      429: '请求过于频繁，请稍后再试',
      500: '服务器内部错误',
      502: '网关错误',
      503: '服务暂不可用',
      504: '网关超时'
    }
    
    return message || defaultMessages[status] || `请求失败 (${status})`
  }

  // 显示信息消息
  private showInfoMessage(message: string) {
    showToast({
      message,
      type: 'success',
      duration: 2000
    })
  }

  // 显示警告消息
  private showWarningMessage(message: string) {
    showToast({
      message,
      type: 'fail',
      duration: 3000
    })
  }

  // 显示错误消息
  private showErrorMessage(message: string) {
    showToast({
      message,
      type: 'fail',
      duration: 4000
    })
  }

  // 显示严重错误
  private showCriticalError(errorInfo: ErrorInfo) {
    showDialog({
      title: '系统错误',
      message: `${errorInfo.message}\n\n如果问题持续存在，请联系技术支持。`,
      confirmButtonText: '确定',
      showCancelButton: false
    })
  }

  // 添加到错误日志
  private addToLog(errorInfo: ErrorInfo) {
    this.errorLog.unshift(errorInfo)
    
    // 限制日志大小
    if (this.errorLog.length > this.maxLogSize) {
      this.errorLog = this.errorLog.slice(0, this.maxLogSize)
    }
    
    // 存储到本地（用于调试）
    if (process.env.NODE_ENV === 'development') {
      localStorage.setItem('error_log', JSON.stringify(this.errorLog.slice(0, 10)))
    }
  }

  // 上报错误到服务器
  private async reportError(errorInfo: ErrorInfo) {
    try {
      // 这里可以调用错误上报API
      // await fetch('/api/error-report', {
      //   method: 'POST',
      //   headers: { 'Content-Type': 'application/json' },
      //   body: JSON.stringify(errorInfo)
      // })
      console.log('Error reported:', errorInfo)
    } catch (error) {
      console.error('Failed to report error:', error)
    }
  }

  // 获取错误日志
  getErrorLog(): ErrorInfo[] {
    return [...this.errorLog]
  }

  // 清除错误日志
  clearErrorLog() {
    this.errorLog = []
    localStorage.removeItem('error_log')
  }

  // 导出错误日志
  exportErrorLog(): string {
    return JSON.stringify(this.errorLog, null, 2)
  }
}

// 导出单例实例
export const errorHandler = GlobalErrorHandler.getInstance()

// 便捷方法
export const handleError = (error: any, type: ErrorType = ErrorType.UNKNOWN) => {
  if (type === ErrorType.API) {
    return errorHandler.handleApiError(error)
  } else {
    errorHandler.handleError({
      type,
      level: ErrorLevel.ERROR,
      message: error.message || '发生未知错误',
      details: error,
      timestamp: Date.now(),
      stack: error.stack
    })
    return error.message || '发生未知错误'
  }
}

export const handleApiError = (error: any) => {
  return errorHandler.handleApiError(error)
}

export const handleValidationError = (field: string, message: string) => {
  errorHandler.handleValidationError(field, message)
}

export const handleAuthError = (message?: string) => {
  errorHandler.handleAuthError(message)
}