import { describe, it, expect, beforeEach, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import ProfileView from '@/views/ProfileView.vue'
import { useAuthStore } from '@/stores/auth'
import DoctorApiService from '@/api/doctor'

// Mock the API service
vi.mock('@/api/doctor', () => ({
  default: vi.fn().mockImplementation(() => ({
    getProfile: vi.fn(),
    updateProfile: vi.fn()
  }))
}))

// Mock the auth store
vi.mock('@/stores/auth', () => ({
  useAuthStore: vi.fn()
}))

// Mock components
vi.mock('@/components/FeedbackButton.vue', () => ({
  default: {
    name: 'FeedbackButton',
    template: '<button @click="$emit(\'click\')"><slot>{{ text }}</slot></button>',
    props: ['text', 'loading', 'type', 'block'],
    emits: ['click']
  }
}))

vi.mock('@/components/ToastMessage.vue', () => ({
  default: {
    name: 'ToastMessage',
    template: '<div v-if="visible">{{ message }}</div>',
    props: ['message', 'type', 'visible'],
    emits: ['close']
  }
}))

vi.mock('@/components/LoadingSpinner.vue', () => ({
  default: {
    name: 'LoadingSpinner',
    template: '<div class="loading-spinner">Loading...</div>'
  }
}))

describe('ProfileView - Loading Fix', () => {
  let mockAuthStore: any
  let mockDoctorApi: any
  let router: any

  beforeEach(() => {
    setActivePinia(createPinia())
    
    // Create router
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/', component: { template: '<div>Home</div>' } },
        { path: '/login', component: { template: '<div>Login</div>' } },
        { path: '/profile', component: ProfileView }
      ]
    })

    // Mock auth store
    mockAuthStore = {
      isLoggedIn: true,
      loginState: {
        doctorId: 123,
        isInitialized: false
      },
      waitForInitialization: vi.fn().mockResolvedValue(undefined)
    }
    vi.mocked(useAuthStore).mockReturnValue(mockAuthStore)

    // Mock API service
    mockDoctorApi = {
      getProfile: vi.fn(),
      updateProfile: vi.fn()
    }
    vi.mocked(DoctorApiService).mockImplementation(() => mockDoctorApi)

    vi.clearAllMocks()
  })

  describe('Initialization', () => {
    it('should wait for auth initialization before fetching profile', async () => {
      const mockProfile = {
        Profile: {
          DId: 123,
          Name: 'Test Doctor',
          Phone: '12345678901'
        }
      }
      
      mockDoctorApi.getProfile.mockResolvedValue(mockProfile)
      
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      // Should show loading initially
      expect(wrapper.find('.loading-container').exists()).toBe(true)
      expect(wrapper.find('.loading-text').text()).toBe('正在初始化...')

      // Wait for initialization
      await wrapper.vm.$nextTick()
      
      expect(mockAuthStore.waitForInitialization).toHaveBeenCalled()
    })

    it('should show error when not logged in', async () => {
      mockAuthStore.isLoggedIn = false
      mockAuthStore.loginState.doctorId = undefined

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      // Should show error toast
      expect(wrapper.vm.toast.message).toBe('请先登录')
      expect(wrapper.vm.toast.type).toBe('error')
    })

    it('should show error when doctorId is missing', async () => {
      mockAuthStore.isLoggedIn = true
      mockAuthStore.loginState.doctorId = undefined

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(wrapper.vm.toast.message).toBe('请先登录')
      expect(wrapper.vm.toast.type).toBe('error')
    })
  })

  describe('Profile Loading', () => {
    beforeEach(() => {
      mockAuthStore.isLoggedIn = true
      mockAuthStore.loginState.doctorId = 123
    })

    it('should load profile successfully', async () => {
      const mockProfile = {
        Profile: {
          DId: 123,
          Name: 'Test Doctor',
          Phone: '12345678901',
          Email: 'test@example.com'
        }
      }
      
      mockDoctorApi.getProfile.mockResolvedValue(mockProfile)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(mockDoctorApi.getProfile).toHaveBeenCalledWith({ doctor_id: 123 })
      expect(wrapper.vm.profileLoaded).toBe(true)
      expect(wrapper.vm.hasError).toBe(false)
    })

    it('should handle API errors gracefully', async () => {
      const error = new Error('Network error')
      mockDoctorApi.getProfile.mockRejectedValue(error)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(wrapper.vm.hasError).toBe(true)
      expect(wrapper.vm.toast.message).toBe('获取个人信息失败，请稍后重试')
      expect(wrapper.vm.toast.type).toBe('error')
    })

    it('should handle 404 errors specifically', async () => {
      const error = new Error('404 - Doctor not found')
      mockDoctorApi.getProfile.mockRejectedValue(error)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(wrapper.vm.toast.message).toBe('医生信息不存在，请联系管理员')
    })

    it('should handle network errors specifically', async () => {
      const error = new Error('网络连接失败')
      mockDoctorApi.getProfile.mockRejectedValue(error)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(wrapper.vm.toast.message).toBe('网络连接失败，请检查网络后重试')
    })
  })

  describe('Retry Functionality', () => {
    it('should show retry button when there is an error', async () => {
      mockDoctorApi.getProfile.mockRejectedValue(new Error('Test error'))

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(wrapper.vm.showRetryButton).toBe(true)
    })

    it('should retry loading when retry button is clicked', async () => {
      mockDoctorApi.getProfile.mockRejectedValueOnce(new Error('First error'))
      mockDoctorApi.getProfile.mockResolvedValueOnce({
        Profile: { DId: 123, Name: 'Test Doctor' }
      })

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      // Should have error initially
      expect(wrapper.vm.hasError).toBe(true)

      // Retry
      await wrapper.vm.retryLoadProfile()
      await new Promise(resolve => setTimeout(resolve, 0))

      expect(mockDoctorApi.getProfile).toHaveBeenCalledTimes(2)
      expect(wrapper.vm.hasError).toBe(false)
      expect(wrapper.vm.profileLoaded).toBe(true)
    })
  })

  describe('Loading States', () => {
    it('should show auth loading initially', () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      expect(wrapper.vm.authLoading).toBe(true)
      expect(wrapper.find('.loading-container').exists()).toBe(true)
    })

    it('should show profile loading after auth is ready', async () => {
      mockDoctorApi.getProfile.mockImplementation(() => 
        new Promise(resolve => setTimeout(() => resolve({ Profile: {} }), 100))
      )

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      // Wait for auth loading to complete
      await wrapper.vm.$nextTick()
      wrapper.vm.authLoading = false
      wrapper.vm.loading = true
      await wrapper.vm.$nextTick()

      expect(wrapper.find('.loading-text').text()).toBe('正在加载个人信息...')
    })

    it('should show error state when loading fails', async () => {
      mockDoctorApi.getProfile.mockRejectedValue(new Error('Test error'))

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 0))

      wrapper.vm.authLoading = false
      await wrapper.vm.$nextTick()

      expect(wrapper.find('.error-container').exists()).toBe(true)
      expect(wrapper.find('.error-text').text()).toBe('加载失败')
    })
  })
})