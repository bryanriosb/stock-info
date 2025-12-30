package container

import (
	"github.com/bryanriosb/stock-info/internal/application"
	"github.com/bryanriosb/stock-info/internal/infrastructure/gateway"
	"github.com/bryanriosb/stock-info/pkg/config"
)

type UseCases struct {
	Stock          application.StockUseCase
	Recommendation application.RecommendationUseCase
}

func newUseCases(repos *Repositories, cfg *config.Config) *UseCases {
	stockAPIClient := gateway.NewStockAPIClient(cfg.StockAPI)

	return &UseCases{
		Stock:          application.NewStockUseCase(repos.Stock, stockAPIClient),
		Recommendation: application.NewRecommendationUseCase(repos.Stock),
	}
}
