import { createApp, App as VueApp } from 'vue'
import ToastMessage, { type ToastType } from '@/components/ToastMessage.vue'

interface ToastOptions {
  message: string
  type?: ToastType
  duration?: number
  showIcon?: boolean
  position?: 'top' | 'center' | 'bottom'
}

class ToastService {
  private toastInstances: VueApp[] = []

  show(options: ToastOptions | string) {
    const config = typeof options === 'string' 
      ? { message: options } 
      : options

    const container = document.createElement('div')
    document.body.appendChild(container)

    const app = createApp(ToastMessage, {
      ...config,
      onClose: () => {
        this.destroy(app, container)
      }
    })

    app.mount(container)
    this.toastInstances.push(app)

    return app
  }

  success(message: string, duration?: number) {
    return this.show({
      message,
      type: 'success',
      duration
    })
  }

  error(message: string, duration?: number) {
    return this.show({
      message,
      type: 'error',
      duration: duration || 4000
    })
  }

  warning(message: string, duration?: number) {
    return this.show({
      message,
      type: 'warning',
      duration
    })
  }

  info(message: string, duration?: number) {
    return this.show({
      message,
      type: 'info',
      duration
    })
  }

  loading(message: string = '加载中...') {
    return this.show({
      message,
      type: 'loading',
      duration: 0 // 不自动关闭
    })
  }

  private destroy(app: VueApp, container: HTMLElement) {
    app.unmount()
    document.body.removeChild(container)
    
    const index = this.toastInstances.indexOf(app)
    if (index > -1) {
      this.toastInstances.splice(index, 1)
    }
  }

  clear() {
    this.toastInstances.forEach(app => {
      app.unmount()
    })
    this.toastInstances = []
    
    // 清理所有toast容器
    const toastContainers = document.querySelectorAll('[data-toast-container]')
    toastContainers.forEach(container => {
      document.body.removeChild(container)
    })
  }
}

export const toast = new ToastService()

// 便捷方法
export const showToast = (options: ToastOptions | string) => toast.show(options)
export const showSuccess = (message: string, duration?: number) => toast.success(message, duration)
export const showError = (message: string, duration?: number) => toast.error(message, duration)
export const showWarning = (message: string, duration?: number) => toast.warning(message, duration)
export const showInfo = (message: string, duration?: number) => toast.info(message, duration)
export const showLoading = (message?: string) => toast.loading(message)
export const clearToast = () => toast.clear()