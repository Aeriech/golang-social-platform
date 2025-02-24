package seeder

import (
	"github.com/aeriech/social/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

const (
	_RANDOM_ORDER = "RANDOM()"
)

func dropAllTables(db *gorm.DB) {
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}

func migrateTables(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{}, &model.Comment{})
}

func seedAll(db *gorm.DB) {
	seedUsers(db)
	seedTags(db)
	seedPosts(db)
	seedComments(db)
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
		db.Order(_RANDOM_ORDER).First(&tag)

		user := model.User{}
		db.Order(_RANDOM_ORDER).First(&user)

		post := model.Post{
			Title:   gofakeit.Sentence(5),
			Content: gofakeit.Sentence(10),
			UserID:  int64(user.Model.ID),
			Tags: []model.Tag{
				tag,
			},
		}

		db.Create(&post)
	}
}

func seedComments(db *gorm.DB) {
	for count := 0; count < 100; count++ {
		post := model.Post{}
		db.Order(_RANDOM_ORDER).First(&post)

		user := model.User{}
		db.Order(_RANDOM_ORDER).First(&user)

		comment := model.Comment{
			Content: gofakeit.Sentence(10),
			UserID:  int64(user.Model.ID),
			PostID:  int64(post.Model.ID),
		}

		db.Create(&comment)
	}
}
