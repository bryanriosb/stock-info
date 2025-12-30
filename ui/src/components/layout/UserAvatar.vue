<script setup lang="ts">
import { useAuthStore } from '@/stores/auth.store'
import { Button } from '@/components/ui/button'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'

const authStore = useAuthStore()
</script>

<template>
  <Button variant="ghost" class="h-9 w-9 rounded-full">
    <Avatar class="h-9 w-9">
      <AvatarFallback class="text-primary-foreground font-bold">
        {{ authStore.userInitials }}
      </AvatarFallback>
    </Avatar>
  </Button>
</template>
