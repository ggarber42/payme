package postegres_repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}
