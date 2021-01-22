package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Map routes with their corresponding handlers.
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1", func(c *fiber.Ctx) error {
			c.JSON(fiber.Map{
				"message": "v1",
			})
			return c.Next()
		})
		{
			v1.Get("/hello", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!")
			})

			v1.Get("/panic", func(c *fiber.Ctx) error {
				panic("this panic is catched by fiber")
			})
		}
		app.Get("/swagger/*", swagger.Handler)
	}
}
