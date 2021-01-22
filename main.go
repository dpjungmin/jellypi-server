package main

import (
	"fmt"
	"log"

	"github.com/dpjungmin/jellypi-server/api"
	"github.com/dpjungmin/jellypi-server/database"
	_ "github.com/dpjungmin/jellypi-server/docs"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	api.Init()
}
