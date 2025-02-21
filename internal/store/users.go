package store

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"-" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersStore struct {
	db *gorm.DB
}

func (s *UsersStore) Create(context context.Context, user *User) error {
	err := validateStruct(user)
	if err != nil {
		return err
	}
	
	result := s.db.Create(&user)

	return result.Error
}
