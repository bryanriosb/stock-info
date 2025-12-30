<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

const props = defineProps<{
  class?: string
  errors?: any[]
}>()

const errorClass = computed(() =>
  cn('text-sm font-medium text-destructive', props.class)
)

const errorMessages = computed(() => {
  if (!props.errors || props.errors.length === 0) return []
  return props.errors
    .map(err => {
      if (typeof err === 'string') return err
      if (err && typeof err === 'object' && 'message' in err) {
        return (err as { message?: string }).message || ''
      }
      return ''
    })
    .filter(Boolean)
})
</script>

<template>
  <div v-if="errorMessages.length > 0 || $slots.default" :class="errorClass" role="alert">
    <slot>
      <p v-for="(message, index) in errorMessages" :key="index">
        {{ message }}
      </p>
    </slot>
  </div>
</template>
