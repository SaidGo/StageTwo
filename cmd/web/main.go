package main

import (
	"log"
	"os"

	"example.com/local/Go2part/internal/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	r := web.NewRouter()
	log.Printf("Starting HTTP server on %s ...", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
