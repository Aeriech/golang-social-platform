package seeder

import (
	"log"
	"testing"

	dbConfig "github.com/aeriech/social/internal/postgressDb"
	"github.com/aeriech/social/internal/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dns := dbConfig.GetDns()

	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error gorm.Open: ", err.Error())
	}

	db.AutoMigrate(&store.User{}, &store.Post{}, &store.Tag{})
}

func TestFreshSeed(t *testing.T)  {
	dropAllTables(db)
	migrateTables(db)

	seedUser(db)
}

func TestMigrate(t *testing.T) {
	db.AutoMigrate(&store.User{}, &store.Post{}, &store.Tag{})
}

func TestDropTable(t *testing.T) {
	dropAllTables(db)
}
