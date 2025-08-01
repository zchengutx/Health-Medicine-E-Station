import React from 'react'
import { render, screen, fireEvent, waitFor, act } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { vi, describe, it, expect, beforeEach, afterEach } from 'vitest'
import { VerificationCodeInput } from '../VerificationCodeInput'
import type { VerificationCodeInputProps } from '../VerificationCodeInput'

// Mock validation utilities
vi.mock('../../../utils/validation', () => ({
  isValidPhone: vi.fn((phone: string) => phone === '13800138000'),
}))

describe('VerificationCodeInput', () => {
  const defaultProps: VerificationCodeInputProps = {
    value: '',
    onChange: vi.fn(),
    onValidationChange: vi.fn(),
    phoneNumber: '13800138000',
    onSendCode: vi.fn(),
    countdown: 0,
  }

  beforeEach(() => {
    vi.clearAllMocks()
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  describe('渲染测试', () => {
    it('应该正确渲染组件', () => {
      render(<VerificationCodeInput {...defaultProps} />)
      
      expect(screen.getByRole('textbox', { name: '验证码输入框' })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: '获取验证码' })).toBeInTheDocument()
    })

    it('应该显示正确的占位符文本', () => {
      render(<VerificationCodeInput {...defaultProps} placeholder="输入验证码" />)
      
      expect(screen.getByPlaceholderText('输入验证码')).toBeInTheDocument()
    })

    it('应该显示默认占位符文本', () => {
      render(<VerificationCodeInput {...defaultProps} />)
      
      expect(screen.getByPlaceholderText('请输入验证码')).toBeInTheDocument()
    })
  })

  describe('输入功能测试', () => {
    it('应该只允许数字输入', async () => {
      const user = userEvent.setup()
      const onChange = vi.fn()
      
      render(<VerificationCodeInput {...defaultProps} onChange={onChange} />)
      
      const input = screen.getByRole('textbox')
      
      await user.type(input, 'abc123def456')
      
      // 检查每次调用都只包含数字
      const calls = onChange.mock.calls
      console.log('onChange calls:', calls.map(call => call[0]))
      
      // 验证所有调用都只包含数字
      expect(calls.every(call => /^\d*$/.test(call[0]))).toBe(true)
      // 验证最终结果是6位数字
      expect(calls[calls.length - 1][0]).toBe('123456')
    })

    it('应该限制输入长度为6位', async () => {
      const user = userEvent.setup()
      const onChange = vi.fn()
      
      render(<VerificationCodeInput {...defaultProps} onChange={onChange} />)
      
      const input = screen.getByRole('textbox')
      
      await user.type(input, '1234567890')
      
      // 检查没有超过6位的调用
      const calls = onChange.mock.calls
      console.log('onChange calls for length test:', calls.map(call => call[0]))
      
      expect(calls.every(call => call[0].length <= 6)).toBe(true)
      expect(calls[calls.length - 1][0]).toBe('123456')
    })

    it('应该格式化显示验证码（3位后添加空格）', () => {
      render(<VerificationCodeInput {...defaultProps} value="123456" />)
      
      const input = screen.getByRole('textbox') as HTMLInputElement
      expect(input.value).toBe('123 456')
    })

    it('应该在输入6位数字时调用验证回调', async () => {
      const user = userEvent.setup()
      const onValidationChange = vi.fn()
      
      render(<VerificationCodeInput {...defaultProps} onValidationChange={onValidationChange} />)
      
      const input = screen.getByRole('textbox')
      
      await user.type(input, '123456')
      
      expect(onValidationChange).toHaveBeenCalledWith(true)
    })
  })

  describe('发送验证码功能测试', () => {
    it('应该在有效手机号时启用发送按钮', () => {
      render(<VerificationCodeInput {...defaultProps} phoneNumber="13800138000" />)
      
      const button = screen.getByRole('button', { name: '获取验证码' })
      expect(button).not.toBeDisabled()
    })

    it('应该在无效手机号时禁用发送按钮', () => {
      render(<VerificationCodeInput {...defaultProps} phoneNumber="invalid" />)
      
      const button = screen.getByRole('button', { name: '获取验证码' })
      expect(button).toBeDisabled()
    })

    it('应该在点击时调用发送验证码回调', async () => {
      const user = userEvent.setup()
      const onSendCode = vi.fn().mockResolvedValue(undefined)
      
      render(<VerificationCodeInput {...defaultProps} onSendCode={onSendCode} />)
      
      const button = screen.getByRole('button', { name: '获取验证码' })
      
      await user.click(button)
      
      expect(onSendCode).toHaveBeenCalledWith('13800138000')
    })

    it('应该在倒计时期间禁用按钮并显示倒计时', () => {
      render(<VerificationCodeInput {...defaultProps} countdown={30} />)
      
      const button = screen.getByRole('button', { name: '30秒后可重新发送' })
      expect(button).toBeDisabled()
      expect(button).toHaveTextContent('30s')
    })

    it('应该在发送中显示加载状态', async () => {
      const user = userEvent.setup()
      const onSendCode = vi.fn().mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
      
      render(<VerificationCodeInput {...defaultProps} onSendCode={onSendCode} />)
      
      const button = screen.getByRole('button')
      
      // 点击按钮开始发送
      await user.click(button)
      
      // 在发送过程中按钮应该显示加载状态
      expect(button).toHaveTextContent('发送中...')
      expect(button).toBeDisabled()
    })
  })

  describe('验证和错误处理测试', () => {
    it('应该在失焦时显示验证错误', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} value="123" />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      await user.tab() // 失焦
      
      await waitFor(() => {
        expect(screen.getByText('请输入6位验证码')).toBeInTheDocument()
      })
    })

    it('应该在空值失焦时显示必填错误', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      await user.tab() // 失焦
      
      await waitFor(() => {
        expect(screen.getByText('请输入验证码')).toBeInTheDocument()
      })
    })

    it('应该在获得焦点时清除错误信息', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} value="123" />)
      
      const input = screen.getByRole('textbox')
      
      // 先失焦显示错误
      await user.click(input)
      await user.tab()
      
      await waitFor(() => {
        expect(screen.getByText('请输入6位验证码')).toBeInTheDocument()
      })
      
      // 再次获得焦点应该清除错误
      await user.click(input)
      
      await waitFor(() => {
        expect(screen.queryByText('请输入6位验证码')).not.toBeInTheDocument()
      })
    })
  })

  describe('无障碍访问测试', () => {
    it('应该有正确的ARIA属性', () => {
      render(<VerificationCodeInput {...defaultProps} />)
      
      const input = screen.getByRole('textbox')
      expect(input).toHaveAttribute('aria-label', '验证码输入框')
      expect(input).toHaveAttribute('inputMode', 'numeric')
      expect(input).toHaveAttribute('autoComplete', 'one-time-code')
    })

    it('应该在有错误时设置aria-invalid', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} value="123" />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      await user.tab()
      
      await waitFor(() => {
        expect(input).toHaveAttribute('aria-invalid', 'true')
        expect(input).toHaveAttribute('aria-describedby', 'code-error')
      })
    })

    it('应该有正确的错误信息关联', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} value="123" />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      await user.tab()
      
      await waitFor(() => {
        const errorElement = screen.getByRole('alert')
        expect(errorElement).toHaveAttribute('id', 'code-error')
        expect(errorElement).toHaveAttribute('aria-live', 'polite')
      })
    })
  })

  describe('禁用状态测试', () => {
    it('应该在禁用时不响应输入', async () => {
      const user = userEvent.setup()
      const onChange = vi.fn()
      
      render(<VerificationCodeInput {...defaultProps} disabled onChange={onChange} />)
      
      const input = screen.getByRole('textbox')
      expect(input).toBeDisabled()
      
      await user.type(input, '123456')
      expect(onChange).not.toHaveBeenCalled()
    })

    it('应该在禁用时禁用发送按钮', () => {
      render(<VerificationCodeInput {...defaultProps} disabled />)
      
      const button = screen.getByRole('button')
      expect(button).toBeDisabled()
    })
  })

  describe('样式类名测试', () => {
    it('应该应用自定义类名', () => {
      render(<VerificationCodeInput {...defaultProps} className="custom-class" />)
      
      const input = screen.getByRole('textbox')
      expect(input).toHaveClass('custom-class')
    })

    it('应该在焦点状态时应用焦点样式', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      
      // CSS modules会生成类似 _focused_xxxxx 的类名
      expect(input.className).toMatch(/focused/)
    })

    it('应该在错误状态时应用错误样式', async () => {
      const user = userEvent.setup()
      
      render(<VerificationCodeInput {...defaultProps} value="123" />)
      
      const input = screen.getByRole('textbox')
      
      await user.click(input)
      await user.tab()
      
      await waitFor(() => {
        // CSS modules会生成类似 _error_xxxxx 的类名
        expect(input.className).toMatch(/error/)
      })
    })
  })

  describe('边界情况测试', () => {
    it('应该处理空的手机号', () => {
      render(<VerificationCodeInput {...defaultProps} phoneNumber="" />)
      
      const button = screen.getByRole('button')
      expect(button).toBeDisabled()
    })

    it('应该处理发送验证码失败的情况', async () => {
      const user = userEvent.setup()
      const onSendCode = vi.fn().mockRejectedValue(new Error('网络错误'))
      const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})
      
      render(<VerificationCodeInput {...defaultProps} onSendCode={onSendCode} />)
      
      const button = screen.getByRole('button')
      
      await user.click(button)
      
      await waitFor(() => {
        expect(consoleSpy).toHaveBeenCalledWith('发送验证码失败:', expect.any(Error))
      })
      
      consoleSpy.mockRestore()
    })

    it('应该处理快速连续点击发送按钮', async () => {
      const user = userEvent.setup()
      const onSendCode = vi.fn().mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
      
      render(<VerificationCodeInput {...defaultProps} onSendCode={onSendCode} />)
      
      const button = screen.getByRole('button')
      
      // 快速点击多次
      await user.click(button)
      await user.click(button)
      await user.click(button)
      
      // 应该只调用一次
      expect(onSendCode).toHaveBeenCalledTimes(1)
    })
  })
})