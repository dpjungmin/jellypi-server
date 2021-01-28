package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dpjungmin/jellypi-server/database"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"github.com/gofiber/fiber/v2"
)

// StartApplication will start listening for incomming requests
func StartApplication() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "JellyPi",
		StrictRouting: true,
		CaseSensitive: true,
	})

	ApplyMiddlewares(app)
	SetupRoutes(app)

	go func() {
		logger.Error("Application is shutting down", app.Listen(":5000"))
		os.Exit(1)
	}()

	logger.Info("Application is starting up")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	_ = <-c
	logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	logger.Info("Running cleanup tasks...")
	cleanup()
}

func cleanup() {
	database.Close()
}
