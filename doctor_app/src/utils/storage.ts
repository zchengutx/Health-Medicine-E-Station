// 简单的加密解密工具
class CryptoUtil {
  private static readonly SECRET_KEY = 'doctor_app_2024'

  // 简单的XOR加密 - 修复Unicode字符问题
  static encrypt(text: string): string {
    try {
      // 先将字符串转换为UTF-8字节
      const utf8Bytes = new TextEncoder().encode(text)
      let result = ''

      for (let i = 0; i < utf8Bytes.length; i++) {
        const keyChar = this.SECRET_KEY.charCodeAt(i % this.SECRET_KEY.length)
        result += String.fromCharCode(utf8Bytes[i] ^ keyChar)
      }

      // 使用更安全的Base64编码
      return btoa(unescape(encodeURIComponent(result)))
    } catch (error) {
      // 如果加密失败，返回原文本（不加密）
      return text
    }
  }

  // 简单的XOR解密 - 修复Unicode字符问题
  static decrypt(encryptedText: string): string {
    try {
      // 使用更安全的Base64解码
      const text = decodeURIComponent(escape(atob(encryptedText)))
      const bytes = new Uint8Array(text.length)

      for (let i = 0; i < text.length; i++) {
        const keyChar = this.SECRET_KEY.charCodeAt(i % this.SECRET_KEY.length)
        bytes[i] = text.charCodeAt(i) ^ keyChar
      }

      // 将字节转换回UTF-8字符串
      return new TextDecoder().decode(bytes)
    } catch (error) {
      // 如果解密失败，返回原文本
      return encryptedText
    }
  }
}

// 本地存储工具类
export class Storage {
  // 设置存储项
  static set(key: string, value: any, encrypt: boolean = false): void {
    try {
      let serializedValue = JSON.stringify(value)

      // 如果需要加密
      if (encrypt) {
        serializedValue = CryptoUtil.encrypt(serializedValue)
      }

      localStorage.setItem(key, serializedValue)
    } catch (error) {
      // 静默处理存储失败
    }
  }

  // 获取存储项
  static get<T = any>(key: string, encrypted: boolean = false): T | null {
    try {
      let item = localStorage.getItem(key)
      if (item === null) {
        return null
      }

      // 如果是加密存储的
      if (encrypted) {
        item = CryptoUtil.decrypt(item)
      }

      return JSON.parse(item) as T
    } catch (error) {
      return null
    }
  }

  // 删除存储项
  static remove(key: string): void {
    try {
      localStorage.removeItem(key)
    } catch (error) {
      // 静默处理移除失败
    }
  }

  // 清空所有存储
  static clear(): void {
    try {
      localStorage.clear()
    } catch (error) {
      // 静默处理清理失败
    }
  }

  // 检查存储项是否存在
  static has(key: string): boolean {
    return localStorage.getItem(key) !== null
  }

  // 获取存储大小（KB）
  static getSize(): number {
    try {
      let total = 0
      for (let key in localStorage) {
        if (localStorage.hasOwnProperty(key)) {
          total += localStorage[key].length + key.length
        }
      }
      return Math.round(total / 1024 * 100) / 100 // 保留两位小数
    } catch (error) {
      return 0
    }
  }

  // 清理过期数据
  static cleanExpired(): void {
    try {
      const now = Date.now()
      const keysToRemove: string[] = []

      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i)
        if (key && key.endsWith('_expires')) {
          const expiryTime = parseInt(localStorage.getItem(key) || '0')
          if (expiryTime < now) {
            const dataKey = key.replace('_expires', '')
            keysToRemove.push(key, dataKey)
          }
        }
      }

      keysToRemove.forEach(key => localStorage.removeItem(key))
    } catch (error) {
      // 静默处理清理失败
    }
  }

  // 设置带过期时间的存储项
  static setWithExpiry(key: string, value: any, expiryMinutes: number, encrypt: boolean = false): void {
    const expiryTime = Date.now() + (expiryMinutes * 60 * 1000)
    this.set(key, value, encrypt)
    this.set(`${key}_expires`, expiryTime)
  }

  // 获取带过期时间的存储项
  static getWithExpiry<T = any>(key: string, encrypted: boolean = false): T | null {
    const expiryTime = this.get<number>(`${key}_expires`)

    if (expiryTime && Date.now() > expiryTime) {
      // 已过期，删除数据
      this.remove(key)
      this.remove(`${key}_expires`)
      return null
    }

    return this.get<T>(key, encrypted)
  }
}

// 存储键名常量
export const STORAGE_KEYS = {
  DOCTOR_TOKEN: 'doctor_token',
  DOCTOR_INFO: 'doctor_info',
  REMEMBER_PHONE: 'remember_phone',
  USER_SETTINGS: 'user_settings',
  LAST_LOGIN_TIME: 'last_login_time'
} as const

// 专门的token存储工具
export class TokenStorage {
  static setToken(token: string): void {
    Storage.set(STORAGE_KEYS.DOCTOR_TOKEN, token, true) // 加密存储token
  }

  static getToken(): string | null {
    return Storage.get<string>(STORAGE_KEYS.DOCTOR_TOKEN, true) // 解密获取token
  }

  static removeToken(): void {
    Storage.remove(STORAGE_KEYS.DOCTOR_TOKEN)
  }

  static hasToken(): boolean {
    return Storage.has(STORAGE_KEYS.DOCTOR_TOKEN)
  }
}

// 用户信息存储工具
export class UserInfoStorage {
  static setUserInfo(userInfo: any): void {
    Storage.set(STORAGE_KEYS.DOCTOR_INFO, userInfo, true) // 加密存储用户信息
  }

  static getUserInfo<T = any>(): T | null {
    return Storage.get<T>(STORAGE_KEYS.DOCTOR_INFO, true) // 解密获取用户信息
  }

  static removeUserInfo(): void {
    Storage.remove(STORAGE_KEYS.DOCTOR_INFO)
  }

  static hasUserInfo(): boolean {
    return Storage.has(STORAGE_KEYS.DOCTOR_INFO)
  }
}