package main

import (
	"backend-test/config"
	"backend-test/server/database"
	"backend-test/server/routes"
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

	database.Connect()
	defer database.Conn.Close()

	routes.Run(port)
}
