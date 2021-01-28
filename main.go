package main

import (
	"github.com/dpjungmin/jellypi-server/api"
	db "github.com/dpjungmin/jellypi-server/database"
	_ "github.com/dpjungmin/jellypi-server/docs"
)

func main() {
	db.Connect()
	db.AutoMigrate()

	api.StartApplication()
}
