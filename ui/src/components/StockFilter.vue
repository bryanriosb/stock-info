<script setup lang="ts">
import { ref, watch, computed, onMounted } from 'vue'
import { useDebounceFn } from '@vueuse/core'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import { 
  Popover, 
  PopoverContent, 
  PopoverTrigger 
} from '@/components/ui/popover'
import { 
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator
} from '@/components/ui/command'
import { Search, FunnelX, PlusCircle } from 'lucide-vue-next'
import { cn } from '@/lib/utils'
import { ratingApi, type RatingOption } from '@/api/rating.api'

interface StockFilters {
  search: string
  rating_from: string
  rating_to: string
}

interface Props {
  loading?: boolean
}

interface Emits {
  (e: 'filter', filters: StockFilters): void
  (e: 'clear'): void
}

const props = withDefaults(defineProps<Props>(), {
  loading: false
})

const emit = defineEmits<Emits>()

const filters = ref<StockFilters>({
  search: '',
  rating_from: '',
  rating_to: ''
})

// Dynamic rating options from API
const ratingOptions = ref<RatingOption[]>([])
const loadingRatingOptions = ref(false)

// Load rating options from API
const loadRatingOptions = async () => {
  loadingRatingOptions.value = true
  try {
    ratingOptions.value = await ratingApi.getRatingOptions()
  } catch (error) {
    console.error('Failed to load rating options:', error)
    ratingOptions.value = []
  } finally {
    loadingRatingOptions.value = false
  }
}

// Load rating options on component mount
onMounted(() => {
  loadRatingOptions()
})

// Debounced search function (300ms like React project)
const debouncedSearch = useDebounceFn(() => {
  emit('filter', filters.value)
}, 300)

// Watch for changes in filters
watch(
  () => filters.value,
  () => {
    debouncedSearch()
  },
  { deep: true }
)

// Computed property to check if any filters are active
const isFiltered = computed(() => {
  return filters.value.search || 
         filters.value.rating_from || 
         filters.value.rating_to
})

// Get selected values for rating filters
const selectedRatingFrom = computed(() => filters.value.rating_from)
const selectedRatingTo = computed(() => filters.value.rating_to)

function handleClear() {
  filters.value = {
    search: '',
    rating_from: '',
    rating_to: ''
  }
  emit('clear')
}

// Helper function to get rating display name
function getRatingDisplay(value: string) {
  const option = ratingOptions.value.find(opt => opt.value === value)
  return option ? option.label : value
}
</script>

