<template>
  <div class="profile-view-scroll">
    <div class="profile-view">
      <div class="header-section">
        <h2 class="title">{{ isFirstTime ? '请如实填写您的个人信息' : '个人信息' }}</h2>
        <!-- 网络状态指示器 -->
        <div v-if="!isOnline" class="network-status offline">
          <span class="status-icon">📶</span>
          <span class="status-text">离线模式</span>
        </div>
      </div>
      
      <!-- 认证状态加载中 -->
      <div v-if="authLoading" class="loading-container">
        <LoadingSpinner />
        <p class="loading-text">正在初始化...</p>
      </div>
      
      <!-- 数据加载中 -->
      <div v-else-if="loading && !profileLoaded && !isFirstTime" class="loading-container">
        <LoadingSpinner />
        <p class="loading-text">正在加载个人信息...</p>
      </div>
      
      <!-- 错误状态 -->
    <div v-else-if="hasError && !profileLoaded && !isFirstTime" class="error-container">
      <div class="error-icon">⚠️</div>
      <p class="error-text">加载失败</p>
      <p class="error-desc">{{ errorMessage || '请检查网络连接或稍后重试' }}</p>
      <FeedbackButton
        class="retry-btn"
        type="secondary"
        block
        text="重新加载"
        @click="retryLoadProfile"
      />
      </div>
      
      <div v-else class="profile-form">
        <!-- 基本信息组 -->
        <div class="form-section">
          <h3 class="section-title">基本信息</h3>
          <div class="form-group" v-for="item in basicFormItems" :key="item.key">
            <label :for="item.key" class="form-label">
              <span v-if="item.required" class="required">*</span>{{ item.label }}
            </label>
            <template v-if="item.type === 'select'">
              <select 
                v-model="form[item.key as keyof typeof form]" 
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @change="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
              >
                <option v-for="option in item.options" :key="option.value" :value="option.value">{{ option.label }}</option>
              </select>
            </template>
            <template v-else-if="item.type === 'date'">
              <input
                v-model="form[item.key as keyof typeof form]"
                :id="item.key"
                :placeholder="item.placeholder"
                type="date"
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :readonly="item.readonly"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @change="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
              />
            </template>
            <template v-else-if="item.type === 'textarea'">
              <textarea
                v-model="form[item.key as keyof typeof form]"
                :id="item.key"
                :placeholder="item.placeholder"
                :class="['form-input', 'form-textarea', { 'form-input-error': fieldErrors[item.key] }]"
                :readonly="item.readonly"
                :required="item.required"
                rows="3"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @input="() => { if (fieldErrors[item.key]) handleFieldValidation(item.key, form[item.key as keyof typeof form]) }"
              ></textarea>
            </template>
            <template v-else>
              <input
                v-model="form[item.key as keyof typeof form]"
                :id="item.key"
                :placeholder="item.placeholder"
                :type="item.type || 'text'"
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :readonly="item.readonly"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @input="() => { if (fieldErrors[item.key]) handleFieldValidation(item.key, form[item.key as keyof typeof form]) }"
              />
            </template>
            <!-- 字段错误提示 -->
            <div v-if="fieldErrors[item.key]" class="field-error">{{ fieldErrors[item.key] }}</div>
          </div>
        </div>

        <!-- 职业信息组 -->
        <div class="form-section">
          <h3 class="section-title">职业信息</h3>
          <div class="form-group" v-for="item in professionalFormItems" :key="item.key">
            <label :for="item.key" class="form-label">
              <span v-if="item.required" class="required">*</span>{{ item.label }}
            </label>
            <template v-if="item.key === 'DepartmentId'">
              <select 
                v-model="form.DepartmentId" 
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form.DepartmentId)"
                @change="handleFieldValidation(item.key, form.DepartmentId)"
              >
                <option v-for="dept in departments" :key="dept.id" :value="dept.id">{{ dept.name }}</option>
              </select>
            </template>
            <template v-else-if="item.key === 'HospitalId'">
              <select 
                v-model="form.HospitalId" 
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form.HospitalId)"
                @change="handleFieldValidation(item.key, form.HospitalId)"
              >
                <option v-for="hos in hospitals" :key="hos.id" :value="hos.id">{{ hos.name }}</option>
              </select>
            </template>
            <template v-else-if="item.type === 'textarea'">
              <textarea
                v-model="form[item.key as keyof typeof form]"
                :id="item.key"
                :placeholder="item.placeholder"
                :class="['form-input', 'form-textarea', { 'form-input-error': fieldErrors[item.key] }]"
                :readonly="item.readonly"
                :required="item.required"
                rows="3"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @input="() => { if (fieldErrors[item.key]) handleFieldValidation(item.key, form[item.key as keyof typeof form]) }"
              ></textarea>
            </template>
            <template v-else>
              <input
                v-model="form[item.key as keyof typeof form]"
                :id="item.key"
                :placeholder="item.placeholder"
                :type="item.type || 'text'"
                :class="['form-input', { 'form-input-error': fieldErrors[item.key] }]"
                :readonly="item.readonly"
                :required="item.required"
                @blur="handleFieldValidation(item.key, form[item.key as keyof typeof form])"
                @input="() => { if (fieldErrors[item.key]) handleFieldValidation(item.key, form[item.key as keyof typeof form]) }"
              />
            </template>
            <!-- 字段错误提示 -->
            <div v-if="fieldErrors[item.key]" class="field-error">{{ fieldErrors[item.key] }}</div>
          </div>
        </div>

        <!-- 保存进度条 -->
        <div v-if="loading && saveProgress > 0" class="save-progress">
          <div class="progress-bar">
            <div class="progress-fill" :style="{ width: saveProgress + '%' }"></div>
          </div>
          <div class="progress-text">保存中... {{ saveProgress }}%</div>
        </div>

        <FeedbackButton
          class="save-btn"
          type="primary"
          block
          :loading="loading"
          :text="isFirstTime ? '下一步' : '保存'"
          @click="handleSave"
        />
        
        <!-- 重试按钮 -->
        <FeedbackButton
          v-if="showRetryButton"
          class="retry-btn"
          type="secondary"
          block
          text="重新加载"
          @click="retryLoadProfile"
        />
        
        <!-- 备用加载按钮 -->
        <FeedbackButton
          v-if="hasError"
          class="fallback-btn"
          type="secondary"
          block
          text="从缓存加载"
          @click="fetchProfileFallback"
        />
        
        <!-- 诊断按钮 -->
        <FeedbackButton
          v-if="hasError"
          class="diagnostic-btn"
          type="secondary"
          block
          text="运行诊断"
          @click="runDiagnostic"
        />
      </div>
      <ToastMessage v-if="toast.visible" :message="toast.message" :type="toast.type" @close="toast.visible=false" />
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import FeedbackButton from '@/components/FeedbackButton.vue'
import ToastMessage from '@/components/ToastMessage.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import DoctorApiService from '@/api/doctor'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { log } from '@/utils/logger'
import { runProfileDiagnostic } from '@/utils/profileDiagnostic'
import { testProfileApi } from '@/utils/apiTest'
import { errorHandler, getUserMessage, retryWithBackoff } from '@/utils/errorHandler'
import { profileCache, cacheProfile, getCachedProfile, clearProfileCache } from '@/utils/cacheManager'

