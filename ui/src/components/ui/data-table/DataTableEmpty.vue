<script setup lang="ts">
import { FileText } from 'lucide-vue-next'

defineProps<{
  icon?: any
  title?: string
  description?: string
}>()
</script>

<template>
  <div class="flex flex-col items-center justify-center space-y-2 text-muted-foreground">
    <slot name="icon">
      <component :is="icon || FileText" class="h-12 w-12 text-muted-foreground/50" />
    </slot>
    <span class="font-medium">{{ title }}</span>
    <span class="text-sm">{{ description }}</span>
  </div>
</template>
