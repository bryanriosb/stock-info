import { defineStore } from 'pinia'
import { ref } from 'vue'
import { stocksApi, type SyncProgress } from '@/api/stocks.api'
import type { Stock, StockQueryParams } from '@/types/stock.types'
import type { PaginationMeta } from '@/types/api.types'

export const useStocksStore = defineStore('stocks', () => {
  const stocks = ref<Stock[]>([])
  const currentStock = ref<Stock | null>(null)
  const loading = ref(false)
  const syncing = ref(false)
  const syncProgress = ref<SyncProgress | null>(null)
  const error = ref<string | null>(null)
  const meta = ref<PaginationMeta | null>(null)
  const queryParams = ref<StockQueryParams>({ page: 1, limit: 20, sort_by: 'id', sort_dir: 'asc' })

  let abortSync: (() => void) | null = null

  async function fetchStocks(params?: StockQueryParams) {
    loading.value = true
    error.value = null
    const merged = { ...queryParams.value, ...params }
    queryParams.value = merged
    try {
      const res = await stocksApi.getAll(merged)
      if (res.data.success) {
        stocks.value = res.data.data
        meta.value = res.data.meta || null
      } else {
        error.value = res.data.error || 'Failed to fetch'
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
    } finally {
      loading.value = false
    }
  }

  async function fetchStockById(id: string) {
    loading.value = true
    error.value = null
    currentStock.value = null
    try {
      const res = await stocksApi.getById(id)
      if (res.data.success) currentStock.value = res.data.data
      else error.value = res.data.error || 'Not found'
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
    } finally {
      loading.value = false
    }
  }

  function syncStocks() {
    if (syncing.value) return

    syncing.value = true
    error.value = null
    syncProgress.value = { current: 0, total: 0, percent: 0, status: 'starting', message: 'Connecting...' }

    abortSync = stocksApi.syncStream(
      // onProgress
      (progress) => {
        syncProgress.value = progress
      },
      // onComplete
      async () => {
        syncing.value = false
        syncProgress.value = null
        abortSync = null
        await fetchStocks()
      },
      // onError
      (err) => {
        error.value = err
        syncing.value = false
        syncProgress.value = null
        abortSync = null
      }
    )
  }

  function cancelSync() {
    if (abortSync) {
      abortSync()
      abortSync = null
      syncing.value = false
      syncProgress.value = null
    }
  }

  function setSort(field: string) {
    if (queryParams.value.sort_by === field) {
      queryParams.value.sort_dir = queryParams.value.sort_dir === 'asc' ? 'desc' : 'asc'
    } else {
      queryParams.value.sort_by = field as any
      queryParams.value.sort_dir = 'asc'
    }
    fetchStocks()
  }

  function setPage(page: number) {
    queryParams.value.page = page
    fetchStocks()
  }

  function setFilters(filters: { ticker?: string; company?: string }) {
    queryParams.value = { ...queryParams.value, ...filters, page: 1 }
    fetchStocks()
  }

  function clearFilters() {
    queryParams.value = { page: 1, limit: 20, sort_by: 'id', sort_dir: 'asc' }
    fetchStocks()
  }

  return { stocks, currentStock, loading, syncing, syncProgress, error, meta, queryParams, fetchStocks, fetchStockById, syncStocks, cancelSync, setSort, setPage, setFilters, clearFilters }
})