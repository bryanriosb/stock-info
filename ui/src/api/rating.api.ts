import axios from './axios'

export interface RatingOption {
  id: string
  label: string
  value: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export const ratingApi = {
  async getRatingOptions(): Promise<RatingOption[]> {
    const response = await axios.get('/rating-options')
    return response.data.data
  }
}