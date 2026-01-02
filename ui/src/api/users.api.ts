import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { User, Role } from '@/types/user.types'

export interface UpdateUserRequest {
  username?: string
  email?: string
  role?: Role
}

export const usersApi = {
  getAll: () => apiClient.get<ApiResponse<User[]>>('/users'),
  update: (id: number, data: UpdateUserRequest) => apiClient.put<ApiResponse<User>>(`/users/${id}`, data),
  delete: (id: number) => apiClient.delete<ApiResponse<{ message: string }>>(`/users/${id}`),
}
