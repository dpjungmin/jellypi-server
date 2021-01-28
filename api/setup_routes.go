package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/dpjungmin/jellypi-server/database"
	h "github.com/dpjungmin/jellypi-server/handlers"
	r "github.com/dpjungmin/jellypi-server/repositories"
	s "github.com/dpjungmin/jellypi-server/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	// Repositories
	userRepo r.UserRepository

	// Services
	userService s.UserService

	// Handlers
	userHandler h.UserHandler
)

func initialize() {
	// The `database.Connect()` function must have been callen
	db = database.DB

	// Repositories
	userRepo = r.NewUserRepository(db)

	// Services
	userService = s.NewUserService(userRepo)

	// Handlers
	userHandler = h.NewUserHandler(userService)
}

// SetupRoutes will map each route with their corresponding handler
func SetupRoutes(app *fiber.App) {
	initialize()

	app.Get("", hello)
	app.Get("/swagger/*", swagger.Handler)
	api := app.Group("/api")
	{
		users := api.Group("/users")
		{
			users.Get("", userHandler.GetUser)
			users.Post("", userHandler.CreateUser)
		}
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello 👻")
}
