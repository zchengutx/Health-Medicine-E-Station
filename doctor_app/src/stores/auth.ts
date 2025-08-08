import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { TokenStorage, UserInfoStorage, Storage, STORAGE_KEYS } from '@/utils/storage'
import { doctorApi } from '@/api/doctor'

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
  lastLoginTime: number | null
  rememberPhone: boolean
  savedPhone: string
}

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref<string | null>(TokenStorage.getToken())
  const doctorInfo = ref<DoctorInfo | null>(UserInfoStorage.getUserInfo<DoctorInfo>())
  const loginState = ref<LoginState>({
    isLoading: false,
    lastLoginTime: Storage.get<number>(STORAGE_KEYS.LAST_LOGIN_TIME),
    rememberPhone: Storage.get<boolean>(STORAGE_KEYS.REMEMBER_PHONE) || false,
    savedPhone: Storage.get<string>('saved_phone') || ''
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
      console.log('开始保存登录状态到store...')
      
      // 验证必要的参数
      if (!newToken || !info || !info.DId) {
        throw new Error('登录参数不完整')
      }
      
      setToken(newToken)
      setDoctorInfo(info)
      setLoginState({
        lastLoginTime: Date.now(),
        isLoading: false
      })
      
      if (phone) {
        savePhone(phone, rememberPhone)
      }
      
      console.log('登录状态保存成功')
    } catch (error) {
      console.error('保存登录状态失败:', error)
      throw error
    }
  }

  const logout = () => {
    token.value = null
    doctorInfo.value = null
    TokenStorage.removeToken()
    UserInfoStorage.removeUserInfo()
    
    // 保留记住手机号的设置
    setLoginState({
      isLoading: false,
      lastLoginTime: null
    })
  }

  const updateDoctorInfo = (updates: Partial<DoctorInfo>) => {
    if (doctorInfo.value) {
      doctorInfo.value = { ...doctorInfo.value, ...updates }
      UserInfoStorage.setUserInfo(doctorInfo.value)
    }
  }

  const initAuth = () => {
    console.log('初始化认证状态...')
    
    // 清理过期数据
    Storage.cleanExpired()
    
    const savedToken = TokenStorage.getToken()
    const savedInfo = UserInfoStorage.getUserInfo<DoctorInfo>()
    
    console.log('从存储中获取的token:', !!savedToken)
    console.log('从存储中获取的用户信息:', !!savedInfo)
    
    if (savedToken) {
      token.value = savedToken
    }
    
    if (savedInfo) {
      doctorInfo.value = savedInfo
    }
    
    // 初始化登录状态
    loginState.value = {
      isLoading: false,
      lastLoginTime: Storage.get<number>(STORAGE_KEYS.LAST_LOGIN_TIME),
      rememberPhone: Storage.get<boolean>(STORAGE_KEYS.REMEMBER_PHONE) || false,
      savedPhone: Storage.get<string>('saved_phone') || ''
    }
    
    console.log('认证状态初始化完成')
    console.log('当前token:', !!token.value)
    console.log('当前用户信息:', !!doctorInfo.value)
    console.log('lastLoginTime:', loginState.value.lastLoginTime)
  }

  const checkTokenExpiry = (): boolean => {
    console.log('检查token有效性...')
    console.log('token存在:', !!token.value)
    console.log('lastLoginTime存在:', !!loginState.value.lastLoginTime)
    
    if (!token.value || !loginState.value.lastLoginTime) {
      console.log('token或lastLoginTime不存在，返回false')
      return false
    }
    
    // 检查token是否过期（假设7天过期）
    const expiryTime = 7 * 24 * 60 * 60 * 1000 // 7天
    const timeSinceLogin = Date.now() - loginState.value.lastLoginTime
    const isExpired = timeSinceLogin > expiryTime
    
    console.log('距离上次登录时间:', timeSinceLogin, 'ms')
    console.log('过期时间:', expiryTime, 'ms')
    console.log('是否过期:', isExpired)
    
    if (isExpired) {
      console.log('token已过期，执行登出操作')
      logout()
      return false
    }
    
    console.log('token有效')
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
      console.error('Failed to refresh user info:', error)
      return false
    }
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
    deleteAccount
  }
})