package http_controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
	"github.com/go-chi/render"
)

type PaymentRequest struct {
	CardData     *entity.CardData `json:"cardData"`
	ShoppingData *entity.CardData `json:"shoppingData"`
}

func (p *PaymentRequest) Bind(r *http.Request) error {
	switch {
	case p.CardData == nil || p.ShoppingData == nil:
		return errors.New("missing required fields")
	}
	return nil
}

type PaymentResponse struct {
	Message string `json:"message"`
}

func newPaymentResponse(message string) *PaymentResponse {
	return &PaymentResponse{Message: message}
}

func (pr *PaymentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (hc *HttpController) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	data := &PaymentRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, NewErrorResponse(http.StatusBadRequest, "Invalid request", err))
		return
	}

	cardData, err := hc.ps.UpsertCard()
	if err != nil{
		render.Status(r, http.StatusInternalServerError)	
		render.Render(w, r, NewErrorResponse(http.StatusInternalServerError, "internal server error", err))
		return
	}

	cardDataJson, err := json.Marshal(cardData)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)	
		render.Render(w, r, NewErrorResponse(http.StatusInternalServerError, "internal server error", err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, newPaymentResponse(string(cardDataJson)))
}
