package user

import (
	"github.com/bryanriosb/stock-info/internal/user/application"
	"github.com/bryanriosb/stock-info/internal/user/infrastructure"
	"github.com/bryanriosb/stock-info/internal/user/interfaces"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(publicRouter, protectedRouter fiber.Router, db *gorm.DB) application.UserUseCase {
	repo := infrastructure.NewUserRepository(db)
	useCase := application.NewUserUseCase(repo)
	handler := interfaces.NewHandler(useCase)

	// Public route (register)
	publicRouter.Post("/users", handler.Create)

	// Protected routes
	users := protectedRouter.Group("/users")
	users.Get("/", handler.GetAll)
	users.Get("/:id", handler.GetByID)
	users.Put("/:id", handler.Update)
	users.Delete("/:id", handler.Delete)

	return useCase
}
