import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import CountdownTimer from '@/components/CountdownTimer.vue'

describe('CountdownTimer', () => {
  it('should render with text', () => {
    const wrapper = mount(CountdownTimer, {
      props: {
        text: '跳过 3'
      }
    })

    expect(wrapper.text()).toContain('跳过 3')
    expect(wrapper.find('.countdown-timer').exists()).toBe(true)
  })

  it('should apply urgent class when countdown is 1', () => {
    const wrapper = mount(CountdownTimer, {
      props: {
        text: '跳过 1',
        countdown: 1
      }
    })

    expect(wrapper.find('.countdown-urgent').exists()).toBe(true)
  })

  it('should not apply urgent class when countdown is greater than 1', () => {
    const wrapper = mount(CountdownTimer, {
      props: {
        text: '跳过 3',
        countdown: 3
      }
    })

    expect(wrapper.find('.countdown-urgent').exists()).toBe(false)
  })

  it('should handle missing countdown prop', () => {
    const wrapper = mount(CountdownTimer, {
      props: {
        text: '跳过'
      }
    })

    expect(wrapper.find('.countdown-urgent').exists()).toBe(false)
  })
})