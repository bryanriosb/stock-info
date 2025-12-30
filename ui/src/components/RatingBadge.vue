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