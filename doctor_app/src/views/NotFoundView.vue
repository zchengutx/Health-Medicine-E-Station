<template>
  <div class="not-found-view">
    <div class="not-found-content">
      <div class="error-icon">
        <van-icon name="warning-o" size="80" color="#CCCCCC" />
      </div>
      <h2 class="error-title">页面不存在</h2>
      <p class="error-message">抱歉，您访问的页面不存在或已被删除</p>
      <div class="error-actions">
        <van-button 
          type="primary" 
          size="large" 
          @click="goHome"
          class="action-button"
        >
          返回首页
        </van-button>
        <van-button 
          size="large" 
          @click="goBack"
          class="action-button"
        >
          返回上页
        </van-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useNavigationUtils } from '@/router'

const router = useRouter()
const navigationUtils = useNavigationUtils()

const goHome = () => {
  navigationUtils.toHome()
}

const goBack = () => {
  navigationUtils.goBack()
}
</script>

<style lang="scss" scoped>
.not-found-view {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: $bg-primary;
}

.not-found-content {
  text-align: center;
  padding: $spacing-xl;
}

.error-icon {
  margin-bottom: $spacing-lg;
}

.error-title {
  font-size: $font-size-xl;
  color: $text-primary;
  margin-bottom: $spacing-md;
  font-weight: 500;
}

.error-message {
  font-size: $font-size-md;
  color: $text-secondary;
  margin-bottom: $spacing-xl;
  line-height: 1.6;
}

.error-actions {
  display: flex;
  flex-direction: column;
  gap: $spacing-md;
  
  .action-button {
    width: 200px;
    margin: 0 auto;
  }
}
</style>