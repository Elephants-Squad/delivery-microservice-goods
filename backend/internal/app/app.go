package app

import (
	"delivery-microservice-goods/backend/config"
	"delivery-microservice-goods/backend/internal/migrate"
	"delivery-microservice-goods/backend/pkg/logger"
	"delivery-microservice-goods/backend/pkg/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(cfg *config.Config) {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	log := logger.New(cfg)

	log.Info("Starting service")

	// Postgres conn
	db, err := postgres.New(cfg.DB.URL, log)
	if err != nil {
		log.Error("Failed to connect to database: %s", err.Error())
		return
	}

	defer db.Close()

	// Migrate up
	if err := migrate.Up(cfg.DB.URL); err != nil {
		log.Error("error with up migrations for database: %s", err.Error())
		return
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	err = app.Listen(cfg.HTTP.Port)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}

	// Shutdown
	err = app.Shutdown()
	if err != nil {
		log.Error("Failed to shutdown app: %s", err.Error())
	}
}
