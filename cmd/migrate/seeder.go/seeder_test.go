package seeder

import (
	"log"
	"testing"

	"github.com/aeriech/social/internal/model"
	dbConfig "github.com/aeriech/social/internal/postgressDb"
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

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

func TestFreshSeed(t *testing.T) {
	dropAllTables(db)
	migrateTables(db)

	seedUsers(db)
	seedTags(db)
	seedPosts(db)
}

func TestMigrate(t *testing.T) {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

func TestDropTable(t *testing.T) {
	dropAllTables(db)
}
