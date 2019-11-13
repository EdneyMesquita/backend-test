package main

import (
	"backend-test/config"
	"backend-test/server/database"
	"backend-test/server/handlers/workflows"
	"backend-test/server/routes"
	"container/list"
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

	workflows.Queue = list.New()
	workflows.MountQueue()

	routes.Run(port)
}
