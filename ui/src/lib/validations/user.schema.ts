import { z } from 'zod'

export const updateUserRoleSchema = z.object({
  role: z.enum(['user', 'admin'], {
    required_error: 'Role is required',
    invalid_type_error: 'Invalid role',
  }),
})

export type UpdateUserRoleFormData = z.infer<typeof updateUserRoleSchema>
