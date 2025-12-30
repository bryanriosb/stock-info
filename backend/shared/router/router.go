package router

import (
	"time"

	"github.com/bryanriosb/stock-info/internal/auth"
	"github.com/bryanriosb/stock-info/internal/recommendation"
	"github.com/bryanriosb/stock-info/internal/stock"
	"github.com/bryanriosb/stock-info/internal/user"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/bryanriosb/stock-info/shared/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *shared.Config) {
	app.Get("/health", healthCheck)
	app.Get("/", root)

	api := app.Group("/api/v1")
	protected := api.Group("", middleware.JWTProtected(cfg.JWT.Secret))

	// User module (public: POST /users, protected: GET/PUT/DELETE)
	userUseCase := user.Register(api, protected, db)

	// Auth (public)
	auth.Register(api, cfg, userUseCase)

	// Protected routes
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
