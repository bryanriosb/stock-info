package stock

import (
	"github.com/bryanriosb/stock-info/internal/stock/application"
	"github.com/bryanriosb/stock-info/internal/stock/infrastructure"
	"github.com/bryanriosb/stock-info/internal/stock/interfaces"
	"github.com/bryanriosb/stock-info/shared"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app fiber.Router, db *gorm.DB, cfg *shared.Config) {
	repo := infrastructure.NewStockRepository(db)
	apiClient := infrastructure.NewStockAPIClient(cfg.StockAPI)
	useCase := application.NewStockUseCase(repo, apiClient)
	handler := interfaces.NewHandler(useCase)

	group := app.Group("/stocks")
	group.Get("/", handler.GetStocks)
	group.Get("/sync-stream", handler.SyncStocksStream) // SSE endpoint - must be before :id
	group.Get("/:id", handler.GetStockByID)
	group.Get("/ticker/:ticker", handler.GetStockByTicker)
}
