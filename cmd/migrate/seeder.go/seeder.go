package seeder

import (
	"github.com/aeriech/social/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func dropAllTables(db *gorm.DB) {
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}

func migrateTables(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

func seedUsers(db *gorm.DB) {
	users := []model.User{
		{
			Name:  "Aeriech Ancheta",
			Email: "aeriech@gmail.com",
		},
	}

	for count := 0; count < 9; count++ {
		users = append(users, model.User{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		})
	}

	db.Create(&users)
}

func seedTags(db *gorm.DB) {
	tags := []model.Tag{}

	for count := 0; count < 10; count++ {
		tags = append(tags, model.Tag{
			Name: gofakeit.EmojiTag(),
		})
	}

	db.Create(&tags)
}

func seedPosts(db *gorm.DB) {
	for count := 0; count < 10; count++ {
		tag := model.Tag{}
		db.Order("RANDOM()").First(&tag)

		user := model.User{}
		db.Order("RANDOM()").First(&user)

		post := model.Post{
			Title:   gofakeit.Sentence(5),
			Content: gofakeit.Sentence(10),
			UserID:  user.ID,
			Tags: []model.Tag{
				tag,
			},
		}

		db.Create(&post)
	}
}
