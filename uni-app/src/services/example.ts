/**
 * API服务使用示例
 * 这个文件展示了如何使用API服务层的各种功能
 */

import { 
  sendVerificationCode, 
  loginWithVerificationCode, 
  logout, 
  validateToken 
} from './authApi'

/**
 * 示例：完整的登录流程
 */
export const exampleLoginFlow = async () => {
  try {
    const phone = '13800138000'
    
    // 1. 发送验证码
    console.log('发送验证码...')
    const codeResponse = await sendVerificationCode({ phone })
    console.log('验证码发送成功:', codeResponse)
    
    // 2. 模拟用户输入验证码（实际应用中由用户输入）
    const verificationCode = '123456'
    
    // 3. 登录
    console.log('登录中...')
    const loginResponse = await loginWithVerificationCode({
      phone,
      verificationCode,
    })
    console.log('登录成功:', loginResponse)
    
    // 4. 验证token是否有效
    const isValid = await validateToken()
    console.log('Token有效性:', isValid)
    
    return loginResponse
  } catch (error) {
    console.error('登录流程失败:', error)
    throw error
  }
}

/**
 * 示例：错误处理
 */
export const exampleErrorHandling = async () => {
  try {
    // 尝试使用无效的手机号
    await sendVerificationCode({ phone: '123' })
  } catch (error) {
    console.log('捕获到验证错误:', (error as Error).message)
  }
  
  try {
    // 尝试使用无效的验证码
    await loginWithVerificationCode({
      phone: '13800138000',
      verificationCode: '123',
    })
  } catch (error) {
    console.log('捕获到登录错误:', (error as Error).message)
  }
}

/**
 * 示例：登出流程
 */
export const exampleLogoutFlow = async () => {
  try {
    console.log('登出中...')
    await logout()
    console.log('登出成功')
    
    // 验证token是否已清除
    const isValid = await validateToken()
    console.log('登出后Token有效性:', isValid) // 应该为false
  } catch (error) {
    console.error('登出失败:', error)
  }
}

/**
 * 示例：检查当前登录状态
 */
export const exampleCheckLoginStatus = async () => {
  const isLoggedIn = await validateToken()
  
  if (isLoggedIn) {
    console.log('用户已登录')
    const userInfo = localStorage.getItem('user_info')
    if (userInfo) {
      console.log('用户信息:', JSON.parse(userInfo))
    }
  } else {
    console.log('用户未登录')
  }
  
  return isLoggedIn
}