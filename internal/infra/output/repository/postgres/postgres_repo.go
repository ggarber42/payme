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

func (pr *PostgresRepo) GetCard(token string) (entity.CardData, error) {
	var card entity.CardData
	query := "SELECT number, cvv FROM card WHERE token = $1"

	ctx := context.Background()
	rows, err := pr.pool.Query(ctx, query, token)
	if err != nil {
		return entity.CardData{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&card.Number,
			&card.Cvv,
		)
		if err != nil {
			return entity.CardData{}, err
		}
	}

	return card, nil
}
