<template>
  <div class="login-view">
    <!-- 医疗主题插图 -->
    <div class="illustration">
      <!-- 药瓶 -->
      <div class="medicine-bottle">
        <div class="bottle-cap"></div>
        <div class="bottle-body">
          <div class="bottle-label">
            <div class="cross-icon">
              <div class="cross-horizontal"></div>
              <div class="cross-vertical"></div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 药片 -->
      <div class="pills">
        <div class="pill pill-1"></div>
        <div class="pill pill-2"></div>
      </div>
      
      <!-- 背景装饰云朵 -->
      <div class="bg-clouds">
        <div class="cloud cloud-1"></div>
        <div class="cloud cloud-2"></div>
        <div class="cloud cloud-3"></div>
      </div>
    </div>
    
    <!-- 登录内容 -->
    <div class="login-content">
      <!-- 标题区域 -->
      <div class="header-section">
        <h1 class="title">手机验证码登录</h1>
        <p class="subtitle">欢迎使用优医在线问诊医生版</p>
      </div>
      
      <!-- 登录表单 -->
      <div class="form-section">
        <!-- 手机号输入 -->
        <div class="input-group phone-group">
          <div class="input-wrapper">
            <span class="country-code">+86</span>
            <div class="divider"></div>
            <input 
              v-model="formData.phone"
              type="tel"
              placeholder="请输入手机号码"
              class="input-field"
              maxlength="11"
              @input="onPhoneInput"
            />
          </div>
        </div>
        
        <!-- 验证码输入 -->
        <div class="input-group code-group">
          <div class="input-wrapper">
            <input 
              v-model="formData.smsCode"
              type="text"
              :placeholder="smsCodePlaceholder"
              class="input-field"
              :class="{ 'highlight': shouldHighlightCodeInput }"
              maxlength="6"
              @input="onSmsCodeInput"
            />
            <button 
              class="code-button"
              :class="{ 'disabled': codeButtonDisabled, 'sending': isSendingSms }"
              :disabled="codeButtonDisabled"
              @click="sendSmsCode"
            >
              {{ codeButtonText }}
            </button>
          </div>
          <!-- 验证码状态提示 -->
          <div class="sms-status" v-if="smsStatusMessage">
            <span :class="smsStatusClass">{{ smsStatusMessage }}</span>
          </div>
        </div>
        
        <!-- 账号密码登录链接 -->
        <div class="password-login-link">
          <span @click="toggleLoginMode">账号密码登陆</span>
        </div>
        

        
        <!-- 登录按钮 -->
        <button 
          class="login-button"
          :disabled="!canLogin || isLoading"
          @click="handleLogin"
          type="button"
        >
          <span v-if="isLoading">登录中...</span>
          <span v-else>登录</span>
        </button>
        

        
        <!-- 注册链接 -->
        <div class="register-link">
          <span @click="goToRegister">注册新用户</span>
        </div>
      </div>
      
      <!-- 底部信息 -->
      <div class="footer-section">
        <p class="footer-text">
          登录即代表您同意
          <span class="link-text">优医在线问诊隐私政策</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNavigationUtils } from '@/router'
import { doctorApi } from '@/api/doctor'
import { validatePhone, getPhoneValidationMessage } from '@/utils/validation'
import { showToast } from 'vant'
import { log } from '@/utils/logger'
// import FeedbackButton from '@/components/FeedbackButton.vue' // 已替换为原生button

const router = useRouter()
const authStore = useAuthStore()
const navigationUtils = useNavigationUtils()

// 表单数据
const formData = reactive({
  phone: '',
  smsCode: ''
})

// 验证码按钮状态
const codeCountdown = ref(0)
const isLoading = ref(false)
const isSendingSms = ref(false)
const smsCodeSent = ref(false)
const smsStatusMessage = ref('')
const smsStatusClass = ref('')

// 计算属性
const codeButtonDisabled = computed(() => {
  return !validatePhone(formData.phone) || codeCountdown.value > 0 || isLoading.value
})

const codeButtonText = computed(() => {
  if (codeCountdown.value > 0) {
    return `获取验证码(${codeCountdown.value})`
  }
  return '获取验证码'
})

