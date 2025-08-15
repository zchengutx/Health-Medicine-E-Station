<template>
  <div class="splash-view">
    <!-- 跳过按钮 -->
    <div class="skip-button" @click="skipSplash" v-if="countdown > 0">
      <CountdownTimer :text="`跳过 ${countdown}`" :countdown="countdown" />
    </div>
    
    <!-- 启动页内容 -->
    <div class="splash-content">
      <!-- 医生患者互动插图 -->
      <div class="illustration">
        <!-- 背景装饰元素 -->
        <div class="bg-elements">
          <div class="cloud cloud-1"></div>
          <div class="cloud cloud-2"></div>
          <div class="cloud cloud-3"></div>
          <div class="cloud cloud-4"></div>
        </div>
        
        <!-- 主要插图 -->
        <div class="main-illustration">
          <!-- 患者 -->
          <div class="patient">
            <div class="patient-body">
              <div class="patient-head"></div>
              <div class="patient-torso"></div>
              <div class="patient-arm-left"></div>
              <div class="patient-arm-right"></div>
              <div class="patient-leg-left"></div>
              <div class="patient-leg-right"></div>
            </div>
          </div>
          
          <!-- 手机屏幕 -->
          <div class="phone-screen">
            <div class="phone-frame">
              <div class="phone-notch"></div>
              <div class="phone-content">
                <!-- 医生头像 -->
                <div class="doctor-avatar">
                  <div class="doctor-head"></div>
                  <div class="doctor-coat"></div>
                  <div class="stethoscope"></div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 装饰植物 -->
          <div class="plant">
            <div class="plant-pot"></div>
            <div class="plant-leaves">
              <div class="leaf leaf-1"></div>
              <div class="leaf leaf-2"></div>
              <div class="leaf leaf-3"></div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 品牌信息 -->
      <div class="brand-info">
        <div class="logo">
          <div class="logo-icon">
            <span class="logo-text">优医</span>
          </div>
        </div>
        <div class="slogan">您身边的健康管理专家</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNavigationUtils } from '@/router'
import CountdownTimer from '@/components/CountdownTimer.vue'
import { log } from '@/utils/logger'

const router = useRouter()
const authStore = useAuthStore()
const navigationUtils = useNavigationUtils()
const countdown = ref(3)
let timer: NodeJS.Timeout | null = null

// 跳转逻辑
const navigateToNextPage = () => {
  log.debug('开始执行启动页跳转逻辑')
  
  // 初始化认证状态
  authStore.initAuth()
  log.debug('认证状态初始化完成', { isLoggedIn: authStore.isLoggedIn })
  
  // 检查登录状态并跳转
  if (authStore.isLoggedIn) {
    log.debug('用户已登录，检查token有效性')
    // 如果已登录，检查token是否有效
    if (authStore.checkTokenExpiry()) {
      log.debug('token有效，跳转到首页')
      navigationUtils.toHome()
    } else {
      log.debug('token过期，跳转到登录页')
      // token过期，跳转到登录页
      navigationUtils.safePush('/login')
    }
  } else {
    log.debug('用户未登录，跳转到登录页')
    // 未登录，跳转到登录页
    navigationUtils.safePush('/login')
  }
}

// 跳过启动页
const skipSplash = () => {
  if (timer) {
    clearInterval(timer)
    timer = null
  }
  countdown.value = 0
  navigateToNextPage()
}

// 启动倒计时
const startCountdown = () => {
  log.debug('开始启动倒计时', { initialValue: countdown.value })
  timer = setInterval(() => {
    countdown.value--
    log.debug('倒计时更新', { countdown: countdown.value })
    if (countdown.value <= 0) {
      log.debug('倒计时结束，准备跳转')
      clearInterval(timer!)
      timer = null
      navigateToNextPage()
    }
  }, 1000)
}

onMounted(() => {
  // 延迟启动倒计时，确保页面渲染完成
  setTimeout(() => {
    startCountdown()
  }, 100)
})

onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
    timer = null
  }
})
</script>

<style lang="scss" scoped>
.splash-view {
  position: relative;
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f8f9ff 0%, #e8f2ff 100%);
  color: $text-primary;
  overflow: hidden;
}

.skip-button {
  position: absolute;
  top: 60px;
  right: 20px;
  z-index: 10;
  cursor: pointer;
  transition: transform 0.2s ease;
  
  &:active {
    transform: scale(0.95);
  }
}

.splash-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: $spacing-xl;
  width: 100%;
  max-width: 400px;
}

