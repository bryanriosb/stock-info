<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useForm } from '@tanstack/vue-form'
import { useAuthStore } from '@/stores/auth.store'
import { registerSchema } from '@/lib/validations/auth.schema'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Field, FieldLabel, FieldError, FieldGroup } from '@/components/ui/field'
import {
  Card, CardContent, CardDescription, CardHeader, CardTitle,
} from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { getErrors, isInvalid } from '@/lib/form'

const router = useRouter()
const authStore = useAuthStore()

const form = useForm({
  defaultValues: {
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
  },
  validators: {
    onSubmit: registerSchema,
  },
  onSubmit: async ({ value }) => {
    const success = await authStore.register({
      username: value.username,
      email: value.email,
      password: value.password,
    })
    if (success) {
      router.push({ name: 'Login', query: { registered: 'true' } })
    }
  },
})
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
        <form id="register-form" @submit.prevent="form.handleSubmit">
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
                    placeholder="Choose a username"
                    autocomplete="username"
                    @blur="field.handleBlur"
                    @input="field.handleChange(($event.target as HTMLInputElement).value)"
                  />
                  <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
                </Field>
              </template>
            </form.Field>

            <form.Field name="email">
              <template #default="{ field }">
                <Field :data-invalid="isInvalid(field)">
                  <FieldLabel :for="field.name">Email</FieldLabel>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="email"
                    :model-value="field.state.value"
                    :aria-invalid="isInvalid(field)"
                    placeholder="Enter your email"
                    autocomplete="email"
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
                    placeholder="Create a password"
                    autocomplete="new-password"
                    @blur="field.handleBlur"
                    @input="field.handleChange(($event.target as HTMLInputElement).value)"
                  />
                  <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
                </Field>
              </template>
            </form.Field>

            <form.Field name="confirmPassword">
              <template #default="{ field }">
                <Field :data-invalid="isInvalid(field)">
                  <FieldLabel :for="field.name">Confirm Password</FieldLabel>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="password"
                    :model-value="field.state.value"
                    :aria-invalid="isInvalid(field)"
                    placeholder="Confirm your password"
                    autocomplete="new-password"
                    @blur="field.handleBlur"
                    @input="field.handleChange(($event.target as HTMLInputElement).value)"
                  />
                  <FieldError v-if="isInvalid(field)" :errors="getErrors(field)" />
                </Field>
              </template>
            </form.Field>

            <Button type="submit" class="w-full" :disabled="authStore.loading">
              <Loader2 v-if="authStore.loading" class="mr-2 h-4 w-4 animate-spin" />
              {{ authStore.loading ? 'Creating account...' : 'Create account' }}
            </Button>

            <p class="text-center text-sm text-muted-foreground">
              Already have an account?
              <router-link to="/login" class="text-primary hover:underline">Sign in</router-link>
            </p>
          </FieldGroup>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
