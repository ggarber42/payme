package postegres_repo

import (
	"context"

	"github.com/ggarber42/payme/internal/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

func (pr *PostgresRepo) GetCard() (entity.CardData, error) {
	var card entity.CardData
	query := "SELECT name, number, token FROM CARD"

	ctx := context.Background()
	rows, err := pr.pool.Query(ctx, query)
	if err != nil {
		return entity.CardData{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&card.CardName,
			&card.CardNumber,
			&card.CardToken,
		)
		if err != nil {
			return entity.CardData{}, err
		}
	}

	return card, nil
}
