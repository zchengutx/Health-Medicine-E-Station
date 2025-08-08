import { describe, it, expect, beforeEach, vi } from 'vitest'
import { Storage, TokenStorage, UserInfoStorage, STORAGE_KEYS } from '@/utils/storage'

describe('storage utils', () => {
  beforeEach(() => {
    localStorage.clear()
    vi.clearAllMocks()
  })

  describe('Storage', () => {
    it('should set and get data', () => {
      const testData = { name: 'test', value: 123 }
      Storage.set('test-key', testData)
      
      const retrieved = Storage.get('test-key')
      expect(retrieved).toEqual(testData)
    })

    it('should return null for non-existent key', () => {
      const result = Storage.get('non-existent')
      expect(result).toBeNull()
    })

    it('should remove data', () => {
      Storage.set('test-key', 'test-value')
      expect(Storage.has('test-key')).toBe(true)
      
      Storage.remove('test-key')
      expect(Storage.has('test-key')).toBe(false)
    })

    it('should clear all data', () => {
      Storage.set('key1', 'value1')
      Storage.set('key2', 'value2')
      
      Storage.clear()
      
      expect(Storage.get('key1')).toBeNull()
      expect(Storage.get('key2')).toBeNull()
    })

    it('should handle encrypted storage', () => {
      const testData = { secret: 'password123' }
      Storage.set('encrypted-key', testData, true)
      
      const retrieved = Storage.get('encrypted-key', true)
      expect(retrieved).toEqual(testData)
    })

    it('should set and get data with expiry', () => {
      const testData = 'test-value'
      Storage.setWithExpiry('expiry-key', testData, 60) // 60 minutes
      
      const retrieved = Storage.getWithExpiry('expiry-key')
      expect(retrieved).toBe(testData)
    })

    it('should return null for expired data', () => {
      const testData = 'test-value'
      Storage.setWithExpiry('expiry-key', testData, -1) // expired
      
      const retrieved = Storage.getWithExpiry('expiry-key')
      expect(retrieved).toBeNull()
    })

    it('should calculate storage size', () => {
      Storage.set('test-key', 'test-value')
      const size = Storage.getSize()
      expect(size).toBeGreaterThan(0)
    })
  })

  describe('TokenStorage', () => {
    it('should set and get token', () => {
      const token = 'test-token-123'
      TokenStorage.setToken(token)
      
      const retrieved = TokenStorage.getToken()
      expect(retrieved).toBe(token)
    })

    it('should check if token exists', () => {
      expect(TokenStorage.hasToken()).toBe(false)
      
      TokenStorage.setToken('test-token')
      expect(TokenStorage.hasToken()).toBe(true)
    })

    it('should remove token', () => {
      TokenStorage.setToken('test-token')
      expect(TokenStorage.hasToken()).toBe(true)
      
      TokenStorage.removeToken()
      expect(TokenStorage.hasToken()).toBe(false)
    })
  })

  describe('UserInfoStorage', () => {
    it('should set and get user info', () => {
      const userInfo = {
        DId: 1,
        Name: 'Test Doctor',
        Phone: '13812345678',
        Email: 'test@example.com',
        Avatar: '',
        Status: 'active'
      }
      
      UserInfoStorage.setUserInfo(userInfo)
      const retrieved = UserInfoStorage.getUserInfo()
      expect(retrieved).toEqual(userInfo)
    })

    it('should check if user info exists', () => {
      expect(UserInfoStorage.hasUserInfo()).toBe(false)
      
      UserInfoStorage.setUserInfo({ name: 'test' })
      expect(UserInfoStorage.hasUserInfo()).toBe(true)
    })

    it('should remove user info', () => {
      UserInfoStorage.setUserInfo({ name: 'test' })
      expect(UserInfoStorage.hasUserInfo()).toBe(true)
      
      UserInfoStorage.removeUserInfo()
      expect(UserInfoStorage.hasUserInfo()).toBe(false)
    })
  })
})