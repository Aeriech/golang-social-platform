package store

import (
	"context"

	"gorm.io/gorm"
)

type UsersStore struct {
	db *gorm.DB
}

func (s *UsersStore) Create(context context.Context, user *User) error {
	result := s.db.Create(&user)

	return result.Error
}
