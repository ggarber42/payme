package mock

import "github.com/ggarber42/payme/internal/domain/entity"

type MockCardService struct{}

func NewCardService() *MockCardService{
	return &MockCardService{}
}

func (mcs MockCardService) EnrichCard(cardData *entity.CardData) error {
	cardData.Number = "4111111111111111"
	cardData.Cvv = "123"
	return nil
}
