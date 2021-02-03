package api

import (
	"time"

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
		ReadTimeout:   time.Second * 5,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	SetupRoutes(app)

	return app
}

func applyMiddlewares(app *fiber.App, mws ...func(*fiber.Ctx) error) {
	for _, md := range mws {
		app.Use(md)
	}
}
