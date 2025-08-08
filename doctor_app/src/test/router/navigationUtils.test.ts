import { describe, it, expect, vi, beforeEach } from 'vitest'
import { useNavigationUtils } from '@/router/navigationUtils'

// Mock vue-router
const mockPush = vi.fn()
const mockReplace = vi.fn()
const mockBack = vi.fn()

vi.mock('vue-router', () => ({
  useRouter: () => ({
    push: mockPush,
    replace: mockReplace,
    back: mockBack,
    options: {
      history: {
        state: {
          back: '/previous-page'
        }
      }
    }
  })
}))

// Mock auth store
vi.mock('@/stores/auth', () => ({
  useAuthStore: () => ({
    isLoggedIn: false,
    doctorInfo: null
  })
}))

// Mock window.location
Object.defineProperty(window, 'location', {
  value: {
    href: ''
  },
  writable: true
})

describe('NavigationUtils', () => {
  let navigationUtils: ReturnType<typeof useNavigationUtils>

  beforeEach(() => {
    navigationUtils = useNavigationUtils()
    vi.clearAllMocks()
    window.location.href = ''
  })

  describe('safePush', () => {
    it('应该成功执行路由跳转', async () => {
      mockPush.mockResolvedValue(undefined)
      
      const result = await navigationUtils.safePush('/home')
      
      expect(mockPush).toHaveBeenCalledWith('/home')
      expect(result).toBe(true)
    })

    it('应该在路由跳转失败时使用备用方案', async () => {
      mockPush.mockRejectedValue(new Error('Navigation failed'))
      
      const result = await navigationUtils.safePush('/home')
      
      expect(mockPush).toHaveBeenCalledWith('/home')
      expect(window.location.href).toBe('/home')
      expect(result).toBe(false)
    })

    it('应该支持replace选项', async () => {
      mockReplace.mockResolvedValue(undefined)
      
      const result = await navigationUtils.safePush('/home', { replace: true })
      
      expect(mockReplace).toHaveBeenCalledWith('/home')
      expect(mockPush).not.toHaveBeenCalled()
      expect(result).toBe(true)
    })

    it('应该支持延迟选项', async () => {
      mockPush.mockResolvedValue(undefined)
      const startTime = Date.now()
      
      await navigationUtils.safePush('/home', { delay: 100 })
      
      const endTime = Date.now()
      expect(endTime - startTime).toBeGreaterThanOrEqual(100)
      expect(mockPush).toHaveBeenCalledWith('/home')
    })

    it('应该在fallback为false时不使用备用方案', async () => {
      mockPush.mockRejectedValue(new Error('Navigation failed'))
      
      const result = await navigationUtils.safePush('/home', { fallback: false })
      
      expect(mockPush).toHaveBeenCalledWith('/home')
      expect(window.location.href).toBe('')
      expect(result).toBe(false)
    })
  })

  describe('便捷导航方法', () => {
    it('toHome应该跳转到首页', () => {
      mockPush.mockResolvedValue(undefined)
      
      navigationUtils.toHome()
      
      expect(mockPush).toHaveBeenCalledWith('/home')
    })

    it('toLogin应该跳转到登录页', () => {
      mockPush.mockResolvedValue(undefined)
      
      navigationUtils.toLogin()
      
      expect(mockPush).toHaveBeenCalledWith('/login')
    })

    it('toRegister应该跳转到注册页', () => {
      mockPush.mockResolvedValue(undefined)
      
      navigationUtils.toRegister()
      
      expect(mockPush).toHaveBeenCalledWith('/register')
    })

    it('toSplash应该跳转到启动页', () => {
      mockPush.mockResolvedValue(undefined)
      
      navigationUtils.toSplash()
      
      expect(mockPush).toHaveBeenCalledWith('/')
    })
  })

  describe('forceToHome', () => {
    it('应该使用多重备用方案跳转到首页', () => {
      mockPush.mockResolvedValue(undefined)
      
      navigationUtils.forceToHome()
      
      expect(mockPush).toHaveBeenCalledWith('/home')
    })

    it('应该在push失败时尝试replace', async () => {
      mockPush.mockRejectedValue(new Error('Push failed'))
      mockReplace.mockResolvedValue(undefined)
      
      navigationUtils.forceToHome()
      
      // 等待异步操作完成
      await new Promise(resolve => setTimeout(resolve, 0))
      
      expect(mockPush).toHaveBeenCalledWith('/home')
      expect(mockReplace).toHaveBeenCalledWith('/home')
    })

    it('应该在replace也失败时使用window.location', async () => {
      mockPush.mockRejectedValue(new Error('Push failed'))
      mockReplace.mockRejectedValue(new Error('Replace failed'))
      
      navigationUtils.forceToHome()
      
      // 等待异步操作完成
      await new Promise(resolve => setTimeout(resolve, 0))
      
      expect(mockPush).toHaveBeenCalledWith('/home')
      expect(mockReplace).toHaveBeenCalledWith('/home')
      expect(window.location.href).toBe('/home')
    })
  })

  describe('goBack', () => {
    it('应该在有历史记录时返回上一页', () => {
      navigationUtils.goBack()
      
      expect(mockBack).toHaveBeenCalled()
    })

    it('应该在没有历史记录时跳转到首页', () => {
      // Mock没有历史记录的情况
      const mockRouter = {
        push: mockPush,
        replace: mockReplace,
        back: mockBack,
        options: {
          history: {
            state: {
              back: null
            }
          }
        }
      }
      
      vi.mocked(vi.importActual('vue-router')).useRouter = () => mockRouter
      
      const navUtils = useNavigationUtils()
      navUtils.goBack()
      
      expect(mockPush).toHaveBeenCalledWith('/home')
    })
  })

  describe('handleNavigationFallback', () => {
    it('应该尝试replace跳转', () => {
      mockReplace.mockResolvedValue(undefined)
      
      navigationUtils.handleNavigationFallback('/home')
      
      expect(mockReplace).toHaveBeenCalledWith('/home')
    })

    it('应该在replace失败时使用window.location', () => {
      mockReplace.mockRejectedValue(new Error('Replace failed'))
      
      navigationUtils.handleNavigationFallback('/home')
      
      expect(mockReplace).toHaveBeenCalledWith('/home')
      expect(window.location.href).toBe('/home')
    })
  })
})