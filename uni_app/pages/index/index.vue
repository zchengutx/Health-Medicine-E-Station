<template>
  <view class="index">
    <view class="header">
      <text class="title">健康医疗驿站</text>
      <text class="subtitle">您的健康管理专家</text>
      <button class="login-button" @click="goToLogin">登录</button>
    </view>
    
    <view class="content">
      <view class="welcome-card">
        <text class="welcome-text">欢迎使用健康医疗驿站</text>
        <text class="welcome-desc">为您提供全方位的健康管理服务</text>
      </view>
      
      <view class="feature-grid">
        <view class="feature-item" @click="navigateToFeature('health')">
          <text class="feature-icon">🏥</text>
          <text class="feature-title">健康咨询</text>
        </view>
        <view class="feature-item" @click="navigateToFeature('medicine')">
          <text class="feature-icon">💊</text>
          <text class="feature-title">药品管理</text>
        </view>
        <view class="feature-item" @click="navigateToFeature('appointment')">
          <text class="feature-icon">📅</text>
          <text class="feature-title">预约挂号</text>
        </view>
        <view class="feature-item" @click="navigateToFeature('record')">
          <text class="feature-icon">📋</text>
          <text class="feature-title">健康档案</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'Index',
  data() {
    return {}
  },
  methods: {
    navigateToFeature(feature) {
      uni.showToast({
        title: `功能开发中: ${feature}`,
        icon: 'none'
      })
    },
    goToLogin() {
      uni.navigateTo({
        url: '/pages/login/login'
      })
    }
  }
}
</script>

<style scoped>
.index {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.header {
  padding: 60rpx 40rpx 40rpx;
  text-align: center;
  position: relative;
}

.title {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: #ffffff;
  margin-bottom: 20rpx;
}

.subtitle {
  display: block;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.8);
}

.login-button {
  position: absolute;
  top: 40rpx;
  right: 40rpx;
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 20rpx;
  padding: 10rpx 20rpx;
  font-size: 24rpx;
}

.content {
  padding: 0 40rpx;
}

.welcome-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20rpx;
  padding: 40rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.1);
}

.welcome-text {
  display: block;
  font-size: 36rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 16rpx;
}

.welcome-desc {
  display: block;
  font-size: 28rpx;
  color: #666;
}

.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30rpx;
}

.feature-item {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20rpx;
  padding: 40rpx;
  text-align: center;
  box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.feature-item:active {
  transform: scale(0.95);
}

.feature-icon {
  display: block;
  font-size: 60rpx;
  margin-bottom: 20rpx;
}

.feature-title {
  display: block;
  font-size: 28rpx;
  color: #333;
  font-weight: 500;
}
</style>
