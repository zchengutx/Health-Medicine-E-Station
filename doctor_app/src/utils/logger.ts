/**
 * 安全日志管理工具
 * 用于控制开发和生产环境的日志输出，防止敏感信息泄露
 */

// 日志级别枚举
export enum LogLevel {
  DEBUG = 0,
  INFO = 1,
  WARN = 2,
  ERROR = 3,
  NONE = 4
}

// 敏感信息字段列表
const SENSITIVE_FIELDS = [
  'token', 'password', 'phone', 'email', 'doctorId', 'DId',
  'licenseNumber', 'birthDate', 'avatar', 'doctorCode'
]

class Logger {
  private level: LogLevel
  private isDevelopment: boolean

  constructor() {
    // 简化环境检查
    this.isDevelopment = true // 暂时设为开发模式，避免类型错误
    // 生产环境只显示错误日志，开发环境显示所有日志
    this.level = this.isDevelopment ? LogLevel.DEBUG : LogLevel.ERROR
  }

  /**
   * 清理敏感信息
   */
  private sanitizeData(data: any): any {
    if (data === null || data === undefined) {
      return data
    }

    if (typeof data === 'string') {
      // 检查是否是敏感字符串（如token）
      if (data.length > 20 && /^[A-Za-z0-9+/=]+$/.test(data)) {
        return '[REDACTED_TOKEN]'
      }
      return data
    }

    if (typeof data === 'number' || typeof data === 'boolean') {
      return data
    }

    if (Array.isArray(data)) {
      return data.map(item => this.sanitizeData(item))
    }

    if (typeof data === 'object') {
      const sanitized: any = {}
      for (const [key, value] of Object.entries(data)) {
        const lowerKey = key.toLowerCase()
        
        // 检查是否是敏感字段
        if (SENSITIVE_FIELDS.some(field => lowerKey.includes(field.toLowerCase()))) {
          if (typeof value === 'string' && value.length > 0) {
            // 对于手机号，只显示前3位和后4位
            if (lowerKey.includes('phone') && typeof value === 'string' && value.length === 11) {
              sanitized[key] = value.substring(0, 3) + '****' + value.substring(7)
            } else if (lowerKey.includes('email') && typeof value === 'string') {
              const [username, domain] = value.split('@')
              if (domain) {
                sanitized[key] = username.substring(0, 2) + '***@' + domain
              } else {
                sanitized[key] = '[REDACTED_EMAIL]'
              }
            } else {
              sanitized[key] = '[REDACTED]'
            }
          } else {
            sanitized[key] = '[REDACTED]'
          }
        } else {
          sanitized[key] = this.sanitizeData(value)
        }
      }
      return sanitized
    }

    return data
  }

  /**
   * 格式化日志消息
   */
  private formatMessage(level: string, message: string, data?: any): string {
    const timestamp = new Date().toISOString()
    const prefix = `[${timestamp}] [${level}]`
    
    if (data !== undefined) {
      const sanitizedData = this.sanitizeData(data)
      return `${prefix} ${message} ${JSON.stringify(sanitizedData, null, 2)}`
    }
    
    return `${prefix} ${message}`
  }

  /**
   * Debug级别日志
   */
  debug(message: string, data?: any): void {
    if (this.level <= LogLevel.DEBUG) {
      console.log(this.formatMessage('DEBUG', message, data))
    }
  }

  /**
   * Info级别日志
   */
  info(message: string, data?: any): void {
    if (this.level <= LogLevel.INFO) {
      console.info(this.formatMessage('INFO', message, data))
    }
  }

  /**
   * Warning级别日志
   */
  warn(message: string, data?: any): void {
    if (this.level <= LogLevel.WARN) {
      console.warn(this.formatMessage('WARN', message, data))
    }
  }

  /**
   * Error级别日志
   */
  error(message: string, error?: any): void {
    if (this.level <= LogLevel.ERROR) {
      if (error instanceof Error) {
        console.error(this.formatMessage('ERROR', message), error.message, error.stack)
      } else {
        console.error(this.formatMessage('ERROR', message, error))
      }
    }
  }

  /**
   * 设置日志级别
   */
  setLevel(level: LogLevel): void {
    this.level = level
  }

  /**
   * 获取当前日志级别
   */
  getLevel(): LogLevel {
    return this.level
  }

  /**
   * 检查是否应该输出指定级别的日志
   */
  shouldLog(level: LogLevel): boolean {
    return this.level <= level
  }
}

// 创建全局日志实例
export const logger = new Logger()

// 导出便捷方法
export const log = {
  debug: (message: string, data?: any) => logger.debug(message, data),
  info: (message: string, data?: any) => logger.info(message, data),
  warn: (message: string, data?: any) => logger.warn(message, data),
  error: (message: string, error?: any) => logger.error(message, error)
}

export default logger