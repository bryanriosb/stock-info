# Plan de Desarrollo UI - Stock Info

## Resumen del Proyecto

Cliente Vue 3 para el servicio backend de Stock Info. La UI permitirá a los usuarios:
- Registrarse y autenticarse
- Visualizar, buscar y ordenar acciones (stocks)
- Sincronizar datos desde la API externa
- Ver recomendaciones de inversión basadas en scoring

**Stack Tecnológico:**
- Vue 3 (Composition API)
- TypeScript
- Pinia (State Management)
- Tailwind CSS
- Vue Router
- **shadcn-vue** (Componentes UI)

---

## Paleta de Colores: "Electric Sunset" - Fintech 2025

Una paleta vanguardista que rompe con los tradicionales azul/verde/negro de fintech:

| Rol | Nombre | Hex | HSL | Uso |
|-----|--------|-----|-----|-----|
| **Primary** | Coral Eléctrico | `#FF6B5B` | `6 100% 68%` | CTAs, acciones principales, brand |
| **Primary Foreground** | Blanco | `#FFFFFF` | `0 0% 100%` | Texto sobre primary |
| **Secondary** | Teal Moderno | `#14B8A6` | `173 80% 40%` | Acciones secundarias, éxito |
| **Accent** | Amber Cálido | `#F59E0B` | `38 92% 50%` | Highlights, badges, warnings |
| **Background** | Gris Perla | `#F8FAFC` | `210 40% 98%` | Fondo principal |
| **Foreground** | Slate Profundo | `#0F172A` | `222 47% 11%` | Texto principal |
| **Success** | Esmeralda | `#10B981` | `160 84% 39%` | Estados positivos, ganancias |
| **Destructive** | Rojo Intenso | `#EF4444` | `0 84% 60%` | Errores, pérdidas |

### Concepto de Diseño
- **Coral como protagonista**: Inspirado en fintechs disruptivas como Monzo
- **Teal como complemento**: Balance entre energía y confianza
- **Gradientes sutiles**: `from-coral-500 to-amber-400` para CTAs premium

---

## Estructura del Plan

El plan está dividido en **6 fases** independientes (~10K tokens cada una).

---

# FASE 1: Setup del Proyecto (~10K tokens)

**Objetivo:** Proyecto Vue 3 configurado con shadcn-vue y paleta "Electric Sunset".

## 1.1 Inicialización

```bash
# Crear proyecto
npm create vite@latest . -- --template vue-ts

# Dependencias core
npm install vue-router@4 pinia @vueuse/core axios class-variance-authority clsx tailwind-merge lucide-vue-next

# Tailwind CSS
npm install -D tailwindcss postcss autoprefixer tailwindcss-animate
npx tailwindcss init -p

# shadcn-vue
npx shadcn-vue@latest init

# Componentes shadcn-vue
npx shadcn-vue@latest add button input card table badge alert dialog dropdown-menu select toast skeleton separator avatar progress
```

## 1.2 Estructura de Carpetas

```
ui/src/
├── api/                 # Clientes HTTP
├── assets/index.css     # Tailwind + CSS variables
├── components/
│   ├── layout/          # AppLayout, etc.
│   └── ui/              # shadcn-vue components
├── lib/utils.ts         # cn() function
├── router/index.ts
├── stores/              # Pinia stores
├── types/               # TypeScript types
├── views/               # Pages
└── main.ts
```

## 1.3 tailwind.config.js

```javascript
const animate = require("tailwindcss-animate")

module.exports = {
  darkMode: ["class"],
  content: ['./src/**/*.{vue,js,ts}'],
  theme: {
    container: { center: true, padding: "2rem" },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
        success: {
          DEFAULT: "hsl(var(--success))",
          foreground: "hsl(var(--success-foreground))",
        },
        coral: {
          500: '#FF6B5B',
          600: '#e85a4a',
        },
        teal: {
          500: '#14B8A6',
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
    },
  },
  plugins: [animate],
}
```

## 1.4 src/assets/index.css (Paleta "Electric Sunset")

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  :root {
    --background: 210 40% 98%;
    --foreground: 222 47% 11%;
    --card: 0 0% 100%;
    --card-foreground: 215 25% 27%;
    --popover: 0 0% 100%;
    --popover-foreground: 222 47% 11%;
    --primary: 6 100% 68%;            /* #FF6B5B Coral */
    --primary-foreground: 0 0% 100%;
    --secondary: 173 80% 40%;         /* #14B8A6 Teal */
    --secondary-foreground: 0 0% 100%;
    --muted: 210 40% 96%;
    --muted-foreground: 215 16% 47%;
    --accent: 38 92% 50%;             /* #F59E0B Amber */
    --accent-foreground: 222 47% 11%;
    --destructive: 0 84% 60%;
    --destructive-foreground: 0 0% 100%;
    --success: 160 84% 39%;           /* #10B981 Esmeralda */
    --success-foreground: 0 0% 100%;
    --border: 214 32% 91%;
    --input: 214 32% 91%;
    --ring: 6 100% 73%;
    --radius: 0.5rem;
  }
}

@layer base {
  * { @apply border-border; }
  body { @apply bg-background text-foreground; }
}

@layer utilities {
  .gradient-coral { @apply bg-gradient-to-r from-coral-500 to-amber-400; }
}
```

## 1.5 Tipos Base

### src/types/api.types.ts
```typescript
export interface ApiResponse<T> {
  success: boolean
  data: T
  error?: string
  meta?: PaginationMeta
}

