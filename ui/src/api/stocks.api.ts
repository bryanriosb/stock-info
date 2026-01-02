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

// SSE direct to backend (bypasses Vite proxy issues)
const SSE_BASE_URL = 'http://localhost:5000/api/v1'

export const stocksApi = {
  getAll: (params?: StockQueryParams) => apiClient.get<ApiResponse<Stock[]>>('/stocks', { params }),
  getById: (id: string) => apiClient.get<ApiResponse<Stock>>(`/stocks/${id}`),

  // SSE sync with progress using fetch + ReadableStream
  syncStream: (
    onProgress: (progress: SyncProgress) => void,
    onComplete: () => void,
    onError: (error: string) => void
  ) => {
    const token = CookieManager.getAccessToken()
    const controller = new AbortController()
    let lastProgress: SyncProgress | null = null
    let retryCount = 0
    const maxRetries = 3

    const connect = async () => {
      try {
        const response = await fetch(`${SSE_BASE_URL}/stocks/sync-stream?token=${token}`, {
          signal: controller.signal,
          headers: { Accept: 'text/event-stream' },
          keepalive: true,
        })

        if (!response.ok) {
          throw new Error(`HTTP error: ${response.status}`)
        }

        const reader = response.body?.getReader()
        if (!reader) {
          throw new Error('No response body')
        }

        const decoder = new TextDecoder()
        let buffer = ''
        retryCount = 0 // Reset on successful connection

        while (true) {
          const { done, value } = await reader.read()
          if (done) break

          buffer += decoder.decode(value, { stream: true })
          const lines = buffer.split('\n')
          buffer = lines.pop() || ''

          for (const line of lines) {
            // Skip heartbeat/ping comments and empty lines
            if (line.startsWith(':') || line.trim() === '') continue

            if (line.startsWith('data: ')) {
              try {
                const data = JSON.parse(line.slice(6)) as SyncProgress
                lastProgress = data
                onProgress(data)

                if (data.status === 'completed') {
                  reader.cancel()
                  onComplete()
                  return
                } else if (data.status === 'error') {
                  reader.cancel()
                  onError(data.message || 'Sync failed')
                  return
                }
              } catch (e) {
                console.error('Failed to parse SSE data:', e)
              }
            }
          }
        }

        // Stream ended unexpectedly - if we had progress, sync might still be running
        throw new Error('Connection closed')
      } catch (e: any) {
        if (e.name === 'AbortError') return

        // Retry if not completed and we have retries left
        if (retryCount < maxRetries && lastProgress?.status !== 'completed') {
          retryCount++
          console.log(`SSE connection lost, retrying (${retryCount}/${maxRetries})...`)
          // Wait before retry
          await new Promise((r) => setTimeout(r, 1000 * retryCount))
          if (!controller.signal.aborted) {
            connect()
          }
          return
        }

        onError(e.message || 'Connection failed')
      }
    }

    connect()
    return () => controller.abort()
  },
}