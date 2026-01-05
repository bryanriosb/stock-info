import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useStocksStore } from '../stocks.store'

vi.mock('@/api/stocks.api', () => ({
  stocksApi: {
    getAll: vi.fn(),
    getById: vi.fn(),
    syncStream: vi.fn(),
  },
}))

import { stocksApi } from '@/api/stocks.api'

describe('stocks.store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  describe('initial state', () => {
    it('has correct initial values', () => {
      const store = useStocksStore()
      expect(store.stocks).toEqual([])
      expect(store.currentStock).toBeNull()
      expect(store.loading).toBe(false)
      expect(store.syncing).toBe(false)
      expect(store.syncProgress).toBeNull()
      expect(store.error).toBeNull()
      expect(store.meta).toBeNull()
      expect(store.hasLoadedOnce).toBe(false)
    })

    it('has default query params', () => {
      const store = useStocksStore()
      expect(store.queryParams).toEqual({
        page: 1,
        limit: 20,
        sort_by: 'id',
        sort_dir: 'asc',
      })
    })
  })

  describe('computed properties', () => {
    it('hasActiveFilters returns false when no filters', () => {
      const store = useStocksStore()
      expect(store.hasActiveFilters).toBe(false)
    })

    it('hasActiveFilters returns true when search is set', () => {
      const store = useStocksStore()
      store.queryParams.search = 'AAPL'
      expect(store.hasActiveFilters).toBe(true)
    })

    it('hasActiveFilters returns true when rating_from is set', () => {
      const store = useStocksStore()
      store.queryParams.rating_from = 'Buy'
      expect(store.hasActiveFilters).toBe(true)
    })

    it('hasActiveFilters returns true when rating_to is set', () => {
      const store = useStocksStore()
      store.queryParams.rating_to = 'Sell'
      expect(store.hasActiveFilters).toBe(true)
    })
  })

  describe('fetchStocks', () => {
    it('sets loading state during fetch', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockImplementation(
        () => new Promise((resolve) => setTimeout(() => resolve({ data: { success: true, data: [] } } as any), 100))
      )

      const fetchPromise = store.fetchStocks()
      expect(store.loading).toBe(true)

      await fetchPromise
      expect(store.loading).toBe(false)
    })

    it('populates stocks on successful fetch', async () => {
      const store = useStocksStore()
      const mockStocks = [
        { id: '1', ticker: 'AAPL', company: 'Apple Inc.' },
        { id: '2', ticker: 'GOOGL', company: 'Alphabet Inc.' },
      ]
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: {
          success: true,
          data: mockStocks,
          meta: { page: 1, limit: 20, total: 2, total_pages: 1 },
        },
      } as any)

      await store.fetchStocks()

      expect(store.stocks).toEqual(mockStocks)
      expect(store.meta).toEqual({ page: 1, limit: 20, total: 2, total_pages: 1 })
      expect(store.hasLoadedOnce).toBe(true)
    })

    it('merges params with existing query params', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.fetchStocks({ search: 'AAPL' })

      expect(store.queryParams.search).toBe('AAPL')
      expect(store.queryParams.page).toBe(1)
      expect(stocksApi.getAll).toHaveBeenCalledWith(
        expect.objectContaining({ search: 'AAPL', page: 1 })
      )
    })

    it('sets error on failed fetch', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: false, error: 'Failed to fetch' },
      } as any)

      await store.fetchStocks()

      expect(store.error).toBe('Failed to fetch')
    })

    it('handles API errors gracefully', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockRejectedValue({
        response: { data: { error: 'Server error' } },
      })

      await store.fetchStocks()

      expect(store.error).toBe('Server error')
    })
  })

  describe('fetchStockById', () => {
    it('sets loading state during fetch', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getById).mockImplementation(
        () => new Promise((resolve) => setTimeout(() => resolve({ data: { success: true, data: {} } } as any), 100))
      )

      const fetchPromise = store.fetchStockById('1')
      expect(store.loading).toBe(true)

      await fetchPromise
      expect(store.loading).toBe(false)
    })

    it('sets currentStock on successful fetch', async () => {
      const store = useStocksStore()
      const mockStock = { id: '1', ticker: 'AAPL', company: 'Apple Inc.' }
      vi.mocked(stocksApi.getById).mockResolvedValue({
        data: { success: true, data: mockStock },
      } as any)

      await store.fetchStockById('1')

      expect(store.currentStock).toEqual(mockStock)
    })

    it('clears currentStock before fetching', async () => {
      const store = useStocksStore()
      store.currentStock = { id: 'old' } as any
      vi.mocked(stocksApi.getById).mockResolvedValue({
        data: { success: true, data: { id: 'new' } },
      } as any)

      const fetchPromise = store.fetchStockById('new')

      expect(store.currentStock).toBeNull()
      await fetchPromise
    })
  })

  describe('setSort', () => {
    it('sets new sort field with asc direction', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.setSort('ticker')

      expect(store.queryParams.sort_by).toBe('ticker')
      expect(store.queryParams.sort_dir).toBe('asc')
    })

    it('toggles direction when clicking same field', async () => {
      const store = useStocksStore()
      store.queryParams.sort_by = 'ticker'
      store.queryParams.sort_dir = 'asc'
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.setSort('ticker')

      expect(store.queryParams.sort_dir).toBe('desc')
    })
  })

  describe('setPage', () => {
    it('updates page and fetches', async () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.setPage(3)

      expect(store.queryParams.page).toBe(3)
      expect(stocksApi.getAll).toHaveBeenCalled()
    })
  })

  describe('setFilters', () => {
    it('updates filters and resets page to 1', async () => {
      const store = useStocksStore()
      store.queryParams.page = 5
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.setFilters({ search: 'AAPL', rating_from: 'Buy' })

      expect(store.queryParams.search).toBe('AAPL')
      expect(store.queryParams.rating_from).toBe('Buy')
      expect(store.queryParams.page).toBe(1)
    })
  })

  describe('clearFilters', () => {
    it('resets all params to defaults', async () => {
      const store = useStocksStore()
      store.queryParams = {
        page: 5,
        limit: 50,
        sort_by: 'ticker',
        sort_dir: 'desc',
        search: 'AAPL',
        rating_from: 'Buy',
      }
      vi.mocked(stocksApi.getAll).mockResolvedValue({
        data: { success: true, data: [] },
      } as any)

      await store.clearFilters()

      expect(store.queryParams).toEqual({
        page: 1,
        limit: 20,
        sort_by: 'id',
        sort_dir: 'asc',
      })
    })
  })

  describe('syncStocks', () => {
    it('sets syncing state and initial progress', () => {
      const store = useStocksStore()
      vi.mocked(stocksApi.syncStream).mockReturnValue(() => {})

      store.syncStocks()

      expect(store.syncing).toBe(true)
      expect(store.syncProgress).toEqual({
        current: 0,
        total: 0,
        percent: 0,
        status: 'starting',
        message: 'Connecting...',
      })
    })

    it('does not start sync if already syncing', () => {
      const store = useStocksStore()
      store.syncing = true
      vi.mocked(stocksApi.syncStream).mockReturnValue(() => {})

      store.syncStocks()

      expect(stocksApi.syncStream).not.toHaveBeenCalled()
    })
  })

  describe('cancelSync', () => {
    it('calls abort function and resets sync state', () => {
      const store = useStocksStore()
      const abortFn = vi.fn()
      vi.mocked(stocksApi.syncStream).mockReturnValue(abortFn)

      store.syncStocks()
      store.cancelSync()

      expect(abortFn).toHaveBeenCalled()
      expect(store.syncing).toBe(false)
      expect(store.syncProgress).toBeNull()
    })
  })
})
