import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createTestingPinia } from '@pinia/testing'
import SplashView from '@/views/SplashView.vue'

// Mock router
const mockPush = vi.fn()
vi.mock('vue-router', () => ({
  useRouter: () => ({
    push: mockPush
  })
}))

// Mock navigation utils
vi.mock('@/router', () => ({
  navigationUtils: {
    toHome: vi.fn(),
    safePush: vi.fn()
  }
}))

describe('SplashView', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.useFakeTimers()
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('should render splash content', () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    expect(wrapper.find('.splash-view').exists()).toBe(true)
    expect(wrapper.find('.brand-info').exists()).toBe(true)
    expect(wrapper.text()).toContain('优医')
    expect(wrapper.text()).toContain('您身边的健康管理专家')
  })

  it('should show countdown timer', () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    expect(wrapper.find('.skip-button').exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'CountdownTimer' }).exists()).toBe(true)
  })

  it('should start countdown on mount', () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    expect(wrapper.vm.countdown).toBe(3)
  })

  it('should skip splash when skip button is clicked', async () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    await wrapper.find('.skip-button').trigger('click')
    expect(wrapper.vm.countdown).toBe(0)
  })

  it('should show illustration elements', () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    expect(wrapper.find('.illustration').exists()).toBe(true)
    expect(wrapper.find('.main-illustration').exists()).toBe(true)
    expect(wrapper.find('.patient').exists()).toBe(true)
    expect(wrapper.find('.phone-screen').exists()).toBe(true)
    expect(wrapper.find('.plant').exists()).toBe(true)
  })

  it('should have proper animation classes', () => {
    const wrapper = mount(SplashView, {
      global: {
        plugins: [createTestingPinia({
          createSpy: vi.fn
        })]
      }
    })

    const content = wrapper.find('.splash-content')
    expect(content.exists()).toBe(true)
  })
})