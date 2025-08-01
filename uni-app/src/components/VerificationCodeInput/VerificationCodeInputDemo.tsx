import React, { useState } from 'react'
import { VerificationCodeInput } from './VerificationCodeInput'

/**
 * VerificationCodeInput 组件演示
 * 用于开发和测试组件功能
 */
export const VerificationCodeInputDemo: React.FC = () => {
  const [code, setCode] = useState('')
  const [isValid, setIsValid] = useState(false)
  const [phoneNumber] = useState('13800138000') // 模拟手机号
  const [countdown, setCountdown] = useState(0)

  // 模拟发送验证码
  const handleSendCode = async (phone: string) => {
    console.log('发送验证码到:', phone)
    
    // 模拟API调用延迟
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // 开始倒计时
    setCountdown(60)
    const timer = setInterval(() => {
      setCountdown(prev => {
        if (prev <= 1) {
          clearInterval(timer)
          return 0
        }
        return prev - 1
      })
    }, 1000)
    
    console.log('验证码发送成功')
  }

  const handleCodeChange = (value: string) => {
    setCode(value)
    console.log('验证码变化:', value)
  }

  const handleValidationChange = (valid: boolean) => {
    setIsValid(valid)
    console.log('验证状态:', valid)
  }

  return (
    <div style={{ 
      padding: '20px', 
      maxWidth: '400px', 
      margin: '0 auto',
      fontFamily: 'system-ui, -apple-system, sans-serif'
    }}>
      <h2>VerificationCodeInput 组件演示</h2>
      
      <div style={{ marginBottom: '20px' }}>
        <p><strong>手机号:</strong> {phoneNumber}</p>
        <p><strong>当前验证码:</strong> {code || '(空)'}</p>
        <p><strong>验证状态:</strong> {isValid ? '✅ 有效' : '❌ 无效'}</p>
        <p><strong>倒计时:</strong> {countdown > 0 ? `${countdown}秒` : '无'}</p>
      </div>

      <VerificationCodeInput
        value={code}
        onChange={handleCodeChange}
        onValidationChange={handleValidationChange}
        phoneNumber={phoneNumber}
        onSendCode={handleSendCode}
        countdown={countdown}
        placeholder="请输入6位验证码"
      />

      <div style={{ marginTop: '20px', fontSize: '14px', color: '#666' }}>
        <h3>测试说明:</h3>
        <ul>
          <li>输入框只允许数字，最多6位</li>
          <li>输入3位后会自动添加空格分隔</li>
          <li>点击"获取验证码"按钮开始60秒倒计时</li>
          <li>倒计时期间按钮不可点击</li>
          <li>输入6位数字后验证状态变为有效</li>
          <li>失焦时会显示验证错误信息</li>
        </ul>
      </div>
    </div>
  )
}

export default VerificationCodeInputDemo