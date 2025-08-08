import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import LoadingSpinner from '@/components/LoadingSpinner.vue'

describe('LoadingSpinner', () => {
  it('should render with default props', () => {
    const wrapper = mount(LoadingSpinner)

    expect(wrapper.find('.loading-spinner').exists()).toBe(true)
    expect(wrapper.find('.spinner').exists()).toBe(true)
    expect(wrapper.find('.spinner-medium').exists()).toBe(true)
  })

  it('should render with text', () => {
    const wrapper = mount(LoadingSpinner, {
      props: {
        text: '加载中...'
      }
    })

    expect(wrapper.text()).toContain('加载中...')
    expect(wrapper.find('.loading-text').exists()).toBe(true)
  })

  it('should apply correct size class', () => {
    const wrapper = mount(LoadingSpinner, {
      props: {
        size: 'large'
      }
    })

    expect(wrapper.find('.spinner-large').exists()).toBe(true)
  })

  it('should apply overlay class when overlay is true', () => {
    const wrapper = mount(LoadingSpinner, {
      props: {
        overlay: true
      }
    })

    expect(wrapper.find('.loading-overlay').exists()).toBe(true)
  })

  it('should apply custom color', () => {
    const wrapper = mount(LoadingSpinner, {
      props: {
        color: '#ff0000'
      }
    })

    const spinner = wrapper.find('.spinner')
    expect(spinner.attributes('style')).toContain('border-top-color: rgb(255, 0, 0)')
  })

  it('should not render text when not provided', () => {
    const wrapper = mount(LoadingSpinner)

    expect(wrapper.find('.loading-text').exists()).toBe(false)
  })
})