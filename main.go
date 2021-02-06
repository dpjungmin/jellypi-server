package main

import (
	"github.com/dpjungmin/jellypi-server/api"
	"github.com/dpjungmin/jellypi-server/config"
	db "github.com/dpjungmin/jellypi-server/database"
)

func init() {
	config.SanityCheck()
	db.GetPGSingleton().AutoMigrate()
}

func main() {
	api.StartApplication()
}
