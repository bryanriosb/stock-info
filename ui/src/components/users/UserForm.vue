<script setup lang="ts">
import { computed, watch } from 'vue'
import { useForm } from '@tanstack/vue-form'
import { z } from 'zod'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Field, FieldLabel, FieldError, FieldGroup } from '@/components/ui/field'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { getErrors, isInvalid } from '@/lib/form'
import { Shield, User as UserIcon, Loader2 } from 'lucide-vue-next'
import type { User, Role } from '@/types/user.types'

export type UserFormMode = 'register' | 'create' | 'edit'

interface Props {
  mode: UserFormMode
  user?: User | null
  loading?: boolean
  error?: string | null
  submitLabel?: string
  showRoleField?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  showRoleField: false,
})

const emit = defineEmits<{
  submit: [data: UserFormData]
}>()

export interface UserFormData {
  username: string
  email: string
  password?: string
  confirmPassword?: string
  role?: Role
}

// Dynamic schema based on mode
const getSchema = () => {
  const baseSchema = {
    username: z
      .string()
      .min(3, 'Username must be at least 3 characters')
      .max(20, 'Username must be at most 20 characters')
      .regex(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters, numbers and underscores'),
    email: z.string().email('Invalid email address'),
  }

  if (props.mode === 'edit') {
    // Edit mode: only role required
    return z.object({
      role: z.enum(['user', 'admin'], {
        required_error: 'Role is required',
        invalid_type_error: 'Invalid role',
      }),
    })
  }

  // Register/Create mode: full validation
  const schema = z.object({
    ...baseSchema,
    password: z.string().min(6, 'Password must be at least 6 characters'),
    confirmPassword: z.string(),
    ...(props.showRoleField ? {
      role: z.enum(['user', 'admin'], {
        required_error: 'Role is required',
      }),
    } : {}),
  })

  return schema.refine((data) => data.password === data.confirmPassword, {
    message: 'Passwords do not match',
    path: ['confirmPassword'],
  })
}

const getDefaultValues = () => {
  if (props.mode === 'edit' && props.user) {
    return {
      username: props.user.username,
      email: props.user.email,
      role: props.user.role,
    }
  }
  return {
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    role: 'user' as Role,
  }
}

const form = useForm({
  defaultValues: getDefaultValues(),
  validators: {
    onSubmit: getSchema(),
  },
  onSubmit: async ({ value }) => {
    emit('submit', value as UserFormData)
  },
})

// Reset form when user prop changes (for edit mode)
watch(() => props.user, (newUser) => {
  if (props.mode === 'edit' && newUser) {
    form.reset()
    form.setFieldValue('role', newUser.role)
  }
}, { immediate: true })

const isEditMode = computed(() => props.mode === 'edit')
const showPasswordFields = computed(() => props.mode !== 'edit')
const buttonLabel = computed(() => {
  if (props.submitLabel) return props.submitLabel
  switch (props.mode) {
    case 'register': return 'Create account'
    case 'create': return 'Create User'
    case 'edit': return 'Save Changes'
    default: return 'Submit'
  }
})
const loadingLabel = computed(() => {
  switch (props.mode) {
    case 'register': return 'Creating account...'
    case 'create': return 'Creating...'
    case 'edit': return 'Saving...'
    default: return 'Loading...'
  }
})
</script>

<template>
  <form @submit.prevent="form.handleSubmit">
    <FieldGroup>
      <Alert v-if="error" variant="destructive">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <!-- Username field (not shown in edit mode) -->
      <form.Field v-if="!isEditMode" name="username">
        <template #default="{ field }">
          <Field :data-invalid="isInvalid(field)">
            <FieldLabel :for="field.name">Username</FieldLabel>
            <Input
              :id="field.name"
              :name="field.name"
              :model-value="field.state.value"
              :aria-invalid="isInvalid(field)"
              placeholder="Choose a username"
              autocomplete="username"
              @blur="field.handleBlur"
              @input="field.handleChange(($event.target as HTMLInputElement).value)"
            />
            <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
          </Field>
        </template>
      </form.Field>

      <!-- Email field (not shown in edit mode) -->
      <form.Field v-if="!isEditMode" name="email">
        <template #default="{ field }">
          <Field :data-invalid="isInvalid(field)">
            <FieldLabel :for="field.name">Email</FieldLabel>
            <Input
              :id="field.name"
              :name="field.name"
              type="email"
              :model-value="field.state.value"
              :aria-invalid="isInvalid(field)"
              placeholder="Enter your email"
              autocomplete="email"
              @blur="field.handleBlur"
              @input="field.handleChange(($event.target as HTMLInputElement).value)"
            />
            <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
          </Field>
        </template>
      </form.Field>

      <!-- Password fields (only for register/create) -->
      <template v-if="showPasswordFields">
        <form.Field name="password">
          <template #default="{ field }">
            <Field :data-invalid="isInvalid(field)">
              <FieldLabel :for="field.name">Password</FieldLabel>
              <Input
                :id="field.name"
                :name="field.name"
                type="password"
                :model-value="field.state.value"
                :aria-invalid="isInvalid(field)"
                placeholder="Create a password"
                autocomplete="new-password"
                @blur="field.handleBlur"
                @input="field.handleChange(($event.target as HTMLInputElement).value)"
              />
              <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
            </Field>
          </template>
        </form.Field>

        <form.Field name="confirmPassword">
          <template #default="{ field }">
            <Field :data-invalid="isInvalid(field)">
              <FieldLabel :for="field.name">Confirm Password</FieldLabel>
              <Input
                :id="field.name"
                :name="field.name"
                type="password"
                :model-value="field.state.value"
                :aria-invalid="isInvalid(field)"
                placeholder="Confirm your password"
                autocomplete="new-password"
                @blur="field.handleBlur"
                @input="field.handleChange(($event.target as HTMLInputElement).value)"
              />
              <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
            </Field>
          </template>
        </form.Field>
      </template>

      <!-- Role field (only when showRoleField is true) -->
      <form.Field v-if="showRoleField" name="role">
        <template #default="{ field }">
          <Field :data-invalid="isInvalid(field)">
            <FieldLabel>Role</FieldLabel>
            <Select
              :model-value="field.state.value"
              @update:model-value="field.handleChange"
            >
              <SelectTrigger :aria-invalid="isInvalid(field)">
                <SelectValue placeholder="Select a role" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="user">
                  <div class="flex items-center gap-2">
                    <UserIcon class="h-4 w-4" />
                    User
                  </div>
                </SelectItem>
                <SelectItem value="admin">
                  <div class="flex items-center gap-2">
                    <Shield class="h-4 w-4" />
                    Admin
                  </div>
                </SelectItem>
              </SelectContent>
            </Select>
            <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
          </Field>
        </template>
      </form.Field>

      <slot name="footer">
        <Button type="submit" class="w-full" :disabled="loading">
          <Loader2 v-if="loading" class="mr-2 h-4 w-4 animate-spin" />
          {{ loading ? loadingLabel : buttonLabel }}
        </Button>
      </slot>
    </FieldGroup>
  </form>
</template>
