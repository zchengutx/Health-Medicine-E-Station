/**
 * 生产环境日志配置
 * 在生产环境中完全禁用敏感信息的日志输出
 */

// 重写console方法，在生产环境中禁用
export const setupProductionLogger = () => {
  const isProduction = process.env.NODE_ENV === 'production'
  
  if (isProduction) {
    // 生产环境中禁用所有console输出
    const noop = () => {}
    
    // 保留error日志用于错误监控，但清理敏感信息
    const originalError = console.error
    console.error = (...args: any[]) => {
      // 过滤敏感信息
      const filteredArgs = args.map(arg => {
        if (typeof arg === 'string') {
          // 移除可能的敏感信息
          return arg
            .replace(/token[:\s]*[A-Za-z0-9+/=]{20,}/gi, 'token: [REDACTED]')
            .replace(/phone[:\s]*1[3-9]\d{9}/gi, 'phone: [REDACTED]')
            .replace(/password[:\s]*\S+/gi, 'password: [REDACTED]')
            .replace(/email[:\s]*\S+@\S+/gi, 'email: [REDACTED]')
        }
        return arg
      })
      
      originalError.apply(console, filteredArgs)
    }
    
    // 禁用其他console方法
    console.log = noop
    console.info = noop
    console.warn = noop
    console.debug = noop
    console.trace = noop
    console.group = noop
    console.groupEnd = noop
    console.table = noop
    
    // 添加全局错误处理
    window.addEventListener('error', (event) => {
      // 只记录错误类型和文件信息，不记录具体内容
      console.error('Application Error:', {
        message: 'An error occurred',
        filename: event.filename,
        lineno: event.lineno,
        colno: event.colno
      })
    })
    
    window.addEventListener('unhandledrejection', (event) => {
      console.error('Unhandled Promise Rejection:', {
        message: 'Promise rejection occurred',
        type: typeof event.reason
      })
    })
  }
}

// 开发环境的console增强
export const setupDevelopmentLogger = () => {
  const isDevelopment = process.env.NODE_ENV === 'development' || !process.env.NODE_ENV
  
  if (isDevelopment) {
    // 开发环境中添加时间戳和样式
    const originalLog = console.log
    const originalError = console.error
    const originalWarn = console.warn
    const originalInfo = console.info
    
    const getTimestamp = () => new Date().toLocaleTimeString()
    
    console.log = (...args: any[]) => {
      originalLog.apply(console, [`%c[${getTimestamp()}] LOG`, 'color: #2196F3', ...args])
    }
    
    console.error = (...args: any[]) => {
      originalError.apply(console, [`%c[${getTimestamp()}] ERROR`, 'color: #F44336; font-weight: bold', ...args])
    }
    
    console.warn = (...args: any[]) => {
      originalWarn.apply(console, [`%c[${getTimestamp()}] WARN`, 'color: #FF9800', ...args])
    }
    
    console.info = (...args: any[]) => {
      originalInfo.apply(console, [`%c[${getTimestamp()}] INFO`, 'color: #4CAF50', ...args])
    }
  }
}