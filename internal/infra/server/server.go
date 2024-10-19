package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ggarber42/payme/internal/common/config"
	commonLogger "github.com/ggarber42/payme/internal/common/logger"
	"github.com/ggarber42/payme/internal/infra/database/postgres"
	"github.com/ggarber42/payme/internal/infra/input/http_controller/router"
	postegres_repo "github.com/ggarber42/payme/internal/infra/output/repository/postgres"
	"github.com/ggarber42/payme/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(config *config.Config, logger commonLogger.ILogger) {

	pgPool, err := postgres.NewPostgresConn(config)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.ServerPort),
		Handler: router.NewRouter(),
	}

	go func() {
		if err != nil {
			logger.PrintFatal(err, nil)
		}
		repo := postegres_repo.NewRepo(pgPool)
		fmt.Println(repo)

		logger.PrintInfo("starting server", nil)
		server.ListenAndServe()
	}()

	gracefulShutDown(&server, pgPool, logger)

}

func gracefulShutDown(server *http.Server, pgPool *pgxpool.Pool, logger commonLogger.ILogger) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	signal := <-shutdownSignal
	logger.PrintInfo(fmt.Sprintf("Initiating graceful shutdown by signal CODE:%d - %s", signal.(syscall.Signal), signal.String()), nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		logger.PrintError(err, nil)
	}
	defer cancel()

	logger.PrintInfo("closing postgres connection...", nil)
	pgPool.Close()

	logger.PrintInfo("shutdown clompeted exit code 0...", nil)
	os.Exit(utils.EXIT_SUCCESS)
}
