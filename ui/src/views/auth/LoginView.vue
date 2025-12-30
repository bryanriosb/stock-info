<script setup lang="ts">
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useAuthStore } from "@/stores/auth.store";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { TrendingUp, Loader2 } from "lucide-vue-next";

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();

const form = ref({ username: "", password: "" });

async function handleSubmit() {
  const success = await authStore.login(form.value);
  if (success) {
    const redirect = (route.query.redirect as string) || "/stocks";
    router.push(redirect);
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background p-4">
    <Card class="w-full max-w-md">
      <CardHeader class="text-center">
        <div
          class="mx-auto mb-4 h-12 w-12 rounded-xl gradient-coral flex items-center justify-center"
        >
          <TrendingUp class="h-6 w-6 text-white" />
        </div>
        <CardTitle class="text-2xl">Welcome back</CardTitle>
        <CardDescription>Sign in to your Stock Info account</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit" class="grid gap-4">
          <Alert v-if="authStore.error" variant="destructive">
            <AlertDescription>{{ authStore.error }}</AlertDescription>
          </Alert>

          <div>
            <label class="text-sm font-medium">Username</label>
            <Input
              v-model="form.username"
              placeholder="Enter your username"
              required
            />
          </div>

          <div>
            <label class="text-sm font-medium">Password</label>
            <Input
              v-model="form.password"
              type="password"
              placeholder="Enter your password"
              required
            />
          </div>

          <Button type="submit" class="w-full" :disabled="authStore.loading">
            <Loader2
              v-if="authStore.loading"
              class="mr-2 h-4 w-4 animate-spin"
            />
            {{ authStore.loading ? "Signing in..." : "Sign in" }}
          </Button>

          <p class="text-center text-sm text-muted-foreground">
            Don't have an account?
            <router-link to="/register" class="text-primary hover:underline"
              >Register</router-link
            >
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
