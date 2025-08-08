import { describe, it, expect } from 'vitest'
import {
  validatePhone,
  validateSmsCode,
  validatePassword,
  validateRequired,
  getPhoneValidationMessage,
  getSmsCodeValidationMessage,
  FormValidator,
  ValidationRules
} from '@/utils/validation'

describe('validation utils', () => {
  describe('validatePhone', () => {
    it('should validate correct phone numbers', () => {
      expect(validatePhone('13812345678')).toBe(true)
      expect(validatePhone('15987654321')).toBe(true)
      expect(validatePhone('18612345678')).toBe(true)
    })

    it('should reject invalid phone numbers', () => {
      expect(validatePhone('12812345678')).toBe(false) // starts with 12
      expect(validatePhone('1381234567')).toBe(false) // too short
      expect(validatePhone('138123456789')).toBe(false) // too long
      expect(validatePhone('abcdefghijk')).toBe(false) // not numbers
      expect(validatePhone('')).toBe(false) // empty
    })
  })

  describe('validateSmsCode', () => {
    it('should validate correct SMS codes', () => {
      expect(validateSmsCode('1234')).toBe(true)
      expect(validateSmsCode('123456')).toBe(true)
      expect(validateSmsCode('5678')).toBe(true)
    })

    it('should reject invalid SMS codes', () => {
      expect(validateSmsCode('123')).toBe(false) // too short
      expect(validateSmsCode('1234567')).toBe(false) // too long
      expect(validateSmsCode('abcd')).toBe(false) // not numbers
      expect(validateSmsCode('')).toBe(false) // empty
    })
  })

  describe('validatePassword', () => {
    it('should validate strong passwords', () => {
      const result = validatePassword('Abc123456!')
      expect(result.isValid).toBe(true)
      expect(result.message).toContain('很强')
    })

    it('should validate good passwords', () => {
      const result = validatePassword('Abc12345')
      expect(result.isValid).toBe(true)
      expect(result.message).toContain('良好')
    })

    it('should validate basic passwords', () => {
      const result = validatePassword('abc123')
      expect(result.isValid).toBe(true)
      expect(result.message).toContain('一般')
    })

    it('should reject invalid passwords', () => {
      expect(validatePassword('').isValid).toBe(false)
      expect(validatePassword('12345').isValid).toBe(false) // too short
      expect(validatePassword('abcdef').isValid).toBe(false) // no numbers
      expect(validatePassword('123456').isValid).toBe(false) // no letters
      expect(validatePassword('abc 123').isValid).toBe(false) // contains space
    })
  })

  describe('validateRequired', () => {
    it('should validate non-empty values', () => {
      const result = validateRequired('test', '测试字段')
      expect(result.isValid).toBe(true)
      expect(result.message).toBe('')
    })

    it('should reject empty values', () => {
      const result = validateRequired('', '测试字段')
      expect(result.isValid).toBe(false)
      expect(result.message).toBe('测试字段不能为空')
    })

    it('should reject whitespace-only values', () => {
      const result = validateRequired('   ', '测试字段')
      expect(result.isValid).toBe(false)
      expect(result.message).toBe('测试字段不能为空')
    })
  })

  describe('getPhoneValidationMessage', () => {
    it('should return empty string for valid phone', () => {
      expect(getPhoneValidationMessage('13812345678')).toBe('')
    })

    it('should return error message for invalid phone', () => {
      expect(getPhoneValidationMessage('')).toBe('请输入手机号')
      expect(getPhoneValidationMessage('123')).toBe('请输入正确的手机号格式')
    })
  })

  describe('getSmsCodeValidationMessage', () => {
    it('should return empty string for valid code', () => {
      expect(getSmsCodeValidationMessage('1234')).toBe('')
    })

    it('should return error message for invalid code', () => {
      expect(getSmsCodeValidationMessage('')).toBe('请输入验证码')
      expect(getSmsCodeValidationMessage('123')).toBe('请输入4-6位数字验证码')
    })
  })

  describe('FormValidator', () => {
    it('should validate form with multiple rules', () => {
      const validator = new FormValidator()
      validator
        .addRule('phone', ValidationRules.required('手机号'))
        .addRule('phone', ValidationRules.phone())
        .addRule('password', ValidationRules.required('密码'))
        .addRule('password', ValidationRules.password())

      const formData = {
        phone: '13812345678',
        password: 'Abc123456'
      }

      const result = validator.validateForm(formData)
      expect(result.isValid).toBe(true)
      expect(Object.keys(result.errors)).toHaveLength(0)
    })

    it('should return errors for invalid form', () => {
      const validator = new FormValidator()
      validator
        .addRule('phone', ValidationRules.required('手机号'))
        .addRule('phone', ValidationRules.phone())

      const formData = {
        phone: '123'
      }

      const result = validator.validateForm(formData)
      expect(result.isValid).toBe(false)
      expect(result.errors.phone).toBeDefined()
    })
  })
})