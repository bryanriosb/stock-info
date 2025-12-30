package auth

import (
	"github.com/bryanriosb/stock-info/internal/auth/interfaces"
	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/gofiber/fiber/v2"
)

func Register(api fiber.Router, cfg *shared.Config, userUseCase application.UserUseCase) {
	handler := interfaces.NewHandler(cfg.JWT, userUseCase)

	group := api.Group("/auth")
	group.Post("/login", handler.Login)
}
