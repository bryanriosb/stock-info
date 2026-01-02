<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { usersApi } from '@/api/users.api'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent } from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { DataTable } from '@/components/ui/data-table'
import type { Column } from '@/components/ui/data-table'
import { UserDeleteDialog, UserFormDialog } from '@/components/users'
import { RefreshCw, Pencil, Trash2, Shield, User as UserIcon, Users } from 'lucide-vue-next'
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

// Table columns definition
const columns: Column<User>[] = [
  { key: 'id', header: 'ID', class: 'font-mono text-muted-foreground' },
  { key: 'username', header: 'Username', class: 'font-medium' },
  { key: 'email', header: 'Email' },
  { key: 'role', header: 'Role' },
  { key: 'created_at', header: 'Created', class: 'text-muted-foreground' },
  { key: 'actions', header: 'Actions', headerClass: 'text-right', class: 'text-right' },
]

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

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return '-'
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function isLastAdmin(user: User): boolean {
  return user.role === 'admin' && adminCount.value === 1
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
        <DataTable
          :data="users"
          :columns="columns"
          :loading="loading"
          :empty-icon="Users"
          empty-title="No users found"
          empty-description="No users have been registered yet"
        >
          <!-- Role badge -->
          <template #cell-role="{ row }">
            <Badge
              :variant="row.role === 'admin' ? 'default' : 'secondary'"
              class="gap-1"
            >
              <Shield v-if="row.role === 'admin'" class="h-3 w-3" />
              <UserIcon v-else class="h-3 w-3" />
              {{ row.role }}
            </Badge>
          </template>

          <!-- Created date -->
          <template #cell-created_at="{ row }">
            {{ formatDate(row.created_at) }}
          </template>

          <!-- Actions -->
          <template #cell-actions="{ row }">
            <div class="space-x-1">
              <Button variant="ghost" size="icon" @click.stop="handleEditUser(row)">
                <Pencil class="h-4 w-4" />
              </Button>
              <Button
                variant="ghost"
                size="icon"
                @click.stop="handleDeleteUser(row)"
                :disabled="isLastAdmin(row)"
                :title="isLastAdmin(row) ? 'Cannot delete the last admin' : 'Delete user'"
              >
                <Trash2 class="h-4 w-4 text-destructive" />
              </Button>
            </div>
          </template>
        </DataTable>
      </CardContent>
    </Card>

    <UserDeleteDialog
      v-model:open="deleteDialogOpen"
      :user="userToDelete"
      @deleted="handleUserDeleted"
    />

    <UserFormDialog
      v-model:open="editDialogOpen"
      :user="userToEdit"
      :admin-count="adminCount"
      mode="edit"
      @updated="handleUserUpdated"
    />
  </div>
</template>
