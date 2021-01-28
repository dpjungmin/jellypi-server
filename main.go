package main

import (
	"github.com/dpjungmin/jellypi-server/api"
	db "github.com/dpjungmin/jellypi-server/database"
	_ "github.com/dpjungmin/jellypi-server/docs"
)

func init() {
	db.Connect()
	db.AutoMigrate()
}

func main() {
	api.StartApplication()
}
