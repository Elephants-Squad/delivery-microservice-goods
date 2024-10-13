package main

import (
	"delivery-microservice-goods/backend/config"
	"delivery-microservice-goods/backend/internal/app"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		return
	}

	app.Run(cfg)
}
