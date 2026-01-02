<script setup lang="ts">
import { ref } from 'vue'
import { Button } from '@/components/ui/button'
import NavLinks from './NavLinks.vue'
import UserMenu from './UserMenu.vue'
import UserProfileDialog from './UserProfileDialog.vue'
import Logo from '../Logo.vue'

const mobileMenuOpen = ref(false)
const profileDialogOpen = ref(false)
</script>

<template>
  <div class="min-h-screen bg-background">
    <nav class="sticky top-0 z-50 w-full border-b bg-card/95 backdrop-blur">
      <div class="container flex h-16 items-center">
        <div class="mr-8 flex items-center space-x-2">
         <Logo />
          <span class="text-xl font-bold">StockInfo</span>
        </div>

        <div class="hidden md:flex md:flex-1 md:space-x-1">
          <NavLinks />
        </div>

        <div class="ml-auto flex items-center space-x-4">
          <UserMenu @open-profile="profileDialogOpen = true" />

          <Button variant="ghost" size="icon" class="md:hidden" @click="mobileMenuOpen = !mobileMenuOpen">
            <Menu v-if="!mobileMenuOpen" class="h-5 w-5" />
            <X v-else class="h-5 w-5" />
          </Button>
        </div>
      </div>

      <div v-if="mobileMenuOpen" class="md:hidden border-t">
        <div class="container py-4 space-y-1">
          <NavLinks mobile @navigate="mobileMenuOpen = false" />
        </div>
      </div>
    </nav>

    <main class="container py-6">
      <slot />
    </main>

    <UserProfileDialog v-model:open="profileDialogOpen" />
  </div>
</template>
