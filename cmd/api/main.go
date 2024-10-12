package main

import (
	"log"
	"net/http"

	"github.com/TheLinuxSham/golang-fortune/internal/server"
)

func main() {
	server := server.NewServer()

	log.Println("Starting server on :3030")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :3030: %v\n", err)
	}
}
