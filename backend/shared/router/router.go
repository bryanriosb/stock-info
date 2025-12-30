package router

import (
	"time"

	"github.com/bryanriosb/stock-info/internal/auth"
	"github.com/bryanriosb/stock-info/internal/recommendation"
	"github.com/bryanriosb/stock-info/internal/stock"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *shared.Config) {
	app.Get("/health", healthCheck)
	app.Get("/", root)

	api := app.Group("/api/v1")

	// Public routes
	auth.Register(api, cfg)

	// Protected routes
	protected := api.Group("", middleware.JWTProtected(cfg.JWT.Secret))
	stock.Register(protected, db, cfg)
	recommendation.Register(protected, db)
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Stock Info API",
		"version": "1.0.0",
	})
}
