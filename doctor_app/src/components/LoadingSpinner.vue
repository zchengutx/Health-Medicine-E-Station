<template>
  <div class="loading-spinner" :class="{ 'loading-overlay': overlay }">
    <div class="spinner-container">
      <div class="spinner" :class="sizeClass" :style="spinnerStyle"></div>
      <p v-if="text" class="loading-text">{{ text }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  text?: string
  size?: 'small' | 'medium' | 'large'
  color?: string
  overlay?: boolean
  backgroundColor?: string
}

const props = withDefaults(defineProps<Props>(), {
  size: 'medium',
  color: '#007AFF',
  overlay: false,
  backgroundColor: 'rgba(255, 255, 255, 0.9)'
})

const sizeClass = computed(() => `spinner-${props.size}`)

const spinnerStyle = computed(() => ({
  borderTopColor: props.color,
  borderRightColor: props.color
}))
</script>

<style lang="scss" scoped>
.loading-spinner {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  
  &.loading-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: v-bind(backgroundColor);
    z-index: 9999;
    backdrop-filter: blur(2px);
  }
}

.spinner-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: $spacing-md;
}

.spinner {
  border: 3px solid rgba(0, 0, 0, 0.1);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  
  &.spinner-small {
    width: 20px;
    height: 20px;
    border-width: 2px;
  }
  
  &.spinner-medium {
    width: 32px;
    height: 32px;
    border-width: 3px;
  }
  
  &.spinner-large {
    width: 48px;
    height: 48px;
    border-width: 4px;
  }
}

.loading-text {
  color: $text-secondary;
  font-size: $font-size-sm;
  font-weight: $font-weight-medium;
  text-align: center;
  margin: 0;
}

@keyframes spin {
  0% { 
    transform: rotate(0deg); 
  }
  100% { 
    transform: rotate(360deg); 
  }
}

// 脉冲动画变体
.spinner-pulse {
  animation: pulse 1.5s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.5;
    transform: scale(1.1);
  }
}

// 弹跳动画变体
.spinner-bounce {
  animation: bounce 1.4s ease-in-out infinite both;
}

@keyframes bounce {
  0%, 80%, 100% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
}
</style>