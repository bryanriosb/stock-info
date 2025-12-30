package repository

import (
	"context"
	"strings"

	"github.com/bryanriosb/stock-info/internal/domain/entity"
	domain "github.com/bryanriosb/stock-info/internal/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) domain.StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) Create(ctx context.Context, stock *entity.Stock) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(stock).Error
}

func (r *stockRepository) CreateBatch(ctx context.Context, stocks []*entity.Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		CreateInBatches(stocks, 100).Error
}

func (r *stockRepository) FindAll(ctx context.Context, params domain.QueryParams) ([]*entity.Stock, int64, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 || params.Limit > 100 {
		params.Limit = 20
	}

	var stocks []*entity.Stock
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Stock{})

	if params.Ticker != "" {
		query = query.Where("ticker = ?", params.Ticker)
	}

	if params.Company != "" {
		query = query.Where("company ILIKE ?", "%"+params.Company+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "id"
	allowed := map[string]bool{"id": true, "ticker": true, "company": true, "target_to": true, "created_at": true}
	if params.SortBy != "" && allowed[params.SortBy] {
		sortBy = params.SortBy
	}

	sortDir := "ASC"
	if strings.ToUpper(params.SortDir) == "DESC" {
		sortDir = "DESC"
	}

	offset := (params.Page - 1) * params.Limit

	err := query.Order(sortBy + " " + sortDir).
		Limit(params.Limit).
		Offset(offset).
		Find(&stocks).Error

	return stocks, total, err
}

func (r *stockRepository) FindByTicker(ctx context.Context, ticker string) ([]*entity.Stock, error) {
	var stocks []*entity.Stock
	err := r.db.WithContext(ctx).
		Where("ticker = ?", ticker).
		Order("created_at DESC").
		Find(&stocks).Error
	return stocks, err
}

func (r *stockRepository) FindByID(ctx context.Context, id int64) (*entity.Stock, error) {
	var stock entity.Stock
	err := r.db.WithContext(ctx).First(&stock, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &stock, nil
}
