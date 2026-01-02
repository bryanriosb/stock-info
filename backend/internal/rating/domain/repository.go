package domain

import "context"

type RatingOptionRepository interface {
	FindByLabel(ctx context.Context, label string) (*RatingOption, error)
	FindAll(ctx context.Context) ([]*RatingOption, error)
	Create(ctx context.Context, option *RatingOption) error
	Upsert(ctx context.Context, option *RatingOption) error
}
