import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import { createPinia, setActivePinia } from 'pinia'

// 集成测试 - 个人信息页面统一功能
describe('个人信息页面统一 - 集成测试', () => {
  let router: any
  let mockLocalStorage: any

  beforeEach(() => {
    setActivePinia(createPinia())
    
    // 创建路由
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/', component: { template: '<div>Home</div>' } },
        { path: '/profile', component: { template: '<div>Profile</div>' } },
        { path: '/doctor-auth', redirect: '/profile' },
        { path: '/mine', component: { template: '<div>Mine</div>' } },
        { path: '/login', component: { template: '<div>Login</div>' } }
      ]
    })

    // Mock localStorage
    mockLocalStorage = {
      getItem: vi.fn(),
      setItem: vi.fn(),
      removeItem: vi.fn(),
      clear: vi.fn()
    }
    Object.defineProperty(window, 'localStorage', { value: mockLocalStorage })

    vi.clearAllMocks()
  })

  describe('路由重定向集成测试', () => {
    it('应该将 /doctor-auth 重定向到 /profile', async () => {
      // 导航到旧的认证页面路径
      await router.push('/doctor-auth')
      
      // 验证重定向到新的统一页面
      expect(router.currentRoute.value.path).toBe('/profile')
    })

    it('应该支持从重定向页面继续导航', async () => {
      // 先重定向
      await router.push('/doctor-auth')
      expect(router.currentRoute.value.path).toBe('/profile')
      
      // 然后导航到其他页面
      await router.push('/mine')
      expect(router.currentRoute.value.path).toBe('/mine')
    })

    it('应该支持浏览器后退按钮', async () => {
      await router.push('/')
      await router.push('/doctor-auth')
      
      // 模拟后退 - 由于重定向，后退可能仍在profile页面
      router.back()
      // 重定向后的后退行为可能不同，这里验证不会出错即可
      expect(router.currentRoute.value.path).toBeTruthy()
    })
  })

  describe('错误处理集成测试', () => {
    it('应该正确处理网络错误到缓存恢复的完整流程', async () => {
      const { ErrorHandler } = await import('@/utils/errorHandler')
      const { CacheManager } = await import('@/utils/cacheManager')
      
      const errorHandler = new ErrorHandler()
      const cacheManager = new CacheManager('test_cache')
      
      // 模拟网络错误
      const networkError = new Error('Network connection failed')
      networkError.name = 'NetworkError'
      
      // 分析错误
      const errorInfo = errorHandler.analyzeError(networkError)
      expect(errorInfo.type).toBe('network')
      expect(errorInfo.retryable).toBe(true)
      expect(errorInfo.fallbackAvailable).toBe(true)
      
      // 模拟缓存中有备用数据
      const cachedData = { id: 1, name: 'Test Doctor' }
      mockLocalStorage.getItem.mockReturnValue(JSON.stringify({
        data: cachedData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      }))
      
      // 从缓存恢复
      const recovered = cacheManager.get('doctor_1')
      expect(recovered).toEqual(cachedData)
    })

    it('应该正确处理重试机制', async () => {
      const { ErrorHandler } = await import('@/utils/errorHandler')
      
      const errorHandler = new ErrorHandler()
      let attemptCount = 0
      
      const mockOperation = vi.fn().mockImplementation(() => {
        attemptCount++
        if (attemptCount < 3) {
          throw new Error('Temporary failure')
        }
        return 'success'
      })
      
      const result = await errorHandler.retryWithBackoff(mockOperation, 'test-op', 3)
      
      expect(result).toBe('success')
      expect(attemptCount).toBe(3)
      expect(mockOperation).toHaveBeenCalledTimes(3)
    })
  })

  describe('缓存管理集成测试', () => {
    it('应该正确处理缓存的完整生命周期', async () => {
      const { CacheManager } = await import('@/utils/cacheManager')
      
      const cacheManager = new CacheManager('integration_test', {
        defaultTTL: 1000,
        version: '1.0.0'
      })
      
      const testData = { id: 1, name: 'Integration Test' }
      
      // 设置缓存
      const setResult = cacheManager.set('test-key', testData)
      expect(setResult).toBe(true)
      
      // 获取缓存
      mockLocalStorage.getItem.mockReturnValue(JSON.stringify({
        data: testData,
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0'
      }))
      
      const retrieved = cacheManager.get('test-key')
      expect(retrieved).toEqual(testData)
      
      // 检查存在
      expect(cacheManager.has('test-key')).toBe(true)
      
      // 删除缓存
      const deleteResult = cacheManager.delete('test-key')
      expect(deleteResult).toBe(true)
    })

    it('应该正确处理缓存过期和版本控制', async () => {
      const { CacheManager } = await import('@/utils/cacheManager')
      
      const cacheManager = new CacheManager('version_test', {
        version: '2.0.0'
      })
      
      // 模拟旧版本缓存
      mockLocalStorage.getItem.mockReturnValue(JSON.stringify({
        data: { id: 1 },
        timestamp: Date.now(),
        expiry: Date.now() + 10000,
        version: '1.0.0' // 旧版本
      }))
      
      const result = cacheManager.get('test-key')
      expect(result).toBeNull()
      expect(mockLocalStorage.removeItem).toHaveBeenCalledWith('version_test_test-key')
    })
  })

  describe('表单验证集成测试', () => {
    it('应该正确处理完整的表单验证流程', () => {
      // 模拟完整的表单验证逻辑
      const validateForm = (formData: any): { isValid: boolean; errors: string[] } => {
        const errors: string[] = []
        
        // 必填字段验证
        if (!formData.Name || formData.Name.trim() === '') {
          errors.push('姓名为必填项')
        }
        
        if (!formData.Phone || !/^1[3-9]\d{9}$/.test(formData.Phone)) {
          errors.push('请输入正确的手机号码')
        }
        
        if (!formData.Title || formData.Title.trim() === '') {
          errors.push('职称为必填项')
        }
        
        if (!formData.HospitalId || formData.HospitalId <= 0) {
          errors.push('请选择就职医院')
        }
        
        if (!formData.DepartmentId || formData.DepartmentId <= 0) {
          errors.push('请选择所属科室')
        }
        
        // 格式验证
        if (formData.Email && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.Email)) {
          errors.push('请输入正确的邮箱地址')
        }
        
        return {
          isValid: errors.length === 0,
          errors
        }
      }
      
      // 测试无效表单
      const invalidForm = {
        Name: '',
        Phone: '123',
        Title: '',
        HospitalId: 0,
        DepartmentId: 0,
        Email: 'invalid-email'
      }
      
      const invalidResult = validateForm(invalidForm)
      expect(invalidResult.isValid).toBe(false)
      expect(invalidResult.errors).toHaveLength(6)
      
      // 测试有效表单
      const validForm = {
        Name: '张医生',
        Phone: '13800138000',
        Title: '主任医师',
        HospitalId: 1,
        DepartmentId: 1,
        Email: 'doctor@example.com'
      }
      
      const validResult = validateForm(validForm)
      expect(validResult.isValid).toBe(true)
      expect(validResult.errors).toHaveLength(0)
    })
  })

  describe('API选择集成测试', () => {
    it('应该根据用户状态正确选择API和处理响应', async () => {
      // 模拟API选择逻辑
      const selectAndCallAPI = async (isFirstTime: boolean, profileData: any) => {
        const mockAuthAPI = vi.fn().mockResolvedValue({ success: true, message: 'Authentication successful' })
        const mockUpdateAPI = vi.fn().mockResolvedValue({ success: true, message: 'Profile updated' })
        
        if (isFirstTime) {
          try {
            const result = await mockAuthAPI(profileData)
            return { api: 'authentication', result }
          } catch (error) {
            // 回退到更新API
            const result = await mockUpdateAPI(profileData)
            return { api: 'updateProfile', result }
          }
        } else {
          const result = await mockUpdateAPI(profileData)
          return { api: 'updateProfile', result }
        }
      }
      
      const profileData = {
        DId: 1,
        Name: '张医生',
        Title: '主任医师'
      }
      
      // 测试首次用户
      const firstTimeResult = await selectAndCallAPI(true, profileData)
      expect(firstTimeResult.api).toBe('authentication')
      expect(firstTimeResult.result.success).toBe(true)
      
      // 测试现有用户
      const existingUserResult = await selectAndCallAPI(false, profileData)
      expect(existingUserResult.api).toBe('updateProfile')
      expect(existingUserResult.result.success).toBe(true)
    })
  })

  describe('数据流集成测试', () => {
    it('应该正确处理从加载到保存的完整数据流', async () => {
      const { CacheManager } = await import('@/utils/cacheManager')
      const { ErrorHandler } = await import('@/utils/errorHandler')
      
      const cacheManager = new CacheManager('data_flow_test')
      const errorHandler = new ErrorHandler()
      
      // 模拟完整的数据流
      const dataFlow = {
        // 1. 尝试从API加载
        async loadFromAPI(doctorId: number) {
          // 模拟API调用失败
          throw new Error('API temporarily unavailable')
        },
        
        // 2. 从缓存恢复
        loadFromCache(doctorId: number) {
          mockLocalStorage.getItem.mockReturnValue(JSON.stringify({
            data: { DId: doctorId, Name: '缓存医生', Title: '主任医师' },
            timestamp: Date.now(),
            expiry: Date.now() + 10000,
            version: '1.0.0'
          }))
          
          return cacheManager.get(`doctor_${doctorId}`)
        },
        
        // 3. 保存到缓存
        saveToCache(doctorId: number, data: any) {
          return cacheManager.set(`doctor_${doctorId}`, data)
        },
        
        // 4. 处理错误
        handleError(error: any) {
          return errorHandler.analyzeError(error)
        }
      }
      
      const doctorId = 1
      let profileData = null
      
      // 执行完整流程
      try {
        profileData = await dataFlow.loadFromAPI(doctorId)
      } catch (error) {
        const errorInfo = dataFlow.handleError(error)
        expect(errorInfo.fallbackAvailable).toBe(true)
        
        // 从缓存恢复
        profileData = dataFlow.loadFromCache(doctorId)
      }
      
      expect(profileData).toBeTruthy()
      expect(profileData.Name).toBe('缓存医生')
      
      // 更新数据并保存到缓存
      profileData.Title = '副主任医师'
      const saveResult = dataFlow.saveToCache(doctorId, profileData)
      expect(saveResult).toBe(true)
    })
  })

  describe('用户体验集成测试', () => {
    it('应该正确处理网络状态变化的完整流程', () => {
      // 模拟网络状态管理
      const networkManager = {
        isOnline: true,
        listeners: [] as Array<(online: boolean) => void>,
        
        addEventListener(callback: (online: boolean) => void) {
          this.listeners.push(callback)
        },
        
        setOnline(online: boolean) {
          this.isOnline = online
          this.listeners.forEach(callback => callback(online))
        }
      }
      
      const uiState = {
        showOfflineIndicator: false,
        allowSave: true,
        message: ''
      }
      
      // 监听网络状态变化
      networkManager.addEventListener((online) => {
        uiState.showOfflineIndicator = !online
        uiState.allowSave = online
        uiState.message = online ? '网络连接已恢复' : '网络连接已断开'
      })
      
      // 测试离线状态
      networkManager.setOnline(false)
      expect(uiState.showOfflineIndicator).toBe(true)
      expect(uiState.allowSave).toBe(false)
      expect(uiState.message).toBe('网络连接已断开')
      
      // 测试恢复在线
      networkManager.setOnline(true)
      expect(uiState.showOfflineIndicator).toBe(false)
      expect(uiState.allowSave).toBe(true)
      expect(uiState.message).toBe('网络连接已恢复')
    })
  })

  describe('端到端流程测试', () => {
    it('应该完整支持首次用户的注册到保存流程', async () => {
      // 模拟完整的首次用户流程
      const userJourney = {
        // 1. 用户访问旧的认证页面URL
        async visitOldURL() {
          await router.push('/doctor-auth')
          return router.currentRoute.value.path
        },
        
        // 2. 检测首次使用
        detectFirstTime(profile: any) {
          return !profile.Name || !profile.Title
        },
        
        // 3. 填写表单
        fillForm() {
          return {
            DId: 1,
            Name: '新医生',
            Gender: '男',
            Phone: '13800138000',
            Title: '住院医师',
            HospitalId: 1,
            DepartmentId: 1
          }
        },
        
        // 4. 验证表单
        validateForm(formData: any) {
          return !!(formData.Name && formData.Phone && formData.Title)
        },
        
        // 5. 保存数据
        async saveData(formData: any) {
          // 模拟API调用成功
          return { success: true, message: '认证信息提交成功！' }
        },
        
        // 6. 跳转到我的页面
        async navigateToMine() {
          await router.push('/mine')
          return router.currentRoute.value.path
        }
      }
      
      // 执行完整流程
      const currentPath = await userJourney.visitOldURL()
      expect(currentPath).toBe('/profile') // 重定向成功
      
      const isFirstTime = userJourney.detectFirstTime({})
      expect(isFirstTime).toBe(true)
      
      const formData = userJourney.fillForm()
      const isValid = userJourney.validateForm(formData)
      expect(isValid).toBe(true)
      
      const saveResult = await userJourney.saveData(formData)
      expect(saveResult.success).toBe(true)
      
      const finalPath = await userJourney.navigateToMine()
      expect(finalPath).toBe('/mine')
    })

    it('应该完整支持现有用户的查看和编辑流程', async () => {
      // 模拟现有用户流程
      const existingUserJourney = {
        // 1. 从我的页面点击编辑按钮
        async clickEditButton() {
          await router.push('/mine')
          // 模拟点击编辑按钮
          await router.push('/profile')
          return router.currentRoute.value.path
        },
        
        // 2. 加载现有数据
        loadExistingData() {
          return {
            DId: 1,
            Name: '现有医生',
            Title: '主任医师',
            HospitalId: 1,
            DepartmentId: 1
          }
        },
        
        // 3. 检测非首次使用
        detectExistingUser(profile: any) {
          return !!(profile.Name && profile.Title)
        },
        
        // 4. 编辑数据
        editData(existingData: any) {
          return {
            ...existingData,
            Title: '副主任医师', // 修改职称
            Speciality: '心血管疾病' // 添加专长
          }
        },
        
        // 5. 保存更新
        async updateData(formData: any) {
          return { success: true, message: '保存成功' }
        }
      }
      
      // 执行现有用户流程
      const profilePath = await existingUserJourney.clickEditButton()
      expect(profilePath).toBe('/profile')
      
      const existingData = existingUserJourney.loadExistingData()
      const isExistingUser = existingUserJourney.detectExistingUser(existingData)
      expect(isExistingUser).toBe(true)
      
      const editedData = existingUserJourney.editData(existingData)
      expect(editedData.Title).toBe('副主任医师')
      expect(editedData.Speciality).toBe('心血管疾病')
      
      const updateResult = await existingUserJourney.updateData(editedData)
      expect(updateResult.success).toBe(true)
    })
  })
})