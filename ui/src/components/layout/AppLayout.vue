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