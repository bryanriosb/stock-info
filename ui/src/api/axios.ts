import axios from 'axios'
import { CookieManager } from '@/lib/cookies'
import { transformBigIntIds } from '@/lib/json'

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  headers: { 'Content-Type': 'application/json' },
  transformResponse: [transformBigIntIds],
})

let isRefreshing = false
let failedQueue: Array<{
  resolve: (token: string) => void
  reject: (error: unknown) => void
}> = []

const processQueue = (error: unknown, token: string | null = null) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error)
    } else {
      prom.resolve(token!)
    }
  })
  failedQueue = []
}

// Request interceptor: Add authentication token
apiClient.interceptors.request.use(
  (config) => {
    const token = CookieManager.getAccessToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor: Handle authentication errors with refresh
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        })
          .then((token) => {
            originalRequest.headers.Authorization = `Bearer ${token}`
            return apiClient(originalRequest)
          })
          .catch((err) => Promise.reject(err))
      }

      originalRequest._retry = true

      const refreshToken = CookieManager.getRefreshToken()

      if (!refreshToken) {
        CookieManager.clearTokens()
        redirectToLogin()
        return Promise.reject(error)
      }

      isRefreshing = true

      try {
        const response = await axios.post(
          `${import.meta.env.VITE_API_URL || '/api/v1'}/auth/refresh`,
          { refresh_token: refreshToken },
          { headers: { 'Content-Type': 'application/json' } }
        )

        if (response.data.success) {
          const { access_token, refresh_token, expires_in, refresh_expires_in } =
            response.data.data

          CookieManager.setTokens(
            access_token,
            refresh_token,
            expires_in,
            refresh_expires_in
          )

          processQueue(null, access_token)

          originalRequest.headers.Authorization = `Bearer ${access_token}`
          return apiClient(originalRequest)
        } else {
          throw new Error('Refresh failed')
        }
      } catch (refreshError) {
        processQueue(refreshError, null)
        CookieManager.clearTokens()
        redirectToLogin()
        return Promise.reject(refreshError)
      } finally {
        isRefreshing = false
      }
    }

    return Promise.reject(error)
  }
)

function redirectToLogin() {
  const currentPath = window.location.pathname
  if (currentPath !== '/login') {
    window.location.href = `/login?redirect=${encodeURIComponent(currentPath)}`
  }
}

export default apiClient
