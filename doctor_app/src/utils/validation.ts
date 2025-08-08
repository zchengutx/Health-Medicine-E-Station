// 表单验证工具函数

// 验证结果接口
export interface ValidationResult {
  isValid: boolean
  message: string
}

// 验证规则类型
export type ValidationRule = (value: string) => ValidationResult

// 手机号验证
export const validatePhone = (phone: string): boolean => {
  // 支持更多手机号段
  const phoneRegex = /^1[3-9]\d{9}$/
  return phoneRegex.test(phone)
}

// 验证码验证
export const validateSmsCode = (code: string): boolean => {
  const codeRegex = /^\d{4,6}$/
  return codeRegex.test(code)
}

// 密码强度验证
export const validatePassword = (password: string): ValidationResult => {
  if (!password) {
    return {
      isValid: false,
      message: '请输入密码'
    }
  }
  
  if (password.length < 6) {
    return {
      isValid: false,
      message: '密码长度不能少于6位'
    }
  }
  
  if (password.length > 20) {
    return {
      isValid: false,
      message: '密码长度不能超过20位'
    }
  }
  
  // 检查是否包含空格
  if (/\s/.test(password)) {
    return {
      isValid: false,
      message: '密码不能包含空格'
    }
  }
  
  // 检查是否包含字母和数字
  const hasLetter = /[a-zA-Z]/.test(password)
  const hasNumber = /\d/.test(password)
  
  if (!hasLetter || !hasNumber) {
    return {
      isValid: false,
      message: '密码必须包含字母和数字'
    }
  }
  
  // 检查是否包含特殊字符（可选，提高安全性）
  const hasSpecialChar = /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password)
  
  if (password.length >= 8 && hasSpecialChar) {
    return {
      isValid: true,
      message: '密码强度很强'
    }
  } else if (password.length >= 8) {
    return {
      isValid: true,
      message: '密码强度良好'
    }
  } else {
    return {
      isValid: true,
      message: '密码强度一般'
    }
  }
}

// 通用非空验证
export const validateRequired = (value: string, fieldName: string): ValidationResult => {
  if (!value || value.trim() === '') {
    return {
      isValid: false,
      message: `${fieldName}不能为空`
    }
  }
  
  return {
    isValid: true,
    message: ''
  }
}

// 邮箱验证
export const validateEmail = (email: string): ValidationResult => {
  if (!email) {
    return {
      isValid: false,
      message: '请输入邮箱地址'
    }
  }
  
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(email)) {
    return {
      isValid: false,
      message: '请输入正确的邮箱格式'
    }
  }
  
  return {
    isValid: true,
    message: ''
  }
}

// 身份证号验证
export const validateIdCard = (idCard: string): ValidationResult => {
  if (!idCard) {
    return {
      isValid: false,
      message: '请输入身份证号'
    }
  }
  
  const idCardRegex = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/
  if (!idCardRegex.test(idCard)) {
    return {
      isValid: false,
      message: '请输入正确的身份证号格式'
    }
  }
  
  return {
    isValid: true,
    message: ''
  }
}

// 验证手机号格式并返回错误信息
export const getPhoneValidationMessage = (phone: string): string => {
  if (!phone) {
    return '请输入手机号'
  }
  
  if (!validatePhone(phone)) {
    return '请输入正确的手机号格式'
  }
  
  return ''
}

// 验证验证码格式并返回错误信息
export const getSmsCodeValidationMessage = (code: string): string => {
  if (!code) {
    return '请输入验证码'
  }
  
  if (!validateSmsCode(code)) {
    return '请输入4-6位数字验证码'
  }
  
  return ''
}

// 表单验证器类
export class FormValidator {
  private rules: Map<string, ValidationRule[]> = new Map()
  
  // 添加验证规则
  addRule(field: string, rule: ValidationRule): FormValidator {
    if (!this.rules.has(field)) {
      this.rules.set(field, [])
    }
    this.rules.get(field)!.push(rule)
    return this
  }
  
  // 验证单个字段
  validateField(field: string, value: string): ValidationResult {
    const fieldRules = this.rules.get(field)
    if (!fieldRules) {
      return { isValid: true, message: '' }
    }
    
    for (const rule of fieldRules) {
      const result = rule(value)
      if (!result.isValid) {
        return result
      }
    }
    
    return { isValid: true, message: '' }
  }
  
  // 验证整个表单
  validateForm(formData: Record<string, string>): {
    isValid: boolean
    errors: Record<string, string>
  } {
    const errors: Record<string, string> = {}
    let isValid = true
    
    for (const [field, value] of Object.entries(formData)) {
      const result = this.validateField(field, value)
      if (!result.isValid) {
        errors[field] = result.message
        isValid = false
      }
    }
    
    return { isValid, errors }
  }
}

// 预定义的验证规则
export const ValidationRules = {
  required: (fieldName: string): ValidationRule => 
    (value: string) => validateRequired(value, fieldName),
    
  phone: (): ValidationRule => 
    (value: string) => ({
      isValid: validatePhone(value),
      message: getPhoneValidationMessage(value)
    }),
    
  smsCode: (): ValidationRule => 
    (value: string) => ({
      isValid: validateSmsCode(value),
      message: getSmsCodeValidationMessage(value)
    }),
    
  password: (): ValidationRule => 
    (value: string) => validatePassword(value),
    
  email: (): ValidationRule => 
    (value: string) => validateEmail(value),
    
  idCard: (): ValidationRule => 
    (value: string) => validateIdCard(value),
    
  minLength: (min: number, fieldName: string): ValidationRule => 
    (value: string) => ({
      isValid: value.length >= min,
      message: value.length < min ? `${fieldName}长度不能少于${min}位` : ''
    }),
    
  maxLength: (max: number, fieldName: string): ValidationRule => 
    (value: string) => ({
      isValid: value.length <= max,
      message: value.length > max ? `${fieldName}长度不能超过${max}位` : ''
    })
}