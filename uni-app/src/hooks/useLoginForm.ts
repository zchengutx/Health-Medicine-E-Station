import { useState, useCallback, useMemo } from 'react'
import { LoginFormData, LoginResponse, User } from '../types'
import { usePhoneValidation } from './usePhoneValidation'
import { useVerificationCode } from './useVerificationCode'
import { useErrorHandler } from './useErrorHandler'
import { loginWithVerificationCode } from '../services/authApi'

/**
 * 登录表单Hook
 * 管理整体表单状态，集成手机号验证、验证码管理等功能
 */
export interface UseLoginFormReturn {
  /** 表单数据 */
  formData: LoginFormData
  /** 表单是否有效 */
  isValid: boolean
  /** 是否正在提交 */
  isSubmitting: boolean
  /** 是否同意用户协议 */
  agreedToTerms: boolean
  /** 手机号验证相关 */
  phone: {
    value: string
    formattedValue: string
    isValid: boolean
    error: string
    isDirty: boolean
    setValue: (value: string) => void
    clear: () => void
    validate: () => boolean
  }
  /** 验证码相关 */
  verificationCode: {
    value: string
    formattedValue: string
    isValid: boolean
    error: string
    isDirty: boolean
    countdown: number
    isCountingDown: boolean
    canSendCode: boolean
    isSending: boolean
    setValue: (value: string) => void
    sendCode: () => Promise<void>
    clear: () => void
    stopCountdown: () => void
    validate: () => boolean
  }
  /** 错误处理 */
  error: string | null
  hasError: boolean
  clearError: () => void
  /** 表单操作 */
  setAgreedToTerms: (agreed: boolean) => void
  submitForm: () => Promise<LoginResponse | null>
  resetForm: () => void
  validateForm: () => boolean
}

export const useLoginForm = (): UseLoginFormReturn => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [agreedToTerms, setAgreedToTerms] = useState(true) // 默认同意协议
  
  // 使用手机号验证Hook
  const phoneValidation = usePhoneValidation()
  
  // 使用验证码Hook
  const verificationCodeHook = useVerificationCode()
  
  // 使用错误处理Hook
  const { error, hasError, handleError, clearError, setRetry } = useErrorHandler()

  // 表单数据
  const formData = useMemo((): LoginFormData => ({
    phone: phoneValidation.cleanedPhone,
    verificationCode: verificationCodeHook.cleanedCode,
  }), [phoneValidation.cleanedPhone, verificationCodeHook.cleanedCode])

  // 表单是否有效
  const isValid = useMemo(() => {
    return (
      phoneValidation.isValid &&
      verificationCodeHook.isValid &&
      agreedToTerms
    )
  }, [phoneValidation.isValid, verificationCodeHook.isValid, agreedToTerms])

  // 发送验证码（使用手机号）
  const sendVerificationCode = useCallback(async () => {
    if (!phoneValidation.isValid) {
      phoneValidation.validate()
      return
    }
    
    try {
      await verificationCodeHook.sendCode(phoneValidation.cleanedPhone)
    } catch (error) {
      handleError(error as Error)
    }
  }, [phoneValidation, verificationCodeHook, handleError])

  // 提交表单
  const submitForm = useCallback(async (): Promise<LoginResponse | null> => {
    // 清除之前的错误
    clearError()
    
    // 验证表单
    if (!validateForm()) {
      return null
    }

    setIsSubmitting(true)
    
    try {
      const response = await loginWithVerificationCode(formData)
      
      // 登录成功，清空表单
      resetForm()
      
      return response
    } catch (error) {
      handleError(error as Error)
      
      // 设置重试函数
      setRetry(() => submitForm)
      
      return null
    } finally {
      setIsSubmitting(false)
    }
  }, [formData, clearError, handleError, setRetry])

  // 验证整个表单
  const validateForm = useCallback((): boolean => {
    const phoneValid = phoneValidation.validate()
    const codeValid = verificationCodeHook.validate()
    
    if (!phoneValid || !codeValid) {
      return false
    }
    
    if (!agreedToTerms) {
      handleError(new Error('请同意用户协议和隐私政策'))
      return false
    }
    
    return true
  }, [phoneValidation, verificationCodeHook, agreedToTerms, handleError])

  // 重置表单
  const resetForm = useCallback(() => {
    phoneValidation.clear()
    verificationCodeHook.clear()
    verificationCodeHook.stopCountdown()
    setAgreedToTerms(true)
    clearError()
    setIsSubmitting(false)
  }, [phoneValidation, verificationCodeHook, clearError])

  return {
    formData,
    isValid,
    isSubmitting,
    agreedToTerms,
    phone: {
      value: phoneValidation.phone,
      formattedValue: phoneValidation.formattedPhone,
      isValid: phoneValidation.isValid,
      error: phoneValidation.error,
      isDirty: phoneValidation.isDirty,
      setValue: phoneValidation.setPhone,
      clear: phoneValidation.clear,
      validate: phoneValidation.validate,
    },
    verificationCode: {
      value: verificationCodeHook.code,
      formattedValue: verificationCodeHook.formattedCode,
      isValid: verificationCodeHook.isValid,
      error: verificationCodeHook.error,
      isDirty: verificationCodeHook.isDirty,
      countdown: verificationCodeHook.countdown,
      isCountingDown: verificationCodeHook.isCountingDown,
      canSendCode: verificationCodeHook.canSendCode,
      isSending: verificationCodeHook.isSending,
      setValue: verificationCodeHook.setCode,
      sendCode: sendVerificationCode,
      clear: verificationCodeHook.clear,
      stopCountdown: verificationCodeHook.stopCountdown,
      validate: verificationCodeHook.validate,
    },
    error,
    hasError,
    clearError,
    setAgreedToTerms,
    submitForm,
    resetForm,
    validateForm,
  }
}