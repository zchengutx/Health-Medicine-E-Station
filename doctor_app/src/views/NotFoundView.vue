<template>
  <div class="not-found-view">
    <div class="not-found-content">
      <div class="error-icon">
        ⚠️
      </div>
      <h2 class="error-title">页面不存在</h2>
      <p class="error-message">抱歉，您访问的页面不存在或已被删除</p>
      <div class="error-actions">
        <button 
          @click="goHome"
          class="action-button primary"
        >
          返回首页
        </button>
        <button 
          @click="goBack"
          class="action-button"
        >
          返回上页
        </button>
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
  background-color: #f8f9fa;
}

.not-found-content {
  text-align: center;
  padding: 32px;
}

.error-icon {
  margin-bottom: 24px;
  font-size: 64px;
}

.error-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 16px;
  font-weight: 500;
}

.error-message {
  font-size: 16px;
  color: #666;
  margin-bottom: 32px;
  line-height: 1.6;
}

.error-actions {
  display: flex;
  flex-direction: column;
  gap: 16px;
  
  .action-button {
    width: 200px;
    margin: 0 auto;
    padding: 12px 24px;
    border: none;
    border-radius: 6px;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s;
    background-color: #f0f0f0;
    color: #333;
    
    &.primary {
      background-color: #007AFF;
      color: white;
    }
    
    &:hover {
      opacity: 0.8;
    }
  }
}
</style>