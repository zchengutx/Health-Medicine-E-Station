import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import { createPinia, setActivePinia } from 'pinia'
import ProfileView from '@/views/ProfileView.vue'

// Mock dependencies
vi.mock('@/stores/auth', () => ({
  useAuthStore: () => ({
    waitForInitialization: vi.fn().mockResolvedValue(undefined),
    isLoggedIn: true,
    loginState: {
      doctorId: 1,
      phone: '13800138000'
    },
    token: 'mock-token'
  })
}))

vi.mock('@/api/doctor', () => ({
  default: class MockDoctorApiService {
    getProfile = vi.fn()
    updateProfile = vi.fn()
    authentication = vi.fn()
  }
}))

vi.mock('@/utils/logger', () => ({
  log: {
    debug: vi.fn(),
    info: vi.fn(),
    warn: vi.fn(),
    error: vi.fn()
  }
}))

vi.mock('@/utils/profileDiagnostic', () => ({
  runProfileDiagnostic: vi.fn().mockResolvedValue('diagnostic report')
}))

vi.mock('@/utils/apiTest', () => ({
  testProfileApi: vi.fn().mockResolvedValue('api test report')
}))

vi.mock('@/utils/errorHandler', () => ({
  errorHandler: {
    analyzeError: vi.fn().mockReturnValue({
      type: 'unknown',
      severity: 'medium',
      message: 'Test error',
      userMessage: '操作失败，请重试',
      retryable: true,
      fallbackAvailable: true
    })
  },
  getUserMessage: vi.fn().mockReturnValue('操作失败，请重试'),
  retryWithBackoff: vi.fn()
}))

vi.mock('@/utils/cacheManager', () => ({
  profileCache: {
    set: vi.fn(),
    get: vi.fn(),
    delete: vi.fn()
  },
  cacheProfile: vi.fn(),
  getCachedProfile: vi.fn(),
  clearProfileCache: vi.fn()
}))

// Mock components
vi.mock('@/components/FeedbackButton.vue', () => ({
  default: {
    name: 'FeedbackButton',
    template: '<button @click="$emit(\'click\')" :disabled="loading"><slot>{{ text }}</slot></button>',
    props: ['text', 'type', 'block', 'loading'],
    emits: ['click']
  }
}))

vi.mock('@/components/ToastMessage.vue', () => ({
  default: {
    name: 'ToastMessage',
    template: '<div class="toast" v-if="visible">{{ message }}</div>',
    props: ['message', 'type'],
    emits: ['close']
  }
}))

vi.mock('@/components/LoadingSpinner.vue', () => ({
  default: {
    name: 'LoadingSpinner',
    template: '<div class="loading-spinner">Loading...</div>'
  }
}))

