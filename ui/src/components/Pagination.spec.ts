import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Pagination from './Pagination.vue'

describe('Pagination', () => {
  const defaultProps = {
    currentPage: 1,
    totalPages: 5,
    totalItems: 100,
    limit: 20,
  }

  it('renders the correct item range', () => {
    const wrapper = mount(Pagination, { props: defaultProps })
    expect(wrapper.text()).toContain('Showing 1 to 20 of 100')
  })

  it('calculates correct range for middle pages', () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 3 },
    })
    expect(wrapper.text()).toContain('Showing 41 to 60 of 100')
  })

  it('calculates correct range for last page with partial items', () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 5, totalItems: 95 },
    })
    expect(wrapper.text()).toContain('Showing 81 to 95 of 95')
  })

  it('disables previous button on first page', () => {
    const wrapper = mount(Pagination, { props: defaultProps })
    const buttons = wrapper.findAll('button')
    expect(buttons[0].attributes('disabled')).toBeDefined()
  })

  it('enables previous button on pages after first', () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 2 },
    })
    const buttons = wrapper.findAll('button')
    expect(buttons[0].attributes('disabled')).toBeUndefined()
  })

  it('disables next button on last page', () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 5 },
    })
    const buttons = wrapper.findAll('button')
    expect(buttons[1].attributes('disabled')).toBeDefined()
  })

  it('enables next button on pages before last', () => {
    const wrapper = mount(Pagination, { props: defaultProps })
    const buttons = wrapper.findAll('button')
    expect(buttons[1].attributes('disabled')).toBeUndefined()
  })

  it('emits page-change event with previous page when clicking previous', async () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 3 },
    })
    const buttons = wrapper.findAll('button')
    await buttons[0].trigger('click')
    expect(wrapper.emitted('page-change')).toEqual([[2]])
  })

  it('emits page-change event with next page when clicking next', async () => {
    const wrapper = mount(Pagination, {
      props: { ...defaultProps, currentPage: 3 },
    })
    const buttons = wrapper.findAll('button')
    await buttons[1].trigger('click')
    expect(wrapper.emitted('page-change')).toEqual([[4]])
  })
})
