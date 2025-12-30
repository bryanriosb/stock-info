package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/health", healthCheck)
	app.Get("/", root)

	// API v1 routes will be added in Phase 4
	// api := app.Group("/api/v1")
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
