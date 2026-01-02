<script setup lang="ts">
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Pencil, Shield, Trash2, User as UserIcon, Users } from 'lucide-vue-next'
import type { User } from '@/types/user.types'

interface Props {
  users: User[]
  loading: boolean
  adminCount: number
}

defineProps<Props>()

const emit = defineEmits<{
  edit: [user: User]
  delete: [user: User]
}>()

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return '-'
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function isLastAdmin(user: User, adminCount: number): boolean {
  return user.role === 'admin' && adminCount === 1
}
</script>

<template>
  <div v-if="loading" class="p-6 space-y-4">
    <Skeleton v-for="i in 5" :key="i" class="h-12 w-full" />
  </div>
  <Table v-else>
    <TableHeader>
      <TableRow>
        <TableHead>ID</TableHead>
        <TableHead>Username</TableHead>
        <TableHead>Email</TableHead>
        <TableHead>Role</TableHead>
        <TableHead>Created</TableHead>
        <TableHead class="text-right">Actions</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-if="users.length === 0">
        <TableCell colspan="6" class="text-center py-12">
          <Users class="mx-auto h-12 w-12 text-muted-foreground" />
          <p class="mt-2 text-muted-foreground">No users found</p>
        </TableCell>
      </TableRow>
      <TableRow v-for="user in users" :key="user.id">
        <TableCell class="font-mono text-muted-foreground">{{ user.id }}</TableCell>
        <TableCell class="font-medium">{{ user.username }}</TableCell>
        <TableCell>{{ user.email }}</TableCell>
        <TableCell>
          <Badge
            :variant="user.role === 'admin' ? 'default' : 'secondary'"
            class="gap-1"
          >
            <Shield v-if="user.role === 'admin'" class="h-3 w-3" />
            <UserIcon v-else class="h-3 w-3" />
            {{ user.role }}
          </Badge>
        </TableCell>
        <TableCell class="text-muted-foreground">{{ formatDate(user.created_at) }}</TableCell>
        <TableCell class="text-right space-x-1">
          <Button variant="ghost" size="icon" @click="emit('edit', user)">
            <Pencil class="h-4 w-4" />
          </Button>
          <Button
            variant="ghost"
            size="icon"
            @click="emit('delete', user)"
            :disabled="isLastAdmin(user, adminCount)"
            :title="isLastAdmin(user, adminCount) ? 'Cannot delete the last admin' : 'Delete user'"
          >
            <Trash2 class="h-4 w-4 text-destructive" />
          </Button>
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
