import { renderHook, act } from '@testing-library/react'
import { vi } from 'vitest'
import { useErrorHandler, ErrorType } from '../useErrorHandler'
import { AppError, HttpError } from '../../types'

describe('useErrorHandler', () => {
  it('should initialize with no error', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    expect(result.current.error).toBeNull()
    expect(result.current.errorType).toBeNull()
    expect(result.current.hasError).toBe(false)
    expect(result.current.retry).toBeNull()
  })

  it('should handle basic Error objects', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    act(() => {
      result.current.handleError(new Error('Test error'))
    })
    
    expect(result.current.error).toBe('Test error')
    expect(result.current.errorType).toBe(ErrorType.SYSTEM_ERROR)
    expect(result.current.hasError).toBe(true)
  })

  it('should handle AppError objects', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    const appError: AppError = {
      code: 'VALIDATION_ERROR',
      message: 'Invalid input',
    }
    
    act(() => {
      result.current.handleError(appError as any)
    })
    
    expect(result.current.error).toBe('Invalid input')
    expect(result.current.errorType).toBe(ErrorType.VALIDATION_ERROR)
  })

  it('should handle HttpError objects with different status codes', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    // Test 400 error
    const httpError400: HttpError = {
      name: 'HttpError',
      message: 'Bad Request',
      status: 400,
      statusText: 'Bad Request',
    }
    
    act(() => {
      result.current.handleError(httpError400)
    })
    
    expect(result.current.error).toBe('请求参数错误，请检查输入信息')
    expect(result.current.errorType).toBe(ErrorType.BUSINESS_ERROR)
    
    // Test 500 error
    const httpError500: HttpError = {
      name: 'HttpError',
      message: 'Internal Server Error',
      status: 500,
      statusText: 'Internal Server Error',
    }
    
    act(() => {
      result.current.handleError(httpError500)
    })
    
    expect(result.current.error).toBe('服务器内部错误，请稍后重试')
    expect(result.current.errorType).toBe(ErrorType.SYSTEM_ERROR)
  })

  it('should detect error types automatically', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    // Network error
    act(() => {
      result.current.handleError(new Error('Network request failed'))
    })
    expect(result.current.errorType).toBe(ErrorType.NETWORK_ERROR)
    
    // Validation error
    act(() => {
      result.current.handleError(new Error('手机号格式错误'))
    })
    expect(result.current.errorType).toBe(ErrorType.VALIDATION_ERROR)
  })

  it('should allow manual error type specification', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    act(() => {
      result.current.handleError(new Error('Test error'), ErrorType.BUSINESS_ERROR)
    })
    
    expect(result.current.errorType).toBe(ErrorType.BUSINESS_ERROR)
  })

  it('should show custom error messages', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    act(() => {
      result.current.showError('Custom error message', ErrorType.VALIDATION_ERROR)
    })
    
    expect(result.current.error).toBe('Custom error message')
    expect(result.current.errorType).toBe(ErrorType.VALIDATION_ERROR)
    expect(result.current.hasError).toBe(true)
  })

  it('should clear errors', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    act(() => {
      result.current.showError('Test error')
    })
    
    expect(result.current.hasError).toBe(true)
    
    act(() => {
      result.current.clearError()
    })
    
    expect(result.current.error).toBeNull()
    expect(result.current.errorType).toBeNull()
    expect(result.current.hasError).toBe(false)
    expect(result.current.retry).toBeNull()
  })

  it('should handle retry functionality', () => {
    const { result } = renderHook(() => useErrorHandler())
    const retryFn = vi.fn()
    
    act(() => {
      result.current.setRetry(retryFn)
    })
    
    expect(result.current.retry).toBe(retryFn)
    
    act(() => {
      result.current.clearError()
    })
    
    expect(result.current.retry).toBeNull()
  })

  it('should provide user-friendly messages for common HTTP status codes', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    const testCases = [
      { status: 401, expected: '登录已过期，请重新登录' },
      { status: 403, expected: '没有权限执行此操作' },
      { status: 404, expected: '请求的资源不存在' },
      { status: 429, expected: '请求过于频繁，请稍后再试' },
      { status: 502, expected: '服务暂时不可用，请稍后重试' },
      { status: 503, expected: '服务暂时不可用，请稍后重试' },
      { status: 504, expected: '服务暂时不可用，请稍后重试' },
    ]
    
    testCases.forEach(({ status, expected }) => {
      const httpError: HttpError = {
        name: 'HttpError',
        message: 'HTTP Error',
        status,
        statusText: 'Error',
      }
      
      act(() => {
        result.current.handleError(httpError)
      })
      
      expect(result.current.error).toBe(expected)
    })
  })

  it('should handle timeout errors', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    act(() => {
      result.current.handleError(new Error('Request timeout'))
    })
    
    expect(result.current.error).toBe('请求超时，请重试')
    expect(result.current.errorType).toBe(ErrorType.NETWORK_ERROR)
  })

  it('should handle validation errors with phone and verification code', () => {
    const { result } = renderHook(() => useErrorHandler())
    
    // Phone validation error
    act(() => {
      result.current.handleError(new Error('手机号格式不正确'))
    })
    
    expect(result.current.error).toBe('手机号格式不正确')
    expect(result.current.errorType).toBe(ErrorType.VALIDATION_ERROR)
    
    // Verification code error
    act(() => {
      result.current.handleError(new Error('验证码错误'))
    })
    
    expect(result.current.error).toBe('验证码错误')
    expect(result.current.errorType).toBe(ErrorType.VALIDATION_ERROR)
  })
})