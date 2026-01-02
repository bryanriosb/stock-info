<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { usersApi } from '@/api/users.api'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { RefreshCw, Trash2, Users } from 'lucide-vue-next'
import type { User } from '@/types/user.types'

const { toast } = useToast()
const users = ref<User[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const deleteDialog = ref(false)
const userToDelete = ref<User | null>(null)
const deleting = ref(false)

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

function confirmDelete(user: User) {
  userToDelete.value = user
  deleteDialog.value = true
}

async function handleDelete() {
  if (!userToDelete.value) return

  deleting.value = true
  try {
    const res = await usersApi.delete(userToDelete.value.id)
    if (res.data.success) {
      users.value = users.value.filter(u => u.id !== userToDelete.value!.id)
      toast({ title: 'User deleted', description: 'User has been deleted successfully' })
    } else {
      toast({ title: 'Error', description: res.data.error || 'Failed to delete user', variant: 'destructive' })
    }
  } catch (err: any) {
    toast({ title: 'Error', description: err.response?.data?.error || 'Failed to delete user', variant: 'destructive' })
  } finally {
    deleting.value = false
    deleteDialog.value = false
    userToDelete.value = null
  }
}

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return '-'
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
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
        <div v-if="loading" class="p-6 space-y-4">
          <Skeleton v-for="i in 5" :key="i" class="h-12 w-full" />
        </div>
        <Table v-else>
          <TableHeader>
            <TableRow>
              <TableHead>ID</TableHead>
              <TableHead>Username</TableHead>
              <TableHead>Email</TableHead>
              <TableHead>Created</TableHead>
              <TableHead class="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-if="users.length === 0">
              <TableCell colspan="5" class="text-center py-12">
                <Users class="mx-auto h-12 w-12 text-muted-foreground" />
                <p class="mt-2 text-muted-foreground">No users found</p>
              </TableCell>
            </TableRow>
            <TableRow v-for="user in users" :key="user.id">
              <TableCell class="font-mono text-muted-foreground">{{ user.id }}</TableCell>
              <TableCell class="font-medium">{{ user.username }}</TableCell>
              <TableCell>{{ user.email }}</TableCell>
              <TableCell class="text-muted-foreground">{{ formatDate(user.created_at) }}</TableCell>
              <TableCell class="text-right">
                <Button variant="ghost" size="icon" @click="confirmDelete(user)">
                  <Trash2 class="h-4 w-4 text-destructive" />
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>

    <Dialog v-model:open="deleteDialog">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Delete User</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete <span class="font-semibold">{{ userToDelete?.username }}</span>?
            This action cannot be undone.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button variant="outline" @click="deleteDialog = false" :disabled="deleting">
            Cancel
          </Button>
          <Button variant="destructive" @click="handleDelete" :disabled="deleting">
            {{ deleting ? 'Deleting...' : 'Delete' }}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
</template>
