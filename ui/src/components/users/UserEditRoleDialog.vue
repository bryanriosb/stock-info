<script setup lang="ts">
import { ref, watch } from 'vue'
import { useForm } from '@tanstack/vue-form'
import { usersApi } from '@/api/users.api'
import { useToast } from '@/components/ui/toast'
import { updateUserRoleSchema } from '@/lib/validations/user.schema'
import { getErrors, isInvalid } from '@/lib/form'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Field, FieldLabel, FieldError, FieldGroup } from '@/components/ui/field'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Loader2, Shield, User as UserIcon } from 'lucide-vue-next'
import type { User, Role } from '@/types/user.types'

interface Props {
  user: User | null
  adminCount: number
}

const props = defineProps<Props>()
const open = defineModel<boolean>('open', { required: true })

const emit = defineEmits<{
  updated: [user: User]
}>()

const { toast } = useToast()
const updating = ref(false)
const formError = ref<string | null>(null)

const form = useForm({
  defaultValues: {
    role: 'user' as Role,
  },
  validators: {
    onSubmit: updateUserRoleSchema,
  },
  onSubmit: async ({ value }) => {
    if (!props.user) return

    // Check if trying to remove last admin
    if (props.user.role === 'admin' && value.role === 'user' && props.adminCount === 1) {
      formError.value = 'Cannot remove admin role from the last admin user.'
      return
    }

    updating.value = true
    formError.value = null

    try {
      const res = await usersApi.update(props.user.id, { role: value.role })
      if (res.data.success) {
        toast({ title: 'Role updated', description: `User role changed to ${value.role}` })
        emit('updated', res.data.data)
        open.value = false
      } else {
        formError.value = res.data.error || 'Failed to update role'
      }
    } catch (err: any) {
      formError.value = err.response?.data?.error || 'Failed to update role'
    } finally {
      updating.value = false
    }
  },
})

// Reset form when dialog opens with user data
watch(open, (isOpen) => {
  if (isOpen && props.user) {
    form.reset()
    form.setFieldValue('role', props.user.role)
    formError.value = null
  }
})
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Edit User Role</DialogTitle>
        <DialogDescription>
          Change the role for <span class="font-semibold">{{ user?.username }}</span>
        </DialogDescription>
      </DialogHeader>

      <form id="edit-role-form" @submit.prevent="form.handleSubmit">
        <FieldGroup>
          <Alert v-if="formError" variant="destructive">
            <AlertDescription>{{ formError }}</AlertDescription>
          </Alert>

          <form.Field name="role">
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
        </FieldGroup>
      </form>

      <DialogFooter>
        <Button variant="outline" @click="open = false" :disabled="updating">
          Cancel
        </Button>
        <Button type="submit" form="edit-role-form" :disabled="updating">
          <Loader2 v-if="updating" class="mr-2 h-4 w-4 animate-spin" />
          {{ updating ? 'Saving...' : 'Save Changes' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
