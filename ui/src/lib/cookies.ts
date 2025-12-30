// Cookie utilities for secure token management
export class CookieManager {
  private static readonly TOKEN_KEY = 'auth_token'
  private static readonly TOKEN_EXPIRY_KEY = 'auth_token_expiry'
  
  /**
   * Set authentication token as a secure cookie
   */
  static setToken(token: string, expiresIn: number) {
    // Store token expiry in localStorage for client-side validation
    const expiry = new Date().getTime() + (expiresIn * 1000)
    localStorage.setItem(this.TOKEN_EXPIRY_KEY, expiry.toString())

    // Store token in cookie
    this.setClientCookie(this.TOKEN_KEY, token, expiresIn)
  }

  /**
   * Get authentication token
   */
  static getToken(): string | null {
    // Check if token has expired
    if (this.isTokenExpired()) {
      this.clearToken()
      return null
    }

    return this.getClientCookie(this.TOKEN_KEY)
  }

  /**
   * Clear authentication token
   */
  static clearToken() {
    localStorage.removeItem(this.TOKEN_EXPIRY_KEY)
    this.deleteClientCookie(this.TOKEN_KEY)
  }

  /**
   * Check if token is expired
   */
  static isTokenExpired(): boolean {
    const expiry = localStorage.getItem(this.TOKEN_EXPIRY_KEY)
    if (!expiry) return true
    
    return new Date().getTime() > parseInt(expiry)
  }

  /**
   * Set a client-side cookie (for development)
   */
  private static setClientCookie(name: string, value: string, maxAge: number) {
    document.cookie = `${name}=${value}; max-age=${maxAge}; path=/; secure=${location.protocol === 'https:'}; samesite=strict`
  }

  /**
   * Get a client-side cookie
   */
  private static getClientCookie(name: string): string | null {
    const nameEQ = name + "="
    const ca = document.cookie.split(';')
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i]
      while (c.charAt(0) === ' ') c = c.substring(1, c.length)
      if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
    }
    return null
  }

  /**
   * Delete a client-side cookie
   */
  private static deleteClientCookie(name: string) {
    document.cookie = `${name}=; max-age=-1; path=/`
  }
}