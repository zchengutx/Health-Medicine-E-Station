import { describe, it, expect, beforeEach, vi } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useAuthStore } from '@/stores/auth'
import { Storage, TokenStorage, UserInfoStorage } from '@/utils/storage'

// Mock storage utilities
vi.mock('@/utils/storage', () => ({
  Storage: {
    get: vi.fn(),
    set: vi.fn(),
    remove: vi.fn(),
    cleanExpired: vi.fn()
  },
  TokenStorage: {
    getToken: vi.fn(),
    setToken: vi.fn(),
    removeToken: vi.fn()
  },
  UserInfoStorage: {
    getUserInfo: vi.fn(),
    setUserInfo: vi.fn(),
    removeUserInfo: vi.fn()
  },
  STORAGE_KEYS: {
    LAST_LOGIN_TIME: 'last_login_time',
    REMEMBER_PHONE: 'remember_phone'
  }
}))

describe('Auth Store - Profile Loading Fix', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  describe('waitForInitialization', () => {
    it('should resolve immediately if already initialized', async () => {
      const authStore = useAuthStore()
      authStore.loginState.isInitialized = true

      const startTime = Date.now()
      await authStore.waitForInitialization()
      const endTime = Date.now()

      expect(endTime - startTime).toBeLessThan(10) // Should be immediate
    })

    it('should wait for initialization to complete', async () => {
      const authStore = useAuthStore()
      authStore.loginState.isInitialized = false

      // Simulate initialization completing after 100ms
      setTimeout(() => {
        authStore.loginState.isInitialized = true
      }, 100)

      const startTime = Date.now()
      await authStore.waitForInitialization()
      const endTime = Date.now()

      expect(endTime - startTime).toBeGreaterThanOrEqual(100)
      expect(authStore.loginState.isInitialized).toBe(true)
    })
  })

  describe('initAuth', () => {
    it('should initialize auth state with stored data', () => {
      const mockToken = 'mock-token'
      const mockUserInfo = { DId: 123, Name: 'Test Doctor', Phone: '12345678901' }
      const mockLastLoginTime = Date.now()

      vi.mocked(TokenStorage.getToken).mockReturnValue(mockToken)
      vi.mocked(UserInfoStorage.getUserInfo).mockReturnValue(mockUserInfo)
      vi.mocked(Storage.get).mockImplementation((key) => {
        if (key === 'last_login_time') return mockLastLoginTime
        if (key === 'remember_phone') return true
        if (key === 'saved_phone') return '12345678901'
        return null
      })

      const authStore = useAuthStore()
      authStore.initAuth()

      expect(authStore.token).toBe(mockToken)
      expect(authStore.doctorInfo).toEqual(mockUserInfo)
      expect(authStore.loginState.isInitialized).toBe(true)
      expect(authStore.loginState.doctorId).toBe(123)
      expect(authStore.loginState.lastLoginTime).toBe(mockLastLoginTime)
    })

    it('should handle initialization errors gracefully', () => {
      vi.mocked(TokenStorage.getToken).mockImplementation(() => {
        throw new Error('Storage error')
      })

      const authStore = useAuthStore()
      
      expect(() => authStore.initAuth()).not.toThrow()
      expect(authStore.loginState.isInitialized).toBe(true) // Should still be marked as initialized
    })

    it('should set isInitialized to true even without stored data', () => {
      vi.mocked(TokenStorage.getToken).mockReturnValue(null)
      vi.mocked(UserInfoStorage.getUserInfo).mockReturnValue(null)

      const authStore = useAuthStore()
      authStore.initAuth()

      expect(authStore.loginState.isInitialized).toBe(true)
      expect(authStore.token).toBeNull()
      expect(authStore.doctorInfo).toBeNull()
    })
  })

  describe('logout', () => {
    it('should reset initialization state on logout', () => {
      const authStore = useAuthStore()
      
      // Set up initial state
      authStore.loginState.isInitialized = true
      authStore.loginState.doctorId = 123
      authStore.token = 'test-token'
      authStore.doctorInfo = { DId: 123, Name: 'Test', Phone: '123' } as any

      authStore.logout()

      expect(authStore.loginState.isInitialized).toBe(false)
      expect(authStore.loginState.doctorId).toBeUndefined()
      expect(authStore.token).toBeNull()
      expect(authStore.doctorInfo).toBeNull()
    })
  })

  describe('login', () => {
    it('should set doctorId in loginState', () => {
      const authStore = useAuthStore()
      const mockToken = 'test-token'
      const mockDoctorInfo = { DId: 456, Name: 'Test Doctor', Phone: '12345678901' } as any

      authStore.login(mockToken, mockDoctorInfo)

      expect(authStore.loginState.doctorId).toBe(456)
      expect(authStore.loginState.isInitialized).toBe(true) // Should be set during login
      expect(authStore.token).toBe(mockToken)
      expect(authStore.doctorInfo).toEqual(mockDoctorInfo)
    })
  })
})