<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStocksStore } from '@/stores/stocks.store'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Separator } from '@/components/ui/separator'
import RatingBadge from '@/components/RatingBadge.vue'
import { ArrowLeft, TrendingUp, TrendingDown, ArrowRight, Calendar, DollarSign } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const store = useStocksStore()

const stockId = computed(() => route.params.id as string)
onMounted(() => store.fetchStockById(stockId.value))

const priceChange = computed(() => {
  if (!store.currentStock) return 0
  const { target_from, target_to } = store.currentStock
  return target_from === 0 ? 0 : ((target_to - target_from) / target_from) * 100
})

const isPositive = computed(() => priceChange.value >= 0)

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function formatDate(d: string) {
  return new Intl.DateTimeFormat('en-US', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(d))
}
</script>

<template>
  <div class="space-y-6">
    <Button variant="ghost" @click="router.back()"><ArrowLeft class="mr-2 h-4 w-4" />Back</Button>

    <div v-if="store.loading" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <Card class="lg:col-span-2"><CardContent class="p-6"><Skeleton class="h-32 w-full" /></CardContent></Card>
      <Card><CardContent class="p-6"><Skeleton class="h-32 w-full" /></CardContent></Card>
    </div>

    <Alert v-else-if="store.error" variant="destructive"><AlertDescription>{{ store.error }}</AlertDescription></Alert>

    <template v-else-if="store.currentStock">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <Card class="lg:col-span-2">
          <CardHeader>
            <div class="flex items-start justify-between">
              <div>
                <CardTitle class="text-3xl">{{ store.currentStock.ticker }}</CardTitle>
                <CardDescription class="text-lg">{{ store.currentStock.company }}</CardDescription>
              </div>
              <div class="text-right">
                <p class="text-sm text-muted-foreground">Target Price</p>
                <p class="text-3xl font-bold" :class="isPositive ? 'text-success' : 'text-destructive'">
                  {{ formatCurrency(store.currentStock.target_to) }}
                </p>
                <div class="flex items-center justify-end gap-1" :class="isPositive ? 'text-success' : 'text-destructive'">
                  <component :is="isPositive ? TrendingUp : TrendingDown" class="h-4 w-4" />
                  <span>{{ isPositive ? '+' : '' }}{{ priceChange.toFixed(2) }}%</span>
                </div>
              </div>
            </div>
          </CardHeader>
          <CardContent>
            <div class="grid grid-cols-2 gap-6">
              <div><p class="text-sm text-muted-foreground">Brokerage</p><p class="font-medium">{{ store.currentStock.brokerage }}</p></div>
              <div><p class="text-sm text-muted-foreground">Action</p><p class="font-medium">{{ store.currentStock.action }}</p></div>
              <div><p class="text-sm text-muted-foreground flex items-center gap-1"><DollarSign class="h-4 w-4" />Previous</p><p class="font-mono">{{ formatCurrency(store.currentStock.target_from) }}</p></div>
              <div><p class="text-sm text-muted-foreground flex items-center gap-1"><DollarSign class="h-4 w-4" />New</p><p class="font-mono">{{ formatCurrency(store.currentStock.target_to) }}</p></div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader><CardTitle>Rating Change</CardTitle></CardHeader>
          <CardContent>
            <div class="flex items-center justify-center gap-4 py-4">
              <div class="text-center space-y-2">
                <p class="text-sm text-muted-foreground">From</p>
                <RatingBadge :rating="store.currentStock.rating_from" />
              </div>
              <ArrowRight class="h-6 w-6 text-muted-foreground" />
              <div class="text-center space-y-2">
                <p class="text-sm text-muted-foreground">To</p>
                <RatingBadge :rating="store.currentStock.rating_to" />
              </div>
            </div>
            <Separator class="my-4" />
            <div class="space-y-2 text-sm text-muted-foreground">
              <p class="flex items-center gap-2"><Calendar class="h-4 w-4" />Created: {{ formatDate(store.currentStock.created_at) }}</p>
              <p class="flex items-center gap-2"><Calendar class="h-4 w-4" />Updated: {{ formatDate(store.currentStock.updated_at) }}</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </template>
  </div>
</template>
