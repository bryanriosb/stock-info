import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useRecommendationsStore } from '../recommendations.store'

vi.mock('@/api/recommendations.api', () => ({
  recommendationsApi: {
    getAll: vi.fn(),
  },
}))

import { recommendationsApi } from '@/api/recommendations.api'

describe('recommendations.store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  describe('initial state', () => {
    it('has correct initial values', () => {
      const store = useRecommendationsStore()
      expect(store.recommendations).toEqual([])
      expect(store.loading).toBe(false)
      expect(store.error).toBeNull()
    })
  })

  describe('fetchRecommendations', () => {
    it('sets loading state during fetch', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockImplementation(
        () =>
          new Promise((resolve) =>
            setTimeout(() => resolve({ data: { success: true, data: [] } } as any), 100)
          )
      )

      const fetchPromise = store.fetchRecommendations()
      expect(store.loading).toBe(true)

      await fetchPromise
      expect(store.loading).toBe(false)
    })

    it('populates recommendations on successful fetch', async () => {
      const store = useRecommendationsStore()
      const mockRecommendations = [
        { id: '1', ticker: 'AAPL', rating: 'Buy', target_price: 200 },
        { id: '2', ticker: 'GOOGL', rating: 'Hold', target_price: 150 },
      ]
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: true, data: mockRecommendations },
      } as any)

      await store.fetchRecommendations()

      expect(store.recommendations).toEqual(mockRecommendations)
    })

    it('uses default limit of 10', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.fetchRecommendations()

      expect(recommendationsApi.getAll).toHaveBeenCalledWith(10)
    })

    it('accepts custom limit parameter', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.fetchRecommendations(25)

      expect(recommendationsApi.getAll).toHaveBeenCalledWith(25)
    })

    it('handles null data response', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: true, data: null },
      } as any)

      await store.fetchRecommendations()

      expect(store.recommendations).toEqual([])
    })

    it('sets error on failed fetch', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: false, error: 'Failed to fetch recommendations' },
      } as any)

      await store.fetchRecommendations()

      expect(store.error).toBe('Failed to fetch recommendations')
      expect(store.recommendations).toEqual([])
    })

    it('handles API errors gracefully', async () => {
      const store = useRecommendationsStore()
      vi.mocked(recommendationsApi.getAll).mockRejectedValue({
        response: { data: { error: 'Server error' } },
      })

      await store.fetchRecommendations()

      expect(store.error).toBe('Server error')
    })

    it('clears error before fetching', async () => {
      const store = useRecommendationsStore()
      store.error = 'Previous error'
      vi.mocked(recommendationsApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.fetchRecommendations()

      expect(store.error).toBeNull()
    })
  })
})
