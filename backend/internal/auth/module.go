package auth

import (
	"github.com/bryanriosb/stock-info/internal/auth/interfaces"
	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(api fiber.Router, db *gorm.DB, cfg *shared.Config, userUseCase application.UserUseCase) {
	handler := interfaces.NewHandler(db, cfg.JWT, userUseCase)

	group := api.Group("/auth")
	group.Post("/login", handler.Login)
	group.Post("/refresh", handler.Refresh)
	group.Post("/logout", handler.Logout)
}
