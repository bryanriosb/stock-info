<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import { usersApi } from '@/api/users.api'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import UserForm from './UserForm.vue'
import type { User } from '@/types/user.types'
import type { UserFormData, UserFormMode } from './UserForm.vue'

interface Props {
  user?: User | null
  adminCount?: number
  mode?: UserFormMode
}

const props = withDefaults(defineProps<Props>(), {
  adminCount: 0,
  mode: 'edit',
})

const open = defineModel<boolean>('open', { required: true })

const emit = defineEmits<{
  updated: [user: User]
  created: [user: User]
}>()

const { toast } = useToast()
const loading = ref(false)
const formError = ref<string | null>(null)

const dialogTitle = computed(() => {
  switch (props.mode) {
    case 'create': return 'Create User'
    case 'edit': return 'Edit User Role'
    default: return 'User'
  }
})

const dialogDescription = computed(() => {
  if (props.mode === 'edit' && props.user) {
    return `Change the role for ${props.user.username}`
  }
  return 'Fill in the user details'
})

async function handleSubmit(data: UserFormData) {
  if (props.mode === 'edit' && props.user) {
    // Check if trying to remove last admin
    if (props.user.role === 'admin' && data.role === 'user' && props.adminCount === 1) {
      formError.value = 'Cannot remove admin role from the last admin user.'
      return
    }

    loading.value = true
    formError.value = null

    try {
      const res = await usersApi.update(props.user.id, { role: data.role })
      if (res.data.success) {
        toast({ title: 'Role updated', description: `User role changed to ${data.role}` })
        emit('updated', res.data.data)
        open.value = false
      } else {
        formError.value = res.data.error || 'Failed to update role'
      }
    } catch (err: any) {
      formError.value = err.response?.data?.error || 'Failed to update role'
    } finally {
      loading.value = false
    }
  }
  // TODO: Add create user logic when needed
}

// Reset error when dialog opens
watch(open, (isOpen) => {
  if (isOpen) {
    formError.value = null
  }
})
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>{{ dialogTitle }}</DialogTitle>
        <DialogDescription>{{ dialogDescription }}</DialogDescription>
      </DialogHeader>

      <UserForm
        :mode="mode"
        :user="user"
        :loading="loading"
        :error="formError"
        :show-role-field="mode === 'edit' || mode === 'create'"
        @submit="handleSubmit"
      >
        <template #footer>
          <DialogFooter class="pt-4">
            <Button variant="outline" type="button" @click="open = false" :disabled="loading">
              Cancel
            </Button>
            <Button type="submit" :disabled="loading">
              {{ loading ? 'Saving...' : 'Save Changes' }}
            </Button>
          </DialogFooter>
        </template>
      </UserForm>
    </DialogContent>
  </Dialog>
</template>
