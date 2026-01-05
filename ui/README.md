# Stock Info UI

[![Vue Version](https://img.shields.io/badge/Vue-3.4+-green.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-blue.svg)](https://www.typescriptlang.org/)
[![Vite](https://img.shields.io/badge/Vite-7.0+-orange.svg)](https://vitejs.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A modern, responsive web application for stock market analysis and recommendations, built with Vue 3, TypeScript, and Tailwind CSS.

## 🚀 Overview

Stock Info UI is a sophisticated frontend application that provides:
- Real-time stock market data visualization
- Algorithmic investment recommendations
- User authentication and management
- Responsive design for all devices
- Modern component-based architecture

## 🎨 Design System

### "Electric Sunset" Theme

A contemporary fintech color palette that breaks away from traditional blue/white schemes:

| Color | Purpose | Hex | HSL |
|-------|---------|-----|-----|
| **Primary (Coral)** | Main actions, brand | `#FF6B5B` | `6 100% 68%` |
| **Secondary (Teal)** | Secondary actions | `#14B8A6` | `173 80% 40%` |
| **Accent (Amber)** | Highlights, warnings | `#F59E0B` | `38 92% 50%` |
| **Background** | Main background | `#F8FAFC` | `210 40% 98%` |
| **Foreground** | Main text | `#0F172A` | `222 47% 11%` |
| **Success** | Positive states | `#10B981` | `160 84% 39%` |
| **Destructive** | Errors, losses | `#EF4444` | `0 84% 60%` |

### Typography & Spacing
- **Font System**: Inter font family for optimal readability
- **Spacing**: Tailwind's consistent spacing scale
- **Border Radius**: 0.5rem for consistent rounded corners
- **Gradients**: Coral to amber gradients for premium elements

## 🏗️ Architecture

### Component Structure

```
ui/
├── src/
│   ├── api/                 # HTTP client and API calls
│   ├── assets/              # Static assets and global styles
│   ├── components/
│   │   ├── layout/         # Layout components
│   │   │   ├── AppLayout.vue
│   │   │   └── Navigation.vue
│   │   └── ui/             # shadcn-vue components
│   │       ├── Button.vue
│   │       ├── Card.vue
│   │       ├── Table.vue
│   │       └── ... (more components)
│   ├── lib/                # Utility functions
│   │   ├── utils.ts
│   │   ├── cookies.ts
│   │   └── json.ts
│   ├── router/             # Vue Router configuration
│   ├── stores/             # Pinia state management
│   │   ├── auth.store.ts
│   │   ├── stocks.store.ts
│   │   └── recommendations.store.ts
│   ├── types/              # TypeScript type definitions
│   ├── views/              # Page components
│   │   ├── auth/
│   │   ├── stocks/
│   │   ├── recommendations/
│   │   └── users/
│   ├── App.vue             # Root component
│   └── main.ts             # Application entry point
├── public/                 # Public assets
├── docs/                   # Documentation
└── tests/                  # Test files
```

### Design Patterns

- **Composition API**: Modern Vue 3 composition patterns
- **TypeScript**: Full type safety throughout the application
- **Component-First**: Reusable, composable components
- **State Management**: Centralized Pinia stores
- **Responsive Design**: Mobile-first approach

## 📋 Features

### Core Functionality
- ✅ **Authentication**: Secure login/register with JWT
- ✅ **Stock Management**: Browse, search, filter stocks
- ✅ **Real-time Updates**: Live sync progress with SSE
- ✅ **Recommendations**: Algorithmic investment suggestions
- ✅ **User Management**: Admin user management interface
- ✅ **Responsive Design**: Optimized for all screen sizes

### Technical Features
- ✅ **Component Library**: shadcn-vue component system
- ✅ **State Management**: Pinia for reactive state
- ✅ **Routing**: Vue Router with route guards
- ✅ **HTTP Client**: Axios with interceptors
- ✅ **Error Handling**: Centralized error management
- ✅ **Loading States**: Skeleton loaders and spinners
- ✅ **Form Validation**: Client-side validation
- ✅ **Toast Notifications**: User feedback system

## 🛠️ Technology Stack

| Category | Technology | Purpose |
|----------|-------------|---------|
| **Framework** | Vue 3.4+ | Progressive JavaScript framework |
| **Language** | TypeScript 5.0+ | Type-safe development |
| **Build Tool** | Vite 7.0+ | Fast development and building |
| **Styling** | Tailwind CSS 3.0+ | Utility-first CSS framework |
| **Components** | shadcn-vue | High-quality UI components |
| **State** | Pinia 2.0+ | Vue state management |
| **Routing** | Vue Router 4.0+ | SPA routing |
| **HTTP** | Axios 1.6+ | HTTP client with interceptors |
| **Icons** | Lucide Vue 3.0+ | Icon library |
| **Forms** | Vueuse Core | Composition utilities |
| **Testing** | Vitest + Vue Test Utils | Unit and component testing |

## 📦 Installation & Setup

### Prerequisites

- Node.js 18+ or Bun
- npm, yarn, or bun
- Git

### Quick Start

1. **Clone the repository**
```bash
git clone <repository-url>
cd stock-info/ui
```

2. **Install dependencies**
```bash
# Using npm
npm install

# Using yarn
yarn install

# Using bun (recommended)
bun install
```

3. **Environment configuration**
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Start development server**
```bash
# Using npm
npm run dev

# Using yarn
yarn dev

# Using bun (recommended)
bun dev
```

5. **Open browser**
Navigate to `http://localhost:5173` (or the URL shown in terminal)

### Environment Variables

```env
# API Configuration
VITE_API_URL=http://localhost:5000/api/v1

# Application
VITE_APP_NAME=Stock Info
VITE_APP_VERSION=1.0.0

# Feature Flags
VITE_ENABLE_ANALYTICS=false
VITE_ENABLE_DEBUG=true
```

## 📚 Component Documentation

### shadcn-vue Components

We use shadcn-vue for consistent, accessible UI components:

#### Button
```vue
<script setup lang="ts">
import { Button } from '@/components/ui/button'
</script>

<template>
  <Button variant="default" size="default">
    Click me
  </Button>
  <Button variant="outline" size="sm">
    Cancel
  </Button>
  <Button variant="destructive" class="gradient-coral">
    Delete
  </Button>
</template>
```

#### Card
```vue
<script setup lang="ts">
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Stock Data</CardTitle>
    </CardHeader>
    <CardContent>
      <!-- Content here -->
    </CardContent>
  </Card>
</template>
```

#### Table
```vue
<script setup lang="ts">
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
</script>

<template>
  <Table>
    <TableHeader>
      <TableRow>
        <TableHead>Ticker</TableHead>
        <TableHead>Company</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow>
        <TableCell>AAPL</TableCell>
        <TableCell>Apple Inc.</TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
```

### Custom Components

#### RatingBadge
Displays stock ratings with color-coded indicators:
```vue
<RatingBadge :rating="stock.rating_to" />
<!-- Shows: Buy (green), Sell (red), Hold (gray) -->
```

#### Pagination
Handles table pagination:
```vue
<Pagination 
  :current-page="1" 
  :total-pages="10" 
  :total-items="200" 
  :limit="20"
  @page-change="handlePageChange" 
/>
```

## 🗂️ State Management

### Auth Store
Manages user authentication and session:
```typescript
import { useAuthStore } from '@/stores/auth.store'

const authStore = useAuthStore()

// Login
await authStore.login({ username: 'admin', password: 'password' })

// Check authentication
if (authStore.isAuthenticated) {
  // User is logged in
}

// Logout
authStore.logout()
```

### Stocks Store
Manages stock data and operations:
```typescript
import { useStocksStore } from '@/stores/stocks.store'

const stocksStore = useStocksStore()

// Fetch stocks
await stocksStore.fetchStocks({ page: 1, limit: 20 })

// Sync stocks
const count = await stocksStore.syncStocks()

// Set filters
stocksStore.setFilters({ ticker: 'AAPL' })
```

### Recommendations Store
Manages recommendation data:
```typescript
import { useRecommendationsStore } from '@/stores/recommendations.store'

const recommendationsStore = useRecommendationsStore()

// Fetch recommendations
await recommendationsStore.fetchRecommendations(10)
```

## 🛣️ Routing

### Route Configuration

```typescript
// router/index.ts
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/stocks' },
    { path: '/login', component: LoginView, meta: { requiresGuest: true } },
    { path: '/stocks', component: StocksListView, meta: { requiresAuth: true } },
    { path: '/stocks/:id', component: StockDetailView, meta: { requiresAuth: true } },
    { path: '/recommendations', component: RecommendationsView, meta: { requiresAuth: true } },
    { path: '/users', component: UsersListView, meta: { requiresAuth: true } },
    { path: '/:pathMatch(.*)*', component: NotFoundView },
  ],
})
```

### Route Guards

Authentication and authorization are handled by router guards:
- `requiresAuth`: Requires authenticated user
- `requiresGuest`: Requires unauthenticated user
- Automatic redirect to login or stocks based on auth status

## 🌐 API Integration

### HTTP Client Configuration

```typescript
// api/axios.ts
import axios from 'axios'

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  headers: { 'Content-Type': 'application/json' },
})

// Request interceptor - Add auth token
apiClient.interceptors.request.use((config) => {
  const token = CookieManager.getAccessToken()
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Response interceptor - Handle auth errors
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Token expired, try refresh
      const refreshToken = CookieManager.getRefreshToken()
      if (refreshToken) {
        try {
          const response = await authApi.refresh({ refresh_token: refreshToken })
          // Update tokens
          CookieManager.setTokens(response.data.data)
          // Retry original request
          return apiClient.request(error.config)
        } catch (refreshError) {
          // Refresh failed, logout user
          CookieManager.clearTokens()
          redirectToLogin()
        }
      }
    }
    return Promise.reject(error)
  }
)
```

### API Services

```typescript
// api/stocks.api.ts
export const stocksApi = {
  getAll: (params?: StockQueryParams) => 
    apiClient.get<ApiResponse<Stock[]>>('/stocks', { params }),
    
  getById: (id: number) => 
    apiClient.get<ApiResponse<Stock>>(`/stocks/${id}`),
    
  sync: () => 
    apiClient.post<ApiResponse<{ message: string; count: number }>>('/stocks/sync'),
}
```

## 🔄 Real-time Features

### Server-Sent Events (SSE)

For real-time synchronization progress:

```typescript
// stocks.api.ts
export const syncStocksStream = (token: string) => {
  return new EventSource(`${SSE_BASE_URL}/stocks/sync-stream?token=${token}`, {
    withCredentials: false,
  })
}

// Usage in component
const eventSource = syncStocksStream(authStore.token!)

eventSource.onmessage = (event) => {
  const data = JSON.parse(event.data) as SyncProgress
  // Handle progress update
  progress.value = data.percent
  status.value = data.status
}

eventSource.onerror = (event) => {
  console.error('SSE error:', event)
  eventSource.close()
}
```

## 🎨 Styling Guide

### Tailwind CSS Configuration

```javascript
// tailwind.config.js
export default {
  darkMode: ['class'],
  content: ['./src/**/*.{vue,js,ts}'],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: 'hsl(var(--primary))',
          foreground: 'hsl(var(--primary-foreground))',
        },
        // ... other colors
      },
      borderRadius: {
        lg: 'var(--radius)',
      },
    },
  },
  plugins: [animate],
}
```

### Custom CSS Variables

```css
/* assets/index.css */
:root {
  --background: 210 40% 98%;
  --foreground: 222 47% 11%;
  --primary: 6 100% 68%;            /* #FF6B5B Coral */
  --secondary: 173 80% 40%;         /* #14B8A6 Teal */
  --accent: 38 92% 50%;             /* #F59E0B Amber */
  --radius: 0.5rem;
}

@layer utilities {
  .gradient-coral {
    @apply bg-gradient-to-r from-coral-500 to-amber-400;
  }
}
```

### Component Styling

```vue
<template>
  <div class="space-y-6">
    <Card class="hover:shadow-lg transition-shadow">
      <CardHeader class="bg-gradient-to-r from-primary/10 to-secondary/10">
        <CardTitle class="text-2xl text-primary">Stock Data</CardTitle>
      </CardHeader>
      <CardContent class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Content -->
        </div>
      </CardContent>
    </Card>
  </div>
</template>
```

## 🧪 Testing

### Test Setup

The project uses **Vitest** with **Vue Test Utils** and **@pinia/testing** for comprehensive testing.

```bash
# Testing dependencies (already installed)
# vitest, @vue/test-utils, @pinia/testing, jsdom, @vitest/coverage-istanbul
```

### Test Structure

```
src/
├── components/
│   ├── RatingBadge.spec.ts      # Component tests
│   ├── Pagination.spec.ts
│   └── Logo.spec.ts
├── stores/
│   └── __tests__/
│       ├── auth.store.spec.ts    # Store tests
│       ├── stocks.store.spec.ts
│       └── recommendations.store.spec.ts
└── test/
    └── setup.ts                  # Test setup configuration
```

### Component Testing Example

```typescript
// src/components/RatingBadge.spec.ts
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import RatingBadge from './RatingBadge.vue'

describe('RatingBadge', () => {
  it('renders the rating text', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Buy' },
    })
    expect(wrapper.text()).toContain('Buy')
  })

  it('applies success styling for buy ratings', () => {
    const wrapper = mount(RatingBadge, {
      props: { rating: 'Strong Buy' },
    })
    expect(wrapper.find('.bg-success\\/10').exists()).toBe(true)
  })
})
```

### Store Testing Example

```typescript
// src/stores/__tests__/auth.store.spec.ts
import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useAuthStore } from '../auth.store'

vi.mock('@/api/auth.api', () => ({
  authApi: { login: vi.fn(), register: vi.fn(), logout: vi.fn() },
}))

describe('Auth Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  it('initializes with no user', () => {
    const authStore = useAuthStore()
    expect(authStore.isAuthenticated).toBe(false)
    expect(authStore.user).toBe(null)
  })
})
```

### Running Tests

```bash
# Run tests in watch mode
bun run test

# Run tests once
bun run test:run

# Run tests with coverage report
bun run test:coverage
```

### Current Test Coverage

- **71 tests** across 6 test files
- **Stores**: ~90% coverage (auth, stocks, recommendations)
- **Components**: 100% coverage for custom components (RatingBadge, Pagination, Logo)

## 🚀 Deployment

### Build for Production

```bash
# Build application
npm run build

# Preview production build
npm run preview
```

### Docker Deployment

```dockerfile
# Dockerfile
FROM node:18-alpine as build

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

### Nginx Configuration

```nginx
# nginx.conf
server {
  listen 80;
  server_name localhost;
  root /usr/share/nginx/html;
  index index.html;

  location / {
    try_files $uri $uri/ /index.html;
  }

  location /api {
    proxy_pass http://backend:5000;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
  }
}
```

### Environment-Specific Builds

```bash
# Development build
npm run build:dev

# Staging build
npm run build:staging

# Production build
npm run build:prod
```

## 📱 Responsive Design

### Breakpoint System

| Breakpoint | Width | Usage |
|-----------|-------|-------|
| `sm` | 640px+ | Small devices (tablets) |
| `md` | 768px+ | Medium devices (laptops) |
| `lg` | 1024px+ | Large devices (desktops) |
| `xl` | 1280px+ | Extra large devices |

### Mobile-First Approach

```vue
<template>
  <!-- Mobile layout (default) -->
  <div class="space-y-4">
    <div class="grid grid-cols-1 gap-4">
      <!-- Content -->
    </div>
    
    <!-- Tablet and up -->
    <div class="md:grid md:grid-cols-2 md:gap-6">
      <!-- Content -->
    </div>
    
    <!-- Desktop and up -->
    <div class="lg:grid lg:grid-cols-3 lg:gap-8">
      <!-- Content -->
    </div>
  </div>
</template>
```

## 🔒 Security Considerations

### Frontend Security
- **Token Storage**: Secure cookie management
- **XSS Prevention**: Vue's built-in XSS protection
- **CSRF Protection**: SameSite cookie attributes
- **Content Security Policy**: Configurable CSP headers
- **HTTPS Only**: Production HTTPS enforcement

### Best Practices
- Input sanitization and validation
- Secure token handling with HttpOnly cookies
- Route guards for protected pages
- Error message sanitization
- Dependency vulnerability scanning

## 🛠️ Development Workflow

### Commands

```bash
# Development
npm run dev              # Start dev server
npm run preview          # Preview production build

# Building
npm run build            # Production build
npm run build:analyze    # Bundle analysis

# Testing
npm run test             # Run tests
npm run test:coverage    # Run with coverage
npm run test:e2e         # End-to-end tests

# Type checking
npm run type-check       # TypeScript checking
```

### Code Quality Tools

- **TypeScript**: Static type checking
- **Vue DevTools**: Vue application debugging
- **Vitest**: Fast unit testing with coverage

### Recommended Workflow

```bash
# Before committing
bun run test:run       # Verify tests pass
bun run type-check     # Verify types (optional)

# Before pushing
bun run build          # Verify build succeeds
```

## 🤝 Contributing

1. **Fork** the repository
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Commit** your changes: `git commit -m 'feat: add amazing feature'`
4. **Push** to the branch: `git push origin feature/amazing-feature`
5. **Open a Pull Request**

### Commit Convention

```
feat:     New feature
fix:      Bug fix
docs:     Documentation changes
style:    Code formatting
refactor: Code refactoring
test:     Adding or updating tests
chore:    Maintenance tasks
```

### Development Guidelines
- Follow Vue 3 Composition API patterns
- Use TypeScript for all new code
- Write tests for new features
- Update documentation
- Maintain code coverage >80%

## 📊 Performance

### Optimization Features
- **Code Splitting**: Lazy-loaded routes and components
- **Tree Shaking**: Unused code elimination
- **Asset Optimization**: Compressed images and fonts
- **Bundle Analysis**: Bundle size monitoring
- **Caching**: HTTP caching strategies

### Performance Metrics
- **First Contentful Paint**: <1.5s
- **Largest Contentful Paint**: <2.5s
- **Cumulative Layout Shift**: <0.1
- **Time to Interactive**: <3.5s

### Bundle Size Analysis

```bash
# Analyze bundle size
npm run build:analyze

# View bundle composition
npm run build:stats
```

## 📄 License

This project is licensed under MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

- **Documentation**: Check `/docs` directory
- **Component Library**: [shadcn-vue](https://www.shadcn-vue.com/)
- **Vue Documentation**: [Vue.js Guide](https://vuejs.org/guide/)
- **Issues**: Create GitHub issues for bugs/features

## 🔗 Related Projects

- **Backend API**: [Stock Info Backend](../backend) - Go REST API
- **Database**: CockroachDB for distributed SQL
- **Deployment**: Docker and Kubernetes configurations

---

Built with ❤️ using Vue 3, TypeScript, and modern web technologies.