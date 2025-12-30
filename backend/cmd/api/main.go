package main

import (
	"log"

	"github.com/bryanriosb/stock-info/internal/container"
	"github.com/bryanriosb/stock-info/internal/domain/entity"
	"github.com/bryanriosb/stock-info/internal/infrastructure/database"
	"github.com/bryanriosb/stock-info/internal/interfaces/http/router"
	"github.com/bryanriosb/stock-info/pkg/config"
)

func main() {
	cfg := config.Load()

	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	if err := database.RunMigrations(database.DB(), &entity.Stock{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected and migrations completed")

	c := container.New(cfg, database.DB())

	app := newFiberApp()
	router.Setup(app, c.RouterHandlers(), c.JWTSecret())

	go startServer(app, cfg.Server.Port)

	gracefulShutdown(app)
}
