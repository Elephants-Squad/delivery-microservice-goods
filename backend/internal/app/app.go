package app

import (
	"delivery-microservice-goods/backend/config"
	"delivery-microservice-goods/backend/pkg/postgres"
	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {

	pg := postgres.New()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
