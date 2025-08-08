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

app.use(createPinia())
app.use(router)

app.mount('#app')