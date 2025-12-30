package auth

import (
	"github.com/bryanriosb/stock-info/internal/auth/interfaces"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router, cfg *shared.Config) {
	handler := interfaces.NewHandler(cfg.JWT)

	group := app.Group("/auth")
	group.Post("/login", handler.Login)
}
