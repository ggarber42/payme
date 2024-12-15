package card_services

import (

	"github.com/ggarber42/payme/internal/domain/entity"
	postegres_repo "github.com/ggarber42/payme/internal/infra/output/repository/postgres"
)

type CardService struct{
	pgRepo *postegres_repo.PostgresRepo
}

func NewCardService(pgRepo *postegres_repo.PostgresRepo) *CardService{
	return &CardService{
		pgRepo,
	}
}


func (cs *CardService) EnrichCard(cardData *entity.CardData) error {
	data, err := cs.pgRepo.GetCard(cardData.Token)
	if err != nil {
		return err
	}

	cardData.Number = data.Number
	cardData.Cvv = data.Cvv
	return nil
}
