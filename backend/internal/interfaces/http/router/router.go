package router

import (
	"time"

	"github.com/bryanriosb/stock-info/internal/interfaces/http/handler"
	"github.com/bryanriosb/stock-info/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Stock          *handler.StockHandler
	Auth           *handler.AuthHandler
	Recommendation *handler.RecommendationHandler
}

func Setup(app *fiber.App, h *Handlers, jwtSecret string) {
	app.Get("/health", healthCheck)
	app.Get("/", root)

	api := app.Group("/api/v1")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/login", h.Auth.Login)

	// Protected routes
	protected := api.Group("", middleware.JWTProtected(jwtSecret))

	stocks := protected.Group("/stocks")
	stocks.Get("/", h.Stock.GetStocks)
	stocks.Get("/:id", h.Stock.GetStockByID)
	stocks.Get("/ticker/:ticker", h.Stock.GetStockByTicker)
	stocks.Post("/sync", h.Stock.SyncStocks)

	protected.Get("/recommendations", h.Recommendation.GetRecommendations)
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
