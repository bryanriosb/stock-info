package repository

import (
	"context"

	"github.com/bryanriosb/stock-info/internal/domain/entity"
)

type QueryParams struct {
	Page    int
	Limit   int
	SortBy  string
	SortDir string
	Ticker  string
	Company string
}

type StockRepository interface {
	Create(ctx context.Context, stock *entity.Stock) error
	CreateBatch(ctx context.Context, stocks []*entity.Stock) error
	FindAll(ctx context.Context, params QueryParams) ([]*entity.Stock, int64, error)
	FindByTicker(ctx context.Context, ticker string) ([]*entity.Stock, error)
	FindByID(ctx context.Context, id int64) (*entity.Stock, error)
}
