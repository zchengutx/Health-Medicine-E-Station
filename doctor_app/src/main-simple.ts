import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'

// 引入Vant样式
import 'vant/lib/index.css'
// 引入全局样式
import '@/assets/styles/index.scss'

// 移动端适配
import '@vant/touch-emulator'

const app = createApp(App)

// 添加全局错误处理
app.config.errorHandler = (err, instance, info) => {
  console.error('Vue应用错误:', err, info)
}

// 安装插件
app.use(createPinia())
app.use(router)

// 挂载应用
app.mount('#app')

console.log('应用启动成功')