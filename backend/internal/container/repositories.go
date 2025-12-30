package container

import (
	"github.com/bryanriosb/stock-info/internal/domain/repository"
	infraRepo "github.com/bryanriosb/stock-info/internal/infrastructure/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	Stock repository.StockRepository
}

func newRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Stock: infraRepo.NewStockRepository(db),
	}
}
