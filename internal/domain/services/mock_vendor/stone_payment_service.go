package vendor_service

import (
	"net/http"

	"github.com/ggarber42/payme/internal/domain/entity"
)

type StoneService struct{}

func (c StoneService) ProcessPayment(paymentData *entity.PaymentRequest) ServiceResponse {

	if paymentData.Purchase.TotalAmount < 1500 {
		return ServiceResponse{
			Message:    "the minimal required value is 1500",
			HttpStatus: http.StatusForbidden,
		}
	}
	return ServiceResponse{
		Message:    "payment accepted",
		HttpStatus: http.StatusAccepted,
	}
}
