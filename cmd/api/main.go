package main

import (
	"log"

	"github.com/aeriech/social/internal/env"
	"github.com/aeriech/social/internal/model"
	dbConfig "github.com/aeriech/social/internal/postgressDb"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	version = "0.0.1"
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
		env:     env.GetString("ENV", "development"),
	}

	db, err := gorm.Open(postgres.Open(newConfig.dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error gorm.Open: ", err.Error())
	}
	log.Println("Connected to database")

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})

	app := &application{
		config: newConfig,
		db:     db,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
