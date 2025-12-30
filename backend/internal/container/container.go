package container

import (
	"github.com/bryanriosb/stock-info/internal/interfaces/http/router"
	"github.com/bryanriosb/stock-info/pkg/config"
	"gorm.io/gorm"
)

type Container struct {
	cfg *config.Config
	db  *gorm.DB

	repositories *Repositories
	useCases     *UseCases
	handlers     *Handlers
}

func New(cfg *config.Config, db *gorm.DB) *Container {
	c := &Container{
		cfg: cfg,
		db:  db,
	}

	c.repositories = newRepositories(db)
	c.useCases = newUseCases(c.repositories, cfg)
	c.handlers = newHandlers(c.useCases, cfg)

	return c
}

func (c *Container) RouterHandlers() *router.Handlers {
	return c.handlers.RouterHandlers()
}

func (c *Container) JWTSecret() string {
	return c.cfg.JWT.Secret
}
