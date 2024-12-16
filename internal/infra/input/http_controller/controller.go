package http_controller

import (
	"github.com/ggarber42/payme/internal/domain/entity"
	card_services "github.com/ggarber42/payme/internal/domain/services/card"
	vendor_service "github.com/ggarber42/payme/internal/domain/services/mock_vendor"
)

type HttpController struct {
	cardService card_services.ICardService
	paymentService map[entity.Vendor]vendor_service.IVendorService
}

func NewHttpController(cs card_services.ICardService) *HttpController {
	return &HttpController{
		cardService: cs,
		paymentService: map[entity.Vendor]vendor_service.IVendorService{
			entity.STONE: &vendor_service.StoneService{},
			entity.CIELO: &vendor_service.CieloService{},
			entity.REDE: &vendor_service.RedeService{},
		},
	}
}
