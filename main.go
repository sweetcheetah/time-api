package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sweetcheetah/time-api/server"
)

func main() {
	router := mux.NewRouter()

	rest := server.NewServer(router)

	log.Printf("Starting time-api...")

	log.Fatal(rest.Start())
}
