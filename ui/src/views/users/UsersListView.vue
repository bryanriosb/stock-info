<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { usersApi } from '@/api/users.api'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { UsersTable, UserDeleteDialog, UserEditRoleDialog } from '@/components/users'
import { RefreshCw } from 'lucide-vue-next'
import type { User } from '@/types/user.types'

const users = ref<User[]>([])
const loading = ref(false)
const error = ref<string | null>(null)

// Delete dialog state
const deleteDialogOpen = ref(false)
const userToDelete = ref<User | null>(null)

// Edit dialog state
const editDialogOpen = ref(false)
const userToEdit = ref<User | null>(null)

// Computed: count admins
const adminCount = computed(() => users.value.filter(u => u.role === 'admin').length)

async function fetchUsers() {
  loading.value = true
  error.value = null
  try {
    const res = await usersApi.getAll()
    if (res.data.success) {
      users.value = res.data.data
    } else {
      error.value = res.data.error || 'Failed to fetch users'
    }
  } catch (err: any) {
    error.value = err.response?.data?.error || 'An error occurred'
  } finally {
    loading.value = false
  }
}

function handleEditUser(user: User) {
  userToEdit.value = user
  editDialogOpen.value = true
}

function handleDeleteUser(user: User) {
  userToDelete.value = user
  deleteDialogOpen.value = true
}

function handleUserUpdated(updatedUser: User) {
  const index = users.value.findIndex(u => u.id === updatedUser.id)
  if (index !== -1) {
    users.value[index] = updatedUser
  }
}

function handleUserDeleted(userId: number) {
  users.value = users.value.filter(u => u.id !== userId)
}

onMounted(fetchUsers)
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Users</h1>
        <p class="text-muted-foreground">Manage system users</p>
      </div>
      <Button @click="fetchUsers" :disabled="loading">
        <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': loading }" />
        Refresh
      </Button>
    </div>

    <Alert v-if="error" variant="destructive">
      <AlertDescription>{{ error }}</AlertDescription>
    </Alert>

    <Card>
      <CardContent class="p-0">
        <UsersTable
          :users="users"
          :loading="loading"
          :admin-count="adminCount"
          @edit="handleEditUser"
          @delete="handleDeleteUser"
        />
      </CardContent>
    </Card>

    <UserDeleteDialog
      v-model:open="deleteDialogOpen"
      :user="userToDelete"
      @deleted="handleUserDeleted"
    />

    <UserEditRoleDialog
      v-model:open="editDialogOpen"
      :user="userToEdit"
      :admin-count="adminCount"
      @updated="handleUserUpdated"
    />
  </div>
</template>
