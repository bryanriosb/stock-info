import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Logo from './Logo.vue'
import { TrendingUp } from 'lucide-vue-next'

describe('Logo', () => {
  it('renders with default props', () => {
    const wrapper = mount(Logo)
    expect(wrapper.find('div').exists()).toBe(true)
  })

  it('applies default classes', () => {
    const wrapper = mount(Logo)
    const div = wrapper.find('div')
    expect(div.classes()).toContain('gradient-coral')
    expect(div.classes()).toContain('rounded-xl')
  })

  it('applies custom class when provided', () => {
    const wrapper = mount(Logo, {
      props: { customClass: 'my-custom-class' },
    })
    const div = wrapper.find('div')
    expect(div.classes()).toContain('my-custom-class')
  })

  it('renders TrendingUp icon', () => {
    const wrapper = mount(Logo)
    expect(wrapper.findComponent(TrendingUp).exists()).toBe(true)
  })

  it('applies custom icon size when provided', () => {
    const wrapper = mount(Logo, {
      props: { iconSize: 32 },
    })
    const svg = wrapper.find('svg')
    expect(svg.attributes('width')).toBe('32')
    expect(svg.attributes('height')).toBe('32')
  })

  it('uses default icon size of 20', () => {
    const wrapper = mount(Logo)
    const svg = wrapper.find('svg')
    expect(svg.attributes('width')).toBe('20')
    expect(svg.attributes('height')).toBe('20')
  })
})
