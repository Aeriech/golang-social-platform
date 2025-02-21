package seeder

import (
	"github.com/aeriech/social/internal/store"
	"gorm.io/gorm"
)

func dropAllTables(db *gorm.DB) {
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}

func migrateTables(db *gorm.DB) {
	db.AutoMigrate(&store.User{}, &store.Post{}, &store.Tag{})
}

func seedUser(db *gorm.DB) {
	users := []*store.User{
		{
			Name:  "Aeriech",
			Email: "aeriech@gmail.com",
		},
		{
			Name:  "Test User",
			Email: "test-user@gmail.com",
		},
	}

	db.Create(users)
}
