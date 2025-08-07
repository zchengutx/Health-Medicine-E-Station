/**
 * 手机号验证工具函数
 */

// 中国大陆手机号正则表达式
const PHONE_REGEX = /^1[3-9]\d{9}$/

/**
 * 验证手机号格式是否正确
 * @param phone 手机号字符串
 * @returns 是否为有效手机号
 */
export const isValidPhone = (phone: string): boolean => {
  if (!phone || typeof phone !== 'string') {
    return false
  }

  // 移除所有空格和特殊字符
  const cleanPhone = phone.replace(/\s|-/g, '')

  return PHONE_REGEX.test(cleanPhone)
}

/**
 * 格式化手机号显示（添加空格分隔）
 * @param phone 原始手机号
 * @returns 格式化后的手机号 (例: 138 0013 8000)
 */
export const formatPhone = (phone: string): string => {
  if (!phone) return ''

  const cleanPhone = phone.replace(/\D/g, '')

  if (cleanPhone.length <= 3) {
    return cleanPhone
  } else if (cleanPhone.length <= 7) {
    return `${cleanPhone.slice(0, 3)} ${cleanPhone.slice(3)}`
  } else {
    return `${cleanPhone.slice(0, 3)} ${cleanPhone.slice(3, 7)} ${cleanPhone.slice(7, 11)}`
  }
}

/**
 * 清理手机号（移除格式化字符）
 * @param phone 格式化的手机号
 * @returns 纯数字手机号
 */
export const cleanPhone = (phone: string): string => {
  return phone.replace(/\D/g, '')
}

/**
 * 获取手机号验证错误信息
 * @param phone 手机号
 * @returns 错误信息，如果有效则返回空字符串
 */
export const getPhoneError = (phone: string): string => {
  if (!phone) {
    return '请输入手机号'
  }

  const cleanedPhone = cleanPhone(phone)

  if (cleanedPhone.length === 0) {
    return '请输入手机号'
  }

  if (cleanedPhone.length < 11) {
    return '手机号长度不足'
  }

  if (cleanedPhone.length > 11) {
    return '手机号长度超出限制'
  }

  if (!isValidPhone(cleanedPhone)) {
    return '请输入正确的手机号格式'
  }

  return ''
}

/**
 * 验证码验证工具函数
 */

// 验证码正则表达式（6位数字）
const VERIFICATION_CODE_REGEX = /^\d{6}$/

/**
 * 验证验证码格式是否正确
 * @param code 验证码字符串
 * @returns 是否为有效验证码
 */
export const isValidVerificationCode = (code: string): boolean => {
  if (!code || typeof code !== 'string') {
    return false
  }

  // 移除所有空格
  const cleanCode = code.replace(/\s/g, '')

  return VERIFICATION_CODE_REGEX.test(cleanCode)
}

/**
 * 格式化验证码显示（添加空格分隔）
 * @param code 原始验证码
 * @returns 格式化后的验证码 (例: 123 456)
 */
export const formatVerificationCode = (code: string): string => {
  if (!code) return ''

  const cleanCode = code.replace(/\D/g, '')

  if (cleanCode.length <= 3) {
    return cleanCode
  } else {
    return `${cleanCode.slice(0, 3)} ${cleanCode.slice(3, 6)}`
  }
}

/**
 * 清理验证码（移除格式化字符）
 * @param code 格式化的验证码
 * @returns 纯数字验证码
 */
export const cleanVerificationCode = (code: string): string => {
  return code.replace(/\D/g, '')
}

/**
 * 获取验证码验证错误信息
 * @param code 验证码
 * @returns 错误信息，如果有效则返回空字符串
 */
export const getVerificationCodeError = (code: string): string => {
  if (!code) {
    return '请输入验证码'
  }

  const cleanedCode = cleanVerificationCode(code)

  if (cleanedCode.length === 0) {
    return '请输入验证码'
  }

  if (cleanedCode.length < 6) {
    return '验证码长度不足'
  }

  if (cleanedCode.length > 6) {
    return '验证码长度超出限制'
  }

  if (!isValidVerificationCode(cleanedCode)) {
    return '请输入正确的验证码格式'
  }

  return ''
}
