package store

import (
	"context"

	"gorm.io/gorm"
)

type PostsStore struct {
	db *gorm.DB
}

func (s *PostsStore) Create(context context.Context, post *Post) error {
	result := s.db.Create(&post)

	return result.Error
}