// 表单字段类型定义
interface FormItem {
  group: 'basic' | 'professional'
  key: string
  label: string
  placeholder: string
  required: boolean
  type: 'text' | 'select' | 'date' | 'email' | 'tel' | 'textarea'
  readonly?: boolean
  options?: { value: string | number; label: string }[]
}

// Toast类型定义
type ToastType = 'info' | 'success' | 'error' | 'warning'
const router = useRouter()
const loading = ref(false)
const authLoading = ref(true)
const hasError = ref(false)
const errorMessage = ref('')
const profileLoaded = ref(false)
const isFirstTime = ref(false)
const toast = reactive({ visible: false, message: '', type: 'info' as ToastType })
const saveProgress = ref(0) // 保存进度
const isOnline = ref(navigator.onLine) // 网络状态
const doctorApi = new DoctorApiService()
const departments = ref([
  { id: 1, name: '内科' },
  { id: 2, name: '外科' },
  { id: 3, name: '儿科' }
])
const hospitals = ref([
  { id: 1, name: '协和医院' },
  { id: 2, name: '同济医院' },
  { id: 3, name: '人民医院' }
])
const form = reactive({
  DId: 1,
  Name: '',
  Gender: '',
  BirthDate: '',
  Phone: '',
  Email: '',
  Avatar: '',
  LicenseNumber: '',
  DepartmentId: 1,
  HospitalId: 1,
  Title: '',
  Speciality: '',
  PracticeScope: ''
})
// 统一的表单字段配置
const unifiedFormItems: FormItem[] = [
  // 基本信息组
  { 
    group: 'basic',
    key: 'Name', 
    label: '真实姓名', 
    placeholder: '请填写您的真实姓名',
    required: true,
    type: 'text'
  },
  {
    group: 'basic',
    key: 'Gender',
    label: '性别',
    placeholder: '请选择性别',
    required: true,
    type: 'select',
    options: [
      { value: '男', label: '男' },
      { value: '女', label: '女' }
    ]
  },
  {
    group: 'basic',
    key: 'BirthDate',
    label: '出生日期',
    placeholder: '请选择出生日期',
    required: false,
    type: 'date'
  },
  {
    group: 'basic',
    key: 'Phone',
    label: '手机号',
    placeholder: '请输入手机号',
    required: true,
    type: 'tel',
    readonly: true  // 手机号通常不允许修改
  },
  {
    group: 'basic',
    key: 'Email',
    label: '邮箱',
    placeholder: '请输入邮箱',
    required: false,
    type: 'email'
  },
  
  // 职业信息组
  {
    group: 'professional',
    key: 'HospitalId',
    label: '就职医院',
    placeholder: '请选择您目前所执业的医院',
    required: true,
    type: 'select'
  },
  {
    group: 'professional',
    key: 'DepartmentId',
    label: '所属科室',
    placeholder: '请选择您所属的科室',
    required: true,
    type: 'select'
  },
  {
    group: 'professional',
    key: 'Title',
    label: '职称',
    placeholder: '请填写您的职称',
    required: true,
    type: 'text'
  },
  {
    group: 'professional',
    key: 'LicenseNumber',
    label: '执业证号',
    placeholder: '请输入执业证号',
    required: false,
    type: 'text'
  },
  {
    group: 'professional',
    key: 'Speciality',
    label: '擅长领域',
    placeholder: '请填写您的擅长领域',
    required: false,
    type: 'textarea'
  },
  {
    group: 'professional',
    key: 'PracticeScope',
    label: '个人简介',
    placeholder: '请填写个人简介',
    required: false,
    type: 'textarea'
  }
]

