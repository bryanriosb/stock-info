package recommendation

import (
	"github.com/bryanriosb/stock-info/internal/recommendation/application"
	"github.com/bryanriosb/stock-info/internal/recommendation/interfaces"
	stockInfra "github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app fiber.Router, db *gorm.DB) {
	repo := stockInfra.NewStockRepository(db)
	useCase := application.NewRecommendationUseCase(repo)
	handler := interfaces.NewHandler(useCase)

	app.Get("/recommendations", handler.GetRecommendations)
}
