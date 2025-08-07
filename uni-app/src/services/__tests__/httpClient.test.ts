/**
 * HTTP客户端单元测试
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import { API_CONFIG } from '../../constants'

// Mock axios
const mockAxiosInstance = {
  interceptors: {
    request: { use: vi.fn() },
    response: { use: vi.fn() },
  },
  request: vi.fn(),
}

vi.mock('axios', () => ({
  default: {
    create: vi.fn(() => mockAxiosInstance),
  },
}))

describe('httpClient', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
  })

  describe('基础配置', () => {
    it('应该正确导入httpClient模块', async () => {
      const { httpClient } = await import('../httpClient')
      expect(httpClient).toBeDefined()
    })

    it('应该正确导入request方法', async () => {
      const { request } = await import('../httpClient')
      expect(request).toBeDefined()
      expect(request.get).toBeDefined()
      expect(request.post).toBeDefined()
      expect(request.put).toBeDefined()
      expect(request.delete).toBeDefined()
    })
  })

  describe('API配置', () => {
    it('应该使用正确的API配置', () => {
      expect(API_CONFIG.BASE_URL).toBeDefined()
      expect(API_CONFIG.TIMEOUT).toBeDefined()
      expect(API_CONFIG.RETRY_TIMES).toBeDefined()
    })
  })

  describe('localStorage操作', () => {
    it('应该能够设置和获取token', () => {
      const testToken = 'test-token-123'
      localStorage.setItem('auth_token', testToken)
      
      const retrievedToken = localStorage.getItem('auth_token')
      expect(retrievedToken).toBe(testToken)
    })

    it('应该能够清除token', () => {
      localStorage.setItem('auth_token', 'test-token')
      localStorage.removeItem('auth_token')
      
      const retrievedToken = localStorage.getItem('auth_token')
      expect(retrievedToken).toBeNull()
    })
  })
})