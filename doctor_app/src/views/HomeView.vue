<template>
  <div class="home-view">
    <!-- 错误提示 -->
    <div v-if="error" class="error-banner">
      <span>{{ error }}</span>
      <button @click="error = null" class="close-btn">×</button>
    </div>
    
    <!-- 加载状态 -->
    <div v-if="isLoading" class="loading-overlay">
      <div class="loading-spinner"></div>
      <p>加载中...</p>
    </div>
    <!-- 顶部医生信息区域 -->
    <div class="doctor-header">
      <div class="doctor-info">
        <div class="avatar-section">
          <div 
            class="doctor-avatar"
            :style="{ backgroundImage: authStore.doctorAvatar ? `url(${authStore.doctorAvatar})` : 'none' }"
          >
            <span v-if="!authStore.doctorAvatar" class="avatar-placeholder">
              {{ authStore.doctorName.charAt(0) }}
            </span>
          </div>
        </div>
        <div class="info-section">
          <div class="doctor-name">{{ authStore.doctorName }}</div>
          <div class="doctor-title" v-if="isVerified">{{ authStore.doctorInfo?.Title || '副主任医师' }}</div>
          <div class="hospital-name" v-if="isVerified">{{ authStore.doctorInfo?.Speciality || '北京积水潭医院' }}</div>
        </div>
      </div>
      <div class="verification-status">
        <van-tag v-if="isVerified" type="primary" size="medium">已认证</van-tag>
        <van-tag v-else type="warning" size="medium">未认证</van-tag>
      </div>
    </div>

    <!-- 统计数据区域 -->
    <div class="stats-section" v-if="isVerified">
      <div class="stat-item">
        <div class="stat-value">90.2%</div>
        <div class="stat-label">好评率</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">1401</div>
        <div class="stat-label">服务次数</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">55</div>
        <div class="stat-label">粉丝</div>
      </div>
    </div>

    <!-- 功能菜单区域 -->
    <div class="menu-section">
      <div class="menu-grid">
        <div class="menu-item" @click="goToConsultation">
          <div class="menu-icon consultation-icon">
            <van-icon name="chat-o" />
          </div>
          <div class="menu-label">我的问诊</div>
        </div>
        <div class="menu-item" @click="goToPatients">
          <div class="menu-icon patients-icon">
            <van-icon name="friends-o" />
          </div>
          <div class="menu-label">私人医生</div>
        </div>
        <div class="menu-item" @click="goToPrescription">
          <div class="menu-icon prescription-icon">
            <van-icon name="records" />
          </div>
          <div class="menu-label">我的药房</div>
        </div>
        <div class="menu-item" @click="goToCreation">
          <div class="menu-icon creation-icon">
            <van-icon name="edit" />
          </div>
          <div class="menu-label">创作中心</div>
        </div>
      </div>
    </div>

    <!-- 抢单中心 -->
    <div class="order-section" v-if="isVerified">
      <div class="section-header">
        <h3>抢单中心</h3>
        <van-button type="primary" size="mini" @click="goToAllOrders">全部</van-button>
      </div>
      <div class="order-card">
        <div class="patient-info">
          <div class="patient-avatar">
            <span class="avatar-placeholder">吴</span>
          </div>
          <div class="patient-details">
            <div class="patient-name">吴珊珊 <span class="patient-gender">女 23岁</span></div>
            <van-tag type="warning" size="small">极速问诊</van-tag>
          </div>
          <div class="order-price">¥29.00</div>
        </div>
        <div class="order-description">
          病情描述：孩子从出生后睡眠就不好，入睡困难，易醒，烦躁不安，近三个月症状加重，尤其夜里睡转难，情绪...
        </div>
        <div class="order-actions">
          <van-button type="primary" size="small" @click="acceptOrder">抢单</van-button>
        </div>
      </div>
    </div>

    <!-- 消息列表 -->
    <div class="message-section">
      <div class="section-header">
        <h3>消息列表</h3>
      </div>
      <div class="message-list">
        <div class="message-item" v-for="message in messages" :key="message.id">
          <div class="message-avatar">
            <div class="avatar-circle" :style="{ backgroundImage: message.avatar ? `url(${message.avatar})` : 'none' }">
              <span v-if="!message.avatar" class="avatar-placeholder">
                {{ message.name.charAt(0) }}
              </span>
            </div>
            <van-badge v-if="message.unread > 0" :content="message.unread" />
          </div>
          <div class="message-content">
            <div class="message-name">{{ message.name }}</div>
            <div class="message-text">{{ message.lastMessage }}</div>
          </div>
          <div class="message-meta">
            <div class="message-time">{{ message.time }}</div>
            <div class="message-status">{{ message.status }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部导航 -->
    <div class="bottom-nav">
      <div class="nav-item active">
        <van-icon name="home-o" />
        <span>首页</span>
      </div>
      <div class="nav-item" @click="goToPatients">
        <van-icon name="friends-o" />
        <span>患者</span>
      </div>
      <div class="nav-item" @click="goToProfile">
        <van-icon name="user-o" />
        <span>我的</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useNavigationUtils } from '@/router'
