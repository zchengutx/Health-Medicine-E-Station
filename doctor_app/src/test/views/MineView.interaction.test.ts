import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import { createPinia, setActivePinia } from 'pinia'

// Mock AuthStore
vi.mock('@/stores/auth', () => ({
  useAuthStore: () => ({
    doctorName: '张医生',
    doctorInfo: {
      Title: '主任医师',
      Speciality: '心内科'
    },
    doctorAvatar: '/default-avatar.svg'
  })
}))

// Mock FeedbackButton component
vi.mock('@/components/FeedbackButton.vue', () => ({
  default: {
    name: 'FeedbackButton',
    template: '<button @click="$emit(\'click\')" :class="type"><slot>{{ text }}</slot></button>',
    props: ['text', 'type', 'block'],
    emits: ['click']
  }
}))

import MineView from '@/views/MineView.vue'

describe('MineView 交互逻辑测试', () => {
  let router: any

  beforeEach(() => {
    // 设置 Pinia
    setActivePinia(createPinia())
    
    // 创建路由
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/', component: { template: '<div>Home</div>' } },
        { path: '/profile', component: { template: '<div>Profile</div>' } },
        { path: '/setting', component: { template: '<div>Setting</div>' } },
        { path: '/patient', component: { template: '<div>Patient</div>' } }
      ]
    })
  })

  it('应该正确渲染医生信息', () => {
    const wrapper = mount(MineView, {
      global: {
        plugins: [router]
      }
    })

    expect(wrapper.find('.name').text()).toBe('张医生')
    expect(wrapper.find('.title').text()).toBe('主任医师')
    expect(wrapper.find('.hospital').text()).toBe('心内科')
  })

  it('医生头像区域不应该有点击事件', () => {
    const wrapper = mount(MineView, {
      global: {
        plugins: [router]
      }
    })

    const profileArea = wrapper.find('.mine-profile')
    expect(profileArea.exists()).toBe(true)
    
    // 检查是否没有点击事件监听器
    const profileElement = profileArea.element as HTMLElement
    expect(profileElement.onclick).toBeNull()
  })

  it('点击编辑个人信息按钮应该跳转到 /profile', async () => {
    const routerPushSpy = vi.spyOn(router, 'push')
    
    const wrapper = mount(MineView, {
      global: {
        plugins: [router]
      }
    })

    // 查找编辑个人信息按钮并点击
    const editButton = wrapper.find('button')
    expect(editButton.text()).toContain('编辑个人信息')
    
    await editButton.trigger('click')
    
    expect(routerPushSpy).toHaveBeenCalledWith('/profile')
  })

  it('点击设置图标应该跳转到 /setting', async () => {
    const routerPushSpy = vi.spyOn(router, 'push')
    
    const wrapper = mount(MineView, {
      global: {
        plugins: [router]
      }
    })

    const settingIcon = wrapper.find('.setting-icon')
    await settingIcon.trigger('click')
    
    expect(routerPushSpy).toHaveBeenCalledWith('/setting')
  })

  it('底部导航应该正常工作', async () => {
    const routerPushSpy = vi.spyOn(router, 'push')
    
    const wrapper = mount(MineView, {
      global: {
        plugins: [router]
      }
    })

    // 测试首页导航
    const homeTab = wrapper.findAll('.tabbar-item')[0]
    await homeTab.trigger('click')
    expect(routerPushSpy).toHaveBeenCalledWith('/')

    // 测试患者导航
    const patientTab = wrapper.findAll('.tabbar-item')[1]
    await patientTab.trigger('click')
    expect(routerPushSpy).toHaveBeenCalledWith('/patient')
  })
})