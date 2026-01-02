package rating

import (
	"github.com/bryanriosb/stock-info/internal/rating/application"
	"github.com/bryanriosb/stock-info/internal/rating/domain"
	"github.com/bryanriosb/stock-info/internal/rating/infrastructure"
	"github.com/bryanriosb/stock-info/internal/rating/interfaces"
	"gorm.io/gorm"
)

type Module struct {
	Handler       *interfaces.Handler
	Repo          domain.RatingOptionRepository
	RatingService *application.RatingService
}

func NewModule(db *gorm.DB) *Module {
	repo := infrastructure.NewRatingOptionRepository(db)
	handler := interfaces.NewHandler(repo)
	ratingService := application.NewRatingService(repo)

	return &Module{
		Handler:       handler,
		Repo:          repo,
		RatingService: ratingService,
	}
}
