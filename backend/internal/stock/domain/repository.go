package domain

import "context"

type QueryParams struct {
	Page       int
	Limit      int
	SortBy     string
	SortDir    string
	Search     string // Combined search for ticker or company
	RatingFrom string // Rating from filter
	RatingTo   string // Rating to filter
}

type StockRepository interface {
	Create(ctx context.Context, stock *Stock) error
	CreateBatch(ctx context.Context, stocks []*Stock) error
	FindAll(ctx context.Context, params QueryParams) ([]*Stock, int64, error)
	FindByID(ctx context.Context, id int64) (*Stock, error)
}
