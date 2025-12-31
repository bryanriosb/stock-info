.PHONY: help \
	backend-build backend-run backend-dev backend-test backend-test-v backend-test-cover \
	backend-test-unit backend-test-integration backend-clean backend-tidy backend-lint \
	backend-fmt backend-fmt-check backend-deps backend-mocks \
	backend-up backend-stop backend-down backend-logs backend-restart backend-rebuild

# Variables
BACKEND_DIR=./backend
BACKEND_APP_NAME=stock-info-api
BACKEND_BUILD_DIR=$(BACKEND_DIR)/bin
BACKEND_MAIN_PATH=$(BACKEND_DIR)/cmd/api

# Default target
all: help

# =============================================================================
# BACKEND COMMANDS
# =============================================================================

## Build the backend application
backend-build:
	@echo "Building $(BACKEND_APP_NAME)..."
	@cd $(BACKEND_DIR) && go build -o bin/$(BACKEND_APP_NAME) ./cmd/api

## Run the backend application
backend-run: backend-build
	@echo "Running $(BACKEND_APP_NAME)..."
	@$(BACKEND_BUILD_DIR)/$(BACKEND_APP_NAME)

## Run backend with hot reload (requires air)
backend-dev:
	@echo "Starting backend development server with hot reload..."
	@docker compose up backend -d

## Run all backend tests
backend-test:
	@echo "Running backend tests..."
	@cd $(BACKEND_DIR) && go test ./... -count=1

## Run backend tests with verbose output
backend-test-v:
	@echo "Running backend tests (verbose)..."
	@cd $(BACKEND_DIR) && go test ./... -v -count=1

## Run backend tests with coverage
backend-test-cover:
	@echo "Running backend tests with coverage..."
	@cd $(BACKEND_DIR) && mkdir -p tmp && go test ./... -coverprofile=./tmp/coverage.out -count=1
	@cd $(BACKEND_DIR) && go tool cover -html=./tmp/coverage.out -o tmp/coverage.html
	@echo "Coverage report generated: $(BACKEND_DIR)/tmp/coverage.html"

## Run backend unit tests only
backend-test-unit:
	@echo "Running backend unit tests..."
	@cd $(BACKEND_DIR) && go test ./internal/*/application/... -v -count=1

## Run backend integration tests only
backend-test-integration:
	@echo "Running backend integration tests..."
	@cd $(BACKEND_DIR) && go test ./internal/*/interfaces/... -v -count=1

## Clean backend build artifacts
backend-clean:
	@echo "Cleaning backend..."
	@rm -rf $(BACKEND_BUILD_DIR)
	@rm -rf $(BACKEND_DIR)/tmp/coverage.out $(BACKEND_DIR)/tmp/coverage.html

## Tidy backend go modules
backend-tidy:
	@echo "Tidying backend go modules..."
	@cd $(BACKEND_DIR) && go mod tidy

## Run backend linter (requires golangci-lint)
backend-lint:
	@echo "Running backend linter..."
	@cd $(BACKEND_DIR) && golangci-lint run ./...

## Format backend code
backend-fmt:
	@echo "Formatting backend code..."
	@cd $(BACKEND_DIR) && go fmt ./...

## Check backend code formatting
backend-fmt-check:
	@echo "Checking backend code format..."
	@cd $(BACKEND_DIR) && test -z "$$(gofmt -l .)" || (echo "Code is not formatted. Run 'make backend-fmt'" && exit 1)

## Download backend dependencies
backend-deps:
	@echo "Downloading backend dependencies..."
	@cd $(BACKEND_DIR) && go mod download

## Generate backend mocks
backend-mocks:
	@echo "Generating backend mocks..."
	@cd $(BACKEND_DIR) && go generate ./...

## Start backend container with compose
backend-up:
	@echo "Starting backend container..."
	@docker compose up -d backend

## Stop backend container
backend-stop:
	@echo "Stopping backend container..."
	@docker compose stop backend

## Stop and remove backend container
backend-down:
	@echo "Stopping and removing backend container..."
	@docker compose rm -sf backend

## View backend container logs
backend-logs:
	@docker compose logs -f backend --tail 100

## Restart backend container
backend-restart:
	@echo "Restarting backend container..."
	@docker compose restart backend

## Rebuild and restart backend container
backend-rebuild:
	@echo "Rebuilding backend container..."
	@docker compose up -d --build backend

db-up:
	@echo "Starting database container..."
	@docker compose up -d cockroachdb

db-down:
	@echo "Stopping and removing database container..."
	@docker compose rm -sf cockroachdb

# =============================================================================
# UI COMMANDS (placeholder for future)
# =============================================================================

frontend-restart:
	@echo "Restarting frontend container..."
	@docker compose restart frontend

# UI commands will be added here

# =============================================================================
# DOCKER COMPOSE COMMANDS
# =============================================================================

## Start all services with docker-compose
up:
	@docker compose up -d

## Stop all services
down:
	@docker compose down

## View logs from all services
logs:
	@docker compose logs -f

## Rebuild and start all services
rebuild:
	@docker compose up -d --build

# =============================================================================
# HELP
# =============================================================================

## Show this help
help:
	@echo "Stock Info - Available Commands"
	@echo "================================"
	@echo ""
	@echo "Backend Commands:"
	@echo "  make backend-build          - Build the backend application"
	@echo "  make backend-run            - Build and run the backend"
	@echo "  make backend-dev            - Run backend with hot reload (requires air)"
	@echo "  make backend-test           - Run all backend tests"
	@echo "  make backend-test-v         - Run backend tests with verbose output"
	@echo "  make backend-test-cover     - Run backend tests with coverage report"
	@echo "  make backend-test-unit      - Run backend unit tests only"
	@echo "  make backend-test-integration - Run backend integration tests only"
	@echo "  make backend-clean          - Clean backend build artifacts"
	@echo "  make backend-tidy           - Tidy backend go modules"
	@echo "  make backend-lint           - Run backend linter"
	@echo "  make backend-fmt            - Format backend code"
	@echo "  make backend-deps           - Download backend dependencies"
	@echo "  make backend-up             - Start backend container"
	@echo "  make backend-stop           - Stop backend container"
	@echo "  make backend-down           - Stop and remove backend container"
	@echo "  make backend-logs           - View backend container logs"
	@echo "  make backend-restart        - Restart backend container"
	@echo "  make backend-rebuild        - Rebuild and restart backend container"
	@echo "	 make db-up                  - Start database container"
	@echo "  make db-down                - Stop and remove database container"
	@echo ""
	@echo "Frontend Commands:"
	@echo "  make frontend-restart         - Restart frontend container"

	@echo "Docker Compose Commands:"
	@echo "  make up                     - Start all services"
	@echo "  make down                   - Stop all services"
	@echo "  make logs                   - View logs from all services"
	@echo "  make rebuild                - Rebuild and start all services"
	@echo ""
	@echo "  make help                   - Show this help"
