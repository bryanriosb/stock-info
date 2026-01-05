# Stock Info UI

[![Vue Version](https://img.shields.io/badge/Vue-3.4+-green.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-blue.svg)](https://www.typescriptlang.org/)
[![Vite](https://img.shields.io/badge/Vite-7.0+-orange.svg)](https://vitejs.dev/)

Modern Vue 3 + TypeScript frontend for stock market analysis with real-time recommendations.

## Tech Stack

**Framework**: Vue 3 + Composition API  
**Language**: TypeScript  
**Build Tool**: Vite  
**Styling**: Tailwind CSS + shadcn-vue  
**State**: Pinia  
**Runtime**: Bun (recommended)

## Quick Start

```bash
cd ui
bun install          # Install dependencies
bun dev             # Start development server
bun run build       # Build for production
bun run test        # Run tests
```

## Features

- 🔐 JWT authentication with refresh tokens
- 📊 Real-time stock data with SSE
- 📱 Responsive design (mobile-first)
- ⚡ Algorithmic investment recommendations
- 🎨 Modern "Electric Sunset" theme

## Project Structure

```
src/
├── api/           # HTTP client & API services
├── components/    # Vue components
├── stores/        # Pinia state management  
├── router/        # Vue Router
├── views/         # Page components
└── types/         # TypeScript definitions
```

## Development

```bash
bun dev              # Development server
bun run build        # Production build  
bun run test         # Run tests
bun run test:coverage # Run with coverage
bun run lint         # Lint code
bun run type-check   # Type checking
```

## Environment Variables

```env
VITE_API_URL=http://localhost:5000/api/v1
VITE_APP_NAME=Stock Info
```

## Key Components

### Auth Store
```typescript
import { useAuthStore } from '@/stores/auth.store'
const authStore = useAuthStore()
await authStore.login({ username: 'admin', password: 'password' })
```

### Stocks Store  
```typescript
import { useStocksStore } from '@/stores/stocks.store'
const stocksStore = useStocksStore()
await stocksStore.fetchStocks({ page: 1, limit: 20 })
```

### Custom Components
- `RatingBadge` - Stock rating displays
- `Pagination` - Table pagination
- shadcn-vue components for UI consistency

## Testing

```bash
bun run test            # Watch mode
bun run test:coverage   # With coverage
```

**Coverage**: 90% stores, 100% custom components

## Deployment

```bash
bun run build           # Production build
bun run preview         # Preview build
```

## License

MIT License - see [LICENSE](LICENSE) file for details.