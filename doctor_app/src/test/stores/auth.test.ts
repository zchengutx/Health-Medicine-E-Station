import { describe, it, expect, beforeEach, vi } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'
import { useAuthStore } from '@/stores/auth'
import type { DoctorInfo } from '@/stores/auth'

// Mock localStorage and storage utilities
vi.mock('@/utils/storage', () => ({
  TokenStorage: {
    getToken: vi.fn(() => null),
    setToken: vi.fn(),
    removeToken: vi.fn()
  },
  UserInfoStorage: {
    getUserInfo: vi.fn(() => null),
    setUserInfo: vi.fn(),
    removeUserInfo: vi.fn()
  },
  Storage: {
    get: vi.fn(() => null),
    set: vi.fn(),
    remove: vi.fn(),
    cleanExpired: vi.fn()
  },
  STORAGE_KEYS: {
    LAST_LOGIN_TIME: 'last_login_time',
    REMEMBER_PHONE: 'remember_phone'
  }
}))

describe('AuthStore', () => {
  let authStore: ReturnType<typeof useAuthStore>

  beforeEach(() => {
    setActivePinia(createPinia())
    authStore = useAuthStore()
    vi.clearAllMocks()
  })

  describe('初始状态', () => {
    it('应该有正确的默认状态', () => {
      expect(authStore.isLoggedIn).toBe(false)
      expect(authStore.isAuthenticated).toBe(false)
      expect(authStore.doctorInfo).toBeNull()
      expect(authStore.doctorName).toBe('医生')
      expect(authStore.doctorStatus).toBe('inactive')
    })
  })

  describe('登录功能', () => {
    const mockDoctorInfo: DoctorInfo = {
      DId: 123,
      Name: '张医生',
      Phone: '13800138000',
      Email: 'doctor@example.com',
      Avatar: 'avatar.jpg',
      Status: 'active'
    }

    it('应该成功登录', () => {
      authStore.login('test_token', mockDoctorInfo, '13800138000', true)

      expect(authStore.isLoggedIn).toBe(true)
      expect(authStore.isAuthenticated).toBe(true)
      expect(authStore.doctorInfo).toEqual(mockDoctorInfo)
      expect(authStore.doctorName).toBe('张医生')
      expect(authStore.doctorStatus).toBe('active')
    })

    it('应该在参数不完整时抛出错误', () => {
      expect(() => {
        authStore.login('', mockDoctorInfo)
      }).toThrow('登录参数不完整')

      expect(() => {
        authStore.login('test_token', { ...mockDoctorInfo, DId: 0 })
      }).toThrow('登录参数不完整')
    })

    it('应该保存手机号', () => {
      authStore.login('test_token', mockDoctorInfo, '13800138000', true)
      
      expect(authStore.loginState.savedPhone).toBe('13800138000')
      expect(authStore.loginState.rememberPhone).toBe(true)
    })
  })

  describe('登出功能', () => {
    it('应该成功登出', () => {
      const mockDoctorInfo: DoctorInfo = {
        DId: 123,
        Name: '张医生',
        Phone: '13800138000',
        Email: 'doctor@example.com',
        Avatar: '',
        Status: 'active'
      }

      // 先登录
      authStore.login('test_token', mockDoctorInfo)
      expect(authStore.isLoggedIn).toBe(true)

      // 再登出
      authStore.logout()
      expect(authStore.isLoggedIn).toBe(false)
      expect(authStore.doctorInfo).toBeNull()
    })
  })

  describe('Token有效性检查', () => {
    it('应该在没有token时返回false', () => {
      expect(authStore.checkTokenExpiry()).toBe(false)
    })

    it('应该在没有lastLoginTime时返回false', () => {
      authStore.setToken('test_token')
      expect(authStore.checkTokenExpiry()).toBe(false)
    })

    it('应该在token有效时返回true', () => {
      authStore.setToken('test_token')
      authStore.setLoginState({ lastLoginTime: Date.now() })
      expect(authStore.checkTokenExpiry()).toBe(true)
    })

    it('应该在token过期时返回false并执行登出', () => {
      authStore.setToken('test_token')
      authStore.setLoginState({ 
        lastLoginTime: Date.now() - 8 * 24 * 60 * 60 * 1000 // 8天前
      })
      
      expect(authStore.checkTokenExpiry()).toBe(false)
      expect(authStore.isLoggedIn).toBe(false)
    })
  })

  describe('用户信息管理', () => {
    const mockDoctorInfo: DoctorInfo = {
      DId: 123,
      Name: '张医生',
      Phone: '13800138000',
      Email: 'doctor@example.com',
      Avatar: '',
      Status: 'active'
    }

    it('应该设置医生信息', () => {
      authStore.setDoctorInfo(mockDoctorInfo)
      expect(authStore.doctorInfo).toEqual(mockDoctorInfo)
    })

    it('应该更新医生信息', () => {
      authStore.setDoctorInfo(mockDoctorInfo)
      
      const updates = { Name: '李医生', Email: 'li@example.com' }
      authStore.updateDoctorInfo(updates)
      
      expect(authStore.doctorInfo?.Name).toBe('李医生')
      expect(authStore.doctorInfo?.Email).toBe('li@example.com')
      expect(authStore.doctorInfo?.Phone).toBe('13800138000') // 保持不变
    })

    it('setUserInfo应该作为setDoctorInfo的别名工作', () => {
      authStore.setUserInfo(mockDoctorInfo)
      expect(authStore.doctorInfo).toEqual(mockDoctorInfo)
    })
  })

  describe('登录状态管理', () => {
    it('应该设置登录状态', () => {
      const newState = {
        isLoading: true,
        lastLoginTime: Date.now(),
        rememberPhone: true,
        savedPhone: '13800138000'
      }
      
      authStore.setLoginState(newState)
      expect(authStore.loginState).toEqual(expect.objectContaining(newState))
    })

    it('应该保存手机号', () => {
      authStore.savePhone('13800138000', true)
      expect(authStore.loginState.savedPhone).toBe('13800138000')
      expect(authStore.loginState.rememberPhone).toBe(true)
    })

    it('应该清除手机号', () => {
      authStore.savePhone('13800138000', true)
      authStore.savePhone('', false)
      expect(authStore.loginState.savedPhone).toBe('')
      expect(authStore.loginState.rememberPhone).toBe(false)
    })
  })

  describe('初始化', () => {
    it('应该正确初始化认证状态', () => {
      authStore.initAuth()
      
      // 验证初始化过程中的日志输出
      expect(console.log).toHaveBeenCalledWith('初始化认证状态...')
      expect(console.log).toHaveBeenCalledWith('认证状态初始化完成')
    })
  })
})