import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { TokenStorage, UserInfoStorage, Storage, STORAGE_KEYS } from '@/utils/storage'
import { doctorApi } from '@/api/doctor'
import { log } from '@/utils/logger'

export interface DoctorInfo {
  DId: number
  DoctorCode?: string
  Name: string
  Gender?: string
  BirthDate?: string
  Phone: string
  Email: string
  Avatar: string
  LicenseNumber?: string
  DepartmentId?: number
  HospitalId?: number
  Title?: string
  Speciality?: string
  PracticeScope?: string
  Status: string
  CreatedAt?: string
  UpdatedAt?: string
}

export interface LoginState {
  isLoading: boolean
  isInitialized: boolean
  lastLoginTime: number | null
  rememberPhone: boolean
  savedPhone: string
  doctorId?: number
}

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref<string | null>(TokenStorage.getToken())
  const doctorInfo = ref<DoctorInfo | null>(UserInfoStorage.getUserInfo<DoctorInfo>())
  const loginState = ref<LoginState>({
    isLoading: false,
    isInitialized: false,
    lastLoginTime: Storage.get<number>(STORAGE_KEYS.LAST_LOGIN_TIME),
    rememberPhone: Storage.get<boolean>(STORAGE_KEYS.REMEMBER_PHONE) || false,
    savedPhone: Storage.get<string>('saved_phone') || '',
    doctorId: doctorInfo.value?.DId
  })

  // 计算属性
  const isLoggedIn = computed(() => !!token.value)
  const isAuthenticated = computed(() => isLoggedIn.value && !!doctorInfo.value)
  const doctorName = computed(() => doctorInfo.value?.Name || '医生')
  const doctorAvatar = computed(() => doctorInfo.value?.Avatar || '')
  const doctorStatus = computed(() => doctorInfo.value?.Status || 'inactive')

  // 方法
  const setToken = (newToken: string) => {
    token.value = newToken
    TokenStorage.setToken(newToken)
  }

  const setDoctorInfo = (info: DoctorInfo) => {
    doctorInfo.value = info
    UserInfoStorage.setUserInfo(info)
  }

  // 为了兼容性，添加setUserInfo别名
  const setUserInfo = setDoctorInfo

  const setLoginState = (state: Partial<LoginState>) => {
    loginState.value = { ...loginState.value, ...state }
    
    // 持久化某些状态
    if (state.lastLoginTime !== undefined) {
      Storage.set(STORAGE_KEYS.LAST_LOGIN_TIME, state.lastLoginTime)
    }
    if (state.rememberPhone !== undefined) {
      Storage.set(STORAGE_KEYS.REMEMBER_PHONE, state.rememberPhone)
    }
  }

  const savePhone = (phone: string, remember: boolean = false) => {
    if (remember) {
      Storage.set('saved_phone', phone)
      loginState.value.savedPhone = phone
      loginState.value.rememberPhone = true
      Storage.set(STORAGE_KEYS.REMEMBER_PHONE, true)
    } else {
      Storage.remove('saved_phone')
      loginState.value.savedPhone = ''
      loginState.value.rememberPhone = false
      Storage.set(STORAGE_KEYS.REMEMBER_PHONE, false)
    }
  }

  const login = (newToken: string, info: DoctorInfo, phone?: string, rememberPhone: boolean = false) => {
    try {
      log.debug('开始保存登录状态到store')
      // 验证必要的参数
      if (!newToken || !info || !info.DId) {
        throw new Error('登录参数不完整')
      }
      setToken(newToken)
      setDoctorInfo(info)
      setLoginState({
        lastLoginTime: Date.now(),
        isLoading: false,
        doctorId: info.DId
      })
      if (phone) {
        savePhone(phone, rememberPhone)
      }
      log.info('登录状态保存成功')
    } catch (error) {
      log.error('保存登录状态失败', error)
      throw error
    }
  }

  const logout = () => {
    token.value = null
    doctorInfo.value = null
    TokenStorage.removeToken()
    UserInfoStorage.removeUserInfo()
    
    // 保留记住手机号的设置，但重置初始化状态
    setLoginState({
      isLoading: false,
      isInitialized: false,
      lastLoginTime: null,
      doctorId: undefined
    })
  }

  const updateDoctorInfo = (updates: Partial<DoctorInfo>) => {
    if (doctorInfo.value) {
      doctorInfo.value = { ...doctorInfo.value, ...updates }
      UserInfoStorage.setUserInfo(doctorInfo.value)
    }
  }

  const initAuth = () => {
    log.debug('开始初始化认证状态')
    
    try {
      // 清理过期数据
      Storage.cleanExpired()
      
      const savedToken = TokenStorage.getToken()
      const savedInfo = UserInfoStorage.getUserInfo<DoctorInfo>()
      
      log.debug('从存储中获取数据', {
        hasToken: !!savedToken,
        hasUserInfo: !!savedInfo,
        savedInfoDId: savedInfo?.DId
      })
      
      if (savedToken) {
        token.value = savedToken
      }
      
      if (savedInfo) {
        doctorInfo.value = savedInfo
      }
      
      // 确保doctorId正确设置
      let doctorId: number | undefined
      if (savedInfo && savedInfo.DId) {
        doctorId = savedInfo.DId
      }
      
      // 初始化登录状态
      loginState.value = {
        isLoading: false,
        isInitialized: true,
        lastLoginTime: Storage.get<number>(STORAGE_KEYS.LAST_LOGIN_TIME),
        rememberPhone: Storage.get<boolean>(STORAGE_KEYS.REMEMBER_PHONE) || false,
        savedPhone: Storage.get<string>('saved_phone') || '',
        doctorId: doctorId
      }
      
      log.info('认证状态初始化完成', {
        isLoggedIn: !!token.value,
        hasUserInfo: !!doctorInfo.value,
        doctorId: doctorId,
        hasDoctorId: !!doctorId
      })
    } catch (error) {
      log.error('认证状态初始化失败', error)
      // 即使初始化失败，也要设置为已初始化状态，避免无限等待
      loginState.value.isInitialized = true
    }
  }

  const checkTokenExpiry = (): boolean => {
    log.debug('检查token有效性')
    
    if (!token.value || !loginState.value.lastLoginTime) {
      log.debug('token或lastLoginTime不存在')
      return false
    }
    
    // 检查token是否过期（假设7天过期）
    const expiryTime = 7 * 24 * 60 * 60 * 1000 // 7天
    const timeSinceLogin = Date.now() - loginState.value.lastLoginTime
    const isExpired = timeSinceLogin > expiryTime
    
    log.debug('token过期检查', {
      timeSinceLogin: Math.floor(timeSinceLogin / 1000 / 60), // 转换为分钟
      expiryTimeMinutes: Math.floor(expiryTime / 1000 / 60), // 转换为分钟
      isExpired
    })
    
    if (isExpired) {
      log.info('token已过期，执行登出操作')
      logout()
      return false
    }
    
    log.debug('token有效')
    return true
  }

  const refreshUserInfo = async (): Promise<boolean> => {
    if (!token.value || !doctorInfo.value?.DId) {
      return false
    }
    
    try {
      // 这里可以调用API刷新用户信息
      // const updatedInfo = await doctorApi.getProfile(doctorInfo.value.DId)
      // setDoctorInfo(updatedInfo)
      return true
    } catch (error) {
      log.error('刷新用户信息失败', error)
      return false
    }
  }

  const waitForInitialization = (): Promise<void> => {
    return new Promise((resolve) => {
      if (loginState.value.isInitialized) {
        resolve()
        return
      }
      
      // 使用轮询检查初始化状态，但设置最大等待时间
      let attempts = 0
      const maxAttempts = 100 // 最多等待5秒(100 * 50ms)
      
      const checkInitialized = () => {
        if (loginState.value.isInitialized || attempts >= maxAttempts) {
          // 如果超过最大尝试次数，强制设置为已初始化
          if (attempts >= maxAttempts && !loginState.value.isInitialized) {
            log.warn('等待初始化超时，强制设置为已初始化')
            loginState.value.isInitialized = true
          }
          resolve()
        } else {
          attempts++
          setTimeout(checkInitialized, 50) // 每50ms检查一次
        }
      }
      
      checkInitialized()
    })
  }

  const deleteAccount = async () => {
    if (!doctorInfo.value?.DId) throw new Error('未找到医生ID')
    await doctorApi.deleteAccount(doctorInfo.value.DId)
    logout()
  }
  
  return {
    // 状态
    token,
    doctorInfo,
    loginState,
    // 计算属性
    isLoggedIn,
    isAuthenticated,
    doctorName,
    doctorAvatar,
    doctorStatus,
    // 方法
    setToken,
    setDoctorInfo,
    setUserInfo,
    setLoginState,
    savePhone,
    login,
    logout,
    updateDoctorInfo,
    initAuth,
    checkTokenExpiry,
    refreshUserInfo,
    waitForInitialization,
    deleteAccount
  }
})