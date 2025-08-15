import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import { useAuthStore } from '@/stores/auth'
import { log } from '@/utils/logger'

// 引入Vant样式
import 'vant/lib/index.css'
// 引入全局样式
import '@/assets/styles/index.scss'

// 移动端适配
import '@vant/touch-emulator'

async function initializeApp() {
  try {
    const app = createApp(App)

    // 添加全局错误处理
    app.config.errorHandler = (err, _instance, info) => {
      console.error('Vue应用错误:', err, info)
    }

    // 先安装Pinia
    app.use(createPinia())

    // 初始化认证状态
    try {
      const authStore = useAuthStore()
      authStore.initAuth()

      // 等待认证状态初始化完成
      await authStore.waitForInitialization()
    } catch (error) {
      console.error('应用认证状态初始化失败:', error)
      // 即使初始化失败，也继续启动应用
    }

    app.use(router)

    // 确保DOM元素存在
    const appElement = document.getElementById('app')
    if (!appElement) {
      throw new Error('找不到应用挂载点 #app')
    }

    app.mount('#app')
    console.log('应用启动成功')
  } catch (error) {
    console.error('应用初始化失败:', error)
    throw error
  }
}

// 启动应用
initializeApp().catch(error => {
  console.error('应用启动失败:', error)

  // 如果初始化失败，尝试基本启动
  try {
    const app = createApp(App)
    app.use(createPinia())
    app.use(router)
    app.mount('#app')
    console.log('应用基本启动成功')
  } catch (fallbackError) {
    console.error('应用基本启动也失败:', fallbackError)
    document.body.innerHTML = '<div style="text-align:center;padding:50px;">应用启动失败，请刷新页面重试</div>'
  }
})