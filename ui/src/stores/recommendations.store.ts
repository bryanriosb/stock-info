import { defineStore } from 'pinia'
import { ref } from 'vue'
import { recommendationsApi } from '@/api/recommendations.api'
import type { StockRecommendation } from '@/types/stock.types'

export const useRecommendationsStore = defineStore('recommendations', () => {
  const recommendations = ref<StockRecommendation[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchRecommendations(limit = 10) {
    loading.value = true
    error.value = null
    try {
      const res = await recommendationsApi.getAll(limit)
      if (res.data.success) {
        recommendations.value = res.data.data ?? []
      } else {
        error.value = res.data.error || 'Failed'
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error fetching recommendations'
    } finally {
      loading.value = false
    }
  }

  return { recommendations, loading, error, fetchRecommendations }
})