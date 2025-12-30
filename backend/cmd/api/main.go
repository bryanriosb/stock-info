package main

import (
	"log"

	"github.com/bryanriosb/stock-info/internal/application"
	"github.com/bryanriosb/stock-info/internal/domain/entity"
	"github.com/bryanriosb/stock-info/internal/infrastructure/database"
	"github.com/bryanriosb/stock-info/internal/infrastructure/gateway"
	infraRepo "github.com/bryanriosb/stock-info/internal/infrastructure/repository"
	"github.com/bryanriosb/stock-info/internal/interfaces/http/handler"
	"github.com/bryanriosb/stock-info/internal/interfaces/http/router"
	"github.com/bryanriosb/stock-info/pkg/config"
)

func main() {
	cfg := config.Load()

	db, err := database.NewCockroachDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := database.RunMigrations(db, &entity.Stock{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected and migrations completed")

	// Initialize dependencies
	stockRepo := infraRepo.NewStockRepository(db)
	stockAPIClient := gateway.NewStockAPIClient(cfg.StockAPI)
	stockUseCase := application.NewStockUseCase(stockRepo, stockAPIClient)

	handlers := &router.Handlers{
		Stock: handler.NewStockHandler(stockUseCase),
		Auth:  handler.NewAuthHandler(cfg.JWT),
	}

	app := newFiberApp()
	router.Setup(app, handlers, cfg.JWT.Secret)

	go startServer(app, cfg.Server.Port)

	gracefulShutdown(app)
}
