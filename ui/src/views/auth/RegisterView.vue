<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
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
const authStore = useAuthStore();

const form = ref({
  username: "",
  email: "",
  password: "",
  confirmPassword: "",
});
const validationError = ref("");

async function handleSubmit() {
  validationError.value = "";
  if (form.value.password !== form.value.confirmPassword) {
    validationError.value = "Passwords do not match";
    return;
  }

  const success = await authStore.register({
    username: form.value.username,
    email: form.value.email,
    password: form.value.password,
  });

  if (success) {
    router.push({ name: "Login", query: { registered: "true" } });
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
        <CardTitle class="text-2xl">Create account</CardTitle>
        <CardDescription>Get started with Stock Info</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit" class="grid gap-4">
          <Alert
            v-if="authStore.error || validationError"
            variant="destructive"
          >
            <AlertDescription>{{
              validationError || authStore.error
            }}</AlertDescription>
          </Alert>

          <div>
            <label class="text-sm font-medium">Username</label>
            <Input
              v-model="form.username"
              placeholder="Choose a username"
              required
            />
          </div>

          <div>
            <label class="text-sm font-medium">Email</label>
            <Input
              v-model="form.email"
              type="email"
              placeholder="Enter your email"
              required
            />
          </div>

          <div>
            <label class="text-sm font-medium">Password</label>
            <Input
              v-model="form.password"
              type="password"
              placeholder="Create a password"
              required
              minlength="6"
            />
          </div>

          <div>
            <label class="text-sm font-medium">Confirm Password</label>
            <Input
              v-model="form.confirmPassword"
              type="password"
              placeholder="Confirm your password"
              required
            />
          </div>

          <Button type="submit" class="w-full" :disabled="authStore.loading">
            <Loader2
              v-if="authStore.loading"
              class="mr-2 h-4 w-4 animate-spin"
            />
            {{ authStore.loading ? "Creating account..." : "Create account" }}
          </Button>

          <p class="text-center text-sm text-muted-foreground">
            Already have an account?
            <router-link to="/login" class="text-primary hover:underline"
              >Sign in</router-link
            >
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
