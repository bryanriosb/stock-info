import type { Role } from '@/types/user.types'

interface JwtPayload {
  sub: string
  email: string
  role: Role
  exp: number
  iat: number
}

export function decodeJwtPayload(token: string): JwtPayload | null {
  try {
    const parts = token.split('.')
    if (parts.length !== 3) return null

    const payload = parts[1]
    const decoded = atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
    return JSON.parse(decoded)
  } catch {
    return null
  }
}
