package application

import (
	"context"
	"fmt"
	"log"

	"github.com/bryanriosb/stock-info/internal/stock/domain"
	"github.com/bryanriosb/stock-info/internal/stock/infrastructure"
)

type StockUseCase interface {
	SyncStocks(ctx context.Context) (int, error)
	SyncStocksWithProgress(ctx context.Context, onProgress infrastructure.ProgressCallback) (int, error)
	GetStocks(ctx context.Context, params domain.QueryParams) ([]*domain.Stock, int64, error)
	GetStockByID(ctx context.Context, id int64) (*domain.Stock, error)
}

type stockUseCase struct {
	repo      domain.StockRepository
	apiClient infrastructure.StockAPIClient
}

func NewStockUseCase(repo domain.StockRepository, apiClient infrastructure.StockAPIClient) StockUseCase {
	return &stockUseCase{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (uc *stockUseCase) SyncStocks(ctx context.Context) (int, error) {
	return uc.SyncStocksWithProgress(ctx, nil)
}

func (uc *stockUseCase) SyncStocksWithProgress(ctx context.Context, onProgress infrastructure.ProgressCallback) (int, error) {
	log.Println("Starting stock sync from external API...")

	stocks, err := uc.apiClient.FetchAllStocksWithProgress(ctx, onProgress)
	if err != nil {
		return 0, err
	}

	log.Printf("Fetched %d stocks from external API", len(stocks))

	// Report saving progress
	if onProgress != nil {
		onProgress(infrastructure.SyncProgress{
			Current: len(stocks),
			Total:   len(stocks),
			Percent: 97,
			Status:  "saving",
			Message: "Saving stocks to database...",
		})
	}

	if err := uc.repo.CreateBatch(ctx, stocks); err != nil {
		if onProgress != nil {
			onProgress(infrastructure.SyncProgress{
				Status:  "error",
				Message: err.Error(),
			})
		}
		return 0, err
	}

	// Report completion
	if onProgress != nil {
		onProgress(infrastructure.SyncProgress{
			Current: len(stocks),
			Total:   len(stocks),
			Percent: 100,
			Status:  "completed",
			Message: fmt.Sprintf("Successfully synced %d stocks", len(stocks)),
		})
	}

	log.Printf("Successfully synced %d stocks to database", len(stocks))
	return len(stocks), nil
}

func (uc *stockUseCase) GetStocks(ctx context.Context, params domain.QueryParams) ([]*domain.Stock, int64, error) {
	return uc.repo.FindAll(ctx, params)
}

func (uc *stockUseCase) GetStockByID(ctx context.Context, id int64) (*domain.Stock, error) {
	return uc.repo.FindByID(ctx, id)
}
