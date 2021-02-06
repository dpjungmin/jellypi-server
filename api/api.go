package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dpjungmin/jellypi-server/config"
	"github.com/dpjungmin/jellypi-server/utils/logger"
)

// StartApplication will start listening for incomming requests
func StartApplication() {
	logger.Info("Application is starting up...")
	app := Bootstrap()

	// Start listening on a different goroutine
	go func() {
		if err := app.Listen(":" + config.API.Port); err != nil {
			logger.Error("Application listen error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	_ = <-c

	logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()
	// go app.Shutdown()
	logger.Info("Application is down")

	cleanup()
}

func cleanup() {
	logger.Info("Running cleanup tasks...")
}
