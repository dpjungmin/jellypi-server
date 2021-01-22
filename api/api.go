package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Init boostraps the application and listens for incomming requests.
func Init() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "JellyPi",
		StrictRouting: true,
		CaseSensitive: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	SetupRoutes(app)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "5000"
		}
		log.Fatal(app.Listen(":" + port))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Cleanup tasks
}
