package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	db "github.com/dpjungmin/jellypi-server/database"
	"github.com/dpjungmin/jellypi-server/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// swagger docs
	_ "github.com/dpjungmin/jellypi-server/docs"
	d "github.com/dpjungmin/jellypi-server/domain"
	h "github.com/dpjungmin/jellypi-server/handler"
	s "github.com/dpjungmin/jellypi-server/service"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

var (
	pgClient = db.GetPGSingleton().Client()

	// Repositories
	userRepo = d.NewUserRepository(pgClient)

	// Services
	authSrv = s.NewAuthService(userRepo)
	userSrv = s.NewUserService(userRepo)

	// Handlers
	authHandler = h.NewAuthHandler(authSrv)
	userHandler = h.NewUserHandler(userSrv)
)

// SetupRoutes will map each route with their corresponding handler
func SetupRoutes(app *fiber.App) {
	app.Get("/health", health)
	app.Get("/swagger/*", swagger.Handler)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	// api
	api := app.Group("/api")

	// auth
	auth := api.Group("/auth")
	auth.Get("check", middleware.Protected(), health)
	auth.Post("login", authHandler.Login)

	// user
	users := api.Group("/users")
	users.Post("", userHandler.CreateUser)
}

func health(c *fiber.Ctx) error {
	return c.JSON(struct {
		Alive bool `json:"alive"`
	}{
		true,
	})
}
