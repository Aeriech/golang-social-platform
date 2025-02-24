package seeder

import (
	"log"
	"testing"

	"github.com/aeriech/social/internal/model"
	"github.com/aeriech/social/internal/postgresDb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dns := postgresDb.GetDns()

	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error gorm.Open: ", err.Error())
	}

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

func TestFreshSeed(t *testing.T) {
	dropAllTables(db)
	migrateTables(db)

	seedAll(db)
}

func TestMigrate(t *testing.T) {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

func TestDropTable(t *testing.T) {
	dropAllTables(db)
}
