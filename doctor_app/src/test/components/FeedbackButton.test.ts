import { describe, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import FeedbackButton from '@/components/FeedbackButton.vue'

describe('FeedbackButton', () => {
  it('should render with text', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        text: '点击按钮'
      }
    })

    expect(wrapper.text()).toContain('点击按钮')
    expect(wrapper.find('.feedback-button').exists()).toBe(true)
  })

  it('should render with slot content', () => {
    const wrapper = mount(FeedbackButton, {
      slots: {
        default: '插槽内容'
      }
    })

    expect(wrapper.text()).toContain('插槽内容')
  })

  it('should apply correct type class', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        type: 'success'
      }
    })

    expect(wrapper.find('.button-success').exists()).toBe(true)
  })

  it('should apply correct size class', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        size: 'large'
      }
    })

    expect(wrapper.find('.button-large').exists()).toBe(true)
  })

  it('should be disabled when disabled prop is true', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        disabled: true
      }
    })

    expect(wrapper.find('.button-disabled').exists()).toBe(true)
    expect(wrapper.element.disabled).toBe(true)
  })

  it('should show loading spinner when loading', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        loading: true,
        text: '加载中'
      }
    })

    expect(wrapper.find('.button-loading').exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'LoadingSpinner' }).exists()).toBe(true)
  })

  it('should emit click event when clicked', async () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        text: '点击'
      }
    })

    await wrapper.trigger('click')
    expect(wrapper.emitted('click')).toBeTruthy()
  })

  it('should not emit click when disabled', async () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        text: '点击',
        disabled: true
      }
    })

    await wrapper.trigger('click')
    expect(wrapper.emitted('click')).toBeFalsy()
  })

  it('should not emit click when loading', async () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        text: '点击',
        loading: true
      }
    })

    await wrapper.trigger('click')
    expect(wrapper.emitted('click')).toBeFalsy()
  })

  it('should apply block class when block is true', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        block: true
      }
    })

    expect(wrapper.find('.button-block').exists()).toBe(true)
  })

  it('should apply round class when round is true', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        round: true
      }
    })

    expect(wrapper.find('.button-round').exists()).toBe(true)
  })

  it('should show icon when icon prop is provided', () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        icon: 'success',
        text: '成功'
      }
    })

    expect(wrapper.find('.button-icon').exists()).toBe(true)
  })

  it('should handle touch events', async () => {
    const wrapper = mount(FeedbackButton, {
      props: {
        text: '触摸'
      }
    })

    await wrapper.trigger('touchstart')
    expect(wrapper.find('.button-pressed').exists()).toBe(true)

    await wrapper.trigger('touchend')
    expect(wrapper.find('.button-pressed').exists()).toBe(false)
  })
})