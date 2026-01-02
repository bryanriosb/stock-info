<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { Button } from '@/components/ui/button'
import {
  Card, CardContent, CardDescription, CardHeader, CardTitle,
} from '@/components/ui/card'
import { UserForm } from '@/components/users'
import type { UserFormData } from '@/components/users/UserForm.vue'

const router = useRouter()
const authStore = useAuthStore()

async function handleSubmit(data: UserFormData) {
  const success = await authStore.register({
    username: data.username,
    email: data.email,
    password: data.password!,
  })
  if (success) {
    router.push({ name: 'Login', query: { registered: 'true' } })
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background p-4">
    <Card class="w-full max-w-md">
      <CardHeader class="text-center">
        <Logo />
        <CardTitle class="text-2xl">Create account</CardTitle>
        <CardDescription>Get started with Stock Info</CardDescription>
      </CardHeader>
      <CardContent>
        <UserForm
          mode="register"
          :loading="authStore.loading"
          :error="authStore.error"
          @submit="handleSubmit"
        >
          <template #footer>
            <Button type="submit" class="w-full" :disabled="authStore.loading">
              <Loader2 v-if="authStore.loading" class="mr-2 h-4 w-4 animate-spin" />
              {{ authStore.loading ? 'Creating account...' : 'Create account' }}
            </Button>

            <p class="text-center text-sm text-muted-foreground">
              Already have an account?
              <router-link to="/login" class="text-primary hover:underline">Sign in</router-link>
            </p>
          </template>
        </UserForm>
      </CardContent>
    </Card>
  </div>
</template>
