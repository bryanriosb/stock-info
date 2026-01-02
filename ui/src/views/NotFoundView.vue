<script setup lang="ts">
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Home } from 'lucide-vue-next'

const router = useRouter()
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background">
    <div class="text-center">
      <h1 class="text-9xl font-bold text-muted">404</h1>
      <p class="text-2xl font-semibold mt-4">Page not found</p>
      <p class="text-muted-foreground mt-2">The page you're looking for doesn't exist.</p>
      <Button class="mt-6" @click="router.push('/stocks')">
        <Home class="mr-2 h-4 w-4" />
        Go Home
      </Button>
    </div>
  </div>
</template>
