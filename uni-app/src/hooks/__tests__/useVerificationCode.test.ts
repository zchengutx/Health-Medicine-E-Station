import { renderHook, act, waitFor } from '@testing-library/react'
import { vi } from 'vitest'
import { useVerificationCode } from '../useVerificationCode'
import { sendVerificationCode } from '../../services/authApi'

// Mock the authApi
vi.mock('../../services/authApi')
const mockSendVerificationCode = sendVerificationCode as any

// Mock the useErrorHandler hook
vi.mock('../useErrorHandler', () => ({
  useErrorHandler: () => ({
    handleError: vi.fn(),
  }),
}))

describe('useVerificationCode', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.useFakeTimers()
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('should initialize with empty verification code', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    expect(result.current.code).toBe('')
    expect(result.current.formattedCode).toBe('')
    expect(result.current.cleanedCode).toBe('')
    expect(result.current.isValid).toBe(false)
    expect(result.current.error).toBe('')
    expect(result.current.isDirty).toBe(false)
    expect(result.current.countdown).toBe(0)
    expect(result.current.isCountingDown).toBe(false)
    expect(result.current.canSendCode).toBe(true)
    expect(result.current.isSending).toBe(false)
  })

  it('should initialize with provided initial code', () => {
    const { result } = renderHook(() => useVerificationCode('123456'))
    
    expect(result.current.code).toBe('123456')
    expect(result.current.formattedCode).toBe('123 456')
    expect(result.current.cleanedCode).toBe('123456')
    expect(result.current.isValid).toBe(true)
    expect(result.current.isDirty).toBe(false)
  })

  it('should set verification code and mark as dirty', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('123')
    })
    
    expect(result.current.code).toBe('123')
    expect(result.current.formattedCode).toBe('123')
    expect(result.current.cleanedCode).toBe('123')
    expect(result.current.isValid).toBe(false)
    expect(result.current.isDirty).toBe(true)
  })

  it('should format verification code correctly', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('123456')
    })
    
    expect(result.current.formattedCode).toBe('123 456')
  })

  it('should validate verification code format', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    // Valid code
    act(() => {
      result.current.setCode('123456')
    })
    expect(result.current.isValid).toBe(true)
    expect(result.current.error).toBe('')
    
    // Invalid code
    act(() => {
      result.current.setCode('123')
    })
    expect(result.current.isValid).toBe(false)
    expect(result.current.error).toBe('验证码长度不足')
  })

  it('should limit verification code length to 6 digits', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('1234567890')
    })
    
    // Should not update if more than 6 digits
    expect(result.current.code).toBe('')
    expect(result.current.cleanedCode).toBe('')
  })

  it('should send verification code successfully', async () => {
    mockSendVerificationCode.mockResolvedValue({
      success: true,
      message: '验证码发送成功',
      countdown: 60,
    })

    const { result } = renderHook(() => useVerificationCode())
    
    await act(async () => {
      await result.current.sendCode('13800138000')
    })
    
    expect(mockSendVerificationCode).toHaveBeenCalledWith({
      phone: '13800138000',
    })
    expect(result.current.countdown).toBe(60)
    expect(result.current.isCountingDown).toBe(true)
    expect(result.current.canSendCode).toBe(false)
  })

  it('should handle countdown timer correctly', async () => {
    mockSendVerificationCode.mockResolvedValue({
      success: true,
      message: '验证码发送成功',
      countdown: 3,
    })

    const { result } = renderHook(() => useVerificationCode())
    
    await act(async () => {
      await result.current.sendCode('13800138000')
    })
    
    expect(result.current.countdown).toBe(3)
    expect(result.current.isCountingDown).toBe(true)
    
    // Fast forward 1 second
    act(() => {
      vi.advanceTimersByTime(1000)
    })
    expect(result.current.countdown).toBe(2)
    
    // Fast forward 2 more seconds
    act(() => {
      vi.advanceTimersByTime(2000)
    })
    expect(result.current.countdown).toBe(0)
    expect(result.current.isCountingDown).toBe(false)
    expect(result.current.canSendCode).toBe(true)
  })

  it('should prevent sending code when already counting down', async () => {
    mockSendVerificationCode.mockResolvedValue({
      success: true,
      message: '验证码发送成功',
      countdown: 60,
    })

    const { result } = renderHook(() => useVerificationCode())
    
    // Send first code
    await act(async () => {
      await result.current.sendCode('13800138000')
    })
    
    expect(mockSendVerificationCode).toHaveBeenCalledTimes(1)
    expect(result.current.isCountingDown).toBe(true)
    
    // Try to send again while counting down
    await act(async () => {
      await result.current.sendCode('13800138000')
    })
    
    // Should not call API again
    expect(mockSendVerificationCode).toHaveBeenCalledTimes(1)
  })

  it('should stop countdown manually', async () => {
    mockSendVerificationCode.mockResolvedValue({
      success: true,
      message: '验证码发送成功',
      countdown: 60,
    })

    const { result } = renderHook(() => useVerificationCode())
    
    await act(async () => {
      await result.current.sendCode('13800138000')
    })
    
    expect(result.current.isCountingDown).toBe(true)
    
    act(() => {
      result.current.stopCountdown()
    })
    
    expect(result.current.countdown).toBe(0)
    expect(result.current.isCountingDown).toBe(false)
    expect(result.current.canSendCode).toBe(true)
  })

  it('should clear verification code and reset dirty state', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('123456')
    })
    
    expect(result.current.code).toBe('123456')
    expect(result.current.isDirty).toBe(true)
    
    act(() => {
      result.current.clear()
    })
    
    expect(result.current.code).toBe('')
    expect(result.current.isDirty).toBe(false)
    expect(result.current.error).toBe('')
  })

  it('should manually validate and set dirty state', () => {
    const { result } = renderHook(() => useVerificationCode('123456'))
    
    expect(result.current.isDirty).toBe(false)
    
    let isValid: boolean
    act(() => {
      isValid = result.current.validate()
    })
    
    expect(isValid).toBe(true)
    expect(result.current.isDirty).toBe(true)
  })

  it('should handle formatted input correctly', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('123 456')
    })
    
    expect(result.current.code).toBe('123 456')
    expect(result.current.cleanedCode).toBe('123456')
    expect(result.current.isValid).toBe(true)
  })

  it('should handle non-numeric characters', () => {
    const { result } = renderHook(() => useVerificationCode())
    
    act(() => {
      result.current.setCode('1a2b3c4d5e6f')
    })
    
    expect(result.current.cleanedCode).toBe('123456')
    expect(result.current.isValid).toBe(true)
  })
})