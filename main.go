package main

import (
	"leva-api/config"
	"leva-api/pkg/orm/connection"
	"leva-api/server/routes"
	"log"
	"os"
)

func main() {
	var port string
	if config.Env == "DEV" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}

	log.Printf("Server started at port %s\n\n", port)

	connection.Connect()
	defer connection.Conn.Close()

	routes.Run(port)
}
