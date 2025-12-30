export type Role = 'user' | 'admin'

export interface User {
  id: number
  username: string
  email: string
  role: Role
  created_at: string
  updated_at: string
}
