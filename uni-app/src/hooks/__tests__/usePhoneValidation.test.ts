import { renderHook, act } from '@testing-library/react'
import { usePhoneValidation } from '../usePhoneValidation'

describe('usePhoneValidation', () => {
  it('should initialize with empty phone number', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    expect(result.current.phone).toBe('')
    expect(result.current.formattedPhone).toBe('')
    expect(result.current.cleanedPhone).toBe('')
    expect(result.current.isValid).toBe(false)
    expect(result.current.error).toBe('')
    expect(result.current.isDirty).toBe(false)
  })

  it('should initialize with provided initial phone', () => {
    const { result } = renderHook(() => usePhoneValidation('13800138000'))
    
    expect(result.current.phone).toBe('13800138000')
    expect(result.current.formattedPhone).toBe('138 0013 8000')
    expect(result.current.cleanedPhone).toBe('13800138000')
    expect(result.current.isValid).toBe(true)
    expect(result.current.isDirty).toBe(false)
  })

  it('should set phone number and mark as dirty', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('138')
    })
    
    expect(result.current.phone).toBe('138')
    expect(result.current.formattedPhone).toBe('138')
    expect(result.current.cleanedPhone).toBe('138')
    expect(result.current.isValid).toBe(false)
    expect(result.current.isDirty).toBe(true)
  })

  it('should format phone number correctly', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('13800138000')
    })
    
    expect(result.current.formattedPhone).toBe('138 0013 8000')
  })

  it('should validate phone number format', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    // Valid phone number
    act(() => {
      result.current.setPhone('13800138000')
    })
    expect(result.current.isValid).toBe(true)
    expect(result.current.error).toBe('')
    
    // Invalid phone number
    act(() => {
      result.current.setPhone('123')
    })
    expect(result.current.isValid).toBe(false)
    expect(result.current.error).toBe('手机号长度不足')
  })

  it('should show error messages when dirty', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    // No error when not dirty
    expect(result.current.error).toBe('')
    
    // Show error when dirty and invalid
    act(() => {
      result.current.setPhone('123')
    })
    expect(result.current.error).toBe('手机号长度不足')
  })

  it('should limit phone number length to 11 digits', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('138001380001234')
    })
    
    // Should not update if more than 11 digits
    expect(result.current.phone).toBe('')
    expect(result.current.cleanedPhone).toBe('')
  })

  it('should clear phone number and reset dirty state', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('13800138000')
    })
    
    expect(result.current.phone).toBe('13800138000')
    expect(result.current.isDirty).toBe(true)
    
    act(() => {
      result.current.clear()
    })
    
    expect(result.current.phone).toBe('')
    expect(result.current.isDirty).toBe(false)
    expect(result.current.error).toBe('')
  })

  it('should manually validate and set dirty state', () => {
    const { result } = renderHook(() => usePhoneValidation('13800138000'))
    
    expect(result.current.isDirty).toBe(false)
    
    let isValid: boolean
    act(() => {
      isValid = result.current.validate()
    })
    
    expect(isValid).toBe(true)
    expect(result.current.isDirty).toBe(true)
  })

  it('should handle formatted input correctly', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('138 0013 8000')
    })
    
    expect(result.current.phone).toBe('138 0013 8000')
    expect(result.current.cleanedPhone).toBe('13800138000')
    expect(result.current.isValid).toBe(true)
  })

  it('should handle non-numeric characters', () => {
    const { result } = renderHook(() => usePhoneValidation())
    
    act(() => {
      result.current.setPhone('138-0013-8000')
    })
    
    expect(result.current.cleanedPhone).toBe('13800138000')
    expect(result.current.isValid).toBe(true)
  })
})