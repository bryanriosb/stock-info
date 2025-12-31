package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/bryanriosb/stock-info/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instance *gorm.DB
	once     sync.Once
	initErr  error
)

func Init(cfg shared.DatabaseConfig) error {
	once.Do(func() {
		instance, initErr = newConnection(cfg)
	})
	return initErr
}

func DB() *gorm.DB {
	if instance == nil {
		panic("database not initialized: call database.Init() first")
	}
	return instance
}

func Close() error {
	if instance == nil {
		return nil
	}
	sqlDB, err := instance.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func newConnection(cfg shared.DatabaseConfig) (*gorm.DB, error) {
	// First, ensure the database exists
	if err := ensureDatabaseExists(cfg); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying db: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	return db, nil
}

// ensureDatabaseExists connects to CockroachDB default database and creates the target database if needed
func ensureDatabaseExists(cfg shared.DatabaseConfig) error {
	// Connect to default database to create our database if it doesn't exist
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=defaultdb sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to default database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying db: %w", err)
	}
	defer sqlDB.Close()

	// Create database if not exists (CockroachDB syntax)
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DBName)
	if err := db.Exec(createDBSQL).Error; err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	log.Printf("Database '%s' is ready", cfg.DBName)
	return nil
}

// RunAutoMigrate runs GORM AutoMigrate (for development only)
// AutoMigrate creates tables, adds missing columns, and creates indexes
// It does NOT delete unused columns or change existing column types
func RunAutoMigrate(db *gorm.DB, models ...interface{}) error {
	log.Println("Running GORM AutoMigrate (development mode)...")

	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to run auto-migrate: %w", err)
	}

	log.Println("AutoMigrate completed successfully")
	return nil
}

// RunProductionMigrations runs versioned SQL migrations using golang-migrate
func RunProductionMigrations(cfg shared.DatabaseConfig, migrationsPath string) error {
	migrator, err := NewMigrator(cfg, migrationsPath)
	if err != nil {
		return err
	}
	defer migrator.Close()

	return migrator.Up()
}
