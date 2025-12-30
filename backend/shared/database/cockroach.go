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

// tableExists checks if a table already exists in the database
func tableExists(db *gorm.DB, tableName string) bool {
	var count int64
	db.Raw("SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = ? AND table_type = 'BASE TABLE'", tableName).Scan(&count)
	return count > 0
}

// RunMigrations runs database migrations only if tables don't exist
func RunMigrations(db *gorm.DB, models ...interface{}) error {
	needsMigration := false

	for _, model := range models {
		tableName := ""
		if tabler, ok := model.(interface{ TableName() string }); ok {
			tableName = tabler.TableName()
		}

		if tableName != "" && !tableExists(db, tableName) {
			needsMigration = true
			log.Printf("Table '%s' does not exist, migration needed", tableName)
			break
		}
	}

	if !needsMigration {
		log.Println("All tables exist, skipping migrations")
		return nil
	}

	log.Println("Running database migrations...")
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	log.Println("Migrations completed successfully")

	return nil
}
