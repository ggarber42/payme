package http_controller

import "github.com/ggarber42/payme/internal/domain/services"

type HttpController struct {
	ps *services.PaymentService
}

func NewHttpController(ps *services.PaymentService) *HttpController {
	return &HttpController{ps: ps}
}
