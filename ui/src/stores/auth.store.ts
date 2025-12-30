import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth.api'
import type { LoginRequest, RegisterRequest } from '@/types/auth.types'
import type { User } from '@/types/user.types'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  async function login(credentials: LoginRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.login(credentials)
      if (response.data.success) {
        token.value = response.data.data.token
        localStorage.setItem('token', response.data.data.token)
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
    localStorage.removeItem('token')
  }

  function checkAuth() {
    const storedToken = localStorage.getItem('token')
    if (storedToken) token.value = storedToken
  }

  return { token, user, loading, error, isAuthenticated, login, register, logout, checkAuth }
})