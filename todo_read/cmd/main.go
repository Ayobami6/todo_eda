package main

import (
	"log"

	"github.com/change_me/go_starter_rest/cmd/api"
)

// main entry point for the application

func main() {
	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
