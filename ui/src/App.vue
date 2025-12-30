<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import AppLayout from '@/components/layout/AppLayout.vue'
import { Toaster } from '@/components/ui/sonner'

const route = useRoute()
const authStore = useAuthStore()

const showLayout = computed(() => {
  const publicRoutes = ['Login', 'Register', 'NotFound']
  return authStore.isAuthenticated && !publicRoutes.includes(route.name as string)
})

onMounted(() => authStore.checkAuth())
</script>

<template>
  <Toaster />
  <AppLayout v-if="showLayout"><router-view /></AppLayout>
  <router-view v-else />
</template>
