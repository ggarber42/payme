package card_services

import "github.com/ggarber42/payme/internal/domain/entity"

type ICardService interface {
	EnrichCard(cardData *entity.CardData) error 
}
