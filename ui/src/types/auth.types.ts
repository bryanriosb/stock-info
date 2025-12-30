export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  expires_in: number
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}
