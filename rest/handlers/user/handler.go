package user

import "github.com/noyandey/go-ddd-starter/config"

type Handler struct {
	cfg *config.Config
	svc Service
}

func NewHandler(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		svc: svc,
		cfg: cfg,
	}
}