// 分组的表单字段
const basicFormItems = computed(() => unifiedFormItems.filter(item => item.group === 'basic'))
const professionalFormItems = computed(() => unifiedFormItems.filter(item => item.group === 'professional'))

// 字段映射函数 - 处理认证页面字段到标准字段的转换
function mapAuthFieldsToProfile(authData: any): any {
  const fieldMapping = {
    // 认证页面字段 -> 标准字段
    'realName': 'Name',
    'hospital': 'HospitalId',  // 需要转换为ID
    'department': 'DepartmentId',  // 需要转换为ID
    'title': 'Title',
    'specialty': 'Speciality',
    'profile': 'PracticeScope',
    'experience': 'PracticeScope'  // 合并到个人简介
  }
  
  const mappedData: any = {}
  
  // 映射已知字段
  Object.entries(authData).forEach(([key, value]) => {
    const mappedKey = fieldMapping[key as keyof typeof fieldMapping]
    if (mappedKey) {
      if (mappedKey === 'PracticeScope' && mappedData[mappedKey]) {
        // 如果是PracticeScope且已有值，则合并
        mappedData[mappedKey] += `\n${value}`
      } else {
        mappedData[mappedKey] = value
      }
    } else {
      // 保持原字段名
      mappedData[key] = value
    }
  })
  
  // 处理医院和科室名称到ID的转换
  if (mappedData.HospitalId && typeof mappedData.HospitalId === 'string') {
    const hospital = hospitals.value.find(h => h.name === mappedData.HospitalId)
    mappedData.HospitalId = hospital ? hospital.id : 1
  }
  
  if (mappedData.DepartmentId && typeof mappedData.DepartmentId === 'string') {
    const department = departments.value.find(d => d.name === mappedData.DepartmentId)
    mappedData.DepartmentId = department ? department.id : 1
  }
  
  return mappedData
}

