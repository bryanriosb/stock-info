export class CookieManager {
  private static readonly ACCESS_TOKEN_KEY = 'access_token'
  private static readonly REFRESH_TOKEN_KEY = 'refresh_token'
  private static readonly ACCESS_EXPIRY_KEY = 'access_token_expiry'
  private static readonly REFRESH_EXPIRY_KEY = 'refresh_token_expiry'

  static setTokens(
    accessToken: string,
    refreshToken: string,
    expiresIn: number,
    refreshExpiresIn: number
  ) {
    const accessExpiry = new Date().getTime() + expiresIn * 1000
    const refreshExpiry = new Date().getTime() + refreshExpiresIn * 1000

    localStorage.setItem(this.ACCESS_EXPIRY_KEY, accessExpiry.toString())
    localStorage.setItem(this.REFRESH_EXPIRY_KEY, refreshExpiry.toString())

    this.setClientCookie(this.ACCESS_TOKEN_KEY, accessToken, expiresIn)
    this.setClientCookie(this.REFRESH_TOKEN_KEY, refreshToken, refreshExpiresIn)
  }

  static getAccessToken(): string | null {
    if (this.isAccessTokenExpired()) {
      return null
    }
    return this.getClientCookie(this.ACCESS_TOKEN_KEY)
  }

  static getRefreshToken(): string | null {
    if (this.isRefreshTokenExpired()) {
      this.clearTokens()
      return null
    }
    return this.getClientCookie(this.REFRESH_TOKEN_KEY)
  }

  static clearTokens() {
    localStorage.removeItem(this.ACCESS_EXPIRY_KEY)
    localStorage.removeItem(this.REFRESH_EXPIRY_KEY)
    this.deleteClientCookie(this.ACCESS_TOKEN_KEY)
    this.deleteClientCookie(this.REFRESH_TOKEN_KEY)
  }

  static isAccessTokenExpired(): boolean {
    const expiry = localStorage.getItem(this.ACCESS_EXPIRY_KEY)
    if (!expiry) return true
    return new Date().getTime() > parseInt(expiry)
  }

  static isRefreshTokenExpired(): boolean {
    const expiry = localStorage.getItem(this.REFRESH_EXPIRY_KEY)
    if (!expiry) return true
    return new Date().getTime() > parseInt(expiry)
  }

  static canRefresh(): boolean {
    return this.isAccessTokenExpired() && !this.isRefreshTokenExpired()
  }

  private static setClientCookie(name: string, value: string, maxAge: number) {
    document.cookie = `${name}=${value}; max-age=${maxAge}; path=/; secure=${location.protocol === 'https:'}; samesite=strict`
  }

  private static getClientCookie(name: string): string | null {
    const nameEQ = name + '='
    const ca = document.cookie.split(';')
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i]
      while (c.charAt(0) === ' ') c = c.substring(1, c.length)
      if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
    }
    return null
  }

  private static deleteClientCookie(name: string) {
    document.cookie = `${name}=; max-age=-1; path=/`
  }
}
