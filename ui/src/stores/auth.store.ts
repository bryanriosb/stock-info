import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth.api'
import { CookieManager } from '@/lib/cookies'
import type { LoginRequest, RegisterRequest } from '@/types/auth.types'
import type { User } from '@/types/user.types'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(CookieManager.getToken())
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value && !CookieManager.isTokenExpired())

  async function login(credentials: LoginRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.login(credentials)
      if (response.data.success) {
        token.value = response.data.data.token
        CookieManager.setToken(response.data.data.token, response.data.data.expires_in)
        return true
      }
      error.value = response.data.error || 'Login failed'
      return false
    } catch (err: any) {
      error.value = err.response?.data?.error || 'An error occurred'
      return false
    } finally {
      loading.value = false
    }
  }

  async function register(data: RegisterRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.register(data)
      if (response.data.success) {
        user.value = response.data.data
        return true
      }
      error.value = response.data.error || 'Registration failed'
      return false
    } catch (err: any) {
      error.value = err.response?.data?.error || 'An error occurred'
      return false
    } finally {
      loading.value = false
    }
  }

  function logout() {
    token.value = null
    user.value = null
    CookieManager.clearToken()
  }

  function checkAuth() {
    const storedToken = CookieManager.getToken()
    if (storedToken && !CookieManager.isTokenExpired()) {
      token.value = storedToken
    } else {
      logout()
    }
  }

  return { token, user, loading, error, isAuthenticated, login, register, logout, checkAuth }
})