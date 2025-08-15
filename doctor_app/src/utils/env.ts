/**
 * 环境检查工具
 * 解决 import.meta.env 的 TypeScript 类型问题
 */

// 安全的环境检查函数
export const isDevelopment = (): boolean => {
  try {
    // 首先检查 Vite 环境变量
    if (typeof import.meta !== 'undefined' && import.meta.env) {
      return import.meta.env.DEV === true
    }
    
    // 备用检查 Node.js 环境变量
    if (typeof process !== 'undefined' && process.env) {
      return process.env.NODE_ENV === 'development'
    }
    
    // 默认为开发环境
    return true
  } catch (error) {
    // 如果检查失败，默认为开发环境
    return true
  }
}

export const isProduction = (): boolean => {
  try {
    // 首先检查 Vite 环境变量
    if (typeof import.meta !== 'undefined' && import.meta.env) {
      return import.meta.env.PROD === true
    }
    
    // 备用检查 Node.js 环境变量
    if (typeof process !== 'undefined' && process.env) {
      return process.env.NODE_ENV === 'production'
    }
    
    // 默认为开发环境
    return false
  } catch (error) {
    // 如果检查失败，默认为开发环境
    return false
  }
}

// 获取环境名称
export const getEnvironment = (): string => {
  if (isProduction()) {
    return 'production'
  }
  return 'development'
}