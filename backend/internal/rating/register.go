package rating

import (
	"github.com/bryanriosb/stock-info/internal/rating/infrastructure"
	"github.com/bryanriosb/stock-info/internal/rating/interfaces"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app fiber.Router, db *gorm.DB) {
	repo := infrastructure.NewRatingOptionRepository(db)
	handler := interfaces.NewHandler(repo)

	app.Get("/rating-options", handler.GetAllRatingOptions)
}
