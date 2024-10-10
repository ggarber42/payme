package main

import (
	"os"

	"github.com/ggarber42/payme/internal/common/config"
	commonLogger "github.com/ggarber42/payme/internal/common/logger"
	"github.com/ggarber42/payme/internal/infra/server"
)


func main(){
	cfg, err := config.LoadConfig(".")
	logger := commonLogger.New(os.Stdout, commonLogger.LevelInfo)
	if err != nil {
		panic(err)
	}

	server.Start(cfg, logger)
}