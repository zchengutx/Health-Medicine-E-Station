import { vi } from 'vitest'
import { config } from '@vue/test-utils'

// Mock Vant components
vi.mock('vant', () => ({
  showToast: vi.fn(),
  showDialog: vi.fn(),
  showConfirmDialog: vi.fn(),
  showLoadingToast: vi.fn(),
  closeToast: vi.fn()
}))

// Mock navigator.vibrate
Object.defineProperty(navigator, 'vibrate', {
  writable: true,
  value: vi.fn()
})

// Mock localStorage
const localStorageMock = {
  getItem: vi.fn(),
  setItem: vi.fn(),
  removeItem: vi.fn(),
  clear: vi.fn()
}
Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
})

// Mock window.location
Object.defineProperty(window, 'location', {
  value: {
    href: 'http://localhost:3000',
    pathname: '/',
    search: '',
    hash: ''
  },
  writable: true
})

// Global test configuration
config.global.mocks = {
  $t: (key: string) => key // Mock i18n if needed
}