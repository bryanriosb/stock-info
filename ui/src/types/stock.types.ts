export interface Stock {
  id: string
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: number
  target_to: number
  time: string
  created_at: string
  updated_at: string
}

export interface StockQueryParams {
  page?: number
  limit?: number
  sort_by?: 'id' | 'ticker' | 'company' | 'target_to' | 'time' | 'created_at'
  sort_dir?: 'asc' | 'desc'
  ticker?: string
  company?: string
}

export interface StockRecommendation {
  stock: Stock
  score: number
  reason: string
  potential_gain_percent: number
}