import { showToast } from 'vant'
import { log } from '@/utils/logger'

const authStore = useAuthStore()
const navigationUtils = useNavigationUtils()

// 添加加载和错误状态
const isLoading = ref(false)
const error = ref<string | null>(null)

// 认证状态判断 - 这里可以根据实际业务逻辑调整
// 假设当医生有LicenseNumber时表示已认证
const isVerified = computed(() => {
  try {
    return authStore.doctorInfo?.LicenseNumber && authStore.doctorInfo.LicenseNumber.trim() !== ''
  } catch (error) {
    log.error('检查认证状态失败', error)
    return false
  }
})

// 消息列表数据
const messages = ref([
  {
    id: 1,
    name: '吴珊珊',
    avatar: '',
    lastMessage: '头还是很痛，吃了药也不管用',
    time: '上午10:13',
    status: '进行中',
    unread: 2
  },
  {
    id: 2,
    name: '李明',
    avatar: '',
    lastMessage: '谢谢医生的建议，我会按时服药',
    time: '上午9:45',
    status: '已完成',
    unread: 0
  }
])

// 导航方法
const goToConsultation = () => {
  showToast('跳转到我的问诊')
}

const goToPatients = () => {
  showToast('跳转到私人医生')
}

const goToPrescription = () => {
  showToast('跳转到我的药房')
}

const goToCreation = () => {
  showToast('跳转到创作中心')
}

const goToAllOrders = () => {
  showToast('查看全部订单')
}

const goToProfile = () => {
  navigationUtils.safePush('/mine')
}

// 抢单操作
const acceptOrder = () => {
  showToast({
    message: '抢单成功！',
    type: 'success'
  })
}

onMounted(() => {
  try {
    log.debug('首页组件挂载，检查认证状态')
    log.debug('认证状态', {
      isLoggedIn: authStore.isLoggedIn,
      hasUserInfo: !!authStore.doctorInfo,
      hasToken: !!authStore.token
    })
    
    // 检查登录状态
    if (!authStore.isLoggedIn) {
      log.debug('用户未登录，重定向到登录页')
      navigationUtils.toLogin()
      return
    }
    
    // 检查token有效性
    if (!authStore.checkTokenExpiry()) {
      log.debug('Token已过期，重定向到登录页')
      showToast({
        message: '登录已过期，请重新登录',
        type: 'fail'
      })
      navigationUtils.toLogin()
      return
    }
    
    log.debug('认证检查通过，用户已认证')
  } catch (error) {
    log.error('首页组件挂载时发生错误', error)
    // 如果出现错误，重定向到登录页
    navigationUtils.toLogin()
  }
})
</script>

<style lang="scss" scoped>
.home-view {
  width: 100%;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding-bottom: 80px; // 为底部导航留空间
  position: relative;
}

// 错误提示样式
.error-banner {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: #ff4d4f;
  color: white;
  padding: 12px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 1000;
  
  .close-btn {
    background: none;
    border: none;
    color: white;
    font-size: 20px;
    cursor: pointer;
    padding: 0;
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}

// 加载状态样式
.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.9);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  z-index: 999;
  
  .loading-spinner {
    width: 40px;
    height: 40px;
    border: 4px solid #f3f3f3;
    border-top: 4px solid #667eea;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin-bottom: 16px;
  }
  
  p {
    color: #666;
    font-size: 14px;
  }
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

// 顶部医生信息区域
.doctor-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px 16px;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  
  .doctor-info {
    display: flex;
    align-items: center;
    flex: 1;
    
    .avatar-section {
      margin-right: 12px;
      
      .doctor-avatar {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        border: 3px solid rgba(255, 255, 255, 0.3);
        background-color: rgba(255, 255, 255, 0.2);
        background-size: cover;
        background-position: center;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .avatar-placeholder {
          color: white;
          font-size: 24px;
          font-weight: bold;
        }
      }
    }
    
    .info-section {
      flex: 1;
      
      .doctor-name {
        font-size: 20px;
        font-weight: bold;
        margin-bottom: 4px;
      }
      
      .doctor-title {
        font-size: 14px;
        opacity: 0.9;
        margin-bottom: 2px;
      }
      
      .hospital-name {
        font-size: 14px;
        opacity: 0.8;
      }
    }
  }
  
  .verification-status {
    margin-top: 8px;
  }
}

