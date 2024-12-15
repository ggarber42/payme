package http_controller

import (
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", hc.HealthCheckerHandler)
	r.Post("/payment/{vendor}", hc.PaymentHandler)


	return r
}
