package server

import (
	"fmt"
	"net/http"

	"github.com/ggarber42/payme/internal/common/config"
	"github.com/ggarber42/payme/internal/infra/input/http_controller/router"
	commonLogger "github.com/ggarber42/payme/internal/common/logger"
)

func Start(config *config.Config, logger commonLogger.ILogger) {

	server := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", config.ServerPort),
		Handler: router.NewRouter(),
	}

	logger.PrintInfo("starting server", nil)
	server.ListenAndServe()
}
