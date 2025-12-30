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