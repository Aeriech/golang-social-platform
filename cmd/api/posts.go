package main

import (
	"net/http"

	"github.com/aeriech/social/internal/store"
)

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostRequest

	err := readJson(w, r, &payload)
	if err != nil {
		errorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID: 1,
		Tags: []store.Tag{
			{ID: 1, Name: "1st tag"},
		},
	}

	context := r.Context()

	err = app.store.Posts.Create(context, post)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = writeJson(w, http.StatusCreated, post)
	if err != nil {
		errorJson(w, http.StatusInternalServerError, err.Error())
		return
	}
}
