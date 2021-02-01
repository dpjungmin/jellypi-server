package main

import (
	app "github.com/dpjungmin/jellypi-server/api"
	db "github.com/dpjungmin/jellypi-server/database"
	_ "github.com/dpjungmin/jellypi-server/docs"
)

func init() {
	db.GetPGSingleton().AutoMigrate()
}

func main() {
	app.Start()
}
