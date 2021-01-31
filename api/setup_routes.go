package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	db "github.com/dpjungmin/jellypi-server/database"
	h "github.com/dpjungmin/jellypi-server/handler"
	r "github.com/dpjungmin/jellypi-server/repository"
	s "github.com/dpjungmin/jellypi-server/service"
	"github.com/gofiber/fiber/v2"
)

var (
	pgClient = db.GetPGSingleton().Client()

	// Repositories
	userRepo = r.NewUserRepository(pgClient)

	// Services
	userSrv = s.NewUserService(userRepo)

	// Handlers
	userHandler = h.NewUserHandler(userSrv)
)

// SetupRoutes will map each route with their corresponding handler
func SetupRoutes(app *fiber.App) {
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
