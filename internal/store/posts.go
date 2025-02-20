package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	Id        int64    `json:"id"`
	Content   string   `json:"content" validate:"required"`
	Title     string   `json:"title" validate:"required"`
	UserId    int64    `json:"user_id" validate:"required"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(context context.Context, post *Post) error {
	err := validateStruct(post)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO posts (content, title, user_id)
	VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`

	err = s.db.QueryRowContext(
		context,
		query,
		post.Content, post.Title, post.UserId, pq.Array(post.Tags),
	).Scan(
		&post.Id, &post.CreatedAt, &post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