// 检测是否为首次使用
function detectFirstTimeUser(profile: any): boolean {
  // 检查关键字段是否为空或默认值
  const requiredFields = ['Name', 'Title', 'HospitalId', 'DepartmentId']
  
  for (const field of requiredFields) {
    const value = profile[field]
    if (!value || 
        (typeof value === 'string' && value.trim() === '') ||
        (typeof value === 'number' && value <= 0)) {
      return true
    }
  }
  
  return false
}
onMounted(async () => {
  // 监听网络状态变化
  window.addEventListener('online', () => {
    isOnline.value = true
    toast.message = '网络连接已恢复'
    toast.type = 'success'
    toast.visible = true
    
    // 网络恢复后，如果有错误状态，尝试重新加载
    if (hasError.value && !profileLoaded.value) {
      setTimeout(() => {
        retryLoadProfile()
      }, 1000)
    }
  })
  
  window.addEventListener('offline', () => {
    isOnline.value = false
    toast.message = '网络连接已断开，将使用缓存数据'
    toast.type = 'warning'
    toast.visible = true
  })
  
  await initializeAndFetchProfile()
})
const authStore = useAuthStore()

// 计算属性
const showRetryButton = computed(() => {
  return !loading.value && !authLoading.value && hasError.value
})

// 重试加载个人信息
function retryLoadProfile() {
  hasError.value = false
  errorMessage.value = ''
  loading.value = false
  authLoading.value = false
  profileLoaded.value = false
  
  // 如果网络已连接，尝试重新加载
  if (isOnline.value) {
    initializeAndFetchProfile()
  } else {
    // 如果网络未连接，提示用户
    toast.message = '网络未连接，请检查网络设置'
    toast.type = 'warning'
    toast.visible = true
    
    // 尝试从缓存加载
    fetchProfileFallback().then(result => {
      if (!result) {
        hasError.value = true
        errorMessage.value = '网络未连接，无法加载个人信息'
      }
    })
  }
}

// 增强的备用获取个人信息方法
function fetchProfileFallback(): boolean {
  try {
    log.debug('尝试从缓存获取个人信息')
    
    const doctorId = authStore.loginState.doctorId
    if (!doctorId) {
      return false
    }

    // 首先尝试从新的缓存管理器获取
    const cachedProfile = getCachedProfile(doctorId)
    if (cachedProfile) {
      log.info('从缓存管理器成功获取个人信息', { doctorId })
      processAndApplyProfile(cachedProfile)
      
      toast.message = '已从缓存加载个人信息'
      toast.type = 'warning'
      toast.visible = true
      return true
    }

    // 回退到旧的localStorage方式
    const userInfo = localStorage.getItem('doctor_info')
    if (userInfo) {
      const parsed = JSON.parse(userInfo)
      if (parsed.DId) {
        log.info('从localStorage成功获取个人信息', { doctorId: parsed.DId })
        processAndApplyProfile(parsed)
        
        // 将数据迁移到新的缓存管理器
        cacheProfile(parsed.DId, parsed)
        
        toast.message = '已从本地缓存加载个人信息'
        toast.type = 'warning'
        toast.visible = true
        return true
      }
    }
    
    log.debug('缓存中没有有效的个人信息')
    return false
  } catch (error) {
    log.error('从缓存获取个人信息失败', error)
    return false
  }
}

// 处理和应用个人信息数据
function processAndApplyProfile(profileData: any) {
  // 处理日期格式
  if (profileData.BirthDate && profileData.BirthDate !== '0001-01-01' && profileData.BirthDate !== '0000-00-00') {
    const date = new Date(profileData.BirthDate)
    if (!isNaN(date.getTime())) {
      profileData.BirthDate = date.toISOString().split('T')[0]
    } else {
      profileData.BirthDate = ''
    }
  } else {
    profileData.BirthDate = ''
  }
  
  Object.assign(form, profileData)
  profileLoaded.value = true
  hasError.value = false
}

