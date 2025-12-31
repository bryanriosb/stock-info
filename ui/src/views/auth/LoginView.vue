<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { useForm } from '@tanstack/vue-form'
import { useAuthStore } from '@/stores/auth.store'
import { loginSchema } from '@/lib/validations/auth.schema'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Field, FieldLabel, FieldError, FieldGroup } from '@/components/ui/field'
import {
  Card, CardContent, CardDescription, CardHeader, CardTitle,
} from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { getErrors, isInvalid } from '@/lib/form'
import Logo from '@/components/Logo.vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const form = useForm({
  defaultValues: {
    username: '',
    password: '',
  },
  validators: {
    onSubmit: loginSchema,
  },
  onSubmit: async ({ value }) => {
    const success = await authStore.login(value)
    if (success) {
      const redirect = (route.query.redirect as string) || '/stocks'
      router.push(redirect)
    }
  },
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background p-4">
    <Card class="w-full max-w-md">
      <CardHeader class="text-center">
        <Logo />
        <CardTitle class="text-2xl">Welcome back</CardTitle>
        <CardDescription>Sign in to your Stock Info account</CardDescription>
      </CardHeader>
      <CardContent>
        <form id="login-form" @submit.prevent="form.handleSubmit">
          <FieldGroup>
            <Alert v-if="authStore.error" variant="destructive">
              <AlertDescription>{{ authStore.error }}</AlertDescription>
            </Alert>

            <form.Field name="username">
              <template #default="{ field }">
                <Field :data-invalid="isInvalid(field)">
                  <FieldLabel :for="field.name">Username</FieldLabel>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    :model-value="field.state.value"
                    :aria-invalid="isInvalid(field)"
                    placeholder="Enter your username"
                    autocomplete="username"
                    @blur="field.handleBlur"
                    @input="field.handleChange(($event.target as HTMLInputElement).value)"
                  />
                  <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
                </Field>
              </template>
            </form.Field>

            <form.Field name="password">
              <template #default="{ field }">
                <Field :data-invalid="isInvalid(field)">
                  <FieldLabel :for="field.name">Password</FieldLabel>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="password"
                    :model-value="field.state.value"
                    :aria-invalid="isInvalid(field)"
                    placeholder="Enter your password"
                    autocomplete="current-password"
                    @blur="field.handleBlur"
                    @input="field.handleChange(($event.target as HTMLInputElement).value)"
                  />
                  <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
                </Field>
              </template>
            </form.Field>

            <Button type="submit" class="w-full" :disabled="authStore.loading">
              <Loader2 v-if="authStore.loading" class="mr-2 h-4 w-4 animate-spin" />
              {{ authStore.loading ? 'Signing in...' : 'Sign in' }}
            </Button>

            <p class="text-center text-sm text-muted-foreground">
              Don't have an account?
              <router-link to="/register" class="text-primary hover:underline">Register</router-link>
            </p>
          </FieldGroup>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
