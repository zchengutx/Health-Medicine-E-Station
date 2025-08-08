<template>
  <button
    :class="buttonClass"
    :disabled="disabled || loading"
    @click="handleClick"
    @touchstart="handleTouchStart"
    @touchend="handleTouchEnd"
    @mousedown="handleMouseDown"
    @mouseup="handleMouseUp"
    @mouseleave="handleMouseLeave"
  >
    <div class="button-content">
      <LoadingSpinner 
        v-if="loading" 
        size="small" 
        :color="loadingColor"
        class="button-loading"
      />
      <van-icon 
        v-else-if="icon" 
        :name="icon" 
        class="button-icon"
      />
      <span v-if="$slots.default || text" class="button-text">
        <slot>{{ text }}</slot>
      </span>
    </div>
    
    <!-- 点击波纹效果 -->
    <div 
      v-if="showRipple && rippleStyle" 
      class="button-ripple"
      :style="rippleStyle"
    ></div>
  </button>
</template>

<script setup lang="ts">
import { computed, ref, nextTick } from 'vue'
import LoadingSpinner from './LoadingSpinner.vue'

interface Props {
  text?: string
  type?: 'primary' | 'secondary' | 'success' | 'warning' | 'error' | 'outline'
  size?: 'small' | 'medium' | 'large'
  icon?: string
  loading?: boolean
  disabled?: boolean
  block?: boolean
  round?: boolean
  hapticFeedback?: boolean
  rippleEffect?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  type: 'primary',
  size: 'medium',
  loading: false,
  disabled: false,
  block: false,
  round: false,
  hapticFeedback: true,
  rippleEffect: true
})

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const isPressed = ref(false)
const showRipple = ref(false)
const rippleStyle = ref<any>(null)

const buttonClass = computed(() => [
  'feedback-button',
  `button-${props.type}`,
  `button-${props.size}`,
  {
    'button-block': props.block,
    'button-round': props.round,
    'button-loading': props.loading,
    'button-disabled': props.disabled,
    'button-pressed': isPressed.value
  }
])

const loadingColor = computed(() => {
  if (props.type === 'outline') return '#007AFF'
  return '#FFFFFF'
})

const handleClick = async (event: MouseEvent) => {
  console.log('FeedbackButton被点击了！', {
    disabled: props.disabled,
    loading: props.loading,
    text: props.text
  })
  
  if (props.disabled || props.loading) {
    console.log('FeedbackButton点击被阻止，原因:', {
      disabled: props.disabled,
      loading: props.loading
    })
    return
  }

  // 触觉反馈
  if (props.hapticFeedback && 'vibrate' in navigator) {
    navigator.vibrate(10)
  }

  // 波纹效果
  if (props.rippleEffect) {
    await createRipple(event)
  }

  console.log('FeedbackButton即将触发click事件')
  emit('click', event)
  console.log('FeedbackButton已触发click事件')
}

const handleTouchStart = () => {
  if (props.disabled || props.loading) return
  isPressed.value = true
}

const handleTouchEnd = () => {
  isPressed.value = false
}

const handleMouseDown = () => {
  if (props.disabled || props.loading) return
  isPressed.value = true
}

const handleMouseUp = () => {
  isPressed.value = false
}

const handleMouseLeave = () => {
  isPressed.value = false
}

const createRipple = async (event: MouseEvent) => {
  const button = event.currentTarget as HTMLElement
  const rect = button.getBoundingClientRect()
  const size = Math.max(rect.width, rect.height)
  const x = event.clientX - rect.left - size / 2
  const y = event.clientY - rect.top - size / 2

  rippleStyle.value = {
    width: `${size}px`,
    height: `${size}px`,
    left: `${x}px`,
    top: `${y}px`
  }

  showRipple.value = true

  await nextTick()

  setTimeout(() => {
    showRipple.value = false
    rippleStyle.value = null
  }, 600)
}
</script>

<style lang="scss" scoped>
.feedback-button {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
  border-radius: $border-radius;
  font-weight: $font-weight-medium;
  cursor: pointer;
  transition: all 0.2s ease;
  user-select: none;
  touch-action: manipulation;
  overflow: hidden;
  
  &:focus {
    outline: none;
  }
  
  // 尺寸变体
  &.button-small {
    height: $button-height-sm;
    padding: 0 $spacing-md;
    font-size: $font-size-sm;
  }
  
  &.button-medium {
    height: $button-height;
    padding: 0 $spacing-lg;
    font-size: $font-size-md;
  }
  
  &.button-large {
    height: $button-height-lg;
    padding: 0 $spacing-xl;
    font-size: $font-size-lg;
  }
  
  // 类型变体
  &.button-primary {
    background: $primary-color;
    color: $text-white;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $primary-dark;
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba($primary-color, 0.3);
    }
  }
  
  &.button-secondary {
    background: $gray-200;
    color: $text-primary;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $gray-300;
      transform: translateY(-1px);
    }
  }
  
  &.button-success {
    background: $success-color;
    color: $text-white;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $success-dark;
      transform: translateY(-1px);
    }
  }
  
  &.button-warning {
    background: $warning-color;
    color: $text-white;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $warning-dark;
      transform: translateY(-1px);
    }
  }
  
  &.button-error {
    background: $error-color;
    color: $text-white;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $error-dark;
      transform: translateY(-1px);
    }
  }
  
  &.button-outline {
    background: transparent;
    color: $primary-color;
    border: 1px solid $primary-color;
    
    &:hover:not(.button-disabled):not(.button-loading) {
      background: $primary-color;
      color: $text-white;
      transform: translateY(-1px);
    }
  }
  
  // 状态变体
  &.button-block {
    width: 100%;
  }
  
  &.button-round {
    border-radius: $border-radius-full;
  }
  
  &.button-pressed {
    transform: scale(0.98);
  }
  
  &.button-loading {
    cursor: not-allowed;
    
    .button-text {
      opacity: 0.7;
    }
  }
  
  &.button-disabled {
    background: $gray-300 !important;
    color: $gray-500 !important;
    cursor: not-allowed !important;
    transform: none !important;
    box-shadow: none !important;
    border-color: $gray-300 !important;
  }
}

.button-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: $spacing-sm;
  position: relative;
  z-index: 1;
}

.button-loading {
  flex-shrink: 0;
}

.button-icon {
  flex-shrink: 0;
  font-size: 1.2em;
}

.button-text {
  white-space: nowrap;
  transition: opacity 0.2s ease;
}

.button-ripple {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  transform: scale(0);
  animation: ripple 0.6s ease-out;
  pointer-events: none;
}

@keyframes ripple {
  to {
    transform: scale(2);
    opacity: 0;
  }
}

// 移动端优化
@media (hover: none) {
  .feedback-button:hover {
    transform: none !important;
    box-shadow: none !important;
  }
}
</style>