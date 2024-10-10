package router

import (
	"github.com/ggarber42/payme/internal/infra/input/http_controller/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", handlers.HealthCheckerHandler)

	return r
}
