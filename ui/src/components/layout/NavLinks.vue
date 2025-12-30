<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { TrendingUp, BarChart3, Users } from 'lucide-vue-next'

interface Props {
  mobile?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  mobile: false
})

const emit = defineEmits<{
  navigate: []
}>()

const authStore = useAuthStore()
const route = useRoute()

const allNavigation = [
  { name: 'Stocks', href: '/stocks', icon: TrendingUp, adminOnly: false },
  { name: 'Recommendations', href: '/recommendations', icon: BarChart3, adminOnly: false },
  { name: 'Users', href: '/users', icon: Users, adminOnly: true },
]

const navigation = computed(() =>
  allNavigation.filter(item => !item.adminOnly || authStore.isAdmin)
)

function isActive(href: string) {
  return route.path.startsWith(href)
}

function handleClick() {
  if (props.mobile) {
    emit('navigate')
  }
}
</script>

<template>
  <router-link
    v-for="item in navigation"
    :key="item.name"
    :to="item.href"
    @click="handleClick"
    class="flex items-center font-medium rounded-md transition-colors"
    :class="[
      mobile ? 'gap-3 px-4 py-3 text-sm' : 'gap-2 px-4 py-2 text-sm',
      isActive(item.href)
        ? 'bg-primary/10 text-primary'
        : mobile ? 'text-muted-foreground' : 'text-muted-foreground hover:bg-muted'
    ]"
  >
    <component :is="item.icon" :class="mobile ? 'h-5 w-5' : 'h-4 w-4'" />
    {{ item.name }}
  </router-link>
</template>
