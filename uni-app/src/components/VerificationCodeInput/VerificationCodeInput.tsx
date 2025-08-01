import React, { useState, useCallback, ChangeEvent, FocusEvent } from 'react'
import { isValidPhone } from '../../utils/validation'
import styles from './VerificationCodeInput.module.css'

export interface VerificationCodeInputProps {
  value: string
  onChange: (value: string) => void
  onValidationChange: (isValid: boolean) => void
  phoneNumber: string
  onSendCode: (phone: string) => Promise<void>
  countdown: number
  disabled?: boolean
  placeholder?: string
  className?: string
}

export const VerificationCodeInput: React.FC<VerificationCodeInputProps> = ({
  value,
  onChange,
  onValidationChange,
  phoneNumber,
  onSendCode,
  countdown,
  disabled = false,
  placeholder = '请输入验证码',
  className = ''
}) => {
  const [isFocused, setIsFocused] = useState(false)
  const [errorMessage, setErrorMessage] = useState('')
  const [hasBlurred, setHasBlurred] = useState(false)
  const [isSending, setIsSending] = useState(false)

  // 处理输入变化
  const handleInputChange = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    const inputValue = event.target.value
    
    // 只允许数字输入，最大6位
    const cleanedValue = inputValue.replace(/\D/g, '').slice(0, 6)
    
    // 更新值
    onChange(cleanedValue)
    
    // 验证并更新状态
    const isValidCode = cleanedValue.length === 6
    onValidationChange(isValidCode)
    
    // 如果已经失焦过，实时显示错误
    if (hasBlurred && cleanedValue.length > 0) {
      const errorMsg = cleanedValue.length < 6 ? '验证码长度不足' : ''
      setErrorMessage(errorMsg)
    }
  }, [onChange, onValidationChange, hasBlurred])

  // 处理焦点获得
  const handleFocus = useCallback((event: FocusEvent<HTMLInputElement>) => {
    setIsFocused(true)
    // 清除错误信息（获得焦点时）
    if (errorMessage) {
      setErrorMessage('')
    }
  }, [errorMessage])

  // 处理焦点失去
  const handleBlur = useCallback((event: FocusEvent<HTMLInputElement>) => {
    setIsFocused(false)
    setHasBlurred(true)
    
    // 失焦时验证并显示错误
    if (value && value.length > 0 && value.length < 6) {
      setErrorMessage('请输入6位验证码')
    } else if (value.length === 0) {
      setErrorMessage('请输入验证码')
    } else {
      setErrorMessage('')
    }
  }, [value])

  // 处理发送验证码
  const handleSendCode = useCallback(async () => {
    const isCountingDown = countdown > 0
    const canSend = !isCountingDown && !isSending && isValidPhone(phoneNumber)
    
    if (!canSend) {
      return
    }

    setIsSending(true)
    try {
      await onSendCode(phoneNumber)
    } catch (error) {
      console.error('发送验证码失败:', error)
    } finally {
      setIsSending(false)
    }
  }, [countdown, phoneNumber, isSending, onSendCode])

  // 计算显示值（格式化）
  const displayValue = value.length <= 3 ? value : `${value.slice(0, 3)} ${value.slice(3, 6)}`
  
  // 计算状态
  const isCountingDown = countdown > 0
  const canSendCode = !isCountingDown && !isSending && isValidPhone(phoneNumber)
  
  // 计算样式类名
  const inputClassName = [
    styles.input,
    isFocused && styles.focused,
    errorMessage && styles.error,
    disabled && styles.disabled,
    className
  ].filter(Boolean).join(' ')

  // 计算按钮样式类名
  const buttonClassName = [
    styles.sendButton,
    !canSendCode && styles.buttonDisabled,
    isCountingDown && styles.buttonCountdown
  ].filter(Boolean).join(' ')

  // 按钮文本
  const getButtonText = () => {
    if (isSending) return '发送中...'
    if (isCountingDown) return `${countdown}s`
    return '获取验证码'
  }

  return (
    <div className={styles.container}>
      <div className={styles.inputWrapper}>
        <input
          type="text"
          value={displayValue}
          onChange={handleInputChange}
          onFocus={handleFocus}
          onBlur={handleBlur}
          disabled={disabled}
          placeholder={placeholder}
          className={inputClassName}
          maxLength={7} // 格式化后的最大长度 (123 456)
          inputMode="numeric"
          autoComplete="one-time-code"
          aria-label="验证码输入框"
          aria-invalid={!!errorMessage}
          aria-describedby={errorMessage ? 'code-error' : undefined}
        />
        
        <button
          type="button"
          onClick={handleSendCode}
          disabled={!canSendCode || disabled}
          className={buttonClassName}
          aria-label={isCountingDown ? `${countdown}秒后可重新发送` : '获取验证码'}
        >
          {getButtonText()}
        </button>
      </div>
      
      {/* 错误信息显示 */}
      {errorMessage && (
        <div 
          id="code-error" 
          className={styles.errorMessage}
          role="alert"
          aria-live="polite"
        >
          {errorMessage}
        </div>
      )}
    </div>
  )
}

export default VerificationCodeInput