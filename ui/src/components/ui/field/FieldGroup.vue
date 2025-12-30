<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  class?: string
}

const props = defineProps<Props>()

const groupClass = computed(() =>
  cn('space-y-4', props.class)
)
</script>

<template>
  <div :class="groupClass">
    <slot />
  </div>
</template>
