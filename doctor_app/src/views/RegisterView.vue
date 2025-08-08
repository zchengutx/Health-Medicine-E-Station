<template>
  <div class="register-view">
    <!-- 注册内容 -->
    <div class="register-content">
      <!-- 标题区域 -->
      <div class="header-section">
        <h1 class="title">注册医生账号</h1>
        <p class="subtitle">欢迎使用优医在线问诊医生版</p>
      </div>
      
      <!-- 注册表单 -->
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
              placeholder="请输入验证码"
              class="input-field"
              maxlength="6"
              inputmode="numeric"
              pattern="[0-9]*"
            />
            <button 
              class="code-button"
              :class="{ 'disabled': codeButtonDisabled }"
              :disabled="codeButtonDisabled"
              @click="sendSmsCode"
            >
              {{ codeButtonText }}
            </button>
          </div>
        </div>
        
        <!-- 密码输入 -->
        <div class="input-group password-group">
          <div class="input-wrapper">
            <input 
              v-model="formData.password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="请输入密码"
              class="input-field"
              maxlength="20"
            />
            <button 
              class="password-toggle"
              @click="togglePasswordVisibility"
              type="button"
            >
              <van-icon :name="showPassword ? 'eye-o' : 'closed-eye'" />
            </button>
          </div>
          <div class="password-strength" v-if="formData.password">
            <div class="strength-bar">
              <div 
                class="strength-fill" 
                :class="passwordStrengthClass"
                :style="{ width: passwordStrengthWidth }"
              ></div>
            </div>
            <span class="strength-text" :class="passwordStrengthClass">
              {{ passwordStrengthText }}
            </span>
          </div>
        </div>
        
        <!-- 注册按钮 -->
        <FeedbackButton
          text="立即注册"
          type="primary"
          size="large"
          block
          :loading="isLoading"
          :disabled="!canRegister"
          @click="handleRegister"
        />
        
        <!-- 登录链接 -->
        <div class="login-link">
          <span @click="goToLogin">已有账号，去登录</span>
        </div>
      </div>
      
      <!-- 底部信息 -->
      <div class="footer-section">
        <p class="footer-text">
          注册即代表您同意
          <span class="link-text">优医在线问诊隐私政策</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useNavigationUtils } from '@/router'
import { doctorApi } from '@/api/doctor'
import { validatePhone, validatePassword, getPhoneValidationMessage } from '@/utils/validation'
import { showToast } from 'vant'
import FeedbackButton from '@/components/FeedbackButton.vue'

const authStore = useAuthStore()
const navigationUtils = useNavigationUtils()

// 表单数据
const formData = reactive({
  phone: '',
  smsCode: '',
  password: ''
})

// 组件状态
const codeCountdown = ref(0)
const isLoading = ref(false)
const showPassword = ref(false)

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

const passwordValidation = computed(() => {
  return validatePassword(formData.password)
})

const passwordStrengthClass = computed(() => {
  if (!formData.password) return ''
  
  const validation = passwordValidation.value
  if (!validation.isValid) return 'weak'
  
  if (validation.message.includes('很强')) return 'strong'
  if (validation.message.includes('良好')) return 'medium'
  return 'weak'
})

const passwordStrengthWidth = computed(() => {
  const strengthClass = passwordStrengthClass.value
  switch (strengthClass) {
    case 'weak': return '33%'
    case 'medium': return '66%'
    case 'strong': return '100%'
    default: return '0%'
  }
})

const passwordStrengthText = computed(() => {
  return passwordValidation.value.message
})

const canRegister = computed(() => {
  return validatePhone(formData.phone) && 
         formData.smsCode.length >= 4 && 
         passwordValidation.value.isValid &&
         !isLoading.value
})

