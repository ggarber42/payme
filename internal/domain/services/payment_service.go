package services

import (
	"fmt"

	"github.com/ggarber42/payme/internal/domain/entity"
	postegres_repo "github.com/ggarber42/payme/internal/infra/output/repository/postgres"
)

type PaymentService struct {
	pgRepo *postegres_repo.PostgresRepo
}

func NewPaymentService(pgRepo *postegres_repo.PostgresRepo) *PaymentService {
	return &PaymentService{
		pgRepo: pgRepo,
	}
}

func (ps *PaymentService) ProcessPayment(data *entity.PaymentRequest) (entity.CardData, error){
	cardData, err := ps.enrichCard(data.CardData.Token)
	if err != nil {
		return entity.CardData{}, err
	}

	fmt.Println(cardData)

	return entity.CardData{}, nil
}

func (ps *PaymentService) enrichCard(token string) (entity.CardData, error) {
	cardData, err := ps.pgRepo.GetCard(token)
	if err != nil {
		return entity.CardData{}, err
	}
	return cardData, nil
}
