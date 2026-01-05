# Stock Info

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue-3.4+-green.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-blue.svg)](https://www.typescriptlang.org/)
[![CockroachDB](https://img.shields.io/badge/CockroachDB-latest-orange.svg)](https://www.cockroachlabs.com/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A full-stack stock market analysis platform featuring algorithmic investment recommendations, real-time data synchronization, and modern user interfaces.

## 🚀 Project Overview

Stock Info is a comprehensive fintech application that:
- Analyzes stock market data with algorithmic recommendations
- Provides real-time synchronization from external APIs
- Offers a modern, responsive web interface
- Implements enterprise-grade security and architecture
- Demonstrates full-stack development best practices

### Key Features

🔥 **Algorithmic Recommendations**
- Multi-factor scoring algorithm (Rating 30%, Target Price 40%, Action 30%)
- Real-time analysis of analyst ratings and price targets
- Quantified investment signals with clear reasoning

⚡ **Real-time Operations**
- Server-Sent Events (SSE) for live sync progress
- Efficient data synchronization from external APIs
- Optimistic UI updates for responsive experience

🎨 **Modern UI/UX**
- "Electric Sunset" design system breaking traditional fintech conventions
- Mobile-first responsive design
- Component-based architecture with shadcn-vue

🏗️ **Enterprise Architecture**
- Clean Architecture principles in Go backend
- Vertical Slice Architecture for maintainability
- TypeScript for type safety across the stack

🔒 **Security First**
- JWT-based authentication with refresh tokens
- SQL injection protection with parameterized queries
- CORS configuration and security headers

## 📁 Project Structure

```
stock-info/
├── backend/                     # Go REST API
│   ├── cmd/                    # Application entry points
│   ├── internal/               # Business logic modules
│   │   ├── auth/              # Authentication
│   │   ├── recommendation/    # Investment analysis
│   │   ├── stock/             # Stock management
│   │   ├── user/              # User management
│   │   └── rating/            # Rating options
│   ├── shared/                # Shared packages
│   ├── docs/                  # API documentation
│   ├── migrations/            # Database migrations
│   └── tests/                # Test suites
├── ui/                        # Vue.js frontend
│   ├── src/
│   │   ├── api/              # HTTP client and API services
│   │   ├── components/       # Vue components
│   │   ├── stores/           # Pinia state management
│   │   ├── router/           # Vue Router configuration
│   │   ├── types/            # TypeScript definitions
│   │   └── views/            # Page components
│   ├── docs/                 # Frontend documentation
│   └── tests/                # Test files
├── docs/                     # Project documentation
├── compose.yml               # Docker Compose
├── .env.example              # Environment template
├── Makefile                  # Build commands
└── README.md                 # This file
```

## 🛠️ Technology Stack

### Backend Technologies

| Category | Technology | Purpose |
|----------|-------------|---------|
| **Framework** | Go 1.21+ | High-performance server |
| **Web Framework** | Fiber v2.52+ | Fast HTTP routing |
| **Database** | CockroachDB | Distributed SQL |
| **ORM** | GORM v1.25+ | Database operations |
| **Authentication** | JWT | Token-based auth |
| **Architecture** | Clean Architecture | Maintainable structure |
| **Testing** | Go testing + Testify | Unit & integration tests |
| **Hot Reload** | Air | Development productivity |

### Frontend Technologies

| Category | Technology | Purpose |
|----------|-------------|---------|
| **Framework** | Vue 3.4+ | Progressive web framework |
| **Language** | TypeScript 5.0+ | Type safety |
| **Build Tool** | Vite 5.0+ | Fast development |
| **Styling** | Tailwind CSS 3.0+ | Utility-first CSS |
| **Components** | shadcn-vue | UI component library |
| **State** | Pinia 2.0+ | State management |
| **Routing** | Vue Router 4.0+ | SPA routing |
| **HTTP Client** | Axios | API communication |
| **Testing** | Vitest + Vue Test Utils | Component testing |

### Infrastructure

| Category | Technology | Purpose |
|----------|-------------|---------|
| **Container** | Docker | Application containerization |
| **Orchestration** | Docker Compose | Multi-container deployment |
| **Database** | CockroachDB | Distributed SQL database |
| **Version Control** | Git | Source code management |
| **Documentation** | Markdown | Technical documentation |

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher (or Bun)
- Docker and Docker Compose
- Git

### One-Command Setup

```bash
# Clone the repository
git clone <repository-url>
cd stock-info

# Start all services (backend, frontend, database)
docker-compose up -d

# Wait for services to be ready (30-60 seconds)
# Frontend: http://localhost:5173
# Backend API: http://localhost:5000
# Database: localhost:26257
```

### Manual Setup

#### Backend Setup

1. **Navigate to backend**
```bash
cd backend
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure environment**
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Start database**
```bash
docker run -d --name cockroachdb -p 26257:26257 cockroachdb/cockroach:v22.2.0 start-single-node --insecure
```

5. **Run migrations**
```bash
make migrate-up
```

6. **Start backend**
```bash
make dev
```

#### Frontend Setup

1. **Navigate to UI**
```bash
cd ui
```

2. **Install dependencies**
```bash
bun install
# or npm install
```

3. **Configure environment**
```bash
cp .env.example .env
# Edit .env with API URL
```

4. **Start frontend**
```bash
bun dev
# or npm run dev
```

### Access Points

- **Frontend Application**: http://localhost:5173
- **Backend API**: http://localhost:5000
- **Health Check**: http://localhost:5000/health
- **Database**: postgres://root@localhost:26257/stock_info?sslmode=disable

## 📚 Documentation

### Backend Documentation
- **[Backend README](backend/README.md)** - Complete API documentation
- **[Recommendation Algorithm](backend/docs/RECOMMENDATION-ALGORITHM.md)** - Investment analysis logic

### Frontend Documentation
- **[UI README](ui/README.md)** - Frontend development guide
- **[Development Plan](ui/docs/PLAN.md)** - UI architecture and design decisions

### API Documentation
- **Authentication**: JWT-based with refresh tokens
- **Stock Management**: CRUD operations with pagination
- **Recommendations**: AI-powered investment suggestions
- **User Management**: Registration, authentication, and administration
- **Real-time Sync**: Server-Sent Events for progress tracking

## 🎯 Core Features

### Algorithmic Recommendation Engine

The recommendation system uses a sophisticated multi-factor algorithm:

```
Score = (Rating_Change × 0.3) + (Target_Price_Change × 0.4) + (_Action_Type × 0.3)
```

**Rating Analysis (30% weight)**:
- Scale: 1 (Strong Sell) to 9 (Strong Buy)
- Tracks analyst rating improvements/downgrades
- Normalized scoring across rating systems

**Target Price Analysis (40% weight)**:
- Percentage change in price targets
- Quantifies analyst confidence changes
- Calculated as: `(New_Target - Old_Target) / Old_Target`

**Action Type (30% weight)**:
- "raised", "upgraded" → +1.0 (Positive)
- "maintained", "reiterated" → +0.5 (Neutral)
- "lowered", "downgraded" → -0.5 (Negative)

### Real-time Synchronization

- **Server-Sent Events**: Live progress updates during data sync
- **Progress Tracking**: Detailed progress through fetch → save → complete
- **Error Handling**: Comprehensive error reporting and recovery
- **Heartbeat System**: Connection monitoring and auto-retry

### Modern User Interface

- **"Electric Sunset" Theme**: Contemporary fintech design
- **Responsive Layout**: Mobile-first, tablet and desktop optimized
- **Component Library**: shadcn-vue for consistent, accessible components
- **Loading States**: Skeleton loaders and progress indicators
- **Toast Notifications**: User feedback for all actions

## 🔒 Security Features

### Authentication & Authorization
- **JWT Tokens**: Access and refresh token system
- **Token Storage**: Secure HTTP-only cookies
- **Route Guards**: Protected routes with automatic redirects
- **Session Management**: Automatic token refresh and logout

### Data Protection
- **SQL Injection Protection**: GORM parameterized queries
- **Input Validation**: Client and server-side validation
- **CORS Configuration**: Configurable cross-origin policies
- **Security Headers**: Content Security Policy and HSTS

### Best Practices
- **Environment Variables**: Secret management through .env files
- **Least Privilege**: Minimal required permissions
- **Error Sanitization**: Secure error message handling
- **HTTPS Ready**: Production SSL/TLS enforcement

## 🧪 Testing Strategy

### Backend Testing
- **Unit Tests**: All business logic and use cases
- **Integration Tests**: API endpoints and database operations
- **Test Coverage**: >80% code coverage target
- **Test Data**: Fixtures and factories for consistent testing

### Frontend Testing
- **Component Tests**: Vue component testing with Vue Test Utils
- **Unit Tests**: Store and service testing with Vitest
- **E2E Tests**: User workflow automation
- **Visual Testing**: Component regression testing

### Quality Assurance
- **Code Coverage**: Coverage reporting with Istanbul (frontend) and Go coverage (backend)
- **Linting**: Go fmt and golangci-lint for backend
- **Type Checking**: Strict TypeScript and Go compilation

## 🚀 Deployment

### Development Deployment
```bash
# Start all services locally
docker-compose up -d

# Individual service management
docker-compose up backend frontend database
```

### Production Deployment

**Using Docker Compose**:
```bash
# Production configuration
docker-compose -f docker-compose.prod.yml up -d
```

**Manual Deployment**:
```bash
# Backend
cd backend
make build
./bin/stock-api

# Frontend
cd ui
npm run build
serve -s dist -l 3000
```

**Environment Variables**:
- `ENVIRONMENT`: development/staging/production
- `DATABASE_URL`: CockroachDB connection string
- `JWT_SECRET`: Secret for token signing
- `VITE_API_URL`: Frontend API endpoint configuration

## 🔄 Development Workflow

### Commands

#### Backend Commands
```bash
make backend-dev          # Start backend with hot reload
make backend-test         # Run backend tests
make backend-test-cover    # Run backend tests with coverage
make backend-build        # Build backend binary
make backend-lint         # Backend linting
make backend-fmt          # Go code formatting
make backend-deps         # Download dependencies
make backend-up           # Start backend container
make backend-logs         # View backend logs
```

#### Frontend Commands
```bash
make frontend-dev        # Start frontend dev server (uses bun)
make frontend-test       # Run frontend tests
make frontend-test-cover # Run frontend tests with coverage
make frontend-build      # Build frontend assets
make frontend-lint       # Frontend linting
make frontend-lint-fix   # Auto-fix linting issues
make frontend-type-check # TypeScript type checking
make frontend-up         # Start frontend container
make frontend-logs       # View frontend logs
```

#### Database Commands
```bash
make migrate-up          # Run all pending migrations
make migrate-down        # Rollback last migration
make migrate-create name=example  # Create new migration
make migrate-version     # Show current migration version
make db-up              # Start database container
make db-down            # Stop and remove database container
```

#### Docker & Combined Operations
```bash
make up                 # Start all services with docker-compose
make down               # Stop all services
make rebuild            # Rebuild and start all services
make logs               # View logs from all services
```

#### Quick Commands
```bash
# Development workflow
make up                 # Start everything
make frontend-dev        # Work on frontend
make backend-dev         # Work on backend

# Testing workflow
make frontend-test && make backend-test  # Run all tests
make frontend-test-cover && make backend-test-cover  # Run with coverage

# Build workflow
make frontend-build && make backend-build  # Build both
```

## 📈 Roadmap

### Version 1.1 (Upcoming)
- [ ] Advanced filtering and search capabilities
- [ ] Portfolio tracking and management
- [ ] Email notifications for recommendations
- [ ] Advanced analytics dashboard

### Version 1.2 (Future)
- [ ] Enhanced algorithmic scoring models with AI
- [ ] Real-time market data streaming
- [ ] Social features and community insights
- [ ] Advanced charting and technical indicators
- [ ] API rate limiting and analytics

### Long-term Vision
- [ ] Multi-market support (international stocks)
- [ ] Automated trading integration
- [ ] Institutional-grade compliance features
- [ ] White-label solution for brokers
- [ ] Mobile-first native applications

## 🤝 Contributing

We welcome contributions from the community! Here's how to get started:

### For Developers
1. **Fork** the repository
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** following our coding standards
4. **Add tests** for your new functionality
5. **Update documentation** as needed
6. **Submit a pull request** with a clear description

### Contribution Guidelines
- **Code Style**: Follow Go and Vue.js best practices
- **Testing**: Maintain >80% test coverage
- **Documentation**: Update README files and inline docs
- **Commits**: Use conventional commit messages
- **Pull Requests**: Include description and testing steps

### Areas for Contribution
- **Frontend**: UI improvements, new components, mobile optimization
- **Backend**: API enhancements, performance optimizations, new features
- **Documentation**: Guides, examples, API documentation
- **Testing**: Test coverage improvements and new test scenarios
- **Infrastructure**: Docker improvements, deployment scripts, monitoring

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support & Community

- **Issues**: [GitHub Issues](https://github.com/bryanriosb/stock-info/issues) for bug reports and feature requests
- **Discussions**: [GitHub Discussions](https://github.com/bryanriosb/stock-info/discussions) for questions and community engagement
- **Documentation**: Check individual README files in backend/ and ui/ directories
- **API Reference**: See [Backend README](backend/README.md) for complete API documentation

## 🏆 Acknowledgments

- **CockroachDB**: For the amazing distributed SQL database
- **Vue.js Team**: For the progressive JavaScript framework
- **Fiber Team**: For the high-performance Go web framework
- **shadcn-vue**: For the beautiful component library
- **Open Source Community**: For all the incredible tools and libraries

---

**Stock Info** - Empowering smarter investment decisions through technology.