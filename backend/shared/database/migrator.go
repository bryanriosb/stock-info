package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/bryanriosb/stock-info/shared"
)

// Migrator handles database migrations using golang-migrate
type Migrator struct {
	migrate *migrate.Migrate
}

// NewMigrator creates a new Migrator instance
func NewMigrator(cfg shared.DatabaseConfig, migrationsPath string) (*Migrator, error) {
	// Build CockroachDB connection string for golang-migrate
	// Format: cockroachdb://user:password@host:port/dbname?sslmode=disable
	dbURL := fmt.Sprintf(
		"cockroachdb://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	sourceURL := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator: %w", err)
	}

	return &Migrator{migrate: m}, nil
}

// Up runs all pending migrations
func (m *Migrator) Up() error {
	log.Println("Running database migrations...")

	err := m.migrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("No new migrations to apply")
	} else {
		log.Println("Migrations applied successfully")
	}

	return nil
}

// Down rolls back the last migration
func (m *Migrator) Down() error {
	log.Println("Rolling back last migration...")

	err := m.migrate.Steps(-1)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Rollback completed successfully")
	return nil
}

// Version returns the current migration version
func (m *Migrator) Version() (uint, bool, error) {
	return m.migrate.Version()
}

// Close closes the migrator
func (m *Migrator) Close() error {
	sourceErr, dbErr := m.migrate.Close()
	if sourceErr != nil {
		return sourceErr
	}
	return dbErr
}
