package shared

import (
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	StockAPI StockAPIConfig
	Admin    AdminConfig
}

type AdminConfig struct {
	Username string
	Email    string
	Password string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret            string
	Expiration        time.Duration
	RefreshExpiration time.Duration
}

type StockAPIConfig struct {
	URL   string
	Token string
}

func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "3000"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "26257"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "stockinfo"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:            getEnv("JWT_SECRET", "default-secret-change-me"),
			Expiration:        parseDuration(getEnv("JWT_EXPIRATION", "15m")),
			RefreshExpiration: parseDuration(getEnv("JWT_REFRESH_EXPIRATION", "7d")),
		},
		StockAPI: StockAPIConfig{
			URL:   getEnv("STOCK_API_URL", "https://api.karenai.click/swechallenge/list"),
			Token: getEnv("STOCK_API_TOKEN", ""),
		},
		Admin: AdminConfig{
			Username: getEnv("ADMIN_USERNAME", "admin"),
			Email:    getEnv("ADMIN_EMAIL", "admin@stockinfo.com"),
			Password: getEnv("ADMIN_PASSWORD", "admin123"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		// Support "7d" format for days
		if len(value) > 0 && value[len(value)-1] == 'd' {
			days, err := time.ParseDuration(value[:len(value)-1] + "h")
			if err == nil {
				return days * 24
			}
		}
		return 24 * time.Hour
	}
	return duration
}
