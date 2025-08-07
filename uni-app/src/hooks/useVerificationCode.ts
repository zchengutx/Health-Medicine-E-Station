import { useState, useCallback, useEffect, useRef, useMemo } from 'react'
import { isValidVerificationCode, getVerificationCodeError, cleanVerificationCode, formatVerificationCode } from '../utils/validation'
import { sendVerificationCode } from '../services/authApi'
import { useErrorHandler } from './useErrorHandler'

/**
 * 验证码Hook
 * 管理验证码状态、倒计时、发送验证码等功能
 */
export interface UseVerificationCodeReturn {
  /** 验证码值 */
  code: string
  /** 格式化后的验证码（用于显示） */
  formattedCode: string
  /** 清理后的验证码（纯数字） */
  cleanedCode: string
  /** 是否为有效验证码 */
  isValid: boolean
  /** 验证错误信息 */
  error: string
  /** 是否已输入内容 */
  isDirty: boolean
  /** 倒计时秒数 */
  countdown: number
  /** 是否正在倒计时 */
  isCountingDown: boolean
  /** 是否可以发送验证码 */
  canSendCode: boolean
  /** 是否正在发送验证码 */
  isSending: boolean
  /** 设置验证码 */
  setCode: (code: string) => void
  /** 发送验证码 */
  sendCode: (phone: string) => Promise<void>
  /** 清空验证码 */
  clear: () => void
  /** 停止倒计时 */
  stopCountdown: () => void
  /** 手动触发验证 */
  validate: () => boolean
}

const DEFAULT_COUNTDOWN = 60 // 默认倒计时60秒

export const useVerificationCode = (initialCode = ''): UseVerificationCodeReturn => {
  const [code, setCodeState] = useState(initialCode)
  const [isDirty, setIsDirty] = useState(false)
  const [countdown, setCountdown] = useState(0)
  const [isSending, setIsSending] = useState(false)
  
  const { handleError } = useErrorHandler()
  const countdownTimerRef = useRef<NodeJS.Timeout | null>(null)

  // 格式化后的验证码
  const formattedCode = useMemo(() => formatVerificationCode(code), [code])
  
  // 清理后的验证码
  const cleanedCode = useMemo(() => cleanVerificationCode(code), [code])
  
  // 验证结果
  const isValid = useMemo(() => isValidVerificationCode(cleanedCode), [cleanedCode])
  
  // 错误信息
  const error = useMemo(() => {
    if (!isDirty) return ''
    return getVerificationCodeError(code)
  }, [code, isDirty])

  // 是否正在倒计时
  const isCountingDown = countdown > 0
  
  // 是否可以发送验证码
  const canSendCode = !isCountingDown && !isSending

  // 设置验证码
  const setCode = useCallback((newCode: string) => {
    // 限制输入长度，最多6位数字
    const cleaned = cleanVerificationCode(newCode)
    if (cleaned.length > 6) return
    
    setCodeState(newCode)
    setIsDirty(true)
  }, [])

  // 开始倒计时
  const startCountdown = useCallback((seconds = DEFAULT_COUNTDOWN) => {
    setCountdown(seconds)
    
    if (countdownTimerRef.current) {
      clearInterval(countdownTimerRef.current)
    }
    
    countdownTimerRef.current = setInterval(() => {
      setCountdown((prev) => {
        if (prev <= 1) {
          if (countdownTimerRef.current) {
            clearInterval(countdownTimerRef.current)
            countdownTimerRef.current = null
          }
          return 0
        }
        return prev - 1
      })
    }, 1000)
  }, [])

  // 停止倒计时
  const stopCountdown = useCallback(() => {
    if (countdownTimerRef.current) {
      clearInterval(countdownTimerRef.current)
      countdownTimerRef.current = null
    }
    setCountdown(0)
  }, [])

  // 发送验证码
  const sendCode = useCallback(async (phone: string) => {
    if (!phone || isSending || isCountingDown) {
      return
    }

    setIsSending(true)
    
    try {
      const response = await sendVerificationCode({ phone })
      
      if (response.success) {
        // 开始倒计时
        startCountdown(response.countdown || DEFAULT_COUNTDOWN)
      } else {
        handleError(new Error(response.message || '发送验证码失败'))
      }
    } catch (error) {
      handleError(error as Error)
    } finally {
      setIsSending(false)
    }
  }, [isSending, isCountingDown, startCountdown, handleError])

  // 清空验证码
  const clear = useCallback(() => {
    setCodeState('')
    setIsDirty(false)
  }, [])

  // 手动验证
  const validate = useCallback(() => {
    setIsDirty(true)
    return isValidVerificationCode(cleanedCode)
  }, [cleanedCode])

  // 组件卸载时清理定时器
  useEffect(() => {
    return () => {
      if (countdownTimerRef.current) {
        clearInterval(countdownTimerRef.current)
      }
    }
  }, [])

  return {
    code,
    formattedCode,
    cleanedCode,
    isValid,
    error,
    isDirty,
    countdown,
    isCountingDown,
    canSendCode,
    isSending,
    setCode,
    sendCode,
    clear,
    stopCountdown,
    validate,
  }
}