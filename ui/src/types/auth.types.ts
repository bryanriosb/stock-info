export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  expires_in: number
  refresh_expires_in: number
}

export interface RefreshRequest {
  refresh_token: string
}

export interface RefreshResponse {
  access_token: string
  refresh_token: string
  expires_in: number
  refresh_expires_in: number
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}
