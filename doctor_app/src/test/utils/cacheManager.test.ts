import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { CacheManager, profileCache, appCache, cacheProfile, getCachedProfile, clearProfileCache } from '@/utils/cacheManager'

// Mock localStorage
const localStorageMock = {
  getItem: vi.fn(),
  setItem: vi.fn(),
  removeItem: vi.fn(),
  clear: vi.fn(),
  length: 0,
  key: vi.fn()
}

Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
})

describe('CacheManager', () => {
  let cacheManager: CacheManager

  beforeEach(() => {
    cacheManager = new CacheManager('test_cache', {
      defaultTTL: 1000, // 1秒用于测试
      version: '1.0.0'
    })
    vi.clearAllMocks()
    localStorageMock.length = 0
  })

  describe('基本缓存操作', () => {
    it('应该能够设置缓存', () => {
      const testData = { id: 1, name: 'test' }
      
      const result = cacheManager.set('test-key', testData)
      
      expect(result).toBe(true)
      expect(localStorageMock.setItem).toHaveBeenCalledWith(
        'test_cache_test-key',
        expect.stringContaining('"data":{"id":1,"name":"test"}')
      )
    })

    it('应该能够获取有效缓存', () => {
      const testData = { id: 1, name: 'test' }
      const cacheItem = {
        data: testData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000, // 10秒后过期
        version: '1.0.0'
      }
      
      localStorageMock.getItem.mockReturnValue(JSON.stringify(cacheItem))
      
      const result = cacheManager.get('test-key')
      
      expect(result).toEqual(testData)
      expect(localStorageMock.getItem).toHaveBeenCalledWith('test_cache_test-key')
    })

    it('应该在缓存过期时返回null', () => {
      const testData = { id: 1, name: 'test' }
      const expiredCacheItem = {
        data: testData,
        timestamp: Date.now() - 20000,
        expiry: Date.now() - 10000, // 已过期
        version: '1.0.0'
      }
      
      localStorageMock.getItem.mockReturnValue(JSON.stringify(expiredCacheItem))
      
      const result = cacheManager.get('test-key')
      
      expect(result).toBeNull()
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_test-key')
    })

    it('应该在版本不匹配时返回null', () => {
      const testData = { id: 1, name: 'test' }
      const oldVersionCacheItem = {
        data: testData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '0.9.0' // 旧版本
      }
      
      localStorageMock.getItem.mockReturnValue(JSON.stringify(oldVersionCacheItem))
      
      const result = cacheManager.get('test-key')
      
      expect(result).toBeNull()
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_test-key')
    })

    it('应该能够删除缓存', () => {
      const result = cacheManager.delete('test-key')
      
      expect(result).toBe(true)
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_test-key')
    })

    it('应该能够检查缓存是否存在', () => {
      const testData = { id: 1, name: 'test' }
      const cacheItem = {
        data: testData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      }
      
      localStorageMock.getItem.mockReturnValue(JSON.stringify(cacheItem))
      
      expect(cacheManager.has('test-key')).toBe(true)
      
      localStorageMock.getItem.mockReturnValue(null)
      
      expect(cacheManager.has('non-existent-key')).toBe(false)
    })
  })

  describe('缓存清理', () => {
    it('应该能够清除所有缓存', () => {
      const mockKeys = ['test_cache_key1', 'test_cache_key2', 'other_cache_key']
      localStorageMock.length = 3
      localStorageMock.key.mockImplementation((index) => mockKeys[index])
      
      const result = cacheManager.clear()
      
      expect(result).toBe(true)
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_key1')
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_key2')
      expect(localStorageMock.removeItem).not.toHaveBeenCalledWith('other_cache_key')
    })

    it('应该能够清除过期缓存', () => {
      const validCacheItem = {
        data: { id: 1 },
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      }
      
      const expiredCacheItem = {
        data: { id: 2 },
        timestamp: Date.now() - 20000,
        expiry: Date.now() - 10000,
        version: '1.0.0'
      }
      
      const mockKeys = ['test_cache_valid', 'test_cache_expired']
      localStorageMock.length = 2
      localStorageMock.key.mockImplementation((index) => mockKeys[index])
      localStorageMock.getItem.mockImplementation((key) => {
        if (key === 'test_cache_valid') return JSON.stringify(validCacheItem)
        if (key === 'test_cache_expired') return JSON.stringify(expiredCacheItem)
        return null
      })
      
      const clearedCount = cacheManager.clearExpired()
      
      expect(clearedCount).toBe(1)
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_expired')
      expect(localStorageMock.removeItem).not.toHaveBeenCalledWith('test_cache_valid')
    })
  })

  describe('缓存统计', () => {
    it('应该能够获取缓存统计信息', () => {
      const validCacheItem = JSON.stringify({
        data: { id: 1 },
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      })
      
      const expiredCacheItem = JSON.stringify({
        data: { id: 2 },
        timestamp: Date.now() - 20000,
        expiry: Date.now() - 10000,
        version: '1.0.0'
      })
      
      const mockKeys = ['test_cache_valid', 'test_cache_expired']
      localStorageMock.length = 2
      localStorageMock.key.mockImplementation((index) => mockKeys[index])
      localStorageMock.getItem.mockImplementation((key) => {
        if (key === 'test_cache_valid') return validCacheItem
        if (key === 'test_cache_expired') return expiredCacheItem
        return null
      })
      
      const stats = cacheManager.getStats()
      
      expect(stats.totalItems).toBe(2)
      expect(stats.validItems).toBe(1)
      expect(stats.expiredItems).toBe(1)
      expect(stats.totalSize).toBe(validCacheItem.length + expiredCacheItem.length)
    })
  })

  describe('getOrSet 方法', () => {
    it('应该在缓存存在时返回缓存数据', async () => {
      const testData = { id: 1, name: 'test' }
      const cacheItem = {
        data: testData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      }
      
      localStorageMock.getItem.mockReturnValue(JSON.stringify(cacheItem))
      const fetchFn = vi.fn()
      
      const result = await cacheManager.getOrSet('test-key', fetchFn)
      
      expect(result).toEqual(testData)
      expect(fetchFn).not.toHaveBeenCalled()
    })

    it('应该在缓存不存在时调用获取函数', async () => {
      const testData = { id: 1, name: 'test' }
      localStorageMock.getItem.mockReturnValue(null)
      const fetchFn = vi.fn().mockResolvedValue(testData)
      
      const result = await cacheManager.getOrSet('test-key', fetchFn)
      
      expect(result).toEqual(testData)
      expect(fetchFn).toHaveBeenCalledTimes(1)
      expect(localStorageMock.setItem).toHaveBeenCalled()
    })

    it('应该在获取函数失败时抛出错误', async () => {
      localStorageMock.getItem.mockReturnValue(null)
      const fetchFn = vi.fn().mockRejectedValue(new Error('Fetch failed'))
      
      await expect(
        cacheManager.getOrSet('test-key', fetchFn)
      ).rejects.toThrow('Fetch failed')
      
      expect(fetchFn).toHaveBeenCalledTimes(1)
    })
  })

  describe('错误处理', () => {
    it('应该在localStorage操作失败时优雅处理', () => {
      localStorageMock.setItem.mockImplementation(() => {
        throw new Error('Storage quota exceeded')
      })
      
      const result = cacheManager.set('test-key', { data: 'test' })
      
      expect(result).toBe(false)
    })

    it('应该在获取损坏的缓存数据时返回null', () => {
      localStorageMock.getItem.mockReturnValue('invalid json')
      
      const result = cacheManager.get('test-key')
      
      expect(result).toBeNull()
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('test_cache_test-key')
    })
  })
})

describe('便捷函数', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('cacheProfile 应该正常工作', () => {
    const profile = { DId: 1, Name: 'Test Doctor' }
    
    cacheProfile(1, profile)
    
    expect(localStorageMock.setItem).toHaveBeenCalledWith(
      'profile_cache_doctor_1',
      expect.stringContaining('"DId":1')
    )
  })

  it('getCachedProfile 应该正常工作', () => {
    const profile = { DId: 1, Name: 'Test Doctor' }
    const cacheItem = {
      data: profile,
      timestamp: Date.now(),
      expiry: Date.now() + 10000,
      version: '1.0.0'
    }
    
    localStorageMock.getItem.mockReturnValue(JSON.stringify(cacheItem))
    
    const result = getCachedProfile(1)
    
    expect(result).toEqual(profile)
    expect(localStorageMock.getItem).toHaveBeenCalledWith('profile_cache_doctor_1')
  })

  it('clearProfileCache 应该正常工作', () => {
    clearProfileCache(1)
    
    expect(localStorageMock.removeItem).toHaveBeenCalledWith('profile_cache_doctor_1')
  })
})