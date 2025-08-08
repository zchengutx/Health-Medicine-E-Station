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
    path: '/doctor-auth',
    name: 'DoctorAuth',
    component: () => import('@/views/DoctorAuthView.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/ProfileView.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFoundView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

// 导出导航工具
export { navigationUtils, useNavigationUtils } from './navigationUtils'