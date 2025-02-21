package store

import (
	"context"

	"gorm.io/gorm"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content" validate:"required"`
	Title     string   `json:"title" validate:"required"`
	UserId    int64    `json:"user_id" validate:"required"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostsStore struct {
	db *gorm.DB
}

func (s *PostsStore) Create(context context.Context, post *Post) error {
	err := validateStruct(post)
	if err != nil {
		return err
	}
	
	result := s.db.Create(&post)

	return result.Error
}
