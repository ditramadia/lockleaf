package handler

import (
	"github.com/ditramadia/lockleaf/internal/config"
	"github.com/ditramadia/lockleaf/internal/manager"
)

type Handler struct {
	cfg *config.Config
	m   *manager.Manager
}

func New(cfg *config.Config, manager *manager.Manager) *Handler {
	return &Handler{
		cfg: cfg,
		m:   manager,
	}
}