.illustration {
  position: relative;
  width: 300px;
  height: 280px;
  margin-bottom: 80px;
  
  .bg-elements {
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
        width: 40px;
        height: 20px;
        top: 20px;
        left: 30px;
        animation: float 6s ease-in-out infinite;
        
        &::before {
          width: 20px;
          height: 20px;
          top: -10px;
          left: 5px;
        }
        
        &::after {
          width: 25px;
          height: 25px;
          top: -12px;
          right: 5px;
        }
      }
      
      &.cloud-2 {
        width: 30px;
        height: 15px;
        top: 40px;
        right: 40px;
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
        width: 35px;
        height: 18px;
        bottom: 60px;
        left: 20px;
        animation: float 7s ease-in-out infinite;
        
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
      
      &.cloud-4 {
        width: 25px;
        height: 12px;
        bottom: 80px;
        right: 30px;
        animation: float 9s ease-in-out infinite reverse;
        
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
  
  .main-illustration {
    position: relative;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  // 患者
  .patient {
    position: absolute;
    left: 20px;
    bottom: 20px;
    z-index: 2;
    
    .patient-body {
      position: relative;
      
      .patient-head {
        width: 35px;
        height: 35px;
        background: #fdbcb4;
        border-radius: 50%;
        position: relative;
        margin-bottom: 5px;
        
        &::before {
          content: '';
          position: absolute;
          width: 20px;
          height: 15px;
          background: #4a4a4a;
          border-radius: 15px 15px 0 0;
          top: -5px;
          left: 50%;
          transform: translateX(-50%);
        }
      }
      
      .patient-torso {
        width: 40px;
        height: 50px;
        background: #6c5ce7;
        border-radius: 20px 20px 5px 5px;
        position: relative;
      }
      
      .patient-arm-left,
      .patient-arm-right {
        position: absolute;
        width: 12px;
        height: 35px;
        background: #fdbcb4;
        border-radius: 6px;
        top: 40px;
      }
      
      .patient-arm-left {
        left: -8px;
        transform: rotate(-20deg);
        animation: wave 2s ease-in-out infinite;
      }
      
      .patient-arm-right {
        right: -8px;
        transform: rotate(20deg);
      }
      
      .patient-leg-left,
      .patient-leg-right {
        position: absolute;
        width: 15px;
        height: 30px;
        background: #2d3436;
        border-radius: 0 0 8px 8px;
        top: 85px;
      }
      
      .patient-leg-left {
        left: 5px;
      }
      
      .patient-leg-right {
        right: 5px;
      }
    }
  }
  
  // 手机屏幕
  .phone-screen {
    position: absolute;
    right: 30px;
    top: 50%;
    transform: translateY(-50%) rotate(-10deg);
    z-index: 3;
    
    .phone-frame {
      width: 120px;
      height: 200px;
      background: #2d3436;
      border-radius: 25px;
      padding: 8px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
      
      .phone-notch {
        width: 60px;
        height: 6px;
        background: #636e72;
        border-radius: 3px;
        margin: 0 auto 8px;
      }
      
      .phone-content {
        width: 100%;
        height: calc(100% - 14px);
        background: #ffffff;
        border-radius: 18px;
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
        overflow: hidden;
        
        .doctor-avatar {
          position: relative;
          
          .doctor-head {
            width: 45px;
            height: 45px;
            background: #fdbcb4;
            border-radius: 50%;
            position: relative;
            margin-bottom: 8px;
            
            &::before {
              content: '';
              position: absolute;
              width: 25px;
              height: 20px;
              background: #2d3436;
              border-radius: 20px 20px 0 0;
              top: -8px;
              left: 50%;
              transform: translateX(-50%);
            }
          }
          
          .doctor-coat {
            width: 50px;
            height: 60px;
            background: #ffffff;
            border: 2px solid #ddd;
            border-radius: 25px 25px 8px 8px;
            position: relative;
            
            &::before {
              content: '';
              position: absolute;
              width: 30px;
              height: 8px;
              background: #007AFF;
              border-radius: 4px;
              top: 15px;
              left: 50%;
              transform: translateX(-50%);
            }
          }
          
          .stethoscope {
            position: absolute;
            top: 35px;
            left: 50%;
            transform: translateX(-50%);
            width: 3px;
            height: 25px;
            background: #636e72;
            border-radius: 2px;
            
            &::before {
              content: '';
              position: absolute;
              width: 8px;
              height: 8px;
              background: #636e72;
              border-radius: 50%;
              bottom: -4px;
              left: 50%;
              transform: translateX(-50%);
            }
          }
        }
      }
    }
  }
  
  // 装饰植物
  .plant {
    position: absolute;
    left: 10px;
    bottom: 10px;
    z-index: 1;
    
    .plant-pot {
      width: 30px;
      height: 25px;
      background: #95a5a6;
      border-radius: 0 0 15px 15px;
      position: relative;
      
      &::before {
        content: '';
        position: absolute;
        width: 35px;
        height: 8px;
        background: #7f8c8d;
        border-radius: 4px;
        top: -4px;
        left: 50%;
        transform: translateX(-50%);
      }
    }
    
    .plant-leaves {
      position: relative;
      
      .leaf {
        position: absolute;
        width: 15px;
        height: 25px;
        background: #00b894;
        border-radius: 50% 10px;
        bottom: 20px;
        
        &.leaf-1 {
          left: 5px;
          transform: rotate(-30deg);
        }
        
        &.leaf-2 {
          left: 12px;
          transform: rotate(0deg);
        }
        
        &.leaf-3 {
          left: 18px;
          transform: rotate(30deg);
        }
      }
    }
  }
}

.brand-info {
  .logo {
    margin-bottom: $spacing-lg;
    
    .logo-icon {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: 80px;
      height: 80px;
      background: $primary-color;
      border-radius: 20px;
      box-shadow: 0 8px 32px rgba(0, 122, 255, 0.3);
      
      .logo-text {
        color: $text-white;
        font-size: 28px;
        font-weight: bold;
      }
    }
  }
  
  .slogan {
    font-size: $font-size-lg;
    color: $text-secondary;
    font-weight: 400;
    letter-spacing: 0.5px;
  }
}

// 动画效果
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes wave {
  0%, 100% {
    transform: rotate(-20deg);
  }
  50% {
    transform: rotate(-10deg);
  }
}

.splash-content {
  animation: fadeInUp 0.8s ease-out;
}

.illustration {
  animation: fadeInUp 0.8s ease-out 0.2s both;
}

.brand-info {
  animation: fadeInUp 0.8s ease-out 0.4s both;
}
</style>