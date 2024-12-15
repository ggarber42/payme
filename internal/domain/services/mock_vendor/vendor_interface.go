package vendor_service

import (

	"github.com/ggarber42/payme/internal/domain/entity"
)

type IVendorService interface {
	ProcessPayment(paymentData *entity.PaymentRequest) ServiceResponse
}
