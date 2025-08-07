import { useState, useCallback, useMemo } from 'react'
import { isValidPhone, getPhoneError, cleanPhone, formatPhone } from '../utils/validation'

/**
 * 手机号验证Hook
 * 处理手机号输入、格式化、验证等功能
 */
export interface UsePhoneValidationReturn {
  /** 原始手机号值 */
  phone: string
  /** 格式化后的手机号（用于显示） */
  formattedPhone: string
  /** 清理后的手机号（纯数字） */
  cleanedPhone: string
  /** 是否为有效手机号 */
  isValid: boolean
  /** 验证错误信息 */
  error: string
  /** 是否已输入内容 */
  isDirty: boolean
  /** 设置手机号 */
  setPhone: (phone: string) => void
  /** 清空手机号 */
  clear: () => void
  /** 手动触发验证 */
  validate: () => boolean
}

export const usePhoneValidation = (initialPhone = ''): UsePhoneValidationReturn => {
  const [phone, setPhoneState] = useState(initialPhone)
  const [isDirty, setIsDirty] = useState(false)

  // 格式化后的手机号
  const formattedPhone = useMemo(() => formatPhone(phone), [phone])
  
  // 清理后的手机号
  const cleanedPhone = useMemo(() => cleanPhone(phone), [phone])
  
  // 验证结果
  const isValid = useMemo(() => isValidPhone(cleanedPhone), [cleanedPhone])
  
  // 错误信息
  const error = useMemo(() => {
    if (!isDirty) return ''
    return getPhoneError(phone)
  }, [phone, isDirty])

  // 设置手机号
  const setPhone = useCallback((newPhone: string) => {
    // 限制输入长度，最多11位数字
    const cleaned = cleanPhone(newPhone)
    if (cleaned.length > 11) return
    
    setPhoneState(newPhone)
    setIsDirty(true)
  }, [])

  // 清空手机号
  const clear = useCallback(() => {
    setPhoneState('')
    setIsDirty(false)
  }, [])

  // 手动验证
  const validate = useCallback(() => {
    setIsDirty(true)
    return isValidPhone(cleanedPhone)
  }, [cleanedPhone])

  return {
    phone,
    formattedPhone,
    cleanedPhone,
    isValid,
    error,
    isDirty,
    setPhone,
    clear,
    validate,
  }
}