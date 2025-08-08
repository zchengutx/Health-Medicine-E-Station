import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import LoginView from '@/views/LoginView.vue'
import { useAuthStore } from '@/stores/auth'
import { doctorApi } from '@/api/doctor'
import { showToast } from 'vant'

// Mock dependencies
vi.mock('@/api/doctor')
vi.mock('vant', () => ({
  showToast: vi.fn()
}))

vi.mock('@/router', () => ({
  useNavigationUtils: () => ({
    safePush: vi.fn(),
    forceToHome: vi.fn(),
    toLogin: vi.fn(),
    toRegister: vi.fn()
  })
}))

vi.mock('vue-router', () => ({
  useRouter: () => ({
    push: vi.fn(),
    replace: vi.fn()
  })
}))

describe('LoginView', () => {
  let wrapper: any
  let authStore: any

  beforeEach(() => {
    setActivePinia(createPinia())
    authStore = useAuthStore()
    
    wrapper = mount(LoginView, {
      global: {
        stubs: {
          'van-icon': true
        }
      }
    })
  })

  afterEach(() => {
    vi.clearAllMocks()
  })

  describe('表单验证', () => {
    it('应该验证手机号格式', async () => {
      const phoneInput = wrapper.find('input[type="tel"]')
      
      // 测试无效手机号
      await phoneInput.setValue('123')
      expect(wrapper.vm.canLogin).toBe(false)
      
      // 测试有效手机号
      await phoneInput.setValue('13800138000')
      await wrapper.find('input[type="number"]').setValue('1234')
      expect(wrapper.vm.canLogin).toBe(true)
    })

    it('应该验证验证码长度', async () => {
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      
      await phoneInput.setValue('13800138000')
      
      // 测试短验证码
      await codeInput.setValue('12')
      expect(wrapper.vm.canLogin).toBe(false)
      
      // 测试有效验证码
      await codeInput.setValue('1234')
      expect(wrapper.vm.canLogin).toBe(true)
    })

    it('应该在加载时禁用登录按钮', async () => {
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      
      wrapper.vm.isLoading = true
      await wrapper.vm.$nextTick()
      
      expect(loginButton.attributes('disabled')).toBeDefined()
    })
  })

  describe('验证码发送', () => {
    it('应该成功发送验证码', async () => {
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      mockSendSms.mockResolvedValue({ Message: 'success', Code: 200 })
      
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeButton = wrapper.find('.code-button')
      
      await phoneInput.setValue('13800138000')
      await codeButton.trigger('click')
      
      expect(mockSendSms).toHaveBeenCalledWith('13800138000')
      expect(showToast).toHaveBeenCalledWith({
        message: '验证码已发送',
        type: 'success'
      })
    })

    it('应该处理验证码发送失败', async () => {
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      mockSendSms.mockRejectedValue(new Error('发送失败'))
      
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeButton = wrapper.find('.code-button')
      
      await phoneInput.setValue('13800138000')
      await codeButton.trigger('click')
      
      expect(mockSendSms).toHaveBeenCalledWith('13800138000')
      expect(showToast).toHaveBeenCalledWith({
        message: '发送失败',
        type: 'fail'
      })
    })

    it('应该在发送后开始倒计时', async () => {
      const mockSendSms = vi.mocked(doctorApi.sendLoginSms)
      mockSendSms.mockResolvedValue({ Message: 'success', Code: 200 })
      
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeButton = wrapper.find('.code-button')
      
      await phoneInput.setValue('13800138000')
      await codeButton.trigger('click')
      
      await wrapper.vm.$nextTick()
      
      expect(wrapper.vm.codeCountdown).toBe(60)
      expect(wrapper.vm.codeButtonDisabled).toBe(true)
    })
  })

  describe('登录流程', () => {
    it('应该成功登录并跳转', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockResolvedValue({ 
        Message: 'success', 
        Code: 200, 
        DId: 123 
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
        message: '登录成功，正在跳转...',
        type: 'success',
        duration: 1500
      })
      
      // 验证store状态更新
      expect(authStore.isLoggedIn).toBe(true)
      expect(authStore.doctorInfo?.DId).toBe(123)
    })

    it('应该处理登录失败', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockRejectedValue(new Error('验证码错误'))
      
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      await loginButton.trigger('click')
      
      expect(showToast).toHaveBeenCalledWith({
        message: '验证码错误，请重新输入',
        type: 'fail',
        duration: 3000
      })
      
      // 验证验证码输入框被清空
      expect(wrapper.vm.formData.smsCode).toBe('')
    })

    it('应该在表单无效时显示错误提示', async () => {
      const loginButton = wrapper.find('.login-button')
      
      // 不填写任何信息直接点击登录
      await loginButton.trigger('click')
      
      expect(showToast).toHaveBeenCalledWith({
        message: '请输入正确的手机号码',
        type: 'fail',
        duration: 2000
      })
    })

    it('应该处理网络错误', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockRejectedValue(new Error('网络连接失败'))
      
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

  describe('用户体验', () => {
    it('应该在登录过程中显示加载状态', async () => {
      const mockLogin = vi.mocked(doctorApi.login)
      mockLogin.mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
      
      const phoneInput = wrapper.find('input[type="tel"]')
      const codeInput = wrapper.find('input[type="number"]')
      const loginButton = wrapper.find('.login-button')
      
      await phoneInput.setValue('13800138000')
      await codeInput.setValue('1234')
      
      const loginPromise = loginButton.trigger('click')
      
      await wrapper.vm.$nextTick()
      expect(wrapper.vm.isLoading).toBe(true)
      expect(loginButton.text()).toBe('登录中...')
      
      await loginPromise
      expect(wrapper.vm.isLoading).toBe(false)
    })

    it('应该在已登录时自动跳转', () => {
      // 设置已登录状态
      authStore.login('test_token', {
        DId: 123,
        Name: '测试医生',
        Phone: '13800138000',
        Email: '',
        Avatar: '',
        Status: 'active'
      })
      
      // 重新挂载组件
      wrapper = mount(LoginView, {
        global: {
          stubs: {
            'van-icon': true
          }
        }
      })
      
      // 验证跳转逻辑被调用
      // 这里需要根据实际的跳转实现来验证
    })
  })

  describe('输入处理', () => {
    it('应该限制手机号只能输入数字', async () => {
      const phoneInput = wrapper.find('input[type="tel"]')
      
      await phoneInput.setValue('abc123def456')
      await phoneInput.trigger('input')
      
      expect(wrapper.vm.formData.phone).toBe('123456')
    })

    it('应该在手机号变化时重置验证码状态', async () => {
      const phoneInput = wrapper.find('input[type="tel"]')
      
      // 设置初始状态
      wrapper.vm.smsCodeSent = true
      wrapper.vm.smsStatusMessage = '验证码已发送'
      wrapper.vm.formData.smsCode = '1234'
      
      await phoneInput.setValue('13800138001')
      await phoneInput.trigger('input')
      
      expect(wrapper.vm.smsCodeSent).toBe(false)
      expect(wrapper.vm.smsStatusMessage).toBe('')
      expect(wrapper.vm.formData.smsCode).toBe('')
    })

    it('应该在输入验证码时清除提示消息', async () => {
      const codeInput = wrapper.find('input[type="number"]')
      
      wrapper.vm.smsStatusMessage = '请输入验证码'
      
      await codeInput.setValue('1')
      await codeInput.trigger('input')
      
      expect(wrapper.vm.smsStatusMessage).toBe('')
    })
  })
})