<script setup lang="ts">
import { ref, watch } from 'vue'
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
import { Loader2 } from 'lucide-vue-next'
import type { User } from '@/types/user.types'

interface Props {
  user: User | null
}

const props = defineProps<Props>()
const open = defineModel<boolean>('open', { required: true })

const emit = defineEmits<{
  deleted: [userId: number]
}>()

const { toast } = useToast()
const deleting = ref(false)

watch(open, (isOpen) => {
  if (!isOpen) {
    deleting.value = false
  }
})

async function handleDelete() {
  if (!props.user) return

  deleting.value = true
  try {
    const res = await usersApi.delete(props.user.id)
    if (res.data.success) {
      toast({ title: 'User deleted', description: 'User has been deleted successfully' })
      emit('deleted', props.user.id)
      open.value = false
    } else {
      toast({ title: 'Error', description: res.data.error || 'Failed to delete user', variant: 'destructive' })
    }
  } catch (err: any) {
    toast({ title: 'Error', description: err.response?.data?.error || 'Failed to delete user', variant: 'destructive' })
  } finally {
    deleting.value = false
  }
}
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Delete User</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete <span class="font-semibold">{{ user?.username }}</span>?
          This action cannot be undone.
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <Button variant="outline" @click="open = false" :disabled="deleting">
          Cancel
        </Button>
        <Button variant="destructive" @click="handleDelete" :disabled="deleting">
          <Loader2 v-if="deleting" class="mr-2 h-4 w-4 animate-spin" />
          {{ deleting ? 'Deleting...' : 'Delete' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
