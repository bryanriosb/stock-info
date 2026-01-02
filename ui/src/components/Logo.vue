<script setup lang="ts">
  import { TrendingUp } from 'lucide-vue-next'
  defineProps({
  customClass: {
    type: String,
    default: ''
  },
  iconSize: {
    type: Number,
    default: 20
  }
})
</script>

<template>
	 <div
    :class="[
      'mx-auto mb-4 h-12 w-12 rounded-xl gradient-coral flex items-center justify-center',
      customClass
    ]">
      <TrendingUp class="text-white" :size="iconSize" />
    </div>
</template>