package user

import (
	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/internal/user/infrastructure"
	"github.com/bryanriosb/stock-info/internal/user/interfaces"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPublicOnly(publicRouter fiber.Router, db *gorm.DB) application.UserUseCase {
	repo := infrastructure.NewUserRepository(db)
	useCase := application.NewUserUseCase(repo)
	handler := interfaces.NewHandler(useCase)

	// Public route (register)
	publicRouter.Post("/users", handler.Create)

	return useCase
}

func RegisterProtected(protectedRouter fiber.Router, useCase application.UserUseCase) {
	handler := interfaces.NewHandler(useCase)

	// Protected routes
	users := protectedRouter.Group("/users")
	users.Get("/", handler.GetAll)
	users.Get("/:id", handler.GetByID)
	users.Put("/:id", handler.Update)
	users.Delete("/:id", handler.Delete)
}

// Legacy function for backward compatibility
func Register(publicRouter, protectedRouter fiber.Router, db *gorm.DB) application.UserUseCase {
	useCase := RegisterPublicOnly(publicRouter, db)
	RegisterProtected(protectedRouter, useCase)
	return useCase
}
