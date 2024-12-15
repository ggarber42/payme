package vendor_service

import (
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
)

type RedeService struct{}

func (c RedeService) ProcessPayment(paymentData *entity.PaymentRequest) ServiceResponse {
	if paymentData.Purchase.Installments.Number > 5 {
		return ServiceResponse{Message: "the maximum number of installments is 5", HttpStatus: http.StatusForbidden}
	}
	return ServiceResponse{Message: "payment accepted", HttpStatus: http.StatusAccepted}
}
