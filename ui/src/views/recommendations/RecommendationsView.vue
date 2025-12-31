<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useRecommendationsStore } from '@/stores/recommendations.store'
import { useToast } from '@/components/ui/toast'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription, CardFooter } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import RatingBadge from '@/components/RatingBadge.vue'
import { RefreshCw, TrendingUp, DollarSign, Percent, ArrowRight } from 'lucide-vue-next'

const router = useRouter()
const store = useRecommendationsStore()
const { toast } = useToast()
const limit = ref('10')

onMounted(() => store.fetchRecommendations(Number(limit.value)))

watch(() => store.error, (error) => {
  if (error) {
    toast({ title: 'Error', description: error, variant: 'destructive' })
    store.error = null
  }
})

function handleLimitChange(v: string) {
  limit.value = v
  store.fetchRecommendations(Number(v))
}

function formatCurrency(v: number) {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(v)
}

function getScoreColor(score: number) {
  if (score >= 0.7) return 'text-success'
  if (score >= 0.4) return 'text-accent'
  return 'text-destructive'
}

function getProgressColor(score: number) {
  if (score >= 0.7) return 'bg-success'
  if (score >= 0.4) return 'bg-accent'
  return 'bg-destructive'
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Recommendations</h1>
        <p class="text-muted-foreground">Stock recommendations</p>
      </div>
      <div class="flex items-center gap-4">
        <Select :model-value="limit" @update:model-value="handleLimitChange">
          <SelectTrigger class="w-24"><SelectValue /></SelectTrigger>
          <SelectContent>
            <SelectItem value="5">5</SelectItem>
            <SelectItem value="10">10</SelectItem>
            <SelectItem value="20">20</SelectItem>
            <SelectItem value="50">50</SelectItem>
          </SelectContent>
        </Select>
        <Button @click="store.fetchRecommendations(Number(limit))" :disabled="store.loading">
          <RefreshCw class="mr-2 h-4 w-4" :class="{ 'animate-spin': store.loading }" />
          Refresh
        </Button>
      </div>
    </div>

    <div v-if="store.loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card v-for="i in 6" :key="i"><CardContent class="p-6"><Skeleton class="h-40 w-full" /></CardContent></Card>
    </div>

    <div v-else-if="store.recommendations.length === 0" class="text-center py-12">
      <TrendingUp class="mx-auto h-12 w-12 text-muted-foreground" />
      <h3 class="mt-4 text-lg font-medium">No recommendations</h3>
      <p class="text-muted-foreground">Sync stocks first to get recommendations</p>
      <Button class="mt-4" @click="router.push('/stocks')">Go to Stocks</Button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card v-for="(rec, index) in store.recommendations" :key="rec.stock.id" class="hover:shadow-lg transition-shadow">
        <CardHeader>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2">
              <span class="text-2xl font-bold text-muted-foreground">#{{ index + 1 }}</span>
              <div>
                <CardTitle class="text-lg text-primary">{{ rec.stock.ticker }}</CardTitle>
                <CardDescription class="truncate max-w-[150px]">{{ rec.stock.company }}</CardDescription>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm text-muted-foreground">Score</p>
              <p class="text-2xl font-bold" :class="getScoreColor(rec.score)">{{ (rec.score * 100).toFixed(0) }}</p>
            </div>
          </div>
        </CardHeader>
        <CardContent class="space-y-4">
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>Recommendation Score</span>
              <span>{{ (rec.score * 100).toFixed(1) }}%</span>
            </div>
            <div class="h-2 bg-muted rounded-full overflow-hidden">
              <div class="h-full rounded-full transition-all" :class="getProgressColor(rec.score)" :style="{ width: `${rec.score * 100}%` }" />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div class="bg-muted/50 rounded-lg p-3">
              <p class="text-xs text-muted-foreground flex items-center gap-1"><DollarSign class="h-3 w-3" />Target</p>
              <p class="font-semibold">{{ formatCurrency(rec.stock.target_to) }}</p>
            </div>
            <div class="bg-muted/50 rounded-lg p-3">
              <p class="text-xs text-muted-foreground flex items-center gap-1"><Percent class="h-3 w-3" />Gain</p>
              <p class="font-semibold" :class="rec.potential_gain_percent >= 0 ? 'text-success' : 'text-destructive'">
                {{ rec.potential_gain_percent >= 0 ? '+' : '' }}{{ rec.potential_gain_percent.toFixed(2) }}%
              </p>
            </div>
          </div>

          <p class="text-sm text-muted-foreground"><span class="font-medium">Analysis:</span> {{ rec.reason }}</p>

          <div class="flex items-center justify-between">
            <RatingBadge :rating="rec.stock.rating_to" />
            <span class="text-xs text-muted-foreground">{{ rec.stock.brokerage }}</span>
          </div>
        </CardContent>
        <CardFooter>
          <Button variant="ghost" class="w-full" @click="router.push(`/stocks/${rec.stock.id}`)">
            View Details <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>
