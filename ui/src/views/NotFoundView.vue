<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background">
    <div class="text-center">
      <h1 class="text-9xl font-bold text-muted">404</h1>
      <p class="text-2xl font-semibold mt-4">Page not found</p>
      <p class="text-muted-foreground mt-2">The page you're looking for doesn't exist.</p>
      <button
        class="mt-6 px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90"
        @click="router.push('/stocks')"
      >
        Go Home
      </button>
    </div>
  </div>
</template>
