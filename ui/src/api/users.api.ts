import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { User } from '@/types/user.types'

export const usersApi = {
  getAll: () => apiClient.get<ApiResponse<User[]>>('/users'),
  delete: (id: number) => apiClient.delete<ApiResponse<{ message: string }>>(`/users/${id}`),
}
