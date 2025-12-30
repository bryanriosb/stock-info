import axios from 'axios'
import { CookieManager } from '@/lib/cookies'
import { transformBigIntIds } from '@/lib/json'

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  headers: { 'Content-Type': 'application/json' },
  transformResponse: [transformBigIntIds],
})

// Request interceptor: Add authentication token
apiClient.interceptors.request.use(
  (config) => {
    const token = CookieManager.getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor: Handle authentication errors
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true

      // Clear expired tokens and redirect to login
      CookieManager.clearToken()
      
      // Redirect to login with return URL
      const currentPath = window.location.pathname
      if (currentPath !== '/login') {
        window.location.href = `/login?redirect=${encodeURIComponent(currentPath)}`
      }
    }

    return Promise.reject(error)
  }
)

export default apiClient
