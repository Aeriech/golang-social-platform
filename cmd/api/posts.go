package main

import (
	"net/http"

	"github.com/aeriech/social/internal/model"
	"github.com/aeriech/social/internal/validate"
)

type CreatePostRequest struct {
	Title   string  `json:"title" validate:"required"`
	Content string  `json:"content" validate:"required"`
	UserId  int64   `json:"user_id" validate:"required"`
	Tags    []int64 `json:"tags" validate:"required"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostRequest

	err := readJson(w, r, &payload)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.ValidateStruct(payload)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	var tags []model.Tag
	findResult := app.db.Find(&tags, payload.Tags)
	if findResult.Error != nil || len(tags) != len(payload.Tags) {
		errorJson(w, http.StatusInternalServerError, "failed to find tags")
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
		errorJson(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	err = writeJson(w, http.StatusCreated, post)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []model.Post
	findResult := app.db.Preload("Tags").Preload("User").Find(&posts)
	if findResult.Error != nil {
		errorJson(w, http.StatusInternalServerError, findResult.Error.Error())
		return
	}

	err := writeJson(w, http.StatusFound, posts)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err.Error())
		return
	}
}
