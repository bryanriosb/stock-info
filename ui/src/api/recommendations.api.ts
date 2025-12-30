import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { StockRecommendation } from '@/types/stock.types'

export const recommendationsApi = {
  getAll: (limit?: number) => apiClient.get<ApiResponse<StockRecommendation[]>>('/recommendations', { params: { limit } }),
}