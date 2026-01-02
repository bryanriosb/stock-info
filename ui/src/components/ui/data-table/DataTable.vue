<script setup lang="ts" generic="TData">
import { computed } from 'vue'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import DataTablePagination from './DataTablePagination.vue'
import DataTableEmpty from './DataTableEmpty.vue'
import { ArrowUpDown, ArrowUp, ArrowDown } from 'lucide-vue-next'

export interface Column<T> {
  key: string
  header: string
  sortable?: boolean
  class?: string
  headerClass?: string
  render?: (row: T) => any
}

export interface SortState {
  field: string
  direction: 'asc' | 'desc'
}

export interface PaginationMeta {
  page: number
  limit: number
  total: number
  total_pages: number
}

interface Props {
  data: TData[]
  columns: Column<TData>[]
  loading?: boolean
  sort?: SortState
  pagination?: PaginationMeta
  rowClickable?: boolean
  emptyIcon?: any
  emptyTitle?: string
  emptyDescription?: string
  skeletonRows?: number
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  rowClickable: false,
  emptyTitle: 'No results found',
  emptyDescription: 'Try adjusting your search or filters',
  skeletonRows: 5,
})

const emit = defineEmits<{
  sort: [field: string]
  'page-change': [page: number]
  'row-click': [row: TData]
}>()

const totalColumns = computed(() => props.columns.length)

function getSortIcon(field: string) {
  if (!props.sort || props.sort.field !== field) return ArrowUpDown
  return props.sort.direction === 'asc' ? ArrowUp : ArrowDown
}

function handleSort(column: Column<TData>) {
  if (column.sortable) {
    emit('sort', column.key)
  }
}

function handleRowClick(row: TData) {
  if (props.rowClickable) {
    emit('row-click', row)
  }
}

function getCellValue(row: TData, column: Column<TData>) {
  if (column.render) {
    return column.render(row)
  }
  return (row as Record<string, any>)[column.key]
}
</script>

<template>
  <div class="space-y-0">
    <!-- Loading skeleton -->
    <div v-if="loading" class="p-6 space-y-4">
      <Skeleton v-for="i in skeletonRows" :key="i" class="h-12 w-full" />
    </div>

    <!-- Table -->
    <Table v-else>
      <TableHeader>
        <TableRow>
          <TableHead
            v-for="column in columns"
            :key="column.key"
            :class="[
              column.headerClass,
              { 'cursor-pointer select-none': column.sortable }
            ]"
            @click="handleSort(column)"
          >
            <div class="flex items-center gap-2" :class="column.headerClass">
              <slot :name="`header-${column.key}`" :column="column">
                {{ column.header }}
              </slot>
              <component
                v-if="column.sortable"
                :is="getSortIcon(column.key)"
                class="h-4 w-4"
              />
            </div>
          </TableHead>
        </TableRow>
      </TableHeader>

      <TableBody>
        <!-- Empty state -->
        <TableRow v-if="data.length === 0">
          <TableCell :colspan="totalColumns" class="text-center py-12">
            <DataTableEmpty
              :icon="emptyIcon"
              :title="emptyTitle"
              :description="emptyDescription"
            >
              <template #icon>
                <slot name="empty-icon" />
              </template>
            </DataTableEmpty>
          </TableCell>
        </TableRow>

        <!-- Data rows -->
        <TableRow
          v-for="(row, index) in data"
          :key="index"
          :class="{ 'cursor-pointer hover:bg-muted/50': rowClickable }"
          @click="handleRowClick(row)"
        >
          <TableCell
            v-for="column in columns"
            :key="column.key"
            :class="column.class"
          >
            <slot :name="`cell-${column.key}`" :row="row" :value="getCellValue(row, column)">
              {{ getCellValue(row, column) }}
            </slot>
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>

    <!-- Pagination -->
    <DataTablePagination
      v-if="pagination && pagination.total_pages > 1"
      :current-page="pagination.page"
      :total-pages="pagination.total_pages"
      :total-items="pagination.total"
      :limit="pagination.limit"
      @page-change="emit('page-change', $event)"
    />
  </div>
</template>
