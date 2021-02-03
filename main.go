package main

import (
	"github.com/dpjungmin/jellypi-server/api"
	"github.com/dpjungmin/jellypi-server/config"
	db "github.com/dpjungmin/jellypi-server/database"
	"github.com/dpjungmin/jellypi-server/utils/prometheus"
)

func init() {
	config.SanityCheck()
	db.GetPGSingleton().AutoMigrate()

	// Start metrics exposition on a different goroutine
	go prometheus.StartExposition()
}

func main() {
	api.StartApplication()
}
