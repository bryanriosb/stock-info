import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth.api'
import { CookieManager } from '@/lib/cookies'
import { decodeJwtPayload } from '@/lib/jwt'
import type { LoginRequest, RegisterRequest } from '@/types/auth.types'
import type { User, Role } from '@/types/user.types'

export const useAuthStore = defineStore('auth', () => {
  const accessToken = ref<string | null>(CookieManager.getAccessToken())
  const user = ref<User | null>(null)
  const username = ref<string | null>(null)
  const email = ref<string | null>(null)
  const role = ref<Role | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => {
    return !!accessToken.value || !!CookieManager.getRefreshToken()
  })
  const isAdmin = computed(() => role.value === 'admin')
  const userInitials = computed(() => {
    if (!username.value) return '?'
    return username.value.slice(0, 2).toUpperCase()
  })

  async function login(credentials: LoginRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.login(credentials)
      if (response.data.success) {
        const data = response.data.data
        accessToken.value = data.access_token
        CookieManager.setTokens(
          data.access_token,
          data.refresh_token,
          data.expires_in,
          data.refresh_expires_in
        )
        const payload = decodeJwtPayload(data.access_token)
        username.value = payload?.sub || null
        email.value = payload?.email || null
        role.value = payload?.role || 'user'
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

  async function logout() {
    const refreshToken = CookieManager.getRefreshToken()
    if (refreshToken) {
      try {
        await authApi.logout({ refresh_token: refreshToken })
      } catch {
        // Ignore logout errors
      }
    }
    accessToken.value = null
    user.value = null
    username.value = null
    email.value = null
    role.value = null
    CookieManager.clearTokens()
  }

  function checkAuth() {
    const storedToken = CookieManager.getAccessToken()
    if (storedToken) {
      accessToken.value = storedToken
      const payload = decodeJwtPayload(storedToken)
      username.value = payload?.sub || null
      email.value = payload?.email || null
      role.value = payload?.role || 'user'
    } else if (CookieManager.getRefreshToken()) {
      // Has refresh token but no access token - will be refreshed on next API call
      accessToken.value = null
    } else {
      logout()
    }
  }

  return {
    accessToken,
    user,
    username,
    email,
    role,
    loading,
    error,
    isAuthenticated,
    isAdmin,
    userInitials,
    login,
    register,
    logout,
    checkAuth,
  }
})