export interface PaginationMeta {
  page: number
  limit: number
  total: number
  total_pages: number
}
```

### src/types/auth.types.ts
```typescript
export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  expires_in: number
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}
```

### src/types/user.types.ts
```typescript
export interface User {
  id: number
  username: string
  email: string
  created_at: string
  updated_at: string
}
```

### src/types/stock.types.ts
```typescript
export interface Stock {
  id: number
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: number
  target_to: number
  created_at: string
  updated_at: string
}

export interface StockQueryParams {
  page?: number
  limit?: number
  sort_by?: 'id' | 'ticker' | 'company' | 'target_to' | 'created_at'
  sort_dir?: 'asc' | 'desc'
  ticker?: string
  company?: string
}

export interface StockRecommendation {
  stock: Stock
  score: number
  reason: string
  potential_gain_percent: number
}
```

## 1.6 Cliente HTTP

### src/api/axios.ts
```typescript
import axios from 'axios'

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  headers: { 'Content-Type': 'application/json' },
})

apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default apiClient
```

## 1.7 Verificación
- [ ] `npm run dev` inicia sin errores
- [ ] Componentes shadcn-vue disponibles en `src/components/ui/`
- [ ] CSS variables aplicadas correctamente

---

# FASE 2: Autenticación (~10K tokens)

**Objetivo:** Login, registro y persistencia de sesión.

## 2.1 API de Auth

### src/api/auth.api.ts
```typescript
import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { LoginRequest, LoginResponse, RegisterRequest } from '@/types/auth.types'
import type { User } from '@/types/user.types'

export const authApi = {
  login: (data: LoginRequest) =>
    apiClient.post<ApiResponse<LoginResponse>>('/auth/login', data),
  register: (data: RegisterRequest) =>
    apiClient.post<ApiResponse<User>>('/users', data),
}
```

## 2.2 Store de Auth

### src/stores/auth.store.ts
```typescript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth.api'
import type { LoginRequest, RegisterRequest } from '@/types/auth.types'
import type { User } from '@/types/user.types'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  async function login(credentials: LoginRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.login(credentials)
      if (response.data.success) {
        token.value = response.data.data.token
        localStorage.setItem('token', response.data.data.token)
        return true
      }
      error.value = response.data.error || 'Login failed'
      return false
    } catch (err: any) {
      error.value = err.response?.data?.error || 'An error occurred'
      return false
    } finally {
      loading.value = false
    }
  }

  async function register(data: RegisterRequest) {
    loading.value = true
    error.value = null
    try {
      const response = await authApi.register(data)
      if (response.data.success) {
        user.value = response.data.data
        return true
      }
      error.value = response.data.error || 'Registration failed'
      return false
    } catch (err: any) {
      error.value = err.response?.data?.error || 'An error occurred'
      return false
    } finally {
      loading.value = false
    }
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  function checkAuth() {
    const storedToken = localStorage.getItem('token')
    if (storedToken) token.value = storedToken
  }

  return { token, user, loading, error, isAuthenticated, login, register, logout, checkAuth }
})
```

## 2.3 Router

### src/router/index.ts
```typescript
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/stocks' },
    { path: '/login', name: 'Login', component: () => import('@/views/auth/LoginView.vue'), meta: { requiresGuest: true } },
    { path: '/register', name: 'Register', component: () => import('@/views/auth/RegisterView.vue'), meta: { requiresGuest: true } },
    { path: '/stocks', name: 'Stocks', component: () => import('@/views/stocks/StocksListView.vue'), meta: { requiresAuth: true } },
    { path: '/stocks/:id', name: 'StockDetail', component: () => import('@/views/stocks/StockDetailView.vue'), meta: { requiresAuth: true } },
    { path: '/recommendations', name: 'Recommendations', component: () => import('@/views/recommendations/RecommendationsView.vue'), meta: { requiresAuth: true } },
    { path: '/users', name: 'Users', component: () => import('@/views/users/UsersListView.vue'), meta: { requiresAuth: true } },
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: () => import('@/views/NotFoundView.vue') },
  ],
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'Login', query: { redirect: to.fullPath } })
  } else if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next({ name: 'Stocks' })
  } else {
    next()
  }
})

export default router
```

## 2.4 Vista Login

### src/views/auth/LoginView.vue
```vue
<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { TrendingUp, Loader2 } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const form = ref({ username: '', password: '' })

