package container

import (
	"github.com/bryanriosb/stock-info/internal/interfaces/http/handler"
	"github.com/bryanriosb/stock-info/internal/interfaces/http/router"
	"github.com/bryanriosb/stock-info/pkg/config"
)

type Handlers struct {
	Stock          *handler.StockHandler
	Auth           *handler.AuthHandler
	Recommendation *handler.RecommendationHandler
}

func newHandlers(useCases *UseCases, cfg *config.Config) *Handlers {
	return &Handlers{
		Stock:          handler.NewStockHandler(useCases.Stock),
		Auth:           handler.NewAuthHandler(cfg.JWT),
		Recommendation: handler.NewRecommendationHandler(useCases.Recommendation),
	}
}

func (h *Handlers) RouterHandlers() *router.Handlers {
	return &router.Handlers{
		Stock:          h.Stock,
		Auth:           h.Auth,
		Recommendation: h.Recommendation,
	}
}
