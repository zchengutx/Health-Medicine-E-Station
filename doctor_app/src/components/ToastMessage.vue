<template>
  <Teleport to="body">
    <Transition name="toast" appear>
      <div v-if="visible" class="toast-message" :class="typeClass">
        <div class="toast-content">
          <div class="toast-icon" v-if="showIcon">
            <van-icon :name="iconName" />
          </div>
          <div class="toast-text">{{ message }}</div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

export type ToastType = 'success' | 'error' | 'warning' | 'info' | 'loading'

interface Props {
  message: string
  type?: ToastType
  duration?: number
  showIcon?: boolean
  position?: 'top' | 'center' | 'bottom'
}

const props = withDefaults(defineProps<Props>(), {
  type: 'info',
  duration: 3000,
  showIcon: true,
  position: 'center'
})

const emit = defineEmits<{
  close: []
}>()

const visible = ref(false)

const typeClass = computed(() => [
  `toast-${props.type}`,
  `toast-${props.position}`
])

const iconName = computed(() => {
  const iconMap: Record<ToastType, string> = {
    success: 'success',
    error: 'close',
    warning: 'warning-o',
    info: 'info-o',
    loading: 'loading'
  }
  return iconMap[props.type]
})

const show = () => {
  visible.value = true
  
  if (props.type !== 'loading' && props.duration > 0) {
    setTimeout(() => {
      hide()
    }, props.duration)
  }
}

const hide = () => {
  visible.value = false
  setTimeout(() => {
    emit('close')
  }, 300)
}

onMounted(() => {
  show()
})

defineExpose({
  show,
  hide
})
</script>

<style lang="scss" scoped>
.toast-message {
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  z-index: 9999;
  pointer-events: none;
  
  &.toast-top {
    top: 20%;
  }
  
  &.toast-center {
    top: 50%;
    transform: translate(-50%, -50%);
  }
  
  &.toast-bottom {
    bottom: 20%;
  }
}

.toast-content {
  display: flex;
  align-items: center;
  gap: $spacing-sm;
  padding: $spacing-md $spacing-lg;
  border-radius: $border-radius-lg;
  box-shadow: $shadow-lg;
  backdrop-filter: blur(10px);
  max-width: 300px;
  min-width: 120px;
}

.toast-icon {
  font-size: 18px;
  flex-shrink: 0;
}

.toast-text {
  font-size: $font-size-sm;
  font-weight: $font-weight-medium;
  line-height: 1.4;
  word-break: break-word;
}

// 不同类型的样式
.toast-success {
  .toast-content {
    background: rgba($success-color, 0.95);
    color: $text-white;
  }
}

.toast-error {
  .toast-content {
    background: rgba($error-color, 0.95);
    color: $text-white;
  }
}

.toast-warning {
  .toast-content {
    background: rgba($warning-color, 0.95);
    color: $text-white;
  }
}

.toast-info {
  .toast-content {
    background: rgba($info-color, 0.95);
    color: $text-white;
  }
}

.toast-loading {
  .toast-content {
    background: rgba($text-primary, 0.9);
    color: $text-white;
  }
  
  .toast-icon {
    animation: spin 1s linear infinite;
  }
}

// 过渡动画
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(-50%) translateY(-20px) scale(0.9);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(-50%) translateY(20px) scale(0.9);
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>