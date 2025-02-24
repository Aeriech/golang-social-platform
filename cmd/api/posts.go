package main

import (
	"errors"
	"net/http"

	"github.com/aeriech/social/internal/model"
	"github.com/aeriech/social/internal/validate"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm/clause"
)

type CreatePostRequest struct {
	Title   string  `json:"title" validate:"required,max=100"`
	Content string  `json:"content" validate:"required,max=200"`
	UserId  int64   `json:"user_id" validate:"required"`
	TagIds  []int64 `json:"tag_ids" validate:"array"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostRequest

	err := readJson(w, r, &payload)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	err = validate.ValidateStruct(payload)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	var tags []model.Tag
	findResult := app.db.Find(&tags, payload.TagIds)
	if findResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	if len(tags) != len(payload.TagIds) {
		app.notFoundError(w, r, errors.New("invalid tags"))
		return
	}

	post := &model.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserId,
		Tags:    tags,
	}

	result := app.db.Create(&post)
	if result.Error != nil {
		app.internalServerError(w, r, result.Error)
		return
	}

	err = writeJson(w, http.StatusCreated, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	var post model.Post
	findResult := app.db.Preload(clause.Associations).Find(&post, postID)
	if findResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	err := writeJson(w, http.StatusFound, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []model.Post
	findResult := app.db.Preload("User").Find(&posts)
	if findResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	err := writeJson(w, http.StatusFound, posts)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
