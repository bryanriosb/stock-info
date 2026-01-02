package infrastructure

import (
	"context"
	"errors"

	"github.com/bryanriosb/stock-info/internal/rating/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ratingOptionRepository struct {
	db *gorm.DB
}

func NewRatingOptionRepository(db *gorm.DB) domain.RatingOptionRepository {
	return &ratingOptionRepository{db: db}
}

func (r *ratingOptionRepository) FindByLabel(ctx context.Context, label string) (*domain.RatingOption, error) {
	var option domain.RatingOption
	err := r.db.WithContext(ctx).Where("label = ?", label).First(&option).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &option, nil
}

func (r *ratingOptionRepository) FindAll(ctx context.Context) ([]*domain.RatingOption, error) {
	var options []*domain.RatingOption
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Order("label ASC").
		Find(&options).Error

	if err != nil {
		return nil, err
	}

	return options, nil
}

func (r *ratingOptionRepository) Create(ctx context.Context, option *domain.RatingOption) error {
	return r.db.WithContext(ctx).Create(option).Error
}

func (r *ratingOptionRepository) Upsert(ctx context.Context, option *domain.RatingOption) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "label"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "is_active", "updated_at"}),
	}).Create(option).Error
}