async function handleSubmit() {
  const success = await authStore.login(form.value)
  if (success) {
    const redirect = route.query.redirect as string || '/stocks'
    router.push(redirect)
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background p-4">
    <Card class="w-full max-w-md">
      <CardHeader class="text-center">
        <div class="mx-auto mb-4 h-12 w-12 rounded-xl gradient-coral flex items-center justify-center">
          <TrendingUp class="h-6 w-6 text-white" />
        </div>
        <CardTitle class="text-2xl">Welcome back</CardTitle>
        <CardDescription>Sign in to your Stock Info account</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <Alert v-if="authStore.error" variant="destructive">
            <AlertDescription>{{ authStore.error }}</AlertDescription>
          </Alert>

          <div class="space-y-2">
            <label class="text-sm font-medium">Username</label>
            <Input v-model="form.username" placeholder="Enter your username" required />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium">Password</label>
            <Input v-model="form.password" type="password" placeholder="Enter your password" required />
          </div>

          <Button type="submit" class="w-full" :disabled="authStore.loading">
            <Loader2 v-if="authStore.loading" class="mr-2 h-4 w-4 animate-spin" />
            {{ authStore.loading ? 'Signing in...' : 'Sign in' }}
          </Button>

          <p class="text-center text-sm text-muted-foreground">
            Don't have an account?
            <router-link to="/register" class="text-primary hover:underline">Register</router-link>
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
```

## 2.5 Vista Register

### src/views/auth/RegisterView.vue
```vue
<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { TrendingUp, Loader2 } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({ username: '', email: '', password: '', confirmPassword: '' })
const validationError = ref('')

async function handleSubmit() {
  validationError.value = ''
  if (form.value.password !== form.value.confirmPassword) {
    validationError.value = 'Passwords do not match'
    return
  }

  const success = await authStore.register({
    username: form.value.username,
    email: form.value.email,
    password: form.value.password,
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
        <div class="mx-auto mb-4 h-12 w-12 rounded-xl gradient-coral flex items-center justify-center">
          <TrendingUp class="h-6 w-6 text-white" />
        </div>
        <CardTitle class="text-2xl">Create account</CardTitle>
        <CardDescription>Get started with Stock Info</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <Alert v-if="authStore.error || validationError" variant="destructive">
            <AlertDescription>{{ validationError || authStore.error }}</AlertDescription>
          </Alert>

          <div class="space-y-2">
            <label class="text-sm font-medium">Username</label>
            <Input v-model="form.username" placeholder="Choose a username" required />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium">Email</label>
            <Input v-model="form.email" type="email" placeholder="Enter your email" required />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium">Password</label>
            <Input v-model="form.password" type="password" placeholder="Create a password" required minlength="6" />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium">Confirm Password</label>
            <Input v-model="form.confirmPassword" type="password" placeholder="Confirm your password" required />
          </div>

          <Button type="submit" class="w-full" :disabled="authStore.loading">
            <Loader2 v-if="authStore.loading" class="mr-2 h-4 w-4 animate-spin" />
            {{ authStore.loading ? 'Creating account...' : 'Create account' }}
          </Button>

          <p class="text-center text-sm text-muted-foreground">
            Already have an account?
            <router-link to="/login" class="text-primary hover:underline">Sign in</router-link>
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
```

## 2.6 Verificación
- [ ] Registro crea usuario y redirige a login
- [ ] Login persiste token y redirige a stocks
- [ ] Rutas protegidas redirigen a login
- [ ] Logout elimina token

---

# FASE 3: Layout y Componentes Base (~8K tokens)

**Objetivo:** Layout con navegación y componentes reutilizables.

## 3.1 Layout Principal

### src/components/layout/AppLayout.vue
```vue
<script setup lang="ts">
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { Button } from '@/components/ui/button'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import {
  DropdownMenu, DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import { TrendingUp, BarChart3, Users, LogOut, Menu, X } from 'lucide-vue-next'

const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()
const mobileMenuOpen = ref(false)

const navigation = [
  { name: 'Stocks', href: '/stocks', icon: TrendingUp },
  { name: 'Recommendations', href: '/recommendations', icon: BarChart3 },
  { name: 'Users', href: '/users', icon: Users },
]

function handleLogout() {
  authStore.logout()
  router.push('/login')
}

function isActive(href: string) {
  return route.path.startsWith(href)
}
</script>

<template>
  <div class="min-h-screen bg-background">
    <nav class="sticky top-0 z-50 w-full border-b bg-card/95 backdrop-blur">
      <div class="container flex h-16 items-center">
        <div class="mr-8 flex items-center space-x-2">
          <div class="h-8 w-8 rounded-lg gradient-coral flex items-center justify-center">
            <TrendingUp class="h-5 w-5 text-white" />
          </div>
          <span class="text-xl font-bold">StockInfo</span>
        </div>

        <div class="hidden md:flex md:flex-1 md:space-x-1">
          <router-link
            v-for="item in navigation"
            :key="item.name"
            :to="item.href"
            class="flex items-center gap-2 px-4 py-2 text-sm font-medium rounded-md transition-colors"
            :class="isActive(item.href) ? 'bg-primary/10 text-primary' : 'text-muted-foreground hover:bg-muted'"
          >
            <component :is="item.icon" class="h-4 w-4" />
            {{ item.name }}
          </router-link>
        </div>

        <div class="ml-auto flex items-center space-x-4">
          <DropdownMenu>
            <DropdownMenuTrigger as-child>
              <Button variant="ghost" class="h-9 w-9 rounded-full">
                <Avatar class="h-9 w-9">
                  <AvatarFallback class="bg-primary text-primary-foreground">U</AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem @click="handleLogout" class="text-destructive cursor-pointer">
                <LogOut class="mr-2 h-4 w-4" />
                Log out
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>

          <Button variant="ghost" size="icon" class="md:hidden" @click="mobileMenuOpen = !mobileMenuOpen">
            <Menu v-if="!mobileMenuOpen" class="h-5 w-5" />
            <X v-else class="h-5 w-5" />
          </Button>
        </div>
      </div>

      <div v-if="mobileMenuOpen" class="md:hidden border-t">
        <div class="container py-4 space-y-1">
          <router-link
            v-for="item in navigation"
            :key="item.name"
            :to="item.href"
            @click="mobileMenuOpen = false"
            class="flex items-center gap-3 px-4 py-3 text-sm font-medium rounded-md"
            :class="isActive(item.href) ? 'bg-primary/10 text-primary' : 'text-muted-foreground'"
          >
            <component :is="item.icon" class="h-5 w-5" />
            {{ item.name }}
          </router-link>
        </div>
      </div>
    </nav>

    <main class="container py-6">
      <slot />
    </main>
  </div>
</template>
```

## 3.2 Componente RatingBadge

### src/components/RatingBadge.vue
```vue
<script setup lang="ts">
import { computed } from 'vue'
import { Badge } from '@/components/ui/badge'
import { TrendingUp, TrendingDown, Minus } from 'lucide-vue-next'

const props = defineProps<{ rating: string }>()

const config = computed(() => {
  const lower = props.rating.toLowerCase()
  if (lower.includes('buy') || lower.includes('outperform')) {
    return { class: 'bg-success/10 text-success border-success/20', icon: TrendingUp }
  }
  if (lower.includes('sell') || lower.includes('underperform')) {
    return { class: 'bg-destructive/10 text-destructive border-destructive/20', icon: TrendingDown }
  }
  return { class: 'bg-accent/10 text-accent-foreground border-accent/20', icon: Minus }
})
</script>

<template>
  <Badge variant="outline" :class="config.class" class="gap-1">
    <component :is="config.icon" class="h-3 w-3" />
    {{ rating }}
  </Badge>
</template>
```

## 3.3 Componente Pagination

### src/components/Pagination.vue
```vue
<script setup lang="ts">
import { computed } from 'vue'
import { Button } from '@/components/ui/button'
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'

const props = defineProps<{
  currentPage: number
  totalPages: number
  totalItems: number
  limit: number
}>()

const emit = defineEmits<{ 'page-change': [page: number] }>()

const startItem = computed(() => (props.currentPage - 1) * props.limit + 1)
const endItem = computed(() => Math.min(props.currentPage * props.limit, props.totalItems))
</script>

<template>
  <div class="flex items-center justify-between px-4 py-4 border-t">
    <p class="text-sm text-muted-foreground">
      Showing {{ startItem }} to {{ endItem }} of {{ totalItems }}
    </p>
    <div class="flex gap-1">
      <Button variant="outline" size="icon" :disabled="currentPage === 1" @click="emit('page-change', currentPage - 1)">
        <ChevronLeft class="h-4 w-4" />
      </Button>
      <Button variant="outline" size="icon" :disabled="currentPage === totalPages" @click="emit('page-change', currentPage + 1)">
        <ChevronRight class="h-4 w-4" />
      </Button>
    </div>
  </div>
</template>
```

## 3.4 App.vue

### src/App.vue
```vue
<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import AppLayout from '@/components/layout/AppLayout.vue'
import { Toaster } from '@/components/ui/toast'

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
```

## 3.5 Verificación
- [ ] Layout con navegación funciona
- [ ] Menú móvil abre/cierra
- [ ] Dropdown de usuario funciona
- [ ] RatingBadge muestra colores correctos

---

# FASE 4: Módulo de Stocks (~12K tokens)

**Objetivo:** Lista, detalle, filtros, paginación y sincronización.

## 4.1 API y Store de Stocks

### src/api/stocks.api.ts
```typescript
import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { Stock, StockQueryParams } from '@/types/stock.types'

export const stocksApi = {
  getAll: (params?: StockQueryParams) => apiClient.get<ApiResponse<Stock[]>>('/stocks', { params }),
  getById: (id: number) => apiClient.get<ApiResponse<Stock>>(`/stocks/${id}`),
  sync: () => apiClient.post<ApiResponse<{ message: string; count: number }>>('/stocks/sync'),
}
```

### src/stores/stocks.store.ts
```typescript
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { stocksApi } from '@/api/stocks.api'
import type { Stock, StockQueryParams } from '@/types/stock.types'
import type { PaginationMeta } from '@/types/api.types'

export const useStocksStore = defineStore('stocks', () => {
  const stocks = ref<Stock[]>([])
  const currentStock = ref<Stock | null>(null)
  const loading = ref(false)
  const syncing = ref(false)
  const error = ref<string | null>(null)
  const meta = ref<PaginationMeta | null>(null)
  const queryParams = ref<StockQueryParams>({ page: 1, limit: 20, sort_by: 'id', sort_dir: 'asc' })

  async function fetchStocks(params?: StockQueryParams) {
    loading.value = true
    error.value = null
    const merged = { ...queryParams.value, ...params }
    queryParams.value = merged
    try {
      const res = await stocksApi.getAll(merged)
      if (res.data.success) {
        stocks.value = res.data.data
        meta.value = res.data.meta || null
      } else {
        error.value = res.data.error || 'Failed to fetch'
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
    } finally {
      loading.value = false
    }
  }

  async function fetchStockById(id: number) {
    loading.value = true
    error.value = null
    currentStock.value = null
    try {
      const res = await stocksApi.getById(id)
      if (res.data.success) currentStock.value = res.data.data
      else error.value = res.data.error || 'Not found'
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
    } finally {
      loading.value = false
    }
  }

  async function syncStocks() {
    syncing.value = true
    error.value = null
    try {
      const res = await stocksApi.sync()
      if (res.data.success) {
        await fetchStocks()
        return res.data.data.count
      }
      error.value = res.data.error || 'Sync failed'
      return 0
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
      return 0
    } finally {
      syncing.value = false
    }
  }

  function setSort(field: string) {
    if (queryParams.value.sort_by === field) {
      queryParams.value.sort_dir = queryParams.value.sort_dir === 'asc' ? 'desc' : 'asc'
    } else {
      queryParams.value.sort_by = field as any
      queryParams.value.sort_dir = 'asc'
    }
    fetchStocks()
  }

  function setPage(page: number) {
    queryParams.value.page = page
    fetchStocks()
  }

  function setFilters(filters: { ticker?: string; company?: string }) {
    queryParams.value = { ...queryParams.value, ...filters, page: 1 }
    fetchStocks()
  }

  function clearFilters() {
    queryParams.value = { page: 1, limit: 20, sort_by: 'id', sort_dir: 'asc' }
    fetchStocks()
  }

  return { stocks, currentStock, loading, syncing, error, meta, queryParams, fetchStocks, fetchStockById, syncStocks, setSort, setPage, setFilters, clearFilters }
})
```

## 4.2 Vista Lista de Stocks

### src/views/stocks/StocksListView.vue
```vue
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useStocksStore } from '@/stores/stocks.store'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import Pagination from '@/components/Pagination.vue'
import RatingBadge from '@/components/RatingBadge.vue'
import { RefreshCw, Search, X, ArrowUpDown, ArrowUp, ArrowDown } from 'lucide-vue-next'
import type { Stock } from '@/types/stock.types'

const router = useRouter()
const store = useStocksStore()
const { toast } = useToast()
const filters = ref({ ticker: '', company: '' })

onMounted(() => store.fetchStocks())

function handleSort(field: string) { store.setSort(field) }
function handlePageChange(page: number) { store.setPage(page) }
function handleSearch() { store.setFilters(filters.value) }
function handleClear() { filters.value = { ticker: '', company: '' }; store.clearFilters() }
function handleRowClick(stock: Stock) { router.push(`/stocks/${stock.id}`) }

async function handleSync() {
  const count = await store.syncStocks()
  if (count > 0) toast({ title: 'Sync Complete', description: `Synced ${count} stocks` })
  else if (store.error) toast({ title: 'Sync Failed', description: store.error, variant: 'destructive' })
}

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function getSortIcon(field: string) {
  if (store.queryParams.sort_by !== field) return ArrowUpDown
  return store.queryParams.sort_dir === 'asc' ? ArrowUp : ArrowDown
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Stocks</h1>
        <p class="text-muted-foreground">Manage and analyze stock recommendations</p>
      </div>
      <Button @click="handleSync" :disabled="store.syncing" class="gradient-coral text-white">
        <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': store.syncing }" />
        {{ store.syncing ? 'Syncing...' : 'Sync from API' }}
      </Button>
    </div>

    <Card>
      <CardHeader><CardTitle class="text-lg">Filters</CardTitle></CardHeader>
      <CardContent>
        <form @submit.prevent="handleSearch" class="grid grid-cols-1 md:grid-cols-4 gap-4">
          <div class="space-y-2">
            <label class="text-sm font-medium">Ticker</label>
            <Input v-model="filters.ticker" placeholder="e.g., AAPL" />
          </div>
          <div class="space-y-2">
            <label class="text-sm font-medium">Company</label>
            <Input v-model="filters.company" placeholder="e.g., Apple" />
          </div>
          <div class="flex items-end gap-2 md:col-span-2">
            <Button type="submit"><Search class="mr-2 h-4 w-4" />Search</Button>
            <Button type="button" variant="outline" @click="handleClear"><X class="mr-2 h-4 w-4" />Clear</Button>
          </div>
        </form>
      </CardContent>
    </Card>

    <Alert v-if="store.error" variant="destructive"><AlertDescription>{{ store.error }}</AlertDescription></Alert>

    <Card>
      <CardContent class="p-0">
        <div v-if="store.loading" class="p-6 space-y-4">
          <Skeleton v-for="i in 5" :key="i" class="h-12 w-full" />
        </div>
        <Table v-else>
          <TableHeader>
            <TableRow>
              <TableHead class="cursor-pointer" @click="handleSort('ticker')">
                <div class="flex items-center gap-2">Ticker<component :is="getSortIcon('ticker')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead class="cursor-pointer" @click="handleSort('company')">
                <div class="flex items-center gap-2">Company<component :is="getSortIcon('company')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead>Brokerage</TableHead>
              <TableHead>Rating From</TableHead>
              <TableHead>Rating To</TableHead>
              <TableHead class="text-right">Target From</TableHead>
              <TableHead class="text-right cursor-pointer" @click="handleSort('target_to')">
                <div class="flex items-center justify-end gap-2">Target To<component :is="getSortIcon('target_to')" class="h-4 w-4" /></div>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-if="store.stocks.length === 0">
              <TableCell colspan="7" class="text-center py-12 text-muted-foreground">No stocks found</TableCell>
            </TableRow>
            <TableRow v-for="stock in store.stocks" :key="stock.id" class="cursor-pointer hover:bg-muted/50" @click="handleRowClick(stock)">
              <TableCell class="font-semibold text-primary">{{ stock.ticker }}</TableCell>
              <TableCell class="max-w-[200px] truncate">{{ stock.company }}</TableCell>
              <TableCell>{{ stock.brokerage }}</TableCell>
              <TableCell><RatingBadge :rating="stock.rating_from" /></TableCell>
              <TableCell><RatingBadge :rating="stock.rating_to" /></TableCell>
              <TableCell class="text-right font-mono">{{ formatCurrency(stock.target_from) }}</TableCell>
              <TableCell class="text-right font-mono font-semibold">{{ formatCurrency(stock.target_to) }}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
        <Pagination v-if="store.meta && store.meta.total_pages > 1" :current-page="store.meta.page" :total-pages="store.meta.total_pages" :total-items="store.meta.total" :limit="store.meta.limit" @page-change="handlePageChange" />
      </CardContent>
    </Card>
  </div>
</template>
```

## 4.3 Vista Detalle de Stock

### src/views/stocks/StockDetailView.vue
```vue
<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStocksStore } from '@/stores/stocks.store'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Separator } from '@/components/ui/separator'
import RatingBadge from '@/components/RatingBadge.vue'
import { ArrowLeft, TrendingUp, TrendingDown, ArrowRight, Calendar, DollarSign } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const store = useStocksStore()

const stockId = computed(() => Number(route.params.id))
onMounted(() => store.fetchStockById(stockId.value))

const priceChange = computed(() => {
  if (!store.currentStock) return 0
  const { target_from, target_to } = store.currentStock
  return target_from === 0 ? 0 : ((target_to - target_from) / target_from) * 100
})

const isPositive = computed(() => priceChange.value >= 0)

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function formatDate(d: string) {
  return new Intl.DateTimeFormat('en-US', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(d))
}
</script>

<template>
  <div class="space-y-6">
    <Button variant="ghost" @click="router.back()"><ArrowLeft class="mr-2 h-4 w-4" />Back</Button>

    <div v-if="store.loading" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <Card class="lg:col-span-2"><CardContent class="p-6"><Skeleton class="h-32 w-full" /></CardContent></Card>
      <Card><CardContent class="p-6"><Skeleton class="h-32 w-full" /></CardContent></Card>
    </div>

    <Alert v-else-if="store.error" variant="destructive"><AlertDescription>{{ store.error }}</AlertDescription></Alert>

    <template v-else-if="store.currentStock">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <Card class="lg:col-span-2">
          <CardHeader>
            <div class="flex items-start justify-between">
              <div>
                <CardTitle class="text-3xl">{{ store.currentStock.ticker }}</CardTitle>
                <CardDescription class="text-lg">{{ store.currentStock.company }}</CardDescription>
              </div>
              <div class="text-right">
                <p class="text-sm text-muted-foreground">Target Price</p>
                <p class="text-3xl font-bold" :class="isPositive ? 'text-success' : 'text-destructive'">
                  {{ formatCurrency(store.currentStock.target_to) }}
                </p>
                <div class="flex items-center justify-end gap-1" :class="isPositive ? 'text-success' : 'text-destructive'">
                  <component :is="isPositive ? TrendingUp : TrendingDown" class="h-4 w-4" />
                  <span>{{ isPositive ? '+' : '' }}{{ priceChange.toFixed(2) }}%</span>
                </div>
              </div>
            </div>
          </CardHeader>
          <CardContent>
            <div class="grid grid-cols-2 gap-6">
              <div><p class="text-sm text-muted-foreground">Brokerage</p><p class="font-medium">{{ store.currentStock.brokerage }}</p></div>
              <div><p class="text-sm text-muted-foreground">Action</p><p class="font-medium">{{ store.currentStock.action }}</p></div>
              <div><p class="text-sm text-muted-foreground flex items-center gap-1"><DollarSign class="h-4 w-4" />Previous</p><p class="font-mono">{{ formatCurrency(store.currentStock.target_from) }}</p></div>
              <div><p class="text-sm text-muted-foreground flex items-center gap-1"><DollarSign class="h-4 w-4" />New</p><p class="font-mono">{{ formatCurrency(store.currentStock.target_to) }}</p></div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader><CardTitle>Rating Change</CardTitle></CardHeader>
          <CardContent>
            <div class="flex items-center justify-center gap-4 py-4">
              <div class="text-center space-y-2">
                <p class="text-sm text-muted-foreground">From</p>
                <RatingBadge :rating="store.currentStock.rating_from" />
              </div>
              <ArrowRight class="h-6 w-6 text-muted-foreground" />
              <div class="text-center space-y-2">
                <p class="text-sm text-muted-foreground">To</p>
                <RatingBadge :rating="store.currentStock.rating_to" />
              </div>
            </div>
            <Separator class="my-4" />
            <div class="space-y-2 text-sm text-muted-foreground">
              <p class="flex items-center gap-2"><Calendar class="h-4 w-4" />Created: {{ formatDate(store.currentStock.created_at) }}</p>
              <p class="flex items-center gap-2"><Calendar class="h-4 w-4" />Updated: {{ formatDate(store.currentStock.updated_at) }}</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </template>
  </div>
</template>
```

## 4.4 Verificación
- [ ] Lista carga correctamente
- [ ] Paginación funciona
- [ ] Filtros funcionan
- [ ] Ordenamiento funciona
- [ ] Detalle muestra información completa
- [ ] Sync funciona con toast

---

# FASE 5: Módulo de Recomendaciones (~8K tokens)

**Objetivo:** Vista de recomendaciones con scoring visual.

## 5.1 API y Store

### src/api/recommendations.api.ts
```typescript
import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { StockRecommendation } from '@/types/stock.types'

export const recommendationsApi = {
  getAll: (limit?: number) => apiClient.get<ApiResponse<StockRecommendation[]>>('/recommendations', { params: { limit } }),
}
```

### src/stores/recommendations.store.ts
```typescript
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { recommendationsApi } from '@/api/recommendations.api'
import type { StockRecommendation } from '@/types/stock.types'

export const useRecommendationsStore = defineStore('recommendations', () => {
  const recommendations = ref<StockRecommendation[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchRecommendations(limit = 10) {
    loading.value = true
    error.value = null
    try {
      const res = await recommendationsApi.getAll(limit)
      if (res.data.success) recommendations.value = res.data.data
      else error.value = res.data.error || 'Failed'
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Error'
    } finally {
      loading.value = false
    }
  }

  return { recommendations, loading, error, fetchRecommendations }
})
```

## 5.2 Vista de Recomendaciones

### src/views/recommendations/RecommendationsView.vue
```vue
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useRecommendationsStore } from '@/stores/recommendations.store'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription, CardFooter } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import RatingBadge from '@/components/RatingBadge.vue'
import { RefreshCw, TrendingUp, DollarSign, Percent, ArrowRight } from 'lucide-vue-next'

const router = useRouter()
const store = useRecommendationsStore()
const limit = ref('10')

onMounted(() => store.fetchRecommendations(Number(limit.value)))

function handleLimitChange(v: string) {
  limit.value = v
  store.fetchRecommendations(Number(v))
}

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function getScoreColor(score: number) {
  if (score >= 0.7) return 'text-success'
  if (score >= 0.4) return 'text-accent'
  return 'text-destructive'
}

function getProgressColor(score: number) {
  if (score >= 0.7) return 'bg-success'
  if (score >= 0.4) return 'bg-accent'
  return 'bg-destructive'
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Recommendations</h1>
        <p class="text-muted-foreground">AI-powered stock recommendations</p>
      </div>
      <div class="flex items-center gap-4">
        <Select :model-value="limit" @update:model-value="handleLimitChange">
          <SelectTrigger class="w-24"><SelectValue /></SelectTrigger>
          <SelectContent>
            <SelectItem value="5">5</SelectItem>
            <SelectItem value="10">10</SelectItem>
            <SelectItem value="20">20</SelectItem>
            <SelectItem value="50">50</SelectItem>
          </SelectContent>
        </Select>
        <Button @click="store.fetchRecommendations(Number(limit))" :disabled="store.loading">
          <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': store.loading }" />
          Refresh
        </Button>
      </div>
    </div>

    <Alert v-if="store.error" variant="destructive"><AlertDescription>{{ store.error }}</AlertDescription></Alert>

    <div v-if="store.loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card v-for="i in 6" :key="i"><CardContent class="p-6"><Skeleton class="h-40 w-full" /></CardContent></Card>
    </div>

    <div v-else-if="store.recommendations.length === 0" class="text-center py-12">
      <TrendingUp class="mx-auto h-12 w-12 text-muted-foreground" />
      <h3 class="mt-4 text-lg font-medium">No recommendations</h3>
      <p class="text-muted-foreground">Sync stocks first to get recommendations</p>
      <Button class="mt-4" @click="router.push('/stocks')">Go to Stocks</Button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card v-for="(rec, index) in store.recommendations" :key="rec.stock.id" class="hover:shadow-lg transition-shadow">
        <CardHeader>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2">
              <span class="text-2xl font-bold text-muted-foreground">#{{ index + 1 }}</span>
              <div>
                <CardTitle class="text-lg text-primary">{{ rec.stock.ticker }}</CardTitle>
                <CardDescription class="truncate max-w-[150px]">{{ rec.stock.company }}</CardDescription>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm text-muted-foreground">Score</p>
              <p class="text-2xl font-bold" :class="getScoreColor(rec.score)">{{ (rec.score * 100).toFixed(0) }}</p>
            </div>
          </div>
        </CardHeader>
        <CardContent class="space-y-4">
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>Recommendation Score</span>
              <span>{{ (rec.score * 100).toFixed(1) }}%</span>
            </div>
            <div class="h-2 bg-muted rounded-full overflow-hidden">
              <div class="h-full rounded-full transition-all" :class="getProgressColor(rec.score)" :style="{ width: `${rec.score * 100}%` }" />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div class="bg-muted/50 rounded-lg p-3">
              <p class="text-xs text-muted-foreground flex items-center gap-1"><DollarSign class="h-3 w-3" />Target</p>
              <p class="font-semibold">{{ formatCurrency(rec.stock.target_to) }}</p>
            </div>
            <div class="bg-muted/50 rounded-lg p-3">
              <p class="text-xs text-muted-foreground flex items-center gap-1"><Percent class="h-3 w-3" />Gain</p>
              <p class="font-semibold" :class="rec.potential_gain_percent >= 0 ? 'text-success' : 'text-destructive'">
                {{ rec.potential_gain_percent >= 0 ? '+' : '' }}{{ rec.potential_gain_percent.toFixed(2) }}%
              </p>
            </div>
          </div>

          <p class="text-sm text-muted-foreground"><span class="font-medium">Analysis:</span> {{ rec.reason }}</p>

          <div class="flex items-center justify-between">
            <RatingBadge :rating="rec.stock.rating_to" />
            <span class="text-xs text-muted-foreground">{{ rec.stock.brokerage }}</span>
          </div>
        </CardContent>
        <CardFooter>
          <Button variant="ghost" class="w-full" @click="router.push(`/stocks/${rec.stock.id}`)">
            View Details <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>
```

## 5.3 Verificación
- [ ] Recomendaciones se cargan
- [ ] Score con barra de progreso y colores
- [ ] Selector de límite funciona
- [ ] Click navega a detalle

---

# FASE 6: Módulo de Usuarios y Refinamientos (~8K tokens)

**Objetivo:** Gestión de usuarios y pulido final.

## 6.1 API de Usuarios

### src/api/users.api.ts
```typescript
import apiClient from './axios'
import type { ApiResponse } from '@/types/api.types'
import type { User } from '@/types/user.types'

export const usersApi = {
  getAll: () => apiClient.get<ApiResponse<User[]>>('/users'),
  delete: (id: number) => apiClient.delete<ApiResponse<{ message: string }>>(`/users/${id}`),
}
```

## 6.2 Vista de Usuarios

### src/views/users/UsersListView.vue
```vue
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { usersApi } from '@/api/users.api'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { RefreshCw, Trash2, Users } from 'lucide-vue-next'
import type { User } from '@/types/user.types'

const { toast } = useToast()
const users = ref<User[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const deleteDialog = ref(false)
const userToDelete = ref<User | null>(null)

async function fetchUsers() {
  loading.value = true
  error.value = null
  try {
    const res = await usersApi.getAll()
    if (res.data.success) users.value = res.data.data
    else error.value = res.data.error || 'Failed'
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Error'
  } finally {
    loading.value = false
  }
}

function confirmDelete(user: User) {
  userToDelete.value = user
  deleteDialog.value = true
}

async function handleDelete() {
  if (!userToDelete.value) return
  try {
    const res = await usersApi.delete(userToDelete.value.id)
    if (res.data.success) {
      users.value = users.value.filter(u => u.id !== userToDelete.value!.id)
      toast({ title: 'User deleted', description: 'User has been deleted successfully' })
    }
  } catch (err: any) {
    toast({ title: 'Error', description: err.response?.data?.error || 'Failed to delete', variant: 'destructive' })
  } finally {
    deleteDialog.value = false
    userToDelete.value = null
  }
}

function formatDate(d: string) {
  return new Intl.DateTimeFormat('en-US', { dateStyle: 'medium' }).format(new Date(d))
}

onMounted(fetchUsers)
</script>

<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Users</h1>
        <p class="text-muted-foreground">Manage system users</p>
      </div>
      <Button @click="fetchUsers" :disabled="loading">
        <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': loading }" />
        Refresh
      </Button>
    </div>

    <Alert v-if="error" variant="destructive"><AlertDescription>{{ error }}</AlertDescription></Alert>

    <Card>
      <CardContent class="p-0">
        <div v-if="loading" class="p-6 space-y-4">
          <Skeleton v-for="i in 5" :key="i" class="h-12 w-full" />
        </div>
        <Table v-else>
          <TableHeader>
            <TableRow>
              <TableHead>ID</TableHead>
              <TableHead>Username</TableHead>
              <TableHead>Email</TableHead>
              <TableHead>Created</TableHead>
              <TableHead class="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-if="users.length === 0">
              <TableCell colspan="5" class="text-center py-12">
                <Users class="mx-auto h-12 w-12 text-muted-foreground" />
                <p class="mt-2 text-muted-foreground">No users found</p>
              </TableCell>
            </TableRow>
            <TableRow v-for="user in users" :key="user.id">
              <TableCell class="font-mono">{{ user.id }}</TableCell>
              <TableCell class="font-medium">{{ user.username }}</TableCell>
              <TableCell>{{ user.email }}</TableCell>
              <TableCell>{{ formatDate(user.created_at) }}</TableCell>
              <TableCell class="text-right">
                <Button variant="ghost" size="icon" @click="confirmDelete(user)">
                  <Trash2 class="h-4 w-4 text-destructive" />
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>

    <Dialog v-model:open="deleteDialog">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Delete User</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete <span class="font-semibold">{{ userToDelete?.username }}</span>?
            This action cannot be undone.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button variant="outline" @click="deleteDialog = false">Cancel</Button>
          <Button variant="destructive" @click="handleDelete">Delete</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
</template>
```

## 6.3 Vista 404

### src/views/NotFoundView.vue
```vue
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
        <Home class="mr-2 h-4 w-4" />Go Home
      </Button>
    </div>
  </div>
</template>
```

## 6.4 main.ts

### src/main.ts
```typescript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/index.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
```

## 6.5 Verificación Final
- [ ] CRUD de usuarios funciona
- [ ] Diálogo de confirmación funciona
- [ ] Página 404 se muestra
- [ ] No hay errores en consola
- [ ] `npm run build` compila sin errores

---

# Resumen de Endpoints

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | /api/v1/auth/login | Login |
| POST | /api/v1/users | Registro |
| GET | /api/v1/users | Listar usuarios |
| DELETE | /api/v1/users/:id | Eliminar usuario |
| GET | /api/v1/stocks | Listar stocks |
| GET | /api/v1/stocks/:id | Detalle stock |
| POST | /api/v1/stocks/sync | Sincronizar |
| GET | /api/v1/recommendations | Recomendaciones |

---

# Notas para Retomar

```
FASE COMPLETADA: 3
SIGUIENTE: Fase 4
PENDIENTE: []
```