// 运行诊断
async function runDiagnostic() {
  try {
    const diagnosticReport = await runProfileDiagnostic()
    const apiTestReport = await testProfileApi(authStore.loginState.doctorId)
    
    console.log('=== 个人信息页面诊断报告 ===')
    console.log(diagnosticReport)
    console.log('\n=== API测试报告 ===')
    console.log(apiTestReport)
    
    toast.message = '诊断完成，请查看控制台'
    toast.type = 'info'
    toast.visible = true
  } catch (error) {
    console.error('诊断失败:', error)
    toast.message = '诊断失败'
    toast.type = 'error'
    toast.visible = true
  }
}
async function initializeAndFetchProfile() {
  try {
    authLoading.value = true
    hasError.value = false
    
    // 设置超时处理
    const initTimeout = setTimeout(() => {
      if (authLoading.value) {
        log.warn('认证状态初始化超时')
        authLoading.value = false
        hasError.value = true
        toast.message = '认证状态初始化超时，请重试'
        toast.type = 'error'
        toast.visible = true
      }
    }, 5000) // 5秒超时
    
    // 等待认证状态初始化完成
    await authStore.waitForInitialization()
    
    // 清除超时定时器
    clearTimeout(initTimeout)
    
    // 检查是否已登录且有医生ID
    if (!authStore.isLoggedIn || !authStore.loginState.doctorId) {
      toast.message = '请先登录'
      toast.type = 'error'
      toast.visible = true
      setTimeout(() => {
        router.replace('/login')
      }, 1500)
      return
    }
    
    await fetchProfile()
  } catch (error) {
    log.error('初始化个人信息页面失败', error)
    hasError.value = true
    toast.message = '页面初始化失败，请重试'
    toast.type = 'error'
    toast.visible = true
  } finally {
    authLoading.value = false
  }
}

// 增强的获取个人信息方法
async function fetchProfile() {
  loading.value = true
  hasError.value = false
  
  const doctorId = authStore.loginState.doctorId
  
  log.debug('开始获取个人信息', {
    doctorId,
    hasToken: !!authStore.token,
    isLoggedIn: authStore.isLoggedIn
  })
  
  if (!doctorId) {
    hasError.value = true
    toast.message = '用户信息异常，请重新登录'
    toast.type = 'error'
    toast.visible = true
    loading.value = false
    
    log.error('获取个人信息失败：缺少医生ID', {
      loginState: authStore.loginState,
      doctorInfo: authStore.doctorInfo
    })
    return
  }

  try {
    // 设置API请求超时
    const fetchTimeout = new Promise((_, reject) => {
      setTimeout(() => reject(new Error('获取个人信息超时')), 10000) // 10秒超时
    })
    
    // 使用重试机制获取个人信息（带超时处理）
    const res = await Promise.race([
      retryWithBackoff(
        () => doctorApi.getProfile({ doctor_id: doctorId }),
        `getProfile_${doctorId}`,
        2 // 最多重试2次
      ),
      fetchTimeout
    ])

    log.debug('获取个人信息API响应', {
      hasResponse: !!res,
      hasProfile: !!res?.Profile,
      profileKeys: res?.Profile ? Object.keys(res.Profile) : []
    })
    
    if (res && res.Profile) {
      const profile = { ...res.Profile }
      
      // 检查是否为首次使用
      const isProfileIncomplete = detectFirstTimeUser(profile)
      if (isProfileIncomplete) {
        log.info('检测到首次使用或信息不完整，切换到编辑模式')
        isFirstTime.value = true
      }
      
      // 处理和应用个人信息
      processAndApplyProfile(profile)
      
      // 缓存个人信息
      cacheProfile(doctorId, profile)
      
      log.info('个人信息加载成功', {
        profileId: profile.DId,
        profileName: profile.Name,
        isFirstTime: isFirstTime.value
      })
    } else {
      throw new Error('数据格式异常：响应中缺少Profile字段')
    }
  } catch (error: any) {
    log.error('获取个人信息失败', {
      error: error.message,
      errorType: error.name,
      doctorId
    })
    
    // 使用错误处理器分析错误
    const errorInfo = errorHandler.analyzeError(error, 'fetchProfile')
    
    // 如果是404错误，可能是首次使用
    if (error.message && error.message.includes('404')) {
      log.info('检测到404错误，可能是首次使用，切换到首次使用模式')
      isFirstTime.value = true
      profileLoaded.value = true
      hasError.value = false
      
      // 设置默认信息
      form.DId = doctorId
      if (authStore.loginState.phone) {
        form.Phone = authStore.loginState.phone
      }
      
      toast.message = '欢迎使用，请填写您的个人信息'
      toast.type = 'info'
      toast.visible = true
      return
    }
    
    // 尝试从缓存获取备用数据
    log.debug('API调用失败，尝试从缓存获取')
    const fallbackSuccess = fetchProfileFallback()
    
    if (!fallbackSuccess) {
      hasError.value = true
      
      // 使用统一的错误消息
      const userMsg = getUserMessage(error, 'fetchProfile')
      errorMessage.value = userMsg + ((!isOnline.value) ? '，请检查网络连接' : '')
      toast.message = userMsg
      toast.type = 'error'
      toast.visible = true
    } else {
      // 从缓存加载成功，清除错误状态
      hasError.value = false
      profileLoaded.value = true
    }
  } finally {
    loading.value = false
  }
}
// 验证规则配置
const validationRules = {
  Name: {
    required: true,
    minLength: 2,
    maxLength: 20,
    pattern: /^[\u4e00-\u9fa5a-zA-Z\s]+$/,
    message: '姓名只能包含中文、英文和空格，长度2-20个字符'
  },
  Gender: {
    required: true,
    options: ['男', '女'],
    message: '请选择正确的性别'
  },
  Phone: {
    required: true,
    pattern: /^1[3-9]\d{9}$/,
    message: '请输入正确的手机号码'
  },
  Email: {
    required: false,
    pattern: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
    message: '请输入正确的邮箱地址'
  },
  BirthDate: {
    required: false,
    pattern: /^\d{4}-\d{2}-\d{2}$/,
    message: '请选择正确的出生日期',
    validator: (value: string) => {
      if (!value) return true
      const date = new Date(value)
      const now = new Date()
      const minDate = new Date(now.getFullYear() - 100, 0, 1)
      const maxDate = new Date(now.getFullYear() - 18, now.getMonth(), now.getDate())
      return date >= minDate && date <= maxDate
    },
    customMessage: '年龄应在18-100岁之间'
  },
  Title: {
    required: true,
    minLength: 2,
    maxLength: 20,
    message: '职称长度应在2-20个字符之间'
  },
  LicenseNumber: {
    required: false,
    pattern: /^[A-Za-z0-9]{10,20}$/,
    message: '执业证号应为10-20位字母或数字'
  },
  Speciality: {
    required: false,
    maxLength: 200,
    message: '擅长领域不能超过200个字符'
  },
  PracticeScope: {
    required: false,
    maxLength: 500,
    message: '个人简介不能超过500个字符'
  }
}

