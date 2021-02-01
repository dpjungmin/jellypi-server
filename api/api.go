package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dpjungmin/jellypi-server/utils/logger"
)

// App defines the interface for an application
type App interface {
	Shutdown() error
}

// Start will start listening for incomming requests
func Start() {
	logger.Info("Application is starting up...")
	app := Bootstrap()

	// Start listening on a different goroutine
	go func() {
		if err := app.Listen(":5000"); err != nil {
			logger.Error("Application listen error", err)
			os.Exit(1)
		}
	}()

	handleGracefulShutdown(app)
}

func handleGracefulShutdown(app App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	_ = <-c

	logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()
	logger.Info("Application is down")

	cleanup()
}

func cleanup() {
	logger.Info("Running cleanup tasks...")
}
