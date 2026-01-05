import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import RatingBadge from './RatingBadge.vue'

describe('RatingBadge', () => {
  it('renders the rating text', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Buy' },
    })
    expect(wrapper.text()).toContain('Buy')
  })

  it('applies success styling for buy ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Strong Buy' },
    })
    expect(wrapper.find('.bg-success\\/10').exists()).toBe(true)
  })

  it('applies success styling for outperform ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Outperform' },
    })
    expect(wrapper.find('.bg-success\\/10').exists()).toBe(true)
  })

  it('applies destructive styling for sell ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Sell' },
    })
    expect(wrapper.find('.bg-destructive\\/10').exists()).toBe(true)
  })

  it('applies destructive styling for underperform ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Underperform' },
    })
    expect(wrapper.find('.bg-destructive\\/10').exists()).toBe(true)
  })

  it('applies neutral styling for hold ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Hold' },
    })
    expect(wrapper.find('.bg-accent\\/10').exists()).toBe(true)
  })

  it('applies neutral styling for unknown ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Unknown' },
    })
    expect(wrapper.find('.bg-accent\\/10').exists()).toBe(true)
  })

  it('handles case insensitive ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'BUY' },
    })
    expect(wrapper.find('.bg-success\\/10').exists()).toBe(true)
  })
})
