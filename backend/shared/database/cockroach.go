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

// RunMigrations runs database migrations (adds missing columns, creates missing tables)
func RunMigrations(db *gorm.DB, models ...interface{}) error {
	log.Println("Running database migrations...")

	// AutoMigrate creates tables, adds missing columns, and creates indexes
	// It does NOT delete unused columns or change existing column types
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}
