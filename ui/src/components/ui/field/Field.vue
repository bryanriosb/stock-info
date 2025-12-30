<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  class?: string
  orientation?: 'vertical' | 'horizontal'
  dataInvalid?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  orientation: 'vertical',
})

const fieldClass = computed(() =>
  cn(
    'space-y-2',
    props.orientation === 'horizontal' && 'flex items-center gap-2 space-y-0',
    props.dataInvalid && '[&_label]:text-destructive',
    props.class
  )
)
</script>

<template>
  <div :class="fieldClass" :data-invalid="dataInvalid || undefined">
    <slot />
  </div>
</template>
