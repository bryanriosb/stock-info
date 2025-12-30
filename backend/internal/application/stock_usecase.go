package application

import (
	"context"
	"log"

	"github.com/bryanriosb/stock-info/internal/domain/entity"
	"github.com/bryanriosb/stock-info/internal/domain/repository"
	"github.com/bryanriosb/stock-info/internal/infrastructure/gateway"
)

type StockUseCase interface {
	SyncStocks(ctx context.Context) (int, error)
	GetStocks(ctx context.Context, params repository.QueryParams) ([]*entity.Stock, int64, error)
	GetStockByTicker(ctx context.Context, ticker string) ([]*entity.Stock, error)
	GetStockByID(ctx context.Context, id int64) (*entity.Stock, error)
}

type stockUseCase struct {
	repo      repository.StockRepository
	apiClient gateway.StockAPIClient
}

func NewStockUseCase(repo repository.StockRepository, apiClient gateway.StockAPIClient) StockUseCase {
	return &stockUseCase{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (uc *stockUseCase) SyncStocks(ctx context.Context) (int, error) {
	log.Println("Starting stock sync from external API...")

	stocks, err := uc.apiClient.FetchAllStocks(ctx)
	if err != nil {
		return 0, err
	}

	log.Printf("Fetched %d stocks from external API", len(stocks))

	if err := uc.repo.CreateBatch(ctx, stocks); err != nil {
		return 0, err
	}

	log.Printf("Successfully synced %d stocks to database", len(stocks))
	return len(stocks), nil
}

func (uc *stockUseCase) GetStocks(ctx context.Context, params repository.QueryParams) ([]*entity.Stock, int64, error) {
	return uc.repo.FindAll(ctx, params)
}

func (uc *stockUseCase) GetStockByTicker(ctx context.Context, ticker string) ([]*entity.Stock, error) {
	return uc.repo.FindByTicker(ctx, ticker)
}

func (uc *stockUseCase) GetStockByID(ctx context.Context, id int64) (*entity.Stock, error) {
	return uc.repo.FindByID(ctx, id)
}
