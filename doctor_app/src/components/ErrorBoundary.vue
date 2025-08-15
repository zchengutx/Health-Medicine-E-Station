<template>
  <div v-if="hasError" class="error-boundary">
    <div class="error-content">
      <div class="error-icon">⚠️</div>
      <h2 class="error-title">出现了一些问题</h2>
      <p class="error-message">{{ errorMessage }}</p>
      <div class="error-actions">
        <button @click="retry" class="retry-button">重试</button>
        <button @click="goHome" class="home-button">返回首页</button>
      </div>
    </div>
  </div>
  <slot v-else />
</template>

<script setup lang="ts">
import { ref, onErrorCaptured } from 'vue'
import { log } from '@/utils/logger'

const hasError = ref(false)
const errorMessage = ref('应用遇到了意外错误，请稍后重试')

// 捕获子组件错误
onErrorCaptured((error, instance, info) => {
  log.error('组件错误被捕获', { error, info })
  
  hasError.value = true
  errorMessage.value = '页面加载失败，请重试'
  
  // 返回false阻止错误继续传播
  return false
})

const retry = () => {
  hasError.value = false
  // 重新加载页面
  window.location.reload()
}

const goHome = () => {
  hasError.value = false
  window.location.href = '/home'
}
</script>

<style scoped>
.error-boundary {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f8f9fa;
}

.error-content {
  text-align: center;
  padding: 32px;
  max-width: 400px;
}

.error-icon {
  font-size: 64px;
  margin-bottom: 16px;
}

.error-title {
  font-size: 20px;
  color: #333;
  margin-bottom: 12px;
  font-weight: 500;
}

.error-message {
  font-size: 14px;
  color: #666;
  margin-bottom: 24px;
  line-height: 1.5;
}

.error-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.retry-button,
.home-button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.retry-button {
  background-color: #007AFF;
  color: white;
}

.retry-button:hover {
  background-color: #0056CC;
}

.home-button {
  background-color: #f0f0f0;
  color: #333;
}

.home-button:hover {
  background-color: #e0e0e0;
}
</style>