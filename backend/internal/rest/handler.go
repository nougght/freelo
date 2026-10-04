package rest

import (
	"backend/internal/config"
	"backend/internal/service"
)

type Handler struct {
}

func newHandler(cfg *config.Config, services *service.Services) *Handler {
	return &Handler{}
}
