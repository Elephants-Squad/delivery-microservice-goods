package app

import (
	"delivery-microservice-goods/backend/config"
	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	err := app.Listen(cfg.HTTP.Port)
	if err != nil {
		return
	}

	// Shutdown
	err = app.Shutdown()
	if err != nil {
	}
}
