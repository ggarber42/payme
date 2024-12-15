package http_controller

import (
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
	vendor_service "github.com/ggarber42/payme/internal/domain/services/mock_vendor"
	"github.com/ggarber42/payme/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PaymentResponse struct {
	Message string `json:"message"`
}

func newPaymentResponse(message string) *PaymentResponse {
	return &PaymentResponse{Message: message}
}

func (pr *PaymentResponse) Render(w http.ResponseWriter, r *http.Request) error { //implements for lib
	return nil
}

func (hc *HttpController) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	vendor := chi.URLParam(r, "vendor")
	if !utils.Contains(entity.VENDOR_OPTIONS, vendor) {
		render.Render(w, r, NewErrorResponse(http.StatusBadRequest, "Invalid request", entity.ErrInvalidVendor))
		return
	}

	data := &entity.PaymentRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, NewErrorResponse(http.StatusBadRequest, "Invalid request", err))
		return
	}

	err := hc.cardService.EnrichCard(data.CardData)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, NewErrorResponse(http.StatusInternalServerError, "internal server error", err))
		return
	}

	var res vendor_service.ServiceResponse

	switch vendor {
	case entity.STONE:
		res = hc.paymentService[entity.STONE].ProcessPayment(data)
	case entity.CIELO:
		res = hc.paymentService[entity.CIELO].ProcessPayment(data)
	case entity.REDE:
		res = hc.paymentService[entity.REDE].ProcessPayment(data)
	}

	render.Status(r, res.HttpStatus)
	render.Render(w, r, newPaymentResponse(res.Message))
}