const canLogin = computed(() => {
  const phoneValid = validatePhone(formData.phone)
  const codeValid = formData.smsCode.length >= 4
  const result = phoneValid && codeValid && !isLoading.value
  
  // 调试信息
  log.debug('登录按钮状态计算', {
    hasPhone: !!formData.phone,
    phoneValid,
    codeLength: formData.smsCode.length,
    codeValid,
    isLoading: isLoading.value,
    result
  })
  
  return result
})

const smsCodePlaceholder = computed(() => {
  // 正在发送中
  if (isSendingSms.value) {
    return '正在发送验证码...'
  }
  
  // 发送失败
  if (smsStatusMessage.value && smsStatusClass.value === 'error') {
    return '验证码发送失败，请重试'
  }
  
  // 发送成功但等待中
  if (!smsCodeSent.value && codeCountdown.value > 0) {
    return '验证码发送中，请稍候...'
  }
  
  // 输入框已启用，可以输入验证码
  if (smsCodeSent.value) {
    return '请输入收到的验证码'
  }
  
  // 默认状态 - 允许直接输入验证码
  return '请输入验证码或点击获取'
})

const shouldHighlightCodeInput = computed(() => {
  return smsCodeSent.value && codeCountdown.value > 50 // 发送后前10秒高亮提示
})

// 方法
const onPhoneInput = () => {
  // 限制只能输入数字
  formData.phone = formData.phone.replace(/\D/g, '')
  // 重置验证码相关状态
  if (smsCodeSent.value) {
    smsCodeSent.value = false
    smsStatusMessage.value = ''
    formData.smsCode = ''
  }
}

const onSmsCodeInput = () => {
  // 限制只能输入数字
  formData.smsCode = formData.smsCode.replace(/\D/g, '')
  
  // 当用户开始输入验证码时，清除提示消息
  if (formData.smsCode.length > 0 && smsStatusMessage.value) {
    smsStatusMessage.value = ''
  }
  
  // 调试信息
  log.debug('验证码输入状态', {
    hasCode: !!formData.smsCode,
    codeLength: formData.smsCode.length
  })
}

const sendSmsCode = async () => {
  if (codeButtonDisabled.value) return
  
  const phoneError = getPhoneValidationMessage(formData.phone)
  if (phoneError) {
    showToast({
      message: phoneError,
      type: 'fail'
    })
    return
  }
  
  try {
    isSendingSms.value = true
    smsCodeSent.value = false // 确保发送开始时输入框保持禁用
    smsStatusMessage.value = '正在发送验证码，请稍候...'
    smsStatusClass.value = 'sending'
    
    await doctorApi.sendLoginSms(formData.phone)
    
    // 发送成功，立即启用输入框
    smsCodeSent.value = true
    smsStatusMessage.value = '验证码已发送至您的手机，请注意查收'
    smsStatusClass.value = 'success'
    
    showToast({
      message: '验证码已发送',
      type: 'success'
    })
    
    // 开始倒计时
    startCountdown()
    
    // 2秒后提示输入验证码
    setTimeout(() => {
      if (formData.smsCode.length === 0) {
        smsStatusMessage.value = '请输入收到的验证码'
        smsStatusClass.value = 'hint'
      }
    }, 2000)
    
  } catch (error: any) {
    // 发送失败时仍然启用输入框，允许用户手动输入验证码
    smsCodeSent.value = true
    smsStatusMessage.value = '验证码发送失败，但您仍可输入验证码'
    smsStatusClass.value = 'error'
    
    showToast({
      message: error.message || '验证码发送失败，请重试或直接输入验证码',
      type: 'fail'
    })
    
    // 3秒后清除错误消息
    setTimeout(() => {
      smsStatusMessage.value = ''
    }, 3000)
  } finally {
    isSendingSms.value = false
  }
}

