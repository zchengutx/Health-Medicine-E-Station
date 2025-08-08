<template>
  <div class="countdown-timer" :class="{ 'countdown-urgent': isUrgent }">
    <span class="countdown-text">{{ text }}</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  text: string
  countdown?: number
}

const props = defineProps<Props>()

// 当倒计时小于等于1秒时显示紧急状态
const isUrgent = computed(() => {
  if (props.countdown !== undefined) {
    return props.countdown <= 1
  }
  return false
})
</script>

<style lang="scss" scoped>
.countdown-timer {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 8px 16px;
  background: rgba(0, 0, 0, 0.4);
  border-radius: 20px;
  color: white;
  font-size: 14px;
  font-weight: 500;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  transition: all 0.3s ease;
  user-select: none;
  
  &:hover {
    background: rgba(0, 0, 0, 0.5);
    transform: scale(1.05);
  }
  
  &:active {
    transform: scale(0.95);
  }
  
  &.countdown-urgent {
    background: rgba(255, 77, 79, 0.8);
    animation: pulse 0.5s ease-in-out infinite alternate;
  }
  
  .countdown-text {
    white-space: nowrap;
    letter-spacing: 0.5px;
  }
}

@keyframes pulse {
  from {
    transform: scale(1);
    opacity: 0.8;
  }
  to {
    transform: scale(1.1);
    opacity: 1;
  }
}
</style>