// 统计数据区域
.stats-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 0 16px 20px;
  display: flex;
  justify-content: space-around;
  
  .stat-item {
    text-align: center;
    color: white;
    
    .stat-value {
      font-size: 24px;
      font-weight: bold;
      margin-bottom: 4px;
    }
    
    .stat-label {
      font-size: 12px;
      opacity: 0.8;
    }
  }
}

// 功能菜单区域
.menu-section {
  background: white;
  margin: 0 16px;
  border-radius: 12px;
  padding: 20px;
  margin-top: -10px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  
  .menu-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 20px;
    
    .menu-item {
      text-align: center;
      cursor: pointer;
      transition: transform 0.2s;
      
      &:active {
        transform: scale(0.95);
      }
      
      .menu-icon {
        width: 48px;
        height: 48px;
        border-radius: 12px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin: 0 auto 8px;
        font-size: 24px;
        
        &.consultation-icon {
          background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
          color: white;
        }
        
        &.patients-icon {
          background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
          color: white;
        }
        
        &.prescription-icon {
          background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
          color: white;
        }
        
        &.creation-icon {
          background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%);
          color: #666;
        }
      }
      
      .menu-label {
        font-size: 12px;
        color: #666;
      }
    }
  }
}

// 抢单中心
.order-section {
  margin: 20px 16px;
  
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
    
    h3 {
      font-size: 16px;
      font-weight: bold;
      color: #333;
      margin: 0;
    }
  }
  
  .order-card {
    background: white;
    border-radius: 12px;
    padding: 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    border: 1px solid #f0f0f0;
    
    .patient-info {
      display: flex;
      align-items: center;
      margin-bottom: 12px;
      
      .patient-avatar {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        margin-right: 12px;
        background-color: #f0f0f0;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .avatar-placeholder {
          color: #666;
          font-size: 16px;
          font-weight: bold;
        }
      }
      
      .patient-details {
        flex: 1;
        
        .patient-name {
          font-size: 14px;
          font-weight: bold;
          margin-bottom: 4px;
          
          .patient-gender {
            font-weight: normal;
            color: #666;
          }
        }
      }
      
      .order-price {
        font-size: 18px;
        font-weight: bold;
        color: #ff6b6b;
      }
    }
    
    .order-description {
      font-size: 13px;
      color: #666;
      line-height: 1.4;
      margin-bottom: 12px;
    }
    
    .order-actions {
      text-align: right;
    }
  }
}

// 消息列表
.message-section {
  margin: 20px 16px;
  
  .section-header {
    margin-bottom: 12px;
    
    h3 {
      font-size: 16px;
      font-weight: bold;
      color: #333;
      margin: 0;
    }
  }
  
  .message-list {
    background: white;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    
    .message-item {
      display: flex;
      align-items: center;
      padding: 12px 16px;
      border-bottom: 1px solid #f0f0f0;
      
      &:last-child {
        border-bottom: none;
      }
      
      .message-avatar {
        position: relative;
        margin-right: 12px;
        
        .avatar-circle {
          width: 40px;
          height: 40px;
          border-radius: 50%;
          background-color: #f0f0f0;
          background-size: cover;
          background-position: center;
          display: flex;
          align-items: center;
          justify-content: center;
          
          .avatar-placeholder {
            color: #666;
            font-size: 16px;
            font-weight: bold;
          }
        }
      }
      
      .message-content {
        flex: 1;
        
        .message-name {
          font-size: 14px;
          font-weight: bold;
          margin-bottom: 4px;
        }
        
        .message-text {
          font-size: 12px;
          color: #666;
        }
      }
      
      .message-meta {
        text-align: right;
        
        .message-time {
          font-size: 11px;
          color: #999;
          margin-bottom: 4px;
        }
        
        .message-status {
          font-size: 11px;
          color: #4CAF50;
        }
      }
    }
  }
}

// 底部导航
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: white;
  display: flex;
  border-top: 1px solid #f0f0f0;
  padding: 8px 0;
  
  .nav-item {
    flex: 1;
    text-align: center;
    padding: 8px;
    cursor: pointer;
    transition: color 0.2s;
    
    .van-icon {
      font-size: 20px;
      margin-bottom: 4px;
      display: block;
    }
    
    span {
      font-size: 10px;
      display: block;
    }
    
    &.active {
      color: #667eea;
    }
    
    &:not(.active) {
      color: #999;
    }
  }
}
</style>