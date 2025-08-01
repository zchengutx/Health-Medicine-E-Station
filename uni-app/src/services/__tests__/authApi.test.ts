/**
 * 认证API服务单元测试
 */

import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import {
  sendVerificationCode,
  loginWithVerificationCode,
  refreshAuthToken,
  logout,
  validateToken,
} from '../authApi'
import { request } from '../httpClient'
import { ERROR_CODES, ERROR_MESSAGES } from '../../constants'

// Mock httpClient
vi.mock('../httpClient', () => ({
  request: {
    post: vi.fn(),
    get: vi.fn(),
  },
}))

const mockRequest = vi.mocked(request)

describe('authApi', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  describe('sendVerificationCode', () => {
    it('应该成功发送验证码', async () => {
      const mockResponse = {
        success: true,
        message: '验证码发送成功',
        countdown: 60,
      }

      mockRequest.post.mockResolvedValue(mockResponse)

      const result = await sendVerificationCode({ phone: '13800138000' })

      expect(mockRequest.post).toHaveBeenCalledWith('/auth/send-code', {
        phone: '13800138000',
        type: 'login',
      })
      expect(result).toEqual(mockResponse)
    })

    it('应该在手机号为空时抛出错误', async () => {
      await expect(sendVerificationCode({ phone: '' })).rejects.toThrow(
        ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS]
      )

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该在手机号格式无效时抛出错误', async () => {
      await expect(sendVerificationCode({ phone: '123456' })).rejects.toThrow(
        ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE]
      )

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该处理频繁请求错误', async () => {
      mockRequest.post.mockRejectedValue(new Error('请求过于频繁'))

      await expect(sendVerificationCode({ phone: '13800138000' })).rejects.toThrow(
        ERROR_MESSAGES[ERROR_CODES.TOO_MANY_REQUESTS]
      )
    })

    it('应该处理手机号相关错误', async () => {
      mockRequest.post.mockRejectedValue(new Error('手机号不存在'))

      await expect(sendVerificationCode({ phone: '13800138000' })).rejects.toThrow(
        ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE]
      )
    })
  })

  describe('loginWithVerificationCode', () => {
    it('应该成功登录并保存token', async () => {
      const mockResponse = {
        user: {
          id: '1',
          phone: '13800138000',
          nickname: 'test',
          avatar: '',
          createdAt: '2024-01-01',
          updatedAt: '2024-01-01',
        },
        token: 'test-token',
        refreshToken: 'test-refresh-token',
        expiresIn: 3600,
      }

      mockRequest.post.mockResolvedValue(mockResponse)

      const result = await loginWithVerificationCode({
        phone: '13800138000',
        verificationCode: '123456',
      })

      expect(mockRequest.post).toHaveBeenCalledWith('/auth/login', {
        phone: '13800138000',
        verificationCode: '123456',
      })
      expect(result).toEqual(mockResponse)

      // 验证token是否保存到localStorage
      expect(localStorage.getItem('auth_token')).toBe('test-token')
      expect(localStorage.getItem('refresh_token')).toBe('test-refresh-token')
      expect(localStorage.getItem('user_info')).toBe(JSON.stringify(mockResponse.user))
    })

    it('应该在参数缺失时抛出错误', async () => {
      await expect(
        loginWithVerificationCode({ phone: '', verificationCode: '123456' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS])

      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.MISSING_PARAMS])

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该在手机号格式无效时抛出错误', async () => {
      await expect(
        loginWithVerificationCode({ phone: '123456', verificationCode: '123456' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.INVALID_PHONE])

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该在验证码格式无效时抛出错误', async () => {
      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '123' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.INVALID_CODE])

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该处理验证码错误', async () => {
      mockRequest.post.mockRejectedValue(new Error('验证码错误'))

      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '123456' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.CODE_INVALID])
    })

    it('应该处理验证码过期错误', async () => {
      mockRequest.post.mockRejectedValue(new Error('过期'))

      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '123456' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.CODE_EXPIRED])
    })

    it('应该处理手机号未注册错误', async () => {
      mockRequest.post.mockRejectedValue(new Error('手机号未注册'))

      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '123456' })
      ).rejects.toThrow(ERROR_MESSAGES[ERROR_CODES.PHONE_NOT_REGISTERED])
    })
  })

  describe('refreshAuthToken', () => {
    it('应该成功刷新token', async () => {
      const mockResponse = {
        token: 'new-token',
        refreshToken: 'new-refresh-token',
        expiresIn: 3600,
      }

      mockRequest.post.mockResolvedValue(mockResponse)

      const result = await refreshAuthToken('old-refresh-token')

      expect(mockRequest.post).toHaveBeenCalledWith('/auth/refresh', {
        refreshToken: 'old-refresh-token',
      })
      expect(result).toEqual(mockResponse)

      // 验证新token是否保存到localStorage
      expect(localStorage.getItem('auth_token')).toBe('new-token')
      expect(localStorage.getItem('refresh_token')).toBe('new-refresh-token')
    })

    it('应该在refreshToken为空时抛出错误', async () => {
      await expect(refreshAuthToken('')).rejects.toThrow('刷新token不能为空')

      expect(mockRequest.post).not.toHaveBeenCalled()
    })

    it('应该在刷新失败时清除本地存储', async () => {
      // 先设置一些本地存储数据
      localStorage.setItem('auth_token', 'old-token')
      localStorage.setItem('refresh_token', 'old-refresh-token')
      localStorage.setItem('user_info', '{}')

      mockRequest.post.mockRejectedValue(new Error('刷新失败'))

      await expect(refreshAuthToken('old-refresh-token')).rejects.toThrow('刷新失败')

      // 验证本地存储是否被清除
      expect(localStorage.getItem('auth_token')).toBeNull()
      expect(localStorage.getItem('refresh_token')).toBeNull()
      expect(localStorage.getItem('user_info')).toBeNull()
    })
  })

  describe('logout', () => {
    it('应该成功登出并清除本地存储', async () => {
      // 先设置一些本地存储数据
      localStorage.setItem('auth_token', 'test-token')
      localStorage.setItem('refresh_token', 'test-refresh-token')
      localStorage.setItem('user_info', '{}')

      mockRequest.post.mockResolvedValue({})

      await logout()

      expect(mockRequest.post).toHaveBeenCalledWith('/auth/logout')

      // 验证本地存储是否被清除
      expect(localStorage.getItem('auth_token')).toBeNull()
      expect(localStorage.getItem('refresh_token')).toBeNull()
      expect(localStorage.getItem('user_info')).toBeNull()
    })

    it('应该在API调用失败时仍然清除本地存储', async () => {
      const consoleSpy = vi.spyOn(console, 'warn').mockImplementation(() => {})

      // 先设置一些本地存储数据
      localStorage.setItem('auth_token', 'test-token')
      localStorage.setItem('refresh_token', 'test-refresh-token')
      localStorage.setItem('user_info', '{}')

      mockRequest.post.mockRejectedValue(new Error('网络错误'))

      await logout()

      expect(consoleSpy).toHaveBeenCalledWith('登出API调用失败:', expect.any(Error))

      // 验证本地存储是否被清除
      expect(localStorage.getItem('auth_token')).toBeNull()
      expect(localStorage.getItem('refresh_token')).toBeNull()
      expect(localStorage.getItem('user_info')).toBeNull()

      consoleSpy.mockRestore()
    })
  })

  describe('validateToken', () => {
    it('应该在token有效时返回true', async () => {
      localStorage.setItem('auth_token', 'valid-token')
      mockRequest.get.mockResolvedValue({})

      const result = await validateToken()

      expect(mockRequest.get).toHaveBeenCalledWith('/auth/validate')
      expect(result).toBe(true)
    })

    it('应该在没有token时返回false', async () => {
      const result = await validateToken()

      expect(mockRequest.get).not.toHaveBeenCalled()
      expect(result).toBe(false)
    })

    it('应该在token无效时返回false并清除本地存储', async () => {
      localStorage.setItem('auth_token', 'invalid-token')
      localStorage.setItem('refresh_token', 'test-refresh-token')
      localStorage.setItem('user_info', '{}')

      mockRequest.get.mockRejectedValue(new Error('Token无效'))

      const result = await validateToken()

      expect(result).toBe(false)

      // 验证本地存储是否被清除
      expect(localStorage.getItem('auth_token')).toBeNull()
      expect(localStorage.getItem('refresh_token')).toBeNull()
      expect(localStorage.getItem('user_info')).toBeNull()
    })
  })
})