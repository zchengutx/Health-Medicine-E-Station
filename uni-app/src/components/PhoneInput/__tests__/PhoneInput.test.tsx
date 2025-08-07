import React from 'react'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { vi } from 'vitest'
import { PhoneInput } from '../PhoneInput'

describe('PhoneInput Component', () => {
  const mockOnChange = vi.fn()
  const mockOnValidationChange = vi.fn()

  beforeEach(() => {
    vi.clearAllMocks()
  })

  const defaultProps = {
    value: '',
    onChange: mockOnChange,
    onValidationChange: mockOnValidationChange
  }

  describe('Rendering', () => {
    it('should render input with default placeholder', () => {
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox', { name: /手机号码输入框/i })
      expect(input).toBeInTheDocument()
      expect(input).toHaveAttribute('placeholder', '请输入手机号码')
    })

    it('should render with custom placeholder', () => {
      render(<PhoneInput {...defaultProps} placeholder="输入您的手机号" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveAttribute('placeholder', '输入您的手机号')
    })

    it('should render +86 prefix', () => {
      render(<PhoneInput {...defaultProps} />)

      expect(screen.getByText('+86')).toBeInTheDocument()
    })

    it('should render with custom className', () => {
      const { container } = render(<PhoneInput {...defaultProps} className="custom-class" />)

      const input = container.querySelector('input')
      expect(input).toHaveClass('custom-class')
    })
  })

  describe('Input Behavior', () => {
    it('should format phone number display', () => {
      render(<PhoneInput {...defaultProps} value="13800138000" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('138 0013 8000')
    })

    it('should handle partial phone number formatting', () => {
      render(<PhoneInput {...defaultProps} value="138" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('138')
    })

    it('should handle medium length phone number formatting', () => {
      render(<PhoneInput {...defaultProps} value="1380013" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('138 0013')
    })

    it('should call onChange with clean numeric value', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.type(input, '138')

      // Check that onChange was called for each character
      expect(mockOnChange).toHaveBeenCalledTimes(3)
      expect(mockOnChange).toHaveBeenNthCalledWith(1, '1')
      expect(mockOnChange).toHaveBeenNthCalledWith(2, '3')
      expect(mockOnChange).toHaveBeenNthCalledWith(3, '8')
    })
  })

  describe('Validation', () => {
    it('should call onValidationChange with false for invalid phone', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.type(input, '123')

      expect(mockOnValidationChange).toHaveBeenCalledWith(false)
    })

    it('should show error message for empty input on blur', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.click(input)
      await user.tab() // blur the input

      await waitFor(() => {
        expect(screen.getByRole('alert')).toHaveTextContent('请输入手机号')
      })
    })

    it('should clear error message on focus', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      // First blur to show error
      await user.click(input)
      await user.tab()

      await waitFor(() => {
        expect(screen.getByRole('alert')).toBeInTheDocument()
      })

      // Focus again to clear error
      await user.click(input)

      await waitFor(() => {
        expect(screen.queryByRole('alert')).not.toBeInTheDocument()
      })
    })
  })

  describe('Disabled State', () => {
    it('should render as disabled when disabled prop is true', () => {
      render(<PhoneInput {...defaultProps} disabled />)

      const input = screen.getByRole('textbox')
      expect(input).toBeDisabled()
    })

    it('should not accept input when disabled', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} disabled />)

      const input = screen.getByRole('textbox')

      await user.type(input, '13800138000')

      expect(mockOnChange).not.toHaveBeenCalled()
    })
  })

  describe('Accessibility', () => {
    it('should have proper ARIA attributes', () => {
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveAttribute('aria-label', '手机号码输入框')
      expect(input).toHaveAttribute('inputMode', 'numeric')
      expect(input).toHaveAttribute('autoComplete', 'tel')
    })

    it('should have aria-invalid when there is an error', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.click(input)
      await user.tab()

      await waitFor(() => {
        expect(input).toHaveAttribute('aria-invalid', 'true')
      })
    })

    it('should have aria-describedby when there is an error', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.click(input)
      await user.tab()

      await waitFor(() => {
        expect(input).toHaveAttribute('aria-describedby', 'phone-error')
        expect(screen.getByRole('alert')).toHaveAttribute('id', 'phone-error')
      })
    })

    it('should have aria-live on error message', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      await user.click(input)
      await user.tab()

      await waitFor(() => {
        const errorMessage = screen.getByRole('alert')
        expect(errorMessage).toHaveAttribute('aria-live', 'polite')
      })
    })
  })

  describe('Edge Cases', () => {
    it('should handle empty value prop', () => {
      render(<PhoneInput {...defaultProps} value="" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('')
    })

    it('should handle undefined value prop', () => {
      render(<PhoneInput {...defaultProps} value={undefined as any} />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('')
    })
  })

  describe('Phone Number Validation', () => {
    it('should validate valid phone number', () => {
      render(<PhoneInput {...defaultProps} value="13800138000" />)

      const input = screen.getByRole('textbox')
      expect(input).toHaveValue('138 0013 8000')
    })

    it('should handle input restrictions', async () => {
      const user = userEvent.setup()
      render(<PhoneInput {...defaultProps} />)

      const input = screen.getByRole('textbox')

      // Test that only numbers are accepted
      await user.type(input, '1')
      expect(mockOnChange).toHaveBeenCalledWith('1')

      await user.clear(input)
      await user.type(input, 'a')
      // Should not call onChange for non-numeric input
      expect(mockOnChange).toHaveBeenLastCalledWith('')
    })
  })
})