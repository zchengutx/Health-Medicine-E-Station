/**
 * API服务层集成测试
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import { sendVerificationCode, loginWithVerificationCode } from '../authApi'
import { request } from '../httpClient'

// Mock httpClient
vi.mock('../httpClient', () => ({
  request: {
    post: vi.fn(),
  },
}))

const mockRequest = vi.mocked(request)

describe('API服务层集成测试', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
  })

  describe('完整登录流程', () => {
    it('应该能够完成发送验证码到登录的完整流程', async () => {
      const phone = '13800138000'
      const verificationCode = '123456'

      // 模拟发送验证码成功
      const sendCodeResponse = {
        success: true,
        message: '验证码发送成功',
        countdown: 60,
      }
      mockRequest.post.mockResolvedValueOnce(sendCodeResponse)

      // 发送验证码
      const codeResult = await sendVerificationCode({ phone })
      expect(codeResult).toEqual(sendCodeResponse)
      expect(mockRequest.post).toHaveBeenCalledWith('/auth/send-code', {
        phone,
        type: 'login',
      })

      // 模拟登录成功
      const loginResponse = {
        user: {
          id: '1',
          phone,
          nickname: 'test',
          avatar: '',
          createdAt: '2024-01-01',
          updatedAt: '2024-01-01',
        },
        token: 'test-token',
        refreshToken: 'test-refresh-token',
        expiresIn: 3600,
      }
      mockRequest.post.mockResolvedValueOnce(loginResponse)

      // 登录
      const loginResult = await loginWithVerificationCode({
        phone,
        verificationCode,
      })

      expect(loginResult).toEqual(loginResponse)
      expect(mockRequest.post).toHaveBeenCalledWith('/auth/login', {
        phone,
        verificationCode,
      })

      // 验证token是否保存
      expect(localStorage.getItem('auth_token')).toBe('test-token')
      expect(localStorage.getItem('refresh_token')).toBe('test-refresh-token')
      expect(localStorage.getItem('user_info')).toBe(JSON.stringify(loginResponse.user))
    })

    it('应该正确处理发送验证码失败的情况', async () => {
      const phone = '13800138000'

      // 模拟发送验证码失败
      mockRequest.post.mockRejectedValueOnce(new Error('请求过于频繁'))

      await expect(sendVerificationCode({ phone })).rejects.toThrow('请求过于频繁，请稍后重试')
    })

    it('应该正确处理登录失败的情况', async () => {
      const phone = '13800138000'
      const verificationCode = '123456'

      // 模拟登录失败
      mockRequest.post.mockRejectedValueOnce(new Error('验证码错误'))

      await expect(
        loginWithVerificationCode({ phone, verificationCode })
      ).rejects.toThrow('验证码错误，请重新输入')
    })
  })

  describe('参数验证', () => {
    it('应该验证手机号格式', async () => {
      await expect(sendVerificationCode({ phone: '123' })).rejects.toThrow('请输入正确的手机号')
      
      await expect(
        loginWithVerificationCode({ phone: '123', verificationCode: '123456' })
      ).rejects.toThrow('请输入正确的手机号')
    })

    it('应该验证验证码格式', async () => {
      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '123' })
      ).rejects.toThrow('请输入正确的验证码')
    })

    it('应该验证必填参数', async () => {
      await expect(sendVerificationCode({ phone: '' })).rejects.toThrow('参数不完整')
      
      await expect(
        loginWithVerificationCode({ phone: '', verificationCode: '123456' })
      ).rejects.toThrow('参数不完整')
      
      await expect(
        loginWithVerificationCode({ phone: '13800138000', verificationCode: '' })
      ).rejects.toThrow('参数不完整')
    })
  })
})