package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content       string    `json:"content,omitempty" gorm:"not null"`
	Title         string    `json:"title,omitempty" gorm:"not null"`
	UserID        int64     `json:"user_id,omitempty" gorm:"not null"`
	User          *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`    // Belongs to a User
	Tags          []Tag     `json:"tags,omitempty" gorm:"many2many:post_tags;"` // Many-to-Many relationship
	Comments      []Comment `json:"comments,omitempty"`                         // One-To-Many relationship
}

type User struct {
	gorm.Model
	Name  string `json:"name,omitempty" gorm:"not null"`
	Email string `json:"email,omitempty" gorm:"uniqueIndex;not null"` // Ensures email uniqueness
	Posts []Post `json:"posts,omitempty"`                             // One-to-Many: A user can have multiple posts
}

type Tag struct {
	gorm.Model
	Name  string `json:"name,omitempty" gorm:"not null"`
	Posts []Post `json:"posts,omitempty" gorm:"many2many:post_tags;"` // Many-to-Many relationship
}

type Comment struct {
	gorm.Model
	PostID  int64  `json:"post_id,omitempty" gorm:"not null"`
	Post    *Post  `json:"post,omitempty" gorm:"foreignKey:PostID"` // Belongs to a Post
	UserID  int64  `json:"user_id,omitempty" gorm:"not null"`
	User    *User  `json:"user,omitempty" gorm:"foreignKey:UserID"` // Belongs to a User
	Content string `json:"content,omitempty" gorm:"not null"`
}
