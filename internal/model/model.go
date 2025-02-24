package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey"`
	Content  string
	Title    string
	UserID   int64
	User     User  `gorm:"foreignKey:UserID"`    // Belongs to a User
	Tags     []Tag `gorm:"many2many:post_tags;"` // Many-to-Many relationship
	Comments []Comment // One-To-Many relationship
}

type User struct {
	gorm.Model
	ID    int64 `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"` // Ensures email uniqueness
	Posts []Post // One-to-Many: A user can have multiple posts
}

type Tag struct {
	gorm.Model
	ID    int64 `gorm:"primaryKey"`
	Name  string
	Posts []Post `gorm:"many2many:post_tags;"` // Many-to-Many relationship
}

type Comment struct {
	gorm.Model
	ID      int64 `gorm:"primaryKey"`
	PostID  int64
	Post    Post `gorm:"foreignKey:PostID"` // Belongs to a Post
	UserID  int64
	User    User `gorm:"foreignKey:UserID"` // Belongs to a User
	Content string
}
