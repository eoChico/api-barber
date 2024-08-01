package main

import (
	"github.com/eoChico/api-barber/config"
	"github.com/eoChico/api-barber/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	//Initialize Router
	router.Initialize()
}
