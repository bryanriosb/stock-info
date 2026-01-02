<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useStocksStore } from '@/stores/stocks.store'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { Progress } from '@/components/ui/progress'
import Pagination from '@/components/Pagination.vue'
import RatingBadge from '@/components/RatingBadge.vue'
import { RefreshCw, Search, X, ArrowUpDown, ArrowUp, ArrowDown, StopCircle, Square } from 'lucide-vue-next'
import type { Stock } from '@/types/stock.types'

const router = useRouter()
const store = useStocksStore()
const { toast } = useToast()

const filters = ref({ ticker: '', company: '' })

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

function handleSort(field: string) { store.setSort(field) }
function handlePageChange(page: number) { store.setPage(page) }
function handleSearch() { store.setFilters(filters.value) }
function handleClear() { filters.value = { ticker: '', company: '' }; store.clearFilters() }
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

function getSortIcon(field: string) {
  if (store.queryParams.sort_by !== field) return ArrowUpDown
  return store.queryParams.sort_dir === 'asc' ? ArrowUp : ArrowDown
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Stocks</h1>
        <p class="text-muted-foreground">Manage and analyze stock recommendations</p>
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

    <Card>
      <CardHeader><CardTitle class="text-lg">Filters</CardTitle></CardHeader>
      <CardContent>
        <form @submit.prevent="handleSearch" class="grid grid-cols-1 md:grid-cols-4 gap-4">
          <div class="space-y-2">
            <label class="text-sm font-medium">Ticker</label>
            <Input v-model="filters.ticker" placeholder="e.g., AAPL" />
          </div>
          <div class="space-y-2">
            <label class="text-sm font-medium">Company</label>
            <Input v-model="filters.company" placeholder="e.g., Apple" />
          </div>
          <div class="flex items-end gap-2 md:col-span-2">
            <Button type="submit"><Search class="mr-2 h-4 w-4" />Search</Button>
            <Button type="button" variant="outline" @click="handleClear"><X class="mr-2 h-4 w-4" />Clear</Button>
          </div>
        </form>
      </CardContent>
    </Card>

    <Card>
      <CardContent class="p-0">
        <div v-if="store.loading" class="p-6 space-y-4">
          <Skeleton v-for="i in 5" :key="i" class="h-12 w-full" />
        </div>
        <Table v-else>
          <TableHeader>
            <TableRow>
              <TableHead class="cursor-pointer" @click="handleSort('ticker')">
                <div class="flex items-center gap-2">Ticker<component :is="getSortIcon('ticker')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead class="cursor-pointer" @click="handleSort('company')">
                <div class="flex items-center gap-2">Company<component :is="getSortIcon('company')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead>Brokerage</TableHead>
              <TableHead>Rating From</TableHead>
              <TableHead>Rating To</TableHead>
              <TableHead class="text-right">Target From</TableHead>
              <TableHead class="text-right cursor-pointer" @click="handleSort('target_to')">
                <div class="flex items-center justify-end gap-2">Target To<component :is="getSortIcon('target_to')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead class="cursor-pointer" @click="handleSort('time')">
                <div class="flex items-center gap-2">Date<component :is="getSortIcon('time')" class="h-4 w-4" /></div>
              </TableHead>
              <TableHead>Synced</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-if="store.stocks.length === 0">
              <TableCell colspan="9" class="text-center py-12 text-muted-foreground">No stocks found</TableCell>
            </TableRow>
            <TableRow v-for="stock in store.stocks" :key="stock.id" class="cursor-pointer hover:bg-muted/50" @click="handleRowClick(stock)">
              <TableCell class="font-semibold text-primary">{{ stock.ticker }}</TableCell>
              <TableCell class="max-w-[200px] truncate">{{ stock.company }}</TableCell>
              <TableCell>{{ stock.brokerage }}</TableCell>
              <TableCell><RatingBadge :rating="stock.rating_from" /></TableCell>
              <TableCell><RatingBadge :rating="stock.rating_to" /></TableCell>
              <TableCell class="text-right font-mono">{{ formatCurrency(stock.target_from) }}</TableCell>
              <TableCell class="text-right font-mono font-semibold">{{ formatCurrency(stock.target_to) }}</TableCell>
              <TableCell class="text-muted-foreground">{{ formatDate(stock.time) }}</TableCell>
              <TableCell class="text-xs text-muted-foreground">{{ formatDateTime(stock.updated_at) }}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
        <Pagination v-if="store.meta && store.meta.total_pages > 1" :current-page="store.meta.page" :total-pages="store.meta.total_pages" :total-items="store.meta.total" :limit="store.meta.limit" @page-change="handlePageChange" />
      </CardContent>
    </Card>
  </div>
</template>
