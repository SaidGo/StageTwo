package main

import (
	"log"

	"example.com/local/Go2part/internal/app"
)

func main() {
	a, err := app.InitApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := a.Run(":8080"); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
