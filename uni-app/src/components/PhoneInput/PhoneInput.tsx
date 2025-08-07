import React, { useState, useCallback, ChangeEvent, FocusEvent } from 'react'
import { isValidPhone, formatPhone, cleanPhone, getPhoneError } from '../../utils/validation'
import styles from './PhoneInput.module.css'

export interface PhoneInputProps {
  value: string
  onChange: (value: string) => void
  onValidationChange: (isValid: boolean) => void
  disabled?: boolean
  placeholder?: string
  className?: string
}

export const PhoneInput: React.FC<PhoneInputProps> = ({
  value,
  onChange,
  onValidationChange,
  disabled = false,
  placeholder = '请输入手机号码',
  className = ''
}) => {
  const [isFocused, setIsFocused] = useState(false)
  const [errorMessage, setErrorMessage] = useState('')
  const [hasBlurred, setHasBlurred] = useState(false)

  // 处理输入变化
  const handleInputChange = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    const inputValue = event.target.value
    
    // 只允许数字输入，最大11位
    const cleanedValue = inputValue.replace(/\D/g, '').slice(0, 11)
    
    // 更新值
    onChange(cleanedValue)
    
    // 验证并更新状态
    const isValid = isValidPhone(cleanedValue)
    onValidationChange(isValid)
    
    // 如果已经失焦过，实时显示错误
    if (hasBlurred) {
      const error = getPhoneError(cleanedValue)
      setErrorMessage(error)
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
    const error = getPhoneError(value)
    setErrorMessage(error)
  }, [value])

  // 计算显示值（格式化）
  const displayValue = formatPhone(value)
  
  // 计算样式类名
  const inputClassName = [
    styles.input,
    isFocused && styles.focused,
    errorMessage && styles.error,
    disabled && styles.disabled,
    className
  ].filter(Boolean).join(' ')

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
          maxLength={13} // 格式化后的最大长度 (138 0013 8000)
          inputMode="numeric"
          autoComplete="tel"
          aria-label="手机号码输入框"
          aria-invalid={!!errorMessage}
          aria-describedby={errorMessage ? 'phone-error' : undefined}
        />
        {/* 手机号前缀显示 */}
        <span className={styles.prefix}>+86</span>
      </div>
      
      {/* 错误信息显示 */}
      {errorMessage && (
        <div 
          id="phone-error" 
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

export default PhoneInput