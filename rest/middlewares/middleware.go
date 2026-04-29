package middlewares

import "github.com/noyandey/go-ddd-starter/config"

type Middlewares struct {
	cfg *config.Config
}

func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{
		cfg: cfg,
	}
}
