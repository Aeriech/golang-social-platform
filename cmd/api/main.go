package main

import (
	"log"

	"github.com/aeriech/social/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newConfig := config{
		address: env.GetString("ADDRESS", ":8080"),
	}

	app := &application{
		config: newConfig,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
