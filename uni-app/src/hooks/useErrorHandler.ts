import { useState, useCallback, useRef } from 'react'
import { AppError, HttpError } from '../types'

/**
 * 错误类型枚举
 */
export enum ErrorType {
  VALIDATION_ERROR = 'VALIDATION_ERROR',
  NETWORK_ERROR = 'NETWORK_ERROR',
  BUSINESS_ERROR = 'BUSINESS_ERROR',
  SYSTEM_ERROR = 'SYSTEM_ERROR'
}

/**
 * 错误处理Hook
 * 统一处理应用中的各种错误类型
 */
export interface UseErrorHandlerReturn {
  /** 当前错误信息 */
  error: string | null
  /** 错误类型 */
  errorType: ErrorType | null
  /** 是否有错误 */
  hasError: boolean
  /** 处理错误 */
  handleError: (error: Error | AppError | HttpError, type?: ErrorType) => void
  /** 显示错误信息 */
  showError: (message: string, type?: ErrorType) => void
  /** 清除错误 */
  clearError: () => void
  /** 重试函数 */
  retry: (() => void) | null
  /** 设置重试函数 */
  setRetry: (retryFn: (() => void) | null) => void
}

/**
 * 根据错误类型和内容生成用户友好的错误信息
 */
const getErrorMessage = (error: Error | AppError | HttpError, type?: ErrorType): string => {
  // 如果是AppError类型，直接使用其message
  if ('code' in error && 'message' in error) {
    return error.message
  }

  // 如果是HttpError类型，根据状态码返回相应信息
  if ('status' in error) {
    const httpError = error as HttpError
    switch (httpError.status) {
      case 400:
        return '请求参数错误，请检查输入信息'
      case 401:
        return '登录已过期，请重新登录'
      case 403:
        return '没有权限执行此操作'
      case 404:
        return '请求的资源不存在'
      case 429:
        return '请求过于频繁，请稍后再试'
      case 500:
        return '服务器内部错误，请稍后重试'
      case 502:
      case 503:
      case 504:
        return '服务暂时不可用，请稍后重试'
      default:
        return '网络请求失败，请检查网络连接'
    }
  }

  // 根据错误类型返回默认信息
  if (type) {
    switch (type) {
      case ErrorType.VALIDATION_ERROR:
        return error.message || '输入信息格式错误'
      case ErrorType.NETWORK_ERROR:
        return error.message.includes('timeout') ? '请求超时，请重试' : '网络连接失败，请检查网络后重试'
      case ErrorType.BUSINESS_ERROR:
        return error.message || '操作失败，请重试'
      case ErrorType.SYSTEM_ERROR:
        // For system errors, return original message if it's a simple error
        return error.message || '系统错误，请稍后重试'
    }
  }

  // 根据错误信息内容判断错误类型
  const message = error.message.toLowerCase()
  
  if (message.includes('network') || message.includes('fetch')) {
    return '网络连接失败，请检查网络后重试'
  }
  
  if (message.includes('timeout')) {
    return '请求超时，请重试'
  }
  
  if (message.includes('手机号') || message.includes('验证码')) {
    return error.message
  }
  
  // 默认返回原始错误信息
  return error.message || '操作失败，请重试'
}

/**
 * 根据错误内容自动判断错误类型
 */
const detectErrorType = (error: Error | AppError | HttpError): ErrorType => {
  // HttpError类型
  if ('status' in error) {
    const status = (error as HttpError).status
    if (status >= 400 && status < 500) {
      return ErrorType.BUSINESS_ERROR
    }
    if (status >= 500) {
      return ErrorType.SYSTEM_ERROR
    }
    return ErrorType.NETWORK_ERROR
  }

  // AppError类型
  if ('code' in error) {
    const code = (error as AppError).code
    if (code.includes('VALIDATION') || code.includes('INVALID')) {
      return ErrorType.VALIDATION_ERROR
    }
    if (code.includes('NETWORK') || code.includes('TIMEOUT')) {
      return ErrorType.NETWORK_ERROR
    }
    return ErrorType.BUSINESS_ERROR
  }

  // 根据错误信息判断
  const message = error.message.toLowerCase()
  
  if (message.includes('手机号') || message.includes('验证码') || message.includes('格式')) {
    return ErrorType.VALIDATION_ERROR
  }
  
  if (message.includes('network') || message.includes('fetch') || message.includes('timeout')) {
    return ErrorType.NETWORK_ERROR
  }
  
  return ErrorType.SYSTEM_ERROR
}

export const useErrorHandler = (): UseErrorHandlerReturn => {
  const [error, setError] = useState<string | null>(null)
  const [errorType, setErrorType] = useState<ErrorType | null>(null)
  const [retry, setRetryState] = useState<(() => void) | null>(null)

  // 处理错误
  const handleError = useCallback((error: Error | AppError | HttpError, type?: ErrorType) => {
    const detectedType = type || detectErrorType(error)
    const errorMessage = getErrorMessage(error, detectedType)
    
    setError(errorMessage)
    setErrorType(detectedType)
    
    // 在开发环境下打印详细错误信息
    if (process.env.NODE_ENV === 'development') {
      console.error('Error handled:', {
        error,
        type: detectedType,
        message: errorMessage,
      })
    }
  }, [])

  // 显示错误信息
  const showError = useCallback((message: string, type = ErrorType.SYSTEM_ERROR) => {
    setError(message)
    setErrorType(type)
  }, [])

  // 清除错误
  const clearError = useCallback(() => {
    setError(null)
    setErrorType(null)
    setRetryState(null)
  }, [])

  // 设置重试函数
  const setRetry = useCallback((retryFn: (() => void) | null) => {
    setRetryState(() => retryFn)
  }, [])

  return {
    error,
    errorType,
    hasError: error !== null,
    handleError,
    showError,
    clearError,
    retry,
    setRetry,
  }
}