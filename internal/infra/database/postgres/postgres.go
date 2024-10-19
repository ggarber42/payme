package postgres

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/ggarber42/payme/internal/common/config"
	commonLogger "github.com/ggarber42/payme/internal/common/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var singleton sync.Once
var pool *pgxpool.Pool

func NewPostgresConn(config *config.Config) (*pgxpool.Pool, error) {
	logger := commonLogger.New(os.Stdout)

	singleton.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB)
		connConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			logger.PrintFatal(err, nil)
		}
		pool , err = pgxpool.NewWithConfig(context.Background(), connConfig)
		if err != nil {
			logger.PrintFatal(err, nil)
		}

	})
	logger.PrintInfo("connected to database", nil)
	return pool, nil
}
