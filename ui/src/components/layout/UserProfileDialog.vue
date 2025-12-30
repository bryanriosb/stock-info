<script setup lang="ts">
import { useAuthStore } from '@/stores/auth.store'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { Badge } from '@/components/ui/badge'
import {
  Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle
} from '@/components/ui/dialog'

const open = defineModel<boolean>('open', { default: false })
const authStore = useAuthStore()
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent class="sm:max-w-md">
      <DialogHeader>
        <DialogTitle>User Profile</DialogTitle>
        <DialogDescription>Your account information</DialogDescription>
      </DialogHeader>
      <div class="space-y-4 py-4">
        <div class="flex items-center gap-4">
          <Avatar class="h-16 w-16">
            <AvatarFallback class="text-primary-foreground text-xl font-bold">
              {{ authStore.userInitials }}
            </AvatarFallback>
          </Avatar>
          <div class="space-y-1">
            <p class="text-lg font-medium">{{ authStore.username }}</p>
            <Badge :variant="authStore.isAdmin ? 'default' : 'secondary'">
              {{ authStore.role }}
            </Badge>
          </div>
        </div>
        <div class="space-y-2 border-t pt-4">
          <div class="flex justify-between text-sm">
            <span class="text-muted-foreground">Username</span>
            <span class="font-medium">{{ authStore.username }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-muted-foreground">Email</span>
            <span class="font-medium">{{ authStore.email }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-muted-foreground">Role</span>
            <span class="font-medium capitalize">{{ authStore.role }}</span>
          </div>
        </div>
      </div>
    </DialogContent>
  </Dialog>
</template>