const startCountdown = () => {
  codeCountdown.value = 60
  const timer = setInterval(() => {
    codeCountdown.value--
    if (codeCountdown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

const handleLogin = async () => {
  log.debug('开始执行登录流程', {
    hasPhone: !!formData.phone,
    hasCode: !!formData.smsCode,
    canLogin: canLogin.value
  })
  
  // 验证表单数据
  if (!canLogin.value) {
    const phoneValid = validatePhone(formData.phone)
    const codeValid = formData.smsCode.length >= 4
    
    let message = '请检查输入信息'
    if (!phoneValid) {
      message = '请输入正确的手机号码'
    } else if (!codeValid) {
      message = '请输入4-6位验证码'
    }
    
    showToast({
      message,
      type: 'fail',
      duration: 2000
    })
    return
  }
  
  try {
    // 设置加载状态
    isLoading.value = true
    
    log.debug('开始调用登录API')
    
    // 调用登录API
    const response = await doctorApi.login({
      Phone: formData.phone,
      Password: '', // 验证码登录不需要密码
      SendSmsCode: formData.smsCode
    })
    
    log.info('登录API调用成功')
    
    // 构建医生信息对象
    const doctorInfo = {
      DId: response.DId,
      Name: '医生', // 临时名称，后续可以通过API获取完整信息
      Phone: formData.phone,
      Email: '',
      Avatar: '',
      Status: 'active'
    }
    
    // 保存登录状态到store
    authStore.login('temp_token', doctorInfo, formData.phone, true)
    
    log.info('登录状态已保存到store')
    
    // 显示成功提示
    showToast({
      message: '登录成功，正在跳转...',
      type: 'success',
      duration: 1500
    })
    
    // 延迟跳转，确保toast显示和状态保存完成
    setTimeout(() => {
      handleLoginSuccess(doctorInfo)
    }, 800)
    
  } catch (error: any) {
    log.error('登录失败', error)
    
    // 处理不同类型的错误
    let errorMessage = '登录失败，请重试'
    
    if (error.message) {
      if (error.message.includes('验证码')) {
        errorMessage = '验证码错误，请重新输入'
      } else if (error.message.includes('手机号')) {
        errorMessage = '手机号不存在，请先注册'
      } else if (error.message.includes('网络')) {
        errorMessage = '网络连接失败，请检查网络'
      } else {
        errorMessage = error.message
      }
    }
    
    showToast({
      message: errorMessage,
      type: 'fail',
      duration: 3000
    })
    
    // 如果是验证码错误，清空验证码输入框
    if (errorMessage.includes('验证码')) {
      formData.smsCode = ''
    }
    
  } finally {
    // 恢复按钮状态
    isLoading.value = false
  }
}

const toggleLoginMode = () => {
  showToast({
    message: '密码登录功能开发中',
    type: 'fail'
  })
}

const goToRegister = () => {
  navigationUtils.safePush('/register')
}





// 处理登录成功后的跳转
const handleLoginSuccess = (userData?: any) => {
  try {
    log.debug('开始处理登录成功跳转')
    
    // 如果有用户数据，更新到store（已经在login方法中处理了）
    if (userData) {
      log.debug('用户数据已保存到store')
    }
    
    // 执行页面跳转，使用多重备用方案
    performNavigation()
    
  } catch (error) {
    log.error('登录成功处理出错', error)
    // 最后的备用方法
    window.location.href = '/home'
  }
}

// 执行页面跳转的方法
const performNavigation = () => {
  log.debug('开始执行页面跳转')
  
  try {
    // 方案1: 使用导航工具
    navigationUtils.forceToHome()
    log.debug('使用导航工具跳转成功')
  } catch (error) {
    log.error('导航工具跳转失败', error)
    
    try {
      // 方案2: 使用路由器直接跳转
      router.push('/home')
      log.debug('使用路由器跳转成功')
    } catch (routerError) {
      log.error('路由器跳转失败', routerError)
      
      // 方案3: 使用window.location强制跳转
      log.debug('使用window.location强制跳转')
      window.location.href = '/home'
    }
  }
}

onMounted(() => {
  // 如果已经登录，直接跳转
  if (authStore.isLoggedIn && authStore.checkTokenExpiry()) {
    handleLoginSuccess()
  }
  
  // 如果有倒计时正在进行，说明之前发送过验证码，启用输入框
  if (codeCountdown.value > 0) {
    smsCodeSent.value = true
  }
})
</script>

<style lang="scss" scoped>
.login-view {
  width: 100%;
  min-height: 100vh;
  background: linear-gradient(135deg, #f8f9ff 0%, #e8f2ff 100%);
  display: flex;
  flex-direction: column;
  position: relative;
  overflow-x: hidden;
}

// 医疗主题插图
.illustration {
  position: relative;
  height: 280px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 60px;
  
  .medicine-bottle {
    position: relative;
    z-index: 3;
    
    .bottle-cap {
      width: 40px;
      height: 15px;
      background: #f39c12;
      border-radius: 20px 20px 0 0;
      margin: 0 auto 2px;
    }
    
    .bottle-body {
      width: 60px;
      height: 100px;
      background: #3498db;
      border-radius: 0 0 30px 30px;
      position: relative;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .bottle-label {
        width: 40px;
        height: 40px;
        background: rgba(255, 255, 255, 0.9);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .cross-icon {
          position: relative;
          
          .cross-horizontal,
          .cross-vertical {
            position: absolute;
            background: #e74c3c;
            border-radius: 2px;
          }
          
          .cross-horizontal {
            width: 16px;
            height: 3px;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
          }
          
          .cross-vertical {
            width: 3px;
            height: 16px;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
          }
        }
      }
    }
  }
  
  .pills {
    position: absolute;
    
    .pill {
      position: absolute;
      width: 20px;
      height: 12px;
      border-radius: 10px;
      
      &.pill-1 {
        background: linear-gradient(45deg, #f39c12 50%, #e67e22 50%);
        top: 50px;
        left: -40px;
        transform: rotate(-20deg);
        animation: float 3s ease-in-out infinite;
      }
      
      &.pill-2 {
        background: linear-gradient(45deg, #9b59b6 50%, #8e44ad 50%);
        bottom: 80px;
        right: -30px;
        transform: rotate(30deg);
        animation: float 3s ease-in-out infinite reverse;
      }
    }
  }
  
  .bg-clouds {
    position: absolute;
    width: 100%;
    height: 100%;
    
    .cloud {
      position: absolute;
      background: rgba(255, 255, 255, 0.6);
      border-radius: 50px;
      
      &::before,
      &::after {
        content: '';
        position: absolute;
        background: rgba(255, 255, 255, 0.6);
        border-radius: 50px;
      }
      
      &.cloud-1 {
        width: 35px;
        height: 18px;
        top: 30px;
        left: 20px;
        animation: float 6s ease-in-out infinite;
        
        &::before {
          width: 18px;
          height: 18px;
          top: -9px;
          left: 4px;
        }
        
        &::after {
          width: 22px;
          height: 22px;
          top: -11px;
          right: 4px;
        }
      }
      
      &.cloud-2 {
        width: 30px;
        height: 15px;
        top: 60px;
        right: 30px;
        animation: float 8s ease-in-out infinite reverse;
        
        &::before {
          width: 15px;
          height: 15px;
          top: -8px;
          left: 3px;
        }
        
        &::after {
          width: 18px;
          height: 18px;
          top: -9px;
          right: 3px;
        }
      }
      
      &.cloud-3 {
        width: 25px;
        height: 12px;
        bottom: 40px;
        left: 50%;
        transform: translateX(-50%);
        animation: float 7s ease-in-out infinite;
        
        &::before {
          width: 12px;
          height: 12px;
          top: -6px;
          left: 2px;
        }
        
        &::after {
          width: 15px;
          height: 15px;
          top: -7px;
          right: 2px;
        }
      }
    }
  }
}

// 登录内容
.login-content {
  flex: 1;
  padding: 0 $spacing-lg;
  display: flex;
  flex-direction: column;
}

.header-section {
  text-align: center;
  margin-bottom: $spacing-xl;
  
  .title {
    font-size: $font-size-xxl;
    font-weight: 600;
    color: $text-primary;
    margin-bottom: $spacing-sm;
  }
  
  .subtitle {
    font-size: $font-size-md;
    color: $text-secondary;
    font-weight: 400;
  }
}

.form-section {
  flex: 1;
  
  .input-group {
    margin-bottom: $spacing-lg;
    
    .input-wrapper {
      display: flex;
      align-items: center;
      background: $bg-secondary;
      border-radius: $border-radius;
      border: 1px solid $border-light;
      overflow: hidden;
      transition: border-color 0.3s ease;
      
      &:focus-within {
        border-color: $primary-color;
      }
      
      .country-code {
        padding: 0 $spacing-md;
        color: $text-primary;
        font-weight: 500;
        font-size: $font-size-md;
        white-space: nowrap;
      }
      
      .divider {
        width: 1px;
        height: 20px;
        background: $border-color;
      }
      
      .input-field {
        flex: 1;
        padding: $spacing-md;
        border: none;
        background: transparent;
        font-size: $font-size-md;
        color: $text-primary;
        
        &::placeholder {
          color: $text-placeholder;
        }
        
        &:focus {
          outline: none;
        }
      }
      
      .code-button {
        padding: $spacing-sm $spacing-md;
        background: transparent;
        border: none;
        color: $primary-color;
        font-size: $font-size-sm;
        font-weight: 500;
        cursor: pointer;
        white-space: nowrap;
        transition: color 0.3s ease;
        
        &:hover:not(.disabled) {
          color: $primary-dark;
        }
        
        &.disabled {
          color: $text-placeholder;
          cursor: not-allowed;
        }
      }
    }
  }
  
  .password-login-link {
    text-align: left;
    margin-bottom: $spacing-xl;
    
    span {
      color: $primary-color;
      font-size: $font-size-sm;
      cursor: pointer;
      transition: color 0.3s ease;
      
      &:hover {
        color: $primary-dark;
      }
    }
  }
  
  .login-button {
    width: 100%;
    height: 50px;
    background: $primary-color;
    border: none;
    border-radius: $border-radius-large;
    color: $text-white;
    font-size: $font-size-md;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    margin-bottom: $spacing-lg;
    
    &:hover:not(:disabled) {
      background: $primary-dark;
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(0, 122, 255, 0.3);
    }
    
    &:active:not(:disabled) {
      transform: translateY(0);
    }
    
    &:disabled {
      background: #CCCCCC !important;
      color: #999999 !important;
      cursor: not-allowed !important;
      transform: none !important;
      box-shadow: none !important;
    }
    
    &.sending {
      background: $primary-light;
      cursor: wait;
      
      &::after {
        content: '';
        position: absolute;
        top: 50%;
        right: 8px;
        width: 12px;
        height: 12px;
        border: 2px solid transparent;
        border-top: 2px solid white;
        border-radius: 50%;
        animation: spin 1s linear infinite;
        transform: translateY(-50%);
      }
    }
  }
  
  .register-link {
    text-align: center;
    
    span {
      color: $text-secondary;
      font-size: $font-size-md;
      cursor: pointer;
      transition: color 0.3s ease;
      
      &:hover {
        color: $primary-color;
      }
    }
  }
  
  // 验证码状态提示样式
  .sms-status {
    margin-top: 8px;
    padding: 0 4px;
    
    span {
      font-size: $font-size-xs;
      line-height: 1.4;
      
      &.sending {
        color: $primary-color;
      }
      
      &.success {
        color: #52c41a;
      }
      
      &.error {
        color: #ff4d4f;
      }
      
      &.hint {
        color: $text-secondary;
      }
    }
  }
  
  // 输入框高亮效果
  .input-field {
    &.highlight {
      border-color: $primary-color;
      box-shadow: 0 0 0 2px rgba(0, 122, 255, 0.1);
      animation: pulse 2s ease-in-out infinite;
    }
    
    &:disabled {
      background-color: #f5f5f5;
      color: #999;
      cursor: not-allowed;
    }
  }
}

.footer-section {
  padding: $spacing-lg 0;
  text-align: center;
  
  .footer-text {
    font-size: $font-size-xs;
    color: $text-secondary;
    line-height: 1.5;
    
    .link-text {
      color: $primary-color;
      cursor: pointer;
      
      &:hover {
        text-decoration: underline;
      }
    }
  }
}

// 动画效果
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}

@keyframes spin {
  0% {
    transform: translateY(-50%) rotate(0deg);
  }
  100% {
    transform: translateY(-50%) rotate(360deg);
  }
}

@keyframes pulse {
  0%, 100% {
    box-shadow: 0 0 0 2px rgba(0, 122, 255, 0.1);
  }
  50% {
    box-shadow: 0 0 0 4px rgba(0, 122, 255, 0.2);
  }
}

// 登录按钮样式
.login-button {
  width: 100%;
  height: 48px;
  background: $primary-color;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  
  &:hover:not(:disabled) {
    background: $primary-dark;
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba($primary-color, 0.3);
  }
  
  &:active:not(:disabled) {
    transform: translateY(0);
  }
  
  &:disabled {
    background: #d1d5db;
    color: #9ca3af;
    cursor: not-allowed;
    transform: none;
    box-shadow: none;
  }
}

// 响应式适配
@media (max-width: 375px) {
  .login-content {
    padding: 0 $spacing-md;
  }
  
  .illustration {
    height: 240px;
    margin-top: 40px;
  }
}
</style>