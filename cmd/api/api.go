package main

import (
	"github.com/ggarber42/payme/internal/common/config"
	"github.com/ggarber42/payme/internal/infra/server"
)


func main(){
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	server.Start(cfg)
}