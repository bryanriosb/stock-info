import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useAuthStore } from '../auth.store'

vi.mock('@/api/auth.api', () => ({
  authApi: {
    login: vi.fn(),
    register: vi.fn(),
    logout: vi.fn(),
  },
}))

vi.mock('@/lib/cookies', () => ({
  CookieManager: {
    getAccessToken: vi.fn(() => null),
    getRefreshToken: vi.fn(() => null),
    setTokens: vi.fn(),
    clearTokens: vi.fn(),
  },
}))

vi.mock('@/lib/jwt', () => ({
  decodeJwtPayload: vi.fn((token: string) => {
    if (token === 'valid-token') {
      return { sub: 'testuser', email: 'test@example.com', role: 'user' }
    }
    if (token === 'admin-token') {
      return { sub: 'admin', email: 'admin@example.com', role: 'admin' }
    }
    return null
  }),
}))

import { authApi } from '@/api/auth.api'
import { CookieManager } from '@/lib/cookies'

describe('auth.store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  describe('initial state', () => {
    it('has correct initial values', () => {
      const store = useAuthStore()
      expect(store.accessToken).toBeNull()
      expect(store.user).toBeNull()
      expect(store.username).toBeNull()
      expect(store.email).toBeNull()
      expect(store.role).toBeNull()
      expect(store.loading).toBe(false)
      expect(store.error).toBeNull()
    })
  })

  describe('computed properties', () => {
    it('isAuthenticated returns false when no tokens', () => {
      const store = useAuthStore()
      expect(store.isAuthenticated).toBe(false)
    })

    it('isAuthenticated returns true when accessToken exists', () => {
      const store = useAuthStore()
      store.accessToken = 'some-token'
      expect(store.isAuthenticated).toBe(true)
    })

    it('isAdmin returns false for regular user', () => {
      const store = useAuthStore()
      store.role = 'user'
      expect(store.isAdmin).toBe(false)
    })

    it('isAdmin returns true for admin role', () => {
      const store = useAuthStore()
      store.role = 'admin'
      expect(store.isAdmin).toBe(true)
    })

    it('userInitials returns ? when no username', () => {
      const store = useAuthStore()
      expect(store.userInitials).toBe('?')
    })

    it('userInitials returns first two characters uppercase', () => {
      const store = useAuthStore()
      store.username = 'john'
      expect(store.userInitials).toBe('JO')
    })
  })

  describe('login', () => {
    it('sets loading state during login', async () => {
      const store = useAuthStore()
      vi.mocked(authApi.login).mockImplementation(
        () => new Promise((resolve) => setTimeout(() => resolve({ data: { success: false } } as any), 100))
      )

      const loginPromise = store.login({ username: 'test', password: 'pass' })
      expect(store.loading).toBe(true)

      await loginPromise
      expect(store.loading).toBe(false)
    })

    it('returns true and sets tokens on successful login', async () => {
      const store = useAuthStore()
      vi.mocked(authApi.login).mockResolvedValue({
        data: {
          success: true,
          data: {
            access_token: 'valid-token',
            refresh_token: 'refresh-token',
            expires_in: 3600,
            refresh_expires_in: 86400,
          },
        },
      } as any)

      const result = await store.login({ username: 'test', password: 'pass' })

      expect(result).toBe(true)
      expect(store.accessToken).toBe('valid-token')
      expect(store.username).toBe('testuser')
      expect(store.email).toBe('test@example.com')
      expect(store.role).toBe('user')
      expect(CookieManager.setTokens).toHaveBeenCalledWith('valid-token', 'refresh-token', 3600, 86400)
    })

    it('returns false and sets error on failed login', async () => {
      const store = useAuthStore()
      vi.mocked(authApi.login).mockResolvedValue({
        data: { success: false, error: 'Invalid credentials' },
      } as any)

      const result = await store.login({ username: 'test', password: 'wrong' })

      expect(result).toBe(false)
      expect(store.error).toBe('Invalid credentials')
    })

    it('handles API errors gracefully', async () => {
      const store = useAuthStore()
      vi.mocked(authApi.login).mockRejectedValue({
        response: { data: { error: 'Server error' } },
      })

      const result = await store.login({ username: 'test', password: 'pass' })

      expect(result).toBe(false)
      expect(store.error).toBe('Server error')
    })
  })

  describe('register', () => {
    it('returns true and sets user on successful registration', async () => {
      const store = useAuthStore()
      const mockUser = { id: '1', username: 'newuser', email: 'new@example.com' }
      vi.mocked(authApi.register).mockResolvedValue({
        data: { success: true, data: mockUser },
      } as any)

      const result = await store.register({
        username: 'newuser',
        email: 'new@example.com',
        password: 'password123',
      })

      expect(result).toBe(true)
      expect(store.user).toEqual(mockUser)
    })

    it('returns false and sets error on failed registration', async () => {
      const store = useAuthStore()
      vi.mocked(authApi.register).mockResolvedValue({
        data: { success: false, error: 'Username taken' },
      } as any)

      const result = await store.register({
        username: 'taken',
        email: 'test@example.com',
        password: 'password123',
      })

      expect(result).toBe(false)
      expect(store.error).toBe('Username taken')
    })
  })

  describe('logout', () => {
    it('clears all auth state', async () => {
      const store = useAuthStore()
      store.accessToken = 'token'
      store.username = 'user'
      store.email = 'email@test.com'
      store.role = 'user'

      vi.mocked(CookieManager.getRefreshToken).mockReturnValue('refresh-token')
      vi.mocked(authApi.logout).mockResolvedValue({} as any)

      await store.logout()

      expect(store.accessToken).toBeNull()
      expect(store.user).toBeNull()
      expect(store.username).toBeNull()
      expect(store.email).toBeNull()
      expect(store.role).toBeNull()
      expect(CookieManager.clearTokens).toHaveBeenCalled()
    })

    it('handles logout API errors gracefully', async () => {
      const store = useAuthStore()
      vi.mocked(CookieManager.getRefreshToken).mockReturnValue('refresh-token')
      vi.mocked(authApi.logout).mockRejectedValue(new Error('Network error'))

      await expect(store.logout()).resolves.not.toThrow()
      expect(CookieManager.clearTokens).toHaveBeenCalled()
    })
  })

  describe('checkAuth', () => {
    it('restores auth state from stored token', () => {
      vi.mocked(CookieManager.getAccessToken).mockReturnValue('valid-token')
      const store = useAuthStore()

      store.checkAuth()

      expect(store.accessToken).toBe('valid-token')
      expect(store.username).toBe('testuser')
      expect(store.email).toBe('test@example.com')
      expect(store.role).toBe('user')
    })

    it('sets accessToken to null when only refresh token exists', () => {
      vi.mocked(CookieManager.getAccessToken).mockReturnValue(null)
      vi.mocked(CookieManager.getRefreshToken).mockReturnValue('refresh-token')
      const store = useAuthStore()

      store.checkAuth()

      expect(store.accessToken).toBeNull()
    })
  })
})
