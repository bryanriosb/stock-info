package main

import (
	"log"

	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/database"
	"github.com/bryanriosb/stock-info/shared/router"
)

func main() {
	cfg := shared.LoadConfig()

	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	if err := database.RunMigrations(database.DB(), &domain.Stock{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected and migrations completed")

	app := newFiberApp()
	router.Setup(app, database.DB(), cfg)

	go startServer(app, cfg.Server.Port)

	gracefulShutdown(app)
}
