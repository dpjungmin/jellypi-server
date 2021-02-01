package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Bootstrap will initialize the fiber application
func Bootstrap() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:  "JellyPi",
		StrictRouting: true,
		CaseSensitive: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	SetupRoutes(app)

	return app
}
