# Stock Info

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue-3.4+-green.svg)](https://vuejs.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Stock market analysis platform with algorithmic recommendations and real-time data sync.

## Features

- **Algorithmic Recommendations**: Multi-factor scoring (Rating 30%, Target 40%, Action 30%)
- **Real-time Sync**: Server-Sent Events for live progress updates
- **Modern UI**: Vue 3 + TypeScript with Tailwind CSS
- **Secure**: JWT authentication with refresh tokens
- **Enterprise Architecture**: Clean Architecture in Go backend

## Quick Start

```bash
# Clone and start all services
git clone <repository-url>
cd stock-info
docker-compose up -d

# Access points
# Frontend: http://localhost:5173
# Backend API: http://localhost:5000
# Database: localhost:26257
```

## Tech Stack

**Backend**: Go 1.21+, Fiber, GORM, CockroachDB, JWT  
**Frontend**: Vue 3, TypeScript, Vite, Tailwind CSS, Pinia  
**Infrastructure**: Docker, Docker Compose

## Project Structure

```
stock-info/
├── backend/          # Go REST API (Clean Architecture)
├── ui/               # Vue.js frontend
├── compose.yml       # Docker Compose configuration
├── Makefile          # Build commands
└── README.md
```

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

## Development

```bash
# Development commands
make up                 # Start all services
make backend-dev         # Backend with hot reload
make frontend-dev        # Frontend dev server
make backend-test        # Run backend tests
make frontend-test       # Run frontend tests
make migrate-up          # Run database migrations
```

## Algorithm Details

**Recommendation Scoring**: `(Rating × 0.3) + (Target_Price × 0.4) + (Action × 0.3)`

- **Rating**: 1-9 scale (Strong Sell to Strong Buy)
- **Target Price**: Percentage change in analyst targets  
- **Action**: Positive (+1.0), Neutral (+0.5), Negative (-0.5)

## Documentation

- [Backend README](backend/README.md) - API documentation
- [Frontend README](ui/README.md) - UI development guide

## Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/name`
3. Add tests and update documentation
4. Submit pull request

## License

MIT License - see [LICENSE](LICENSE) file for details.