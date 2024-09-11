package main

import (
	"log"

	"github.com/athulmekkoth/go_server.git/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000", nil)
	if err := server.RUN(); err != nil {
		log.Fatal()
	}

}
