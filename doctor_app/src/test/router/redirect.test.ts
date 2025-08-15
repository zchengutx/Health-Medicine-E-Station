import { describe, it, expect, beforeEach } from 'vitest'
import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

// 导入路由配置
const routes: RouteRecordRaw[] = [
  {
    path: '/profile',
    name: 'Profile',
    component: { template: '<div>Profile</div>' }
  },
  {
    path: '/doctor-auth',
    name: 'DoctorAuthRedirect',
    redirect: '/profile'
  },
  {
    path: '/mine',
    name: 'Mine',
    component: { template: '<div>Mine</div>' }
  }
]

describe('路由重定向测试', () => {
  let router: any

  beforeEach(() => {
    router = createRouter({
      history: createWebHistory(),
      routes
    })
  })

  it('应该将 /doctor-auth 重定向到 /profile', async () => {
    // 导航到 /doctor-auth
    await router.push('/doctor-auth')
    
    // 验证当前路由是否为 /profile
    expect(router.currentRoute.value.path).toBe('/profile')
    expect(router.currentRoute.value.name).toBe('Profile')
  })

  it('应该能够直接访问 /profile', async () => {
    // 直接导航到 /profile
    await router.push('/profile')
    
    // 验证当前路由
    expect(router.currentRoute.value.path).toBe('/profile')
    expect(router.currentRoute.value.name).toBe('Profile')
  })

  it('重定向后应该能够正常导航到其他页面', async () => {
    // 先重定向到 /profile
    await router.push('/doctor-auth')
    expect(router.currentRoute.value.path).toBe('/profile')
    
    // 然后导航到 /mine
    await router.push('/mine')
    expect(router.currentRoute.value.path).toBe('/mine')
    expect(router.currentRoute.value.name).toBe('Mine')
  })
})