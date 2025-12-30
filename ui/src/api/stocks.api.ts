import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { Stock, StockQueryParams } from '@/types/stock.types'

export const stocksApi = {
  getAll: (params?: StockQueryParams) => apiClient.get<ApiResponse<Stock[]>>('/stocks', { params }),
  getById: (id: number) => apiClient.get<ApiResponse<Stock>>(`/stocks/${id}`),
  sync: () => apiClient.post<ApiResponse<{ message: string; count: number }>>('/stocks/sync'),
}