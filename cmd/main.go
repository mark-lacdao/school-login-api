package main

import (
	"log"
	"net/http"
	"school-login-api/internal/database"
	"school-login-api/internal/routes"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %s\n", err)
	}
	defer db.Close()

	r := routes.InitializeRoutes(db)
	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
