import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { Stock, StockQueryParams } from '@/types/stock.types'

export interface SyncProgress {
  current: number
  total: number
  percent: number
  status: 'starting' | 'fetching' | 'saving' | 'completed' | 'error'
  message?: string
}

export const stocksApi = {
  getAll: (params?: StockQueryParams) => apiClient.get<ApiResponse<Stock[]>>('/stocks', { params }),
  getById: (id: string) => apiClient.get<ApiResponse<Stock>>(`/stocks/${id}`),

  // SSE sync with progress using fetch streaming (supports auth headers)
  syncStream: (onProgress: (progress: SyncProgress) => void, onComplete: () => void, onError: (error: string) => void) => {
    const baseURL = apiClient.defaults.baseURL || ''
    const token = localStorage.getItem('access_token')
    const controller = new AbortController()

    fetch(`${baseURL}/stocks/sync-stream`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Accept': 'text/event-stream',
      },
      signal: controller.signal,
    }).then(async (response) => {
      if (!response.ok) {
        onError(`HTTP error: ${response.status}`)
        return
      }

      const reader = response.body?.getReader()
      if (!reader) {
        onError('No response body')
        return
      }

      const decoder = new TextDecoder()
      let buffer = ''

      while (true) {
        const { done, value } = await reader.read()
        if (done) break

        buffer += decoder.decode(value, { stream: true })
        const lines = buffer.split('\n\n')
        buffer = lines.pop() || ''

        for (const line of lines) {
          if (line.startsWith('data: ')) {
            try {
              const data = JSON.parse(line.slice(6)) as SyncProgress
              onProgress(data)

              if (data.status === 'completed') {
                onComplete()
              } else if (data.status === 'error') {
                onError(data.message || 'Sync failed')
              }
            } catch (e) {
              console.error('Failed to parse SSE data:', e)
            }
          }
        }
      }
    }).catch((err) => {
      if (err.name !== 'AbortError') {
        onError(err.message || 'Connection failed')
      }
    })

    return () => controller.abort()
  },
}