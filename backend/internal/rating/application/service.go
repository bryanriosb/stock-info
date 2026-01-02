package application

import (
	"context"
	"log"

	"github.com/bryanriosb/stock-info/internal/rating/domain"
	stockDomain "github.com/bryanriosb/stock-info/internal/stock/domain"
)

type RatingService struct {
	repo domain.RatingOptionRepository
}

func NewRatingService(repo domain.RatingOptionRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) ExtractAndSaveRatingOptions(ctx context.Context, stocks []*stockDomain.Stock) error {
	ratingSet := make(map[string]string)

	for _, stock := range stocks {
		if stock.RatingFrom != "" {
			ratingSet[stock.RatingFrom] = stock.RatingFrom
		}
		if stock.RatingTo != "" {
			ratingSet[stock.RatingTo] = stock.RatingTo
		}
	}

	for label, value := range ratingSet {
		option := &domain.RatingOption{
			Label:    label,
			Value:    value,
			IsActive: true,
		}

		err := s.repo.Upsert(ctx, option)
		if err != nil {
			log.Printf("Failed to upsert rating option %s: %v", label, err)
			continue
		}
	}

	return nil
}