describe('ProfileView - 增强功能测试', () => {
  let router: any
  let mockDoctorApi: any

  beforeEach(() => {
    setActivePinia(createPinia())
    
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/profile', component: { template: '<div>Profile</div>' } },
        { path: '/mine', component: { template: '<div>Mine</div>' } },
        { path: '/login', component: { template: '<div>Login</div>' } }
      ]
    })

    // Reset mocks
    vi.clearAllMocks()
    
    // Mock navigator.onLine
    Object.defineProperty(navigator, 'onLine', {
      writable: true,
      value: true
    })
  })

  describe('首次使用检测', () => {
    it('应该检测到首次使用用户', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      const mockProfile = {
        DId: 1,
        Name: '', // 空名称表示首次使用
        Title: '',
        HospitalId: 0,
        DepartmentId: 0
      }

      retryWithBackoff.mockResolvedValue({
        Profile: mockProfile
      })

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))

      expect(wrapper.vm.isFirstTime).toBe(true)
      expect(wrapper.find('.title').text()).toContain('请如实填写您的个人信息')
    })

    it('应该检测到现有用户', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      const mockProfile = {
        DId: 1,
        Name: '张医生',
        Title: '主任医师',
        HospitalId: 1,
        DepartmentId: 1
      }

      retryWithBackoff.mockResolvedValue({
        Profile: mockProfile
      })

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))

      expect(wrapper.vm.isFirstTime).toBe(false)
      expect(wrapper.find('.title').text()).toBe('个人信息')
    })
  })

  describe('表单验证', () => {
    it('应该验证必填字段', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      // 设置为首次使用模式
      wrapper.vm.isFirstTime = true
      wrapper.vm.profileLoaded = true
      await wrapper.vm.$nextTick()

      // 尝试保存空表单
      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      expect(wrapper.vm.toast.visible).toBe(true)
      expect(wrapper.vm.toast.type).toBe('error')
      expect(wrapper.vm.toast.message).toContain('必填项')
    })

    it('应该验证邮箱格式', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isFirstTime = true
      wrapper.vm.profileLoaded = true
      
      // 设置必填字段
      wrapper.vm.form.Name = '张医生'
      wrapper.vm.form.Gender = '男'
      wrapper.vm.form.Phone = '13800138000'
      wrapper.vm.form.Title = '主任医师'
      wrapper.vm.form.HospitalId = 1
      wrapper.vm.form.DepartmentId = 1
      
      // 设置无效邮箱
      wrapper.vm.form.Email = 'invalid-email'
      
      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      expect(wrapper.vm.toast.visible).toBe(true)
      expect(wrapper.vm.toast.message).toContain('邮箱格式')
    })

    it('应该验证手机号格式', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isFirstTime = true
      wrapper.vm.profileLoaded = true
      
      // 设置必填字段
      wrapper.vm.form.Name = '张医生'
      wrapper.vm.form.Gender = '男'
      wrapper.vm.form.Phone = '123' // 无效手机号
      wrapper.vm.form.Title = '主任医师'
      wrapper.vm.form.HospitalId = 1
      wrapper.vm.form.DepartmentId = 1
      
      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      expect(wrapper.vm.toast.visible).toBe(true)
      expect(wrapper.vm.toast.message).toContain('手机号格式')
    })
  })

  describe('实时字段验证', () => {
    it('应该在字段失焦时进行验证', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.profileLoaded = true
      await wrapper.vm.$nextTick()

      // 模拟字段验证
      const error = wrapper.vm.validateField('Email', 'invalid-email')
      expect(error).toContain('邮箱格式')

      const validResult = wrapper.vm.validateField('Email', 'test@example.com')
      expect(validResult).toBeNull()
    })

    it('应该显示字段错误状态', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.profileLoaded = true
      wrapper.vm.fieldErrors.Email = '邮箱格式不正确'
      await wrapper.vm.$nextTick()

      const errorDiv = wrapper.find('.field-error')
      expect(errorDiv.exists()).toBe(true)
      expect(errorDiv.text()).toBe('邮箱格式不正确')
    })
  })

  describe('网络状态处理', () => {
    it('应该显示离线状态', async () => {
      Object.defineProperty(navigator, 'onLine', {
        writable: true,
        value: false
      })

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isOnline = false
      await wrapper.vm.$nextTick()

      const offlineIndicator = wrapper.find('.network-status.offline')
      expect(offlineIndicator.exists()).toBe(true)
      expect(offlineIndicator.text()).toContain('离线模式')
    })

    it('应该在离线时阻止保存操作', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isOnline = false
      wrapper.vm.profileLoaded = true
      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      expect(wrapper.vm.toast.visible).toBe(true)
      expect(wrapper.vm.toast.message).toContain('网络连接不可用')
    })
  })

  describe('保存进度显示', () => {
    it('应该在保存时显示进度条', async () => {
      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.loading = true
      wrapper.vm.saveProgress = 50
      await wrapper.vm.$nextTick()

      const progressBar = wrapper.find('.save-progress')
      expect(progressBar.exists()).toBe(true)
      
      const progressFill = wrapper.find('.progress-fill')
      expect(progressFill.attributes('style')).toContain('width: 50%')
      
      const progressText = wrapper.find('.progress-text')
      expect(progressText.text()).toContain('50%')
    })
  })

  describe('缓存功能', () => {
    it('应该在API失败时尝试从缓存加载', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      const { getCachedProfile } = await import('@/utils/cacheManager')
      
      const cachedProfile = {
        DId: 1,
        Name: '张医生',
        Title: '主任医师'
      }

      retryWithBackoff.mockRejectedValue(new Error('Network error'))
      getCachedProfile.mockReturnValue(cachedProfile)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))

      expect(getCachedProfile).toHaveBeenCalledWith(1)
      expect(wrapper.vm.form.Name).toBe('张医生')
      expect(wrapper.vm.toast.message).toContain('缓存')
    })

    it('应该在保存成功后更新缓存', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      const { cacheProfile } = await import('@/utils/cacheManager')

      retryWithBackoff.mockResolvedValue(undefined)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isFirstTime = false
      wrapper.vm.profileLoaded = true
      wrapper.vm.form = {
        DId: 1,
        Name: '张医生',
        Gender: '男',
        Phone: '13800138000',
        Title: '主任医师',
        HospitalId: 1,
        DepartmentId: 1
      }

      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      await new Promise(resolve => setTimeout(resolve, 100))

      expect(cacheProfile).toHaveBeenCalledWith(1, expect.objectContaining({
        Name: '张医生'
      }))
    })
  })

  describe('智能API调用', () => {
    it('应该为首次用户优先使用authentication API', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      
      retryWithBackoff.mockResolvedValue(undefined)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isFirstTime = true
      wrapper.vm.profileLoaded = true
      wrapper.vm.form = {
        DId: 1,
        Name: '张医生',
        Gender: '男',
        Phone: '13800138000',
        Title: '主任医师',
        HospitalId: 1,
        DepartmentId: 1
      }

      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      await new Promise(resolve => setTimeout(resolve, 100))

      // 应该调用authentication API的重试版本
      expect(retryWithBackoff).toHaveBeenCalledWith(
        expect.any(Function),
        expect.stringContaining('auth'),
        1
      )
    })

    it('应该为现有用户使用updateProfile API', async () => {
      const { retryWithBackoff } = await import('@/utils/errorHandler')
      
      retryWithBackoff.mockResolvedValue(undefined)

      const wrapper = mount(ProfileView, {
        global: {
          plugins: [router]
        }
      })

      wrapper.vm.isFirstTime = false
      wrapper.vm.profileLoaded = true
      wrapper.vm.form = {
        DId: 1,
        Name: '张医生',
        Gender: '男',
        Phone: '13800138000',
        Title: '主任医师',
        HospitalId: 1,
        DepartmentId: 1
      }

      await wrapper.vm.$nextTick()

      const saveButton = wrapper.find('button')
      await saveButton.trigger('click')

      await new Promise(resolve => setTimeout(resolve, 100))

      // 应该直接调用updateProfile API的重试版本
      expect(retryWithBackoff).toHaveBeenCalledWith(
        expect.any(Function),
        expect.stringMatching(/^saveProfile_1_\d+$/),
        2
      )
    })
  })
})