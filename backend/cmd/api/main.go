package main

import (
	"log"
	"path/filepath"
	"runtime"

	authDomain "github.com/bryanriosb/stock-info/internal/auth/domain"
	ratingDomain "github.com/bryanriosb/stock-info/internal/rating/domain"
	stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"
	userDomain "github.com/bryanriosb/stock-info/internal/user/domain"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/database"
	"github.com/bryanriosb/stock-info/shared/router"
)

func main() {
	cfg := shared.LoadConfig()

	log.Printf("Starting application in %s mode", cfg.Env)

	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run migrations based on environment
	if err := runMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := database.SeedAdmin(database.DB(), cfg.Admin); err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	log.Println("Database connected and migrations completed")

	app := newFiberApp()
	router.Setup(app, database.DB(), cfg)

	go startServer(app, cfg.Server.Port)

	gracefulShutdown(app)
}

// runMigrations executes database migrations based on the environment
// - Development: uses GORM AutoMigrate (automatic, flexible)
// - Production: uses golang-migrate with versioned SQL files (controlled, safe)
func runMigrations(cfg *shared.Config) error {
	if cfg.IsProduction() {
		// Production: use versioned SQL migrations
		_, currentFile, _, _ := runtime.Caller(0)
		migrationsPath := filepath.Join(filepath.Dir(currentFile), "..", "..", "migrations")
		return database.RunProductionMigrations(cfg.Database, migrationsPath)
	}

	// Development: use GORM AutoMigrate
	return database.RunAutoMigrate(
		database.DB(),
		&stockDomain.Stock{},
		&userDomain.User{},
		&authDomain.RefreshToken{},
		&ratingDomain.RatingOption{},
	)
}
