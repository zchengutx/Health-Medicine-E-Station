import { log } from './logger'

// 缓存项接口
interface CacheItem<T> {
  data: T
  timestamp: number
  expiry: number
  version: string
}

// 缓存配置
interface CacheConfig {
  defaultTTL: number // 默认过期时间（毫秒）
  maxSize: number    // 最大缓存项数
  version: string    // 缓存版本
}

// 默认配置
const defaultConfig: CacheConfig = {
  defaultTTL: 30 * 60 * 1000, // 30分钟
  maxSize: 100,
  version: '1.0.0'
}

/**
 * 本地缓存管理器
 */
export class CacheManager {
  private config: CacheConfig
  private keyPrefix: string

  constructor(keyPrefix: string = 'app_cache', config: Partial<CacheConfig> = {}) {
    this.keyPrefix = keyPrefix
    this.config = { ...defaultConfig, ...config }
  }

  /**
   * 设置缓存
   */
  set<T>(key: string, data: T, ttl?: number): boolean {
    try {
      const expiry = Date.now() + (ttl || this.config.defaultTTL)
      const cacheItem: CacheItem<T> = {
        data,
        timestamp: Date.now(),
        expiry,
        version: this.config.version
      }

      const cacheKey = this.getCacheKey(key)
      localStorage.setItem(cacheKey, JSON.stringify(cacheItem))
      
      log.debug('缓存设置成功', { key, expiry: new Date(expiry) })
      return true
    } catch (error) {
      log.error('设置缓存失败', { key, error })
      return false
    }
  }

  /**
   * 获取缓存
   */
  get<T>(key: string): T | null {
    try {
      const cacheKey = this.getCacheKey(key)
      const cached = localStorage.getItem(cacheKey)
      
      if (!cached) {
        return null
      }

      const cacheItem: CacheItem<T> = JSON.parse(cached)
      
      // 检查版本
      if (cacheItem.version !== this.config.version) {
        log.debug('缓存版本不匹配，删除缓存', { key, cached: cacheItem.version, current: this.config.version })
        this.delete(key)
        return null
      }

      // 检查是否过期
      if (Date.now() > cacheItem.expiry) {
        log.debug('缓存已过期，删除缓存', { key, expiry: new Date(cacheItem.expiry) })
        this.delete(key)
        return null
      }

      log.debug('缓存命中', { key, age: Date.now() - cacheItem.timestamp })
      return cacheItem.data
    } catch (error) {
      log.error('获取缓存失败', { key, error })
      this.delete(key) // 删除损坏的缓存
      return null
    }
  }

  /**
   * 删除缓存
   */
  delete(key: string): boolean {
    try {
      const cacheKey = this.getCacheKey(key)
      localStorage.removeItem(cacheKey)
      log.debug('缓存删除成功', { key })
      return true
    } catch (error) {
      log.error('删除缓存失败', { key, error })
      return false
    }
  }

  /**
   * 检查缓存是否存在且有效
   */
  has(key: string): boolean {
    return this.get(key) !== null
  }

  /**
   * 清除所有缓存
   */
  clear(): boolean {
    try {
      const keys = this.getAllKeys()
      keys.forEach(key => {
        localStorage.removeItem(key)
      })
      log.info('所有缓存已清除', { count: keys.length })
      return true
    } catch (error) {
      log.error('清除缓存失败', error)
      return false
    }
  }

  /**
   * 清除过期缓存
   */
  clearExpired(): number {
    let clearedCount = 0
    try {
      const keys = this.getAllKeys()
      const now = Date.now()

      keys.forEach(fullKey => {
        try {
          const cached = localStorage.getItem(fullKey)
          if (cached) {
            const cacheItem: CacheItem<any> = JSON.parse(cached)
            if (now > cacheItem.expiry || cacheItem.version !== this.config.version) {
              localStorage.removeItem(fullKey)
              clearedCount++
            }
          }
        } catch (error) {
          // 删除损坏的缓存项
          localStorage.removeItem(fullKey)
          clearedCount++
        }
      })

      log.info('过期缓存清理完成', { clearedCount })
      return clearedCount
    } catch (error) {
      log.error('清理过期缓存失败', error)
      return clearedCount
    }
  }

  /**
   * 获取缓存统计信息
   */
  getStats(): {
    totalItems: number
    totalSize: number
    expiredItems: number
    validItems: number
  } {
    const keys = this.getAllKeys()
    const now = Date.now()
    let totalSize = 0
    let expiredItems = 0
    let validItems = 0

    keys.forEach(fullKey => {
      try {
        const cached = localStorage.getItem(fullKey)
        if (cached) {
          totalSize += cached.length
          const cacheItem: CacheItem<any> = JSON.parse(cached)
          if (now > cacheItem.expiry || cacheItem.version !== this.config.version) {
            expiredItems++
          } else {
            validItems++
          }
        }
      } catch (error) {
        expiredItems++
      }
    })

    return {
      totalItems: keys.length,
      totalSize,
      expiredItems,
      validItems
    }
  }

  /**
   * 缓存或获取数据（如果缓存不存在则执行获取函数）
   */
  async getOrSet<T>(
    key: string,
    fetchFn: () => Promise<T>,
    ttl?: number
  ): Promise<T> {
    // 先尝试从缓存获取
    const cached = this.get<T>(key)
    if (cached !== null) {
      return cached
    }

    // 缓存不存在，执行获取函数
    try {
      const data = await fetchFn()
      this.set(key, data, ttl)
      return data
    } catch (error) {
      log.error('获取数据失败', { key, error })
      throw error
    }
  }

  // 私有方法

  private getCacheKey(key: string): string {
    return `${this.keyPrefix}_${key}`
  }

  private getAllKeys(): string[] {
    const keys: string[] = []
    const prefix = `${this.keyPrefix}_`
    
    for (let i = 0; i < localStorage.length; i++) {
      const key = localStorage.key(i)
      if (key && key.startsWith(prefix)) {
        keys.push(key)
      }
    }
    
    return keys
  }
}

// 创建默认实例
export const profileCache = new CacheManager('profile_cache', {
  defaultTTL: 15 * 60 * 1000, // 15分钟
  version: '1.0.0'
})

export const appCache = new CacheManager('app_cache', {
  defaultTTL: 30 * 60 * 1000, // 30分钟
  version: '1.0.0'
})

// 导出便捷函数
export const cacheProfile = (doctorId: number, profile: any) => 
  profileCache.set(`doctor_${doctorId}`, profile)

export const getCachedProfile = (doctorId: number) => 
  profileCache.get(`doctor_${doctorId}`)

export const clearProfileCache = (doctorId: number) => 
  profileCache.delete(`doctor_${doctorId}`)

// 定期清理过期缓存
if (typeof window !== 'undefined') {
  // 页面加载时清理一次
  setTimeout(() => {
    profileCache.clearExpired()
    appCache.clearExpired()
  }, 1000)

  // 每30分钟清理一次
  setInterval(() => {
    profileCache.clearExpired()
    appCache.clearExpired()
  }, 30 * 60 * 1000)
}