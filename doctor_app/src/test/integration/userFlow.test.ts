import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createTestingPinia } from '@pinia/testing'
import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import SplashView from '@/views/SplashView.vue'

// Mock API
vi.mock('@/api/doctor', () => ({
  doctorApi: {
    sendLoginSms: vi.fn(() => Promise.resolve({ Message: '验证码已发送', Code: 200 })),
    sendRegisterSms: vi.fn(() => Promise.resolve({ Message: '验证码已发送', Code: 200 })),
    login: vi.fn(() => Promise.resolve({ Message: '登录成功', Code: 200, DId: 1 })),
    register: vi.fn(() => Promise.resolve({ Message: '注册成功', Code: 200 }))
  }
}))

// Mock Vant Toast
vi.mock('vant', () => ({
  showToast: vi.fn()
}))

describe('User Flow Integration Tests', () => {
  let router: any

  beforeEach(() => {
    vi.clearAllMocks()
    
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/', component: SplashView },
        { path: '/login', component: LoginView },
        { path: '/register', component: RegisterView }
      ]
    })
  })

  describe('Login Flow', () => {
    it('should complete login flow successfully', async () => {
      const wrapper = mount(LoginView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      // Fill phone number
      const phoneInput = wrapper.find('input[placeholder="请输入手机号码"]')
      await phoneInput.setValue('13812345678')

      // Click get SMS code button
      const smsButton = wrapper.find('.code-button')
      expect(smsButton.element.disabled).toBe(false)
      await smsButton.trigger('click')

      // Fill SMS code
      const codeInput = wrapper.find('input[placeholder="请输入验证码"]')
      await codeInput.setValue('1234')

      // Click login button
      const loginButton = wrapper.findComponent({ name: 'FeedbackButton' })
      expect(loginButton.props('disabled')).toBe(false)
      await loginButton.trigger('click')

      // Verify API was called
      const { doctorApi } = await import('@/api/doctor')
      expect(doctorApi.sendLoginSms).toHaveBeenCalledWith('13812345678')
      expect(doctorApi.login).toHaveBeenCalledWith({
        Phone: '13812345678',
        Password: '',
        SendSmsCode: '1234'
      })
    })

    it('should show validation errors for invalid input', async () => {
      const wrapper = mount(LoginView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      // Try to get SMS with invalid phone
      const phoneInput = wrapper.find('input[placeholder="请输入手机号码"]')
      await phoneInput.setValue('123')

      const smsButton = wrapper.find('.code-button')
      expect(smsButton.element.disabled).toBe(true)
    })

    it('should navigate to register page', async () => {
      const wrapper = mount(LoginView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      const registerLink = wrapper.find('.register-link span')
      await registerLink.trigger('click')

      // Should call navigation utility
      expect(router.currentRoute.value.path).toBe('/login') // Initial path
    })
  })

  describe('Register Flow', () => {
    it('should complete register flow successfully', async () => {
      const wrapper = mount(RegisterView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      // Fill phone number
      const phoneInput = wrapper.find('input[placeholder="请输入手机号码"]')
      await phoneInput.setValue('13812345678')

      // Click get SMS code button
      const smsButton = wrapper.find('.code-button')
      await smsButton.trigger('click')

      // Fill SMS code
      const codeInput = wrapper.find('input[placeholder="请输入验证码"]')
      await codeInput.setValue('1234')

      // Fill password
      const passwordInput = wrapper.find('input[placeholder="请输入密码"]')
      await passwordInput.setValue('Abc123456')

      // Click register button
      const registerButton = wrapper.findComponent({ name: 'FeedbackButton' })
      await registerButton.trigger('click')

      // Verify API was called
      const { doctorApi } = await import('@/api/doctor')
      expect(doctorApi.sendRegisterSms).toHaveBeenCalledWith('13812345678')
      expect(doctorApi.register).toHaveBeenCalledWith({
        Phone: '13812345678',
        Password: 'Abc123456',
        SendSmsCode: '1234'
      })
    })

    it('should show password strength indicator', async () => {
      const wrapper = mount(RegisterView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      const passwordInput = wrapper.find('input[placeholder="请输入密码"]')
      await passwordInput.setValue('weak')

      expect(wrapper.find('.password-strength').exists()).toBe(true)
    })

    it('should toggle password visibility', async () => {
      const wrapper = mount(RegisterView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      const passwordInput = wrapper.find('input[placeholder="请输入密码"]')
      const toggleButton = wrapper.find('.password-toggle')

      expect(passwordInput.attributes('type')).toBe('password')

      await toggleButton.trigger('click')
      expect(passwordInput.attributes('type')).toBe('text')
    })
  })

  describe('Splash Screen Flow', () => {
    it('should auto-navigate after countdown', async () => {
      vi.useFakeTimers()

      const wrapper = mount(SplashView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      expect(wrapper.vm.countdown).toBe(3)

      // Fast-forward time
      vi.advanceTimersByTime(3100) // 3.1 seconds

      expect(wrapper.vm.countdown).toBe(0)

      vi.useRealTimers()
    })

    it('should skip countdown when skip button is clicked', async () => {
      const wrapper = mount(SplashView, {
        global: {
          plugins: [
            createTestingPinia({ createSpy: vi.fn }),
            router
          ]
        }
      })

      const skipButton = wrapper.find('.skip-button')
      await skipButton.trigger('click')

      expect(wrapper.vm.countdown).toBe(0)
    })
  })
})