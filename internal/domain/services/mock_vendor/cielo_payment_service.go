package vendor_service

import (
	"fmt"
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
)

type CieloService struct {}

func (c CieloService) ProcessPayment(paymentData *entity.PaymentRequest) ServiceResponse {
	if paymentData.CardData.Number == "4111111111111111" && paymentData.CardData.Cvv == "123" {
		msg := fmt.Sprintf("payment denied for %s", paymentData.CardData.Name)
		return ServiceResponse{Message: msg, HttpStatus: http.StatusAccepted}
	}
	return ServiceResponse{Message: "payment accepted", HttpStatus: http.StatusAccepted}
}
