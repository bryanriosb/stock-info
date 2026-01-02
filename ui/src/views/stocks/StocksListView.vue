<script setup lang="ts">
import { onMounted, watch, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStocksStore } from '@/stores/stocks.store'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'
import { DataTable } from '@/components/ui/data-table'
import type { Column, SortState } from '@/components/ui/data-table'
import RatingBadge from '@/components/RatingBadge.vue'
import StockFilter from '@/components/StockFilter.vue'
import { RefreshCw, Square, TrendingUp } from 'lucide-vue-next'
import type { Stock } from '@/types/stock.types'

const router = useRouter()
const store = useStocksStore()
const { toast } = useToast()

onMounted(() => store.fetchStocks())

// Watch for sync completion to show toast
watch(() => store.syncProgress, (progress, oldProgress) => {
  if (oldProgress?.status === 'saving' && progress === null) {
    toast({ title: 'Sync Complete', description: 'Stocks synced successfully' })
  }
})

watch(() => store.error, (error) => {
  if (error) {
    toast({ title: 'Error', description: error, variant: 'destructive' })
    store.error = null
  }
})

// Table columns definition
const columns: Column<Stock>[] = [
  { key: 'ticker', header: 'Ticker', sortable: true, class: 'font-semibold text-primary' },
  { key: 'company', header: 'Company', sortable: true, class: 'max-w-[200px] truncate' },
  { key: 'brokerage', header: 'Brokerage' },
  { key: 'rating_from', header: 'Rating From' },
  { key: 'rating_to', header: 'Rating To' },
  { key: 'target_from', header: 'Target From', class: 'text-right font-mono', headerClass: 'text-right' },
  { key: 'target_to', header: 'Target To', sortable: true, class: 'text-right font-mono font-semibold', headerClass: 'text-right justify-end' },
  { key: 'time', header: 'Date', sortable: true, class: 'text-muted-foreground' },
  { key: 'updated_at', header: 'Synced', class: 'text-xs text-muted-foreground' },
]

// Sort state from store
const sortState = computed<SortState | undefined>(() => {
  if (!store.queryParams.sort_by) return undefined
  return {
    field: store.queryParams.sort_by,
    direction: store.queryParams.sort_dir as 'asc' | 'desc',
  }
})

// Pagination meta from store
const paginationMeta = computed(() => store.meta ?? undefined)

function handleSort(field: string) { store.setSort(field) }
function handlePageChange(page: number) { store.setPage(page) }
function handleFilter(filters: { search?: string; rating_from?: string; rating_to?: string; ticker?: string; company?: string }) {
  store.setFilters(filters)
}
function handleClearFilters() { store.clearFilters() }
function handleRowClick(stock: Stock) { router.push(`/stocks/${stock.id}`) }

function handleSync() {
  store.syncStocks()
}

function handleCancelSync() {
  store.cancelSync()
  toast({ title: 'Sync Cancelled', description: 'Stock sync was cancelled' })
}

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return '-'
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function formatDateTime(dateStr: string) {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return '-'
  return date.toLocaleString('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Stocks</h1>
        <p class="text-muted-foreground">Analyze stock performance</p>
      </div>
      <div class="flex flex-col items-end gap-2">
        <div class="flex gap-2">
          <Button @click="handleSync" :disabled="store.syncing" class="gradient-coral text-white">
            <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': store.syncing }" />
            {{ store.syncing ? 'Syncing...' : 'Sync from API' }}
          </Button>
          <Button v-if="store.syncing" variant="outline" @click="handleCancelSync">
            <Square class="mr-2 h-4 w-4" />
            Cancel
          </Button>
        </div>
        <!-- Progress bar -->
        <div v-if="store.syncProgress" class="w-full space-y-1">
          <div class="flex justify-between text-xs text-muted-foreground">
            <span>{{ store.syncProgress.message }}</span>
            <span>{{ store.syncProgress.percent }}%</span>
          </div>
          <Progress :model-value="store.syncProgress.percent" class="h-2" />
        </div>
      </div>
    </div>

    <!-- StockFilter component - show when loaded or has active filters -->
    <Card v-if="store.hasLoadedOnce || store.hasActiveFilters">
      <CardContent class="pt-6">
        <StockFilter
          :loading="store.loading"
          @filter="handleFilter"
          @clear="handleClearFilters"
        />
      </CardContent>
    </Card>

    <Card>
      <CardContent class="p-0">
        <DataTable
          :data="store.stocks"
          :columns="columns"
          :loading="store.loading"
          :sort="sortState"
          :pagination="paginationMeta"
          :row-clickable="true"
          :empty-icon="TrendingUp"
          empty-title="No stocks found"
          empty-description="Try adjusting your filters or sync from API"
          @sort="handleSort"
          @page-change="handlePageChange"
          @row-click="handleRowClick"
        >
          <!-- Custom cell renderers -->
          <template #cell-rating_from="{ row }">
            <RatingBadge :rating="row.rating_from" />
          </template>
          <template #cell-rating_to="{ row }">
            <RatingBadge :rating="row.rating_to" />
          </template>
          <template #cell-target_from="{ row }">
            {{ formatCurrency(row.target_from) }}
          </template>
          <template #cell-target_to="{ row }">
            {{ formatCurrency(row.target_to) }}
          </template>
          <template #cell-time="{ row }">
            {{ formatDate(row.time) }}
          </template>
          <template #cell-updated_at="{ row }">
            {{ formatDateTime(row.updated_at) }}
          </template>
        </DataTable>
      </CardContent>
    </Card>
  </div>
</template>