// 方法
const onPhoneInput = () => {
  // 限制只能输入数字
  formData.phone = formData.phone.replace(/\D/g, '')
}

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value
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
    isLoading.value = true
    await doctorApi.sendRegisterSms(formData.phone)
    
    showToast({
      message: '验证码已发送',
      type: 'success'
    })
    
    // 开始倒计时
    startCountdown()
  } catch (error: any) {
    showToast({
      message: error.message || '发送验证码失败',
      type: 'fail'
    })
  } finally {
    isLoading.value = false
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

const handleRegister = async () => {
  if (!canRegister.value) return
  
  // 验证密码强度
  const passwordCheck = validatePassword(formData.password)
  if (!passwordCheck.isValid) {
    showToast({
      message: passwordCheck.message,
      type: 'fail'
    })
    return
  }
  
  try {
    isLoading.value = true
    
    await doctorApi.register({
      Phone: formData.phone,
      Password: formData.password,
      SendSmsCode: formData.smsCode
    })
    
    showToast({
      message: '注册成功，请登录',
      type: 'success'
    })
    
    // 注册成功后跳转到登录页
    setTimeout(() => {
      handleRegisterSuccess()
    }, 1500)
    
  } catch (error: any) {
    showToast({
      message: error.message || '注册失败',
      type: 'fail'
    })
  } finally {
    isLoading.value = false
  }
}

const goToLogin = () => {
  navigationUtils.safePush('/login')
}

// 处理注册成功后的跳转
const handleRegisterSuccess = () => {
  // 注册成功后跳转到登录页，并预填手机号
  navigationUtils.safePush('/login')
}

onMounted(() => {
  // 如果已经登录，直接跳转到首页
  if (authStore.isLoggedIn && authStore.checkTokenExpiry()) {
    navigationUtils.toHome()
  }
})
</script>

<style lang="scss" scoped>
.register-view {
  width: 100%;
  min-height: 100vh;
  background: linear-gradient(135deg, #f8f9ff 0%, #e8f2ff 100%);
  display: flex;
  flex-direction: column;
  position: relative;
  overflow-x: hidden;
}

// 注册内容
.register-content {
  flex: 1;
  padding: 60px $spacing-lg $spacing-lg;
  display: flex;
  flex-direction: column;
  justify-content: center;
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
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
  
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
      
      .password-toggle {
        padding: $spacing-sm $spacing-md;
        background: transparent;
        border: none;
        color: $text-secondary;
        cursor: pointer;
        display: flex;
        align-items: center;
        transition: color 0.3s ease;
        
        &:hover {
          color: $primary-color;
        }
      }
    }
    
    // 密码强度指示器
    .password-strength {
      margin-top: $spacing-sm;
      display: flex;
      align-items: center;
      gap: $spacing-sm;
      
      .strength-bar {
        flex: 1;
        height: 4px;
        background: $border-light;
        border-radius: 2px;
        overflow: hidden;
        
        .strength-fill {
          height: 100%;
          border-radius: 2px;
          transition: all 0.3s ease;
          
          &.weak {
            background: $error-color;
          }
          
          &.medium {
            background: $warning-color;
          }
          
          &.strong {
            background: $success-color;
          }
        }
      }
      
      .strength-text {
        font-size: $font-size-xs;
        font-weight: 500;
        white-space: nowrap;
        
        &.weak {
          color: $error-color;
        }
        
        &.medium {
          color: $warning-color;
        }
        
        &.strong {
          color: $success-color;
        }
      }
    }
  }
  
  .register-button {
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
    
    &:hover:not(.disabled) {
      background: $primary-dark;
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(0, 122, 255, 0.3);
    }
    
    &:active:not(.disabled) {
      transform: translateY(0);
    }
    
    &.disabled {
      background: #CCCCCC;
      cursor: not-allowed;
      transform: none;
      box-shadow: none;
    }
  }
  
  .login-link {
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

// 响应式适配
@media (max-width: 375px) {
  .register-content {
    padding: 40px $spacing-md $spacing-md;
  }
  
  .form-section {
    .input-group {
      margin-bottom: $spacing-md;
    }
  }
}

// 动画效果
.register-content {
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>