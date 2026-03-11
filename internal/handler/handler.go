package handler

import (
	"github.com/ditramadia/lockleaf/internal/config"
	"github.com/ditramadia/lockleaf/internal/service"
)

type Handler struct {
	cfg *config.Config
	s   *service.Service
}

func New(config *config.Config, service *service.Service) *Handler {
	return &Handler{
		cfg: config,
		s:   service,
	}
}
