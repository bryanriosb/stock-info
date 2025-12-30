package domain

import "context"

type QueryParams struct {
	Page    int
	Limit   int
	SortBy  string
	SortDir string
	Ticker  string
	Company string
}

type StockRepository interface {
	Create(ctx context.Context, stock *Stock) error
	CreateBatch(ctx context.Context, stocks []*Stock) error
	FindAll(ctx context.Context, params QueryParams) ([]*Stock, int64, error)
	FindByTicker(ctx context.Context, ticker string) ([]*Stock, error)
	FindByID(ctx context.Context, id int64) (*Stock, error)
}
