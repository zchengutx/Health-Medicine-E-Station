import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Splash',
    component: () => import('@/views/SplashView.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/LoginView.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/RegisterView.vue')
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/HomeView.vue')
  },
  {
    path: '/mine',
    name: 'Mine',
    component: () => import('@/views/MineView.vue')
  },
  {
    path: '/setting',
    name: 'Setting',
    component: () => import('@/views/SettingView.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/ProfileView.vue')
  },
  {
    path: '/doctor-auth',
    name: 'DoctorAuthRedirect',
    redirect: '/profile'
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/views/NotFoundView.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 添加全局路由守卫
router.beforeEach((to, from, next) => {
  try {
    // 记录路由重定向
    if (from.path === '/doctor-auth' || to.path === '/doctor-auth') {
      console.log('路由重定向:', {
        from: from.path,
        to: to.path,
        name: to.name,
        redirected: to.path === '/profile' && from.path === '/doctor-auth'
      })
    }
    
    // 基本的路由验证
    if (to.matched.length === 0) {
      // 如果没有匹配的路由，重定向到404页面
      next('/404')
      return
    }
    
    next()
  } catch (error) {
    console.error('Router navigation error:', error)
    next('/')
  }
})

// 添加路由错误处理
router.onError((error) => {
  console.error('Router error:', error)
})

export default router

// 导出导航工具
export { navigationUtils, useNavigationUtils } from './navigationUtils'