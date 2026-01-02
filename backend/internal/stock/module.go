package stock

import (
	"github.com/bryanriosb/stock-info/internal/rating/application"
	"github.com/bryanriosb/stock-info/internal/rating/infrastructure"
	stockApp "github.com/bryanriosb/stock-info/internal/stock/application"
	stockInfra "github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/bryanriosb/stock-info/internal/stock/interfaces"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app fiber.Router, db *gorm.DB, cfg *shared.Config) {
	// Initialize rating service
	ratingRepo := infrastructure.NewRatingOptionRepository(db)
	ratingService := application.NewRatingService(ratingRepo)

	repo := stockInfra.NewStockRepository(db)
	apiClient := stockInfra.NewStockAPIClient(cfg.StockAPI)
	useCase := stockApp.NewStockUseCase(repo, apiClient, ratingService)
	handler := interfaces.NewHandler(useCase)

	group := app.Group("/stocks")
	group.Get("/", handler.GetStocks)
	group.Get("/sync-stream", handler.SyncStocksStream) // SSE endpoint - must be before :id
	group.Get("/:id", handler.GetStockByID)
}
