import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { Stock, StockQueryParams } from '@/types/stock.types'
import { CookieManager } from '@/lib/cookies'

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
    const token = CookieManager.getAccessToken()
    const controller = new AbortController()
    let completed = false

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

      try {
        while (true) {
          const { done, value } = await reader.read()
          if (done) {
            // Stream ended - if we got completed status, it's fine
            if (!completed) {
              // Process any remaining buffer
              if (buffer.trim()) {
                const lines = buffer.split('\n')
                for (const line of lines) {
                  if (line.startsWith('data: ')) {
                    try {
                      const data = JSON.parse(line.slice(6)) as SyncProgress
                      onProgress(data)
                      if (data.status === 'completed') {
                        completed = true
                        onComplete()
                      }
                    } catch { /* ignore parse errors */ }
                  }
                }
              }
            }
            break
          }

          buffer += decoder.decode(value, { stream: true })
          const lines = buffer.split('\n\n')
          buffer = lines.pop() || ''

          for (const line of lines) {
            // Skip heartbeat comments
            if (line.startsWith(':')) continue

            if (line.startsWith('data: ')) {
              try {
                const data = JSON.parse(line.slice(6)) as SyncProgress
                onProgress(data)

                if (data.status === 'completed') {
                  completed = true
                  onComplete()
                } else if (data.status === 'error') {
                  onError(data.message || 'Sync failed')
                  return
                }
              } catch (e) {
                console.error('Failed to parse SSE data:', e)
              }
            }
          }
        }
      } catch (readError) {
        // Stream read error - only report if not already completed
        if (!completed && (readError as Error).name !== 'AbortError') {
          console.error('Stream read error:', readError)
          // Don't call onError here - the stream might have completed successfully
        }
      }
    }).catch((err) => {
      if (err.name !== 'AbortError' && !completed) {
        onError(err.message || 'Connection failed')
      }
    })

    return () => controller.abort()
  },
}