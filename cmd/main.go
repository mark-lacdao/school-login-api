package main

import (
	"log"
	"net/http"
	"school/internal/routes"
)

func main() {
	r := routes.InitializeRoutes()
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
