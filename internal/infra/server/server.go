package server

import (
	"fmt"
	"net/http"

	"github.com/ggarber42/payme/internal/common/config"
	"github.com/ggarber42/payme/internal/infra/input/http_controller/router"
)

func Start(config *config.Config) {

	server := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", config.ServerPort),
		Handler: router.NewRouter(),
	}

	fmt.Println("stating server...")
	server.ListenAndServe()
}
