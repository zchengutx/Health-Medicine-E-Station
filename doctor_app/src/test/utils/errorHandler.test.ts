import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { ErrorHandler, ErrorType, ErrorSeverity, analyzeError, getUserMessage, retryWithBackoff } from '@/utils/errorHandler'

describe('ErrorHandler', () => {
  let errorHandler: ErrorHandler

  beforeEach(() => {
    errorHandler = new ErrorHandler()
    vi.clearAllMocks()
  })

  describe('错误分析', () => {
    it('应该正确识别网络错误', () => {
      const networkError = new Error('网络连接失败')
      networkError.name = 'NetworkError'
      
      const errorInfo = errorHandler.analyzeError(networkError)
      
      expect(errorInfo.type).toBe(ErrorType.NETWORK)
      expect(errorInfo.severity).toBe(ErrorSeverity.MEDIUM)
      expect(errorInfo.retryable).toBe(true)
      expect(errorInfo.fallbackAvailable).toBe(true)
      expect(errorInfo.userMessage).toContain('网络')
    })

    it('应该正确识别认证错误', () => {
      const authError = { status: 401, message: 'Unauthorized' }
      
      const errorInfo = errorHandler.analyzeError(authError)
      
      expect(errorInfo.type).toBe(ErrorType.AUTHENTICATION)
      expect(errorInfo.severity).toBe(ErrorSeverity.HIGH)
      expect(errorInfo.retryable).toBe(false)
      expect(errorInfo.fallbackAvailable).toBe(false)
      expect(errorInfo.userMessage).toContain('登录')
    })

    it('应该正确识别验证错误', () => {
      const validationError = { status: 400, message: 'Validation failed' }
      
      const errorInfo = errorHandler.analyzeError(validationError)
      
      expect(errorInfo.type).toBe(ErrorType.VALIDATION)
      expect(errorInfo.severity).toBe(ErrorSeverity.LOW)
      expect(errorInfo.retryable).toBe(false)
    })

    it('应该正确识别服务器错误', () => {
      const serverError = { status: 500, message: 'Internal Server Error' }
      
      const errorInfo = errorHandler.analyzeError(serverError)
      
      expect(errorInfo.type).toBe(ErrorType.SERVER)
      expect(errorInfo.severity).toBe(ErrorSeverity.HIGH)
      expect(errorInfo.retryable).toBe(true)
      expect(errorInfo.fallbackAvailable).toBe(true)
    })

    it('应该处理未知错误', () => {
      const unknownError = new Error('Something went wrong')
      
      const errorInfo = errorHandler.analyzeError(unknownError)
      
      expect(errorInfo.type).toBe(ErrorType.UNKNOWN)
      expect(errorInfo.severity).toBe(ErrorSeverity.MEDIUM)
      expect(errorInfo.retryable).toBe(true)
    })
  })

  describe('重试机制', () => {
    it('应该在成功时不重试', async () => {
      const mockOperation = vi.fn().mockResolvedValue('success')
      
      const result = await errorHandler.retryWithBackoff(mockOperation, 'test-op')
      
      expect(result).toBe('success')
      expect(mockOperation).toHaveBeenCalledTimes(1)
    })

    it('应该在失败时重试', async () => {
      const mockOperation = vi.fn()
        .mockRejectedValueOnce(new Error('Network error'))
        .mockRejectedValueOnce(new Error('Network error'))
        .mockResolvedValue('success')
      
      const result = await errorHandler.retryWithBackoff(mockOperation, 'test-op', 3)
      
      expect(result).toBe('success')
      expect(mockOperation).toHaveBeenCalledTimes(3)
    })

    it('应该在达到最大重试次数后抛出错误', async () => {
      const mockOperation = vi.fn().mockRejectedValue(new Error('Persistent error'))
      
      await expect(
        errorHandler.retryWithBackoff(mockOperation, 'test-op', 2)
      ).rejects.toThrow('Persistent error')
      
      expect(mockOperation).toHaveBeenCalledTimes(3) // 初始调用 + 2次重试
    })

    it('应该对不可重试的错误立即失败', async () => {
      const authError = { status: 401, message: 'Unauthorized' }
      const mockOperation = vi.fn().mockRejectedValue(authError)
      
      await expect(
        errorHandler.retryWithBackoff(mockOperation, 'test-op', 3)
      ).rejects.toEqual(authError)
      
      expect(mockOperation).toHaveBeenCalledTimes(1)
    })
  })

  describe('便捷函数', () => {
    it('analyzeError 应该正常工作', () => {
      const error = new Error('Test error')
      const errorInfo = analyzeError(error, 'test-context')
      
      expect(errorInfo.type).toBe(ErrorType.UNKNOWN)
      expect(errorInfo.details?.context).toBe('test-context')
    })

    it('getUserMessage 应该返回用户友好的消息', () => {
      const networkError = new Error('网络连接失败')
      networkError.name = 'NetworkError'
      
      const message = getUserMessage(networkError)
      
      expect(message).toContain('网络')
      expect(message).not.toContain('NetworkError')
    })

    it('retryWithBackoff 应该正常工作', async () => {
      const mockOperation = vi.fn().mockResolvedValue('success')
      
      const result = await retryWithBackoff(mockOperation, 'test-op')
      
      expect(result).toBe('success')
      expect(mockOperation).toHaveBeenCalledTimes(1)
    })
  })

  describe('重试计数', () => {
    it('应该正确跟踪重试次数', async () => {
      const mockOperation = vi.fn()
        .mockRejectedValueOnce(new Error('Network error'))
        .mockResolvedValue('success')
      
      await errorHandler.retryWithBackoff(mockOperation, 'test-op')
      
      // 成功后重试计数应该被清除
      expect(errorHandler.getRetryCount('test-op')).toBe(0)
    })

    it('应该在失败时保持重试计数', async () => {
      const mockOperation = vi.fn().mockRejectedValue(new Error('Persistent error'))
      
      try {
        await errorHandler.retryWithBackoff(mockOperation, 'test-op', 2)
      } catch (error) {
        // 失败后重试计数应该被清除
        expect(errorHandler.getRetryCount('test-op')).toBe(0)
      }
    })
  })

  describe('错误检测方法', () => {
    it('canRetry 应该正确判断是否可重试', () => {
      const networkError = new Error('网络错误')
      const authError = { status: 401, message: 'Unauthorized' }
      
      expect(errorHandler.canRetry(networkError)).toBe(true)
      expect(errorHandler.canRetry(authError)).toBe(false)
    })

    it('hasFallback 应该正确判断是否有备用方案', () => {
      const networkError = new Error('网络错误')
      const authError = { status: 401, message: 'Unauthorized' }
      
      expect(errorHandler.hasFallback(networkError)).toBe(true)
      expect(errorHandler.hasFallback(authError)).toBe(false)
    })
  })
})