// 增强的表单验证函数
function validateForm(): boolean {
  const errors: string[] = []
  
  for (const item of unifiedFormItems) {
    const value = form[item.key as keyof typeof form]
    const rule = validationRules[item.key as keyof typeof validationRules]
    
    if (!rule) continue
    
    // 必填字段验证
    if (rule.required && (!value || (typeof value === 'string' && value.trim() === ''))) {
      errors.push(`${item.label}为必填项`)
      continue
    }
    
    // 如果字段为空且非必填，跳过其他验证
    if (!value || (typeof value === 'string' && value.trim() === '')) {
      continue
    }
    
    const stringValue = String(value).trim()
    
    // 长度验证
    if (rule.minLength && stringValue.length < rule.minLength) {
      errors.push(`${item.label}长度不能少于${rule.minLength}个字符`)
      continue
    }
    
    if (rule.maxLength && stringValue.length > rule.maxLength) {
      errors.push(`${item.label}长度不能超过${rule.maxLength}个字符`)
      continue
    }
    
    // 选项验证
    if (rule.options && !rule.options.includes(stringValue)) {
      errors.push(rule.message || `${item.label}选项不正确`)
      continue
    }
    
    // 正则表达式验证
    if (rule.pattern && !rule.pattern.test(stringValue)) {
      errors.push(rule.message || `${item.label}格式不正确`)
      continue
    }
    
    // 自定义验证器
    if (rule.validator && !rule.validator(stringValue)) {
      errors.push(rule.customMessage || rule.message || `${item.label}验证失败`)
      continue
    }
  }
  
  // 特殊业务逻辑验证
  if (form.HospitalId <= 0) {
    errors.push('请选择就职医院')
  }
  
  if (form.DepartmentId <= 0) {
    errors.push('请选择所属科室')
  }
  
  if (errors.length > 0) {
    toast.message = errors[0] // 显示第一个错误
    toast.type = 'error'
    toast.visible = true
    return false
  }
  
  return true
}

// 字段错误状态
const fieldErrors = reactive<Record<string, string>>({})

// 增强的实时验证单个字段
function validateField(key: string, value: any): string | null {
  const item = unifiedFormItems.find(item => item.key === key)
  const rule = validationRules[key as keyof typeof validationRules]
  
  if (!item || !rule) return null
  
  // 必填字段验证
  if (rule.required && (!value || (typeof value === 'string' && value.trim() === ''))) {
    return `${item.label}为必填项`
  }
  
  // 如果字段为空且非必填，清除错误
  if (!value || (typeof value === 'string' && value.trim() === '')) {
    return null
  }
  
  const stringValue = String(value).trim()
  
  // 长度验证
  if (rule.minLength && stringValue.length < rule.minLength) {
    return `${item.label}长度不能少于${rule.minLength}个字符`
  }
  
  if (rule.maxLength && stringValue.length > rule.maxLength) {
    return `${item.label}长度不能超过${rule.maxLength}个字符`
  }
  
  // 选项验证
  if (rule.options && !rule.options.includes(stringValue)) {
    return rule.message || `${item.label}选项不正确`
  }
  
  // 正则表达式验证
  if (rule.pattern && !rule.pattern.test(stringValue)) {
    return rule.message || `${item.label}格式不正确`
  }
  
  // 自定义验证器
  if (rule.validator && !rule.validator(stringValue)) {
    return rule.customMessage || rule.message || `${item.label}验证失败`
  }
  
  return null
}

// 实时验证处理函数
function handleFieldValidation(key: string, value: any) {
  const error = validateField(key, value)
  if (error) {
    fieldErrors[key] = error
  } else {
    delete fieldErrors[key]
  }
}

// 清除所有字段错误
function clearFieldErrors() {
  Object.keys(fieldErrors).forEach(key => {
    delete fieldErrors[key]
  })
}

// 增强的智能API调用 - 根据是否为首次使用选择合适的API
async function saveProfile() {
  // 处理出生日期，确保格式正确或为空
  let birthDate = form.BirthDate
  if (birthDate && birthDate !== '0001-01-01' && birthDate !== '0000-00-00') {
    // 验证日期格式
    const dateRegex = /^\d{4}-\d{2}-\d{2}$/
    if (!dateRegex.test(birthDate)) {
      birthDate = ''
    }
  } else {
    birthDate = ''
  }

  const profileData = {
    DId: form.DId,
    Name: form.Name,
    Gender: form.Gender,
    BirthDate: birthDate,
    Email: form.Email,
    Avatar: form.Avatar,
    Title: form.Title,
    Speciality: form.Speciality,
    PracticeScope: form.PracticeScope,
    LicenseNumber: form.LicenseNumber,
    DepartmentId: form.DepartmentId,
    HospitalId: form.HospitalId
  }

  const operationId = `saveProfile_${form.DId}_${Date.now()}`

  if (isFirstTime.value) {
    // 首次使用，尝试使用authentication API，如果失败则使用updateProfile
    try {
      await retryWithBackoff(
        () => doctorApi.authentication(profileData),
        `${operationId}_auth`,
        1 // 认证API只重试1次
      )
      log.info('使用authentication API保存成功')
    } catch (authError) {
      log.warn('authentication API失败，尝试使用updateProfile API', authError)
      // 如果authentication失败，回退到updateProfile
      await retryWithBackoff(
        () => doctorApi.updateProfile(profileData),
        `${operationId}_update`,
        2 // updateProfile API重试2次
      )
      log.info('使用updateProfile API保存成功')
    }
  } else {
    // 非首次使用，直接使用updateProfile API
    await retryWithBackoff(
      () => doctorApi.updateProfile(profileData),
      operationId,
      2 // 重试2次
    )
    log.info('使用updateProfile API更新成功')
  }

  // 保存成功后更新缓存
  cacheProfile(form.DId, profileData)
  log.debug('个人信息已缓存', { doctorId: form.DId })
}

// 增强的保存处理方法
async function handleSave() {
  if (loading.value) return
  
  // 检查网络状态
  if (!isOnline.value) {
    toast.message = '网络连接不可用，请检查网络设置'
    toast.type = 'error'
    toast.visible = true
    return
  }
  
  // 清除之前的字段错误
  clearFieldErrors()
  
  // 表单验证
  if (!validateForm()) {
    return
  }
  
  loading.value = true
  saveProgress.value = 0
  
  try {
    // 模拟保存进度
    const progressInterval = setInterval(() => {
      if (saveProgress.value < 90) {
        saveProgress.value += 10
      }
    }, 100)
    
    await saveProfile()
    
    // 完成进度
    clearInterval(progressInterval)
    saveProgress.value = 100
    
    const successMessage = isFirstTime.value ? '认证信息提交成功！' : '保存成功'
    toast.message = successMessage
    toast.type = 'success'
    toast.visible = true
    
    // 如果是首次使用，保存成功后切换状态
    if (isFirstTime.value) {
      isFirstTime.value = false
      profileLoaded.value = true
    }
    
    // 延迟跳转，让用户看到成功提示
    setTimeout(() => {
      router.replace('/mine')
    }, 1200)
  } catch (error: any) {
    log.error('保存个人信息失败', error)
    
    // 重置进度
    saveProgress.value = 0
    
    // 使用统一的错误处理
    const userMessage = getUserMessage(error, 'saveProfile')
    toast.message = userMessage
    toast.type = 'error'
    toast.visible = true
    
    // 如果是认证错误，跳转到登录页
    const errorInfo = errorHandler.analyzeError(error)
    if (errorInfo.type === 'authentication') {
      setTimeout(() => {
        router.replace('/login')
      }, 2000)
    }
  } finally {
    loading.value = false
    saveProgress.value = 0
  }
}
</script>
<style scoped>
.profile-view-scroll {
  height: 100vh;
  overflow-y: auto;
  background: #fff;
}
.profile-view {
  padding: 20px 16px;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 18px;
}

.title {
  font-size: 18px;
  font-weight: bold;
  text-align: left;
  margin: 0;
}

.network-status {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.network-status.offline {
  background-color: #fff2f0;
  color: #ff4d4f;
  border: 1px solid #ffccc7;
}

.status-icon {
  font-size: 14px;
}

.status-text {
  font-weight: 500;
}
.profile-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.form-label {
  font-size: 15px;
  color: #333;
}
.form-input {
  padding: 8px 12px;
  border: 1px solid #e5e5e5;
  border-radius: 6px;
  font-size: 15px;
  outline: none;
  transition: border-color 0.3s ease;
}

.form-input:focus {
  border-color: #4a90e2;
  box-shadow: 0 0 0 2px rgba(74, 144, 226, 0.1);
}

.form-input-error {
  border-color: #ff4d4f !important;
  box-shadow: 0 0 0 2px rgba(255, 77, 79, 0.1) !important;
}

.field-error {
  color: #ff4d4f;
  font-size: 12px;
  margin-top: 4px;
  line-height: 1.4;
}
.form-section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
  padding-bottom: 8px;
  border-bottom: 1px solid #f0f0f0;
}

.required {
  color: #ff4d4f;
  margin-right: 2px;
}

.form-textarea {
  min-height: 80px;
  resize: vertical;
}

.save-progress {
  margin: 16px 0;
}

.progress-bar {
  width: 100%;
  height: 4px;
  background-color: #f0f0f0;
  border-radius: 2px;
  overflow: hidden;
  margin-bottom: 8px;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #4a90e2 0%, #357ae8 100%);
  border-radius: 2px;
  transition: width 0.3s ease;
}

.progress-text {
  text-align: center;
  font-size: 12px;
  color: #666;
}

.save-btn {
  margin-top: 24px;
}

.retry-btn {
  margin-top: 12px;
}

.fallback-btn {
  margin-top: 8px;
}

.diagnostic-btn {
  margin-top: 8px;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
}

.loading-text {
  margin-top: 16px;
  color: #666;
  font-size: 14px;
}

.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
}

.error-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.error-text {
  color: #333;
  font-size: 16px;
  font-weight: 500;
  margin: 0 0 8px 0;
}

.error-desc {
  color: #666;
  font-size: 14px;
  margin: 0 0 24px 0;
  text-align: center;
}
.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-bottom: 16px;
}
.avatar-img {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  object-fit: cover;
  border: 1px solid #e5e5e5;
}
</style>
