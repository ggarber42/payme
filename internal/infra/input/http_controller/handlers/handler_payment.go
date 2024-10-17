package handlers

import (
	"errors"
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
	http_errors "github.com/ggarber42/payme/internal/infra/input/http_controller/errors"
	"github.com/go-chi/render"
)

type PaymentRequest struct {
	CardData     *entity.CardData `json:"cardData"`
	CardToken    string           `json:"cardToken"`
	ShoppingData *entity.CardData `json:"shoppingData"`
}

func (p *PaymentRequest) Bind(r *http.Request) error {
	switch {
	case p.CardData == nil || p.CardToken == "" || p.ShoppingData == nil:
		return errors.New("missing required fields")
	}
	return nil
}

type PaymentResponse struct {
	Message string `json:"message"`
	Elapsed int64 `json:"elapsed"`
}

func newPaymentResponse(message string) *PaymentResponse {
	return &PaymentResponse{Message: message}
}

func (pr *PaymentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	pr.Elapsed = 10
	return nil
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	data := &PaymentRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, http_errors.NewErrorResponse(http.StatusBadRequest, "Invalid request", err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, newPaymentResponse("payment created"))
}
