import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RefreshRequest,
  RefreshResponse,
} from '@/types/auth.types'
import type { User } from '@/types/user.types'

export const authApi = {
  login: (data: LoginRequest) =>
    apiClient.post<ApiResponse<LoginResponse>>('/auth/login', data),
  register: (data: RegisterRequest) =>
    apiClient.post<ApiResponse<User>>('/users', data),
  refresh: (data: RefreshRequest) =>
    apiClient.post<ApiResponse<RefreshResponse>>('/auth/refresh', data),
  logout: (data: RefreshRequest) =>
    apiClient.post<ApiResponse<{ message: string }>>('/auth/logout', data),
}