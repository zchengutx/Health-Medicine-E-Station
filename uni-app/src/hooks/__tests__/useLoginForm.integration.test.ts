import { renderHook, act } from '@testing-library/react'
import { vi } from 'vitest'
import { useLoginForm } from '../useLoginForm'
import { loginWithVerificationCode, sendVerificationCode } from '../../services/authApi'
import { LoginResponse } from '../../types'

// Mock the authApi
vi.mock('../../services/authApi')
const mockLoginWithVerificationCode = loginWithVerificationCode as any
const mockSendVerificationCode = sendVerificationCode as any

describe('useLoginForm Integration', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.useFakeTimers()
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('should initialize with default values', () => {
    const { result } = renderHook(() => useLoginForm())
    
    expect(result.current.formData).toEqual({
      phone: '',
      verificationCode: '',
    })
    expect(result.current.isValid).toBe(false)
    expect(result.current.isSubmitting).toBe(false)
    expect(result.current.agreedToTerms).toBe(true)
    expect(result.current.hasError).toBe(false)
  })

  it('should handle phone input and validation', () => {
    const { result } = renderHook(() => useLoginForm())
    
    // Test invalid phone
    act(() => {
      result.current.phone.setValue('123')
    })
    
    expect(result.current.phone.value).toBe('123')
    expect(result.current.phone.isValid).toBe(false)
    expect(result.current.phone.error).toBe('手机号长度不足')
    expect(result.current.isValid).toBe(false)
    
    // Test valid phone
    act(() => {
      result.current.phone.setValue('13800138000')
    })
    
    expect(result.current.phone.value).toBe('13800138000')
    expect(result.current.phone.isValid).toBe(true)
    expect(result.current.phone.error).toBe('')
  })

  it('should handle verification code input and validation', () => {
    const { result } = renderHook(() => useLoginForm())
    
    // Test invalid code
    act(() => {
      result.current.verificationCode.setValue('123')
    })
    
    expect(result.current.verificationCode.value).toBe('123')
    expect(result.current.verificationCode.isValid).toBe(false)
    expect(result.current.verificationCode.error).toBe('验证码长度不足')
    
    // Test valid code
    act(() => {
      result.current.verificationCode.setValue('123456')
    })
    
    expect(result.current.verificationCode.value).toBe('123456')
    expect(result.current.verificationCode.isValid).toBe(true)
    expect(result.current.verificationCode.error).toBe('')
  })

  it('should handle terms agreement', () => {
    const { result } = renderHook(() => useLoginForm())
    
    expect(result.current.agreedToTerms).toBe(true)
    
    act(() => {
      result.current.setAgreedToTerms(false)
    })
    
    expect(result.current.agreedToTerms).toBe(false)
    expect(result.current.isValid).toBe(false)
  })

  it('should validate form correctly', () => {
    const { result } = renderHook(() => useLoginForm())
    
    // Initially invalid
    expect(result.current.isValid).toBe(false)
    
    // Set valid phone
    act(() => {
      result.current.phone.setValue('13800138000')
    })
    expect(result.current.isValid).toBe(false) // Still need code
    
    // Set valid code
    act(() => {
      result.current.verificationCode.setValue('123456')
    })
    expect(result.current.isValid).toBe(true) // Now valid
    
    // Disagree with terms
    act(() => {
      result.current.setAgreedToTerms(false)
    })
    expect(result.current.isValid).toBe(false) // Invalid again
  })

  it('should send verification code successfully', async () => {
    mockSendVerificationCode.mockResolvedValue({
      success: true,
      message: '验证码发送成功',
      countdown: 60,
    })

    const { result } = renderHook(() => useLoginForm())
    
    // Set valid phone first
    act(() => {
      result.current.phone.setValue('13800138000')
    })
    
    // Send verification code
    await act(async () => {
      await result.current.verificationCode.sendCode()
    })
    
    expect(mockSendVerificationCode).toHaveBeenCalledWith({
      phone: '13800138000',
    })
    expect(result.current.verificationCode.countdown).toBe(60)
    expect(result.current.verificationCode.isCountingDown).toBe(true)
    expect(result.current.verificationCode.canSendCode).toBe(false)
  })

  it('should submit form successfully', async () => {
    const mockLoginResponse: LoginResponse = {
      user: {
        id: '1',
        phone: '13800138000',
        nickname: 'Test User',
        createdAt: '2023-01-01',
        updatedAt: '2023-01-01',
      },
      token: 'test-token',
      refreshToken: 'test-refresh-token',
      expiresIn: 3600,
    }

    mockLoginWithVerificationCode.mockResolvedValue(mockLoginResponse)

    const { result } = renderHook(() => useLoginForm())
    
    // Set up valid form
    act(() => {
      result.current.phone.setValue('13800138000')
      result.current.verificationCode.setValue('123456')
    })
    
    expect(result.current.isValid).toBe(true)
    
    // Submit form
    let response: LoginResponse | null = null
    await act(async () => {
      response = await result.current.submitForm()
    })
    
    expect(mockLoginWithVerificationCode).toHaveBeenCalledWith({
      phone: '13800138000',
      verificationCode: '123456',
    })
    expect(response).toEqual(mockLoginResponse)
    
    // Form should be reset after successful login
    expect(result.current.phone.value).toBe('')
    expect(result.current.verificationCode.value).toBe('')
  })

  it('should handle form submission errors', async () => {
    const error = new Error('Login failed')
    mockLoginWithVerificationCode.mockRejectedValue(error)

    const { result } = renderHook(() => useLoginForm())
    
    // Set up valid form
    act(() => {
      result.current.phone.setValue('13800138000')
      result.current.verificationCode.setValue('123456')
    })
    
    // Submit form
    let response: LoginResponse | null = null
    await act(async () => {
      response = await result.current.submitForm()
    })
    
    expect(response).toBeNull()
    expect(result.current.hasError).toBe(true)
    expect(result.current.error).toBe('Login failed')
  })

  it('should reset form correctly', () => {
    const { result } = renderHook(() => useLoginForm())
    
    // Set up form with data
    act(() => {
      result.current.phone.setValue('13800138000')
      result.current.verificationCode.setValue('123456')
      result.current.setAgreedToTerms(false)
    })
    
    expect(result.current.phone.value).toBe('13800138000')
    expect(result.current.verificationCode.value).toBe('123456')
    expect(result.current.agreedToTerms).toBe(false)
    
    // Reset form
    act(() => {
      result.current.resetForm()
    })
    
    expect(result.current.phone.value).toBe('')
    expect(result.current.verificationCode.value).toBe('')
    expect(result.current.agreedToTerms).toBe(true)
    expect(result.current.hasError).toBe(false)
    expect(result.current.isSubmitting).toBe(false)
  })
})