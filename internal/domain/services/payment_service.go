package services

import (
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

func (ps *PaymentService) UpsertCard() (entity.CardData, error) {
	cardData, err := ps.pgRepo.GetCard()
	if err != nil {
		return entity.CardData{}, err
	}
	return cardData, nil
}
