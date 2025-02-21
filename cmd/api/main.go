package main

import (
	"log"

	"github.com/aeriech/social/internal/env"
	dbConfig "github.com/aeriech/social/internal/postgressDb"
	"github.com/aeriech/social/internal/store"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newConfig := config{
		address: env.GetString("ADDRESS", ":8080"),
		dns:     dbConfig.GetDns(),
	}

	db, err := gorm.Open(postgres.Open(newConfig.dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error gorm.Open: ", err.Error())
	}
	log.Println("Connected to database")

	store := store.NewStorage(db)

	app := &application{
		config: newConfig,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
