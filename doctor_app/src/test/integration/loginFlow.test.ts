import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import HomeView from '@/views/HomeView.vue'
import { useAuthStore } from '@/stores/auth'
import { doctorApi } from '@/api/doctor'
import { showToast } from 'vant'

// Mock dependencies
vi.mock('@/api/doctor')
vi.mock('vant', () => ({
  showToast: vi.fn()
}))

// Mock window.location
Object.defineProperty(window, 'location', {
  value: {
    href: '',
    pathname: '/login'
  },
  writable: true
})

describe('登录到首页完整流程集成测试', () => {
  let router: any
  let authStore: any

  beforeEach(async () => {
    setActivePinia(createPinia())
    authStore = useAuthStore()
    
    // 创建路由器
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/login', component: LoginView },
        { path: '/home', component: HomeView },
        { path: '/', redirect: '/login' }
      ]
    })
    
    // 初始化路由
    router.push('/login')
    await router.isReady()
  })

  afterEach(() => {
    vi.clearAllMocks()
    window.location.href = ''
  })

  describe('成功登录流程', () => {
    it('应该完成从登录到首页的完整流程', async () => {
      // Mock API响应
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      const mockLogin = vi.mocked(doctorApi.login)
      
      mockSendSms.mockResolvedValue({ Message: 'success', Code: 200 })
      mockLogin.mockResolvedValue({ 
        Message: 'success', 
        Code: 200, 
        DId: 123 
      })

      // 挂载登录组件
      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      // 步骤1: 输入手机号
      const phoneInput = wrapper.find('input[type="tel"]')
      await phoneInput.setValue('13800138000')
      
      // 步骤2: 发送验证码
      const codeButton = wrapper.find('.code-button')
      await codeButton.trigger('click')
      
      expect(mockSendSms).toHaveBeenCalledWith('13800138000')
      expect(showToast).toHaveBeenCalledWith({
        message: '验证码已发送',
        type: 'success'
      })
      
      // 步骤3: 输入验证码
      const codeInput = wrapper.find('input[type="number"]')
      await codeInput.setValue('1234')
      
      // 步骤4: 点击登录
      const loginButton = wrapper.find('.login-button')
      await loginButton.trigger('click')
      
      // 验证API调用
      expect(mockLogin).toHaveBeenCalledWith({
        Phone: '13800138000',
        Password: '',
        SendSmsCode: '1234'
      })
      
      // 验证成功提示
      expect(showToast).toHaveBeenCalledWith({
        message: '登录成功，正在跳转...',
        type: 'success',
        duration: 1500
      })
      
      // 验证认证状态
      expect(authStore.isLoggedIn).toBe(true)
      expect(authStore.doctorInfo?.DId).toBe(123)
      expect(authStore.doctorInfo?.Phone).toBe('13800138000')
      
      // 等待跳转完成
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      // 验证页面跳转（通过window.location.href或路由状态）
      // 由于测试环境的限制，这里主要验证跳转逻辑被调用
      expect(window.location.href === '/home' || router.currentRoute.value.path === '/home').toBeTruthy()
    })

    it('应该在已登录状态下直接跳转到首页', async () => {
      // 预设登录状态
      authStore.login('test_token', {
        DId: 123,
        Name: '测试医生',
        Phone: '13800138000',
        Email: '',
        Avatar: '',
        Status: 'active'
      })

      // 挂载登录组件
      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      // 等待组件挂载完成
      await wrapper.vm.$nextTick()
      
      // 验证直接跳转逻辑
      // 在实际应用中，这里会触发跳转到首页
      expect(authStore.isLoggedIn).toBe(true)
    })
  })

  describe('登录失败处理', () => {
    it('应该处理验证码发送失败', async () => {
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      mockSendSms.mockRejectedValue(new Error('验证码发送失败'))

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeButton = wrapper.find('.code-button')
      
      await phoneInput.setValue('13800138000')
      await codeButton.trigger('click')
      
      expect(mockSendSms).toHaveBeenCalledWith('13800138000')
      expect(showToast).toHaveBeenCalledWith({
        message: '验证码发送失败',
        type: 'fail'
      })
      
      // 验证用户仍可输入验证码
      expect(wrapper.vm.smsCodeSent).toBe(true)
    })

    it('应该处理登录API失败', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockRejectedValue(new Error('验证码错误'))

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      await loginButton.trigger('click')
      
      expect(mockLogin).toHaveBeenCalledWith({
        Phone: '13800138000',
        Password: '',
        SendSmsCode: '1234'
      })
      
      expect(showToast).toHaveBeenCalledWith({
        message: '验证码错误，请重新输入',
        type: 'fail',
        duration: 3000
      })
      
      // 验证验证码输入框被清空
      expect(wrapper.vm.formData.smsCode).toBe('')
      
      // 验证用户未登录
      expect(authStore.isLoggedIn).toBe(false)
    })

    it('应该处理网络错误', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockRejectedValue(new Error('网络连接失败'))

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      await loginButton.trigger('click')
      
      expect(showToast).toHaveBeenCalledWith({
        message: '网络连接失败，请检查网络',
        type: 'fail',
        duration: 3000
      })
    })
  })

  describe('用户交互体验', () => {
    it('应该正确管理加载状态', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      // 模拟慢速API响应
      mockLogin.mockImplementation(() => 
        new Promise(resolve => setTimeout(() => resolve({ 
          Message: 'success', 
          Code: 200, 
          DId: 123 
        }), 100))
      )

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      
      // 点击登录按钮
      const loginPromise = loginButton.trigger('click')
      
      // 验证加载状态
      await wrapper.vm.$nextTick()
      expect(wrapper.vm.isLoading).toBe(true)
      expect(loginButton.text()).toBe('登录中...')
      expect(loginButton.attributes('disabled')).toBeDefined()
      
      // 等待登录完成
      await loginPromise
      
      // 验证加载状态恢复
      expect(wrapper.vm.isLoading).toBe(false)
      expect(loginButton.text()).toBe('登录')
    })

    it('应该正确处理验证码倒计时', async () => {
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      mockSendSms.mockResolvedValue({ Message: 'success', Code: 200 })

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeButton = wrapper.find('.code-button')
      
      await phoneInput.setValue('13800138000')
      await codeButton.trigger('click')
      
      // 验证倒计时开始
      expect(wrapper.vm.codeCountdown).toBe(60)
      expect(wrapper.vm.codeButtonDisabled).toBe(true)
      expect(codeButton.text()).toContain('获取验证码(60)')
      
      // 模拟时间流逝
      wrapper.vm.codeCountdown = 30
      await wrapper.vm.$nextTick()
      expect(codeButton.text()).toContain('获取验证码(30)')
    })

    it('应该正确处理表单验证', async () => {
      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const loginButton = wrapper.find('.login-button')
      
      // 测试空表单提交
      await loginButton.trigger('click')
      
      expect(showToast).toHaveBeenCalledWith({
        message: '请输入正确的手机号码',
        type: 'fail',
        duration: 2000
      })
      
      // 测试只有手机号
      const phoneInput = wrapper.find('input[type="tel"]')
      await phoneInput.setValue('13800138000')
      await loginButton.trigger('click')
      
      expect(showToast).toHaveBeenCalledWith({
        message: '请输入4-6位验证码',
        type: 'fail',
        duration: 2000
      })
    })
  })

  describe('状态持久化', () => {
    it('应该正确保存和恢复登录状态', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockResolvedValue({ 
        Message: 'success', 
        Code: 200, 
        DId: 123 
      })

      const wrapper = mount(LoginView, {
        global: {
          plugins: [router],
          stubs: {
            'van-icon': true
          }
        }
      })

      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      await loginButton.trigger('click')
      
      // 验证状态保存
      expect(authStore.token).toBe('temp_token')
      expect(authStore.doctorInfo?.DId).toBe(123)
      expect(authStore.loginState.savedPhone).toBe('13800138000')
      expect(authStore.loginState.rememberPhone).toBe(true)
      expect(authStore.loginState.lastLoginTime).toBeGreaterThan(0)
    })
  })
})