<template>
  <div class="flex items-center justify-between">
    <div class="flex flex-1 items-center space-x-2">
      <!-- Combined Search Input -->
      <div class="relative">
        <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input
          v-model="filters.search"
          placeholder="Search ticker or company..."
          class="h-8 w-[150px] lg:w-[250px] pl-9"
          @keyup.enter="debouncedSearch"
        />
      </div>

      <!-- Rating From Faceted Filter -->
      <Popover>
        <PopoverTrigger asChild>
          <Button 
            variant="outline" 
            size="sm" 
            class="btn-trigger-popover"
            :disabled="loadingRatingOptions || ratingOptions.length === 0"
          >
            <PlusCircle class="mr-2 h-4 w-4" />
            Rating From
            <Separator v-if="selectedRatingFrom" orientation="vertical" class="mx-2 h-4" />
            <Badge
              v-if="selectedRatingFrom"
              variant="secondary"
              class="rounded-sm px-1 font-normal lg:hidden"
            >
              1
            </Badge>
            <div v-if="selectedRatingFrom" class="hidden space-x-1 lg:flex">
              <Badge
                variant="secondary"
                class="rounded-sm px-1 font-normal"
              >
                {{ getRatingDisplay(selectedRatingFrom) }}
              </Badge>
            </div>
          </Button>
        </PopoverTrigger>
        <PopoverContent class="w-[200px] p-0" align="start">
          <Command>
            <CommandInput placeholder="Rating From" />
            <CommandList>
              <CommandEmpty v-if="!loadingRatingOptions">No rating options available.</CommandEmpty>
              <CommandEmpty v-if="loadingRatingOptions">Loading rating options...</CommandEmpty>
              <CommandGroup>
                <CommandItem
                  v-for="option in ratingOptions"
                  :key="option.value"
                  :value="option.value"
                  :class="cn(
                    'command-item-hover',
                    selectedRatingFrom === option.value && 'command-item-selected'
                  )"
                  @select="filters.rating_from = selectedRatingFrom === option.value ? '' : option.value"
                >
                  <div
                    :class="cn(
                      'mr-2 flex h-4 w-4 items-center justify-center rounded-sm border',
                      selectedRatingFrom === option.value
                        ? 'checkbox-selected'
                        : 'checkbox-unselected [&_svg]:invisible'
                    )"
                  >
                    <svg v-if="selectedRatingFrom === option.value" class="h-3 w-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <polyline points="20 6 9 17 4 12"></polyline>
                    </svg>
                  </div>
                  <span>{{ option.label }}</span>
                </CommandItem>
              </CommandGroup>
              <CommandSeparator v-if="selectedRatingFrom" />
              <CommandGroup v-if="selectedRatingFrom">
                <CommandItem
                  @select="filters.rating_from = ''"
                  class="justify-center text-center command-item-hover"
                >
                  Clear
                </CommandItem>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>

      <!-- Rating To Faceted Filter -->
      <Popover>
        <PopoverTrigger asChild>
          <Button 
            variant="outline" 
            size="sm" 
            class="btn-trigger-popover"
            :disabled="loadingRatingOptions || ratingOptions.length === 0"
          >
            <PlusCircle class="mr-2 h-4 w-4" />
            Rating To
            <Separator v-if="selectedRatingTo" orientation="vertical" class="mx-2 h-4" />
            <Badge
              v-if="selectedRatingTo"
              variant="secondary"
              class="rounded-sm px-1 font-normal lg:hidden"
            >
              1
            </Badge>
            <div v-if="selectedRatingTo" class="hidden space-x-1 lg:flex">
              <Badge
                variant="secondary"
                class="rounded-sm px-1 font-normal"
              >
                {{ getRatingDisplay(selectedRatingTo) }}
              </Badge>
            </div>
          </Button>
        </PopoverTrigger>
        <PopoverContent class="w-[200px] p-0" align="start">
          <Command>
            <CommandInput placeholder="Rating To" />
            <CommandList>
              <CommandEmpty v-if="!loadingRatingOptions">No rating options available.</CommandEmpty>
              <CommandEmpty v-if="loadingRatingOptions">Loading rating options...</CommandEmpty>
              <CommandGroup>
                <CommandItem
                  v-for="option in ratingOptions"
                  :key="option.value"
                  :value="option.value"
                  :class="cn(
                    'command-item-hover',
                    selectedRatingTo === option.value && 'command-item-selected'
                  )"
                  @select="filters.rating_to = selectedRatingTo === option.value ? '' : option.value"
                >
                  <div
                    :class="cn(
                      'mr-2 flex h-4 w-4 items-center justify-center rounded-sm border',
                      selectedRatingTo === option.value
                        ? 'checkbox-selected'
                        : 'checkbox-unselected [&_svg]:invisible'
                    )"
                  >
                    <svg v-if="selectedRatingTo === option.value" class="h-3 w-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <polyline points="20 6 9 17 4 12"></polyline>
                    </svg>
                  </div>
                  <span>{{ option.label }}</span>
                </CommandItem>
              </CommandGroup>
              <CommandSeparator v-if="selectedRatingTo" />
              <CommandGroup v-if="selectedRatingTo">
                <CommandItem
                  @select="filters.rating_to = ''"
                  class="justify-center text-center command-item-hover"
                >
                  Clear
                </CommandItem>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>

      <!-- Clear Filters Button -->
      <Button
        v-if="isFiltered"
        variant="outline"
        @click="handleClear"
        class="h-8 px-2 border-dashed border-destructive"
      >
        <FunnelX class="text-destructive" :size="20" />
      </Button>
    </div>
  </div>
</template>