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

type PostWithCommentsCount struct {
	model.Post
	CommentsCount int64 `json:"comments_count"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required,max=200"`
}

var (
	errPostNotFound = errors.New("post not found")
)

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

	if post.ID == 0 {
		app.notFoundError(w, r, errPostNotFound)
		return
	}

	err := writeJson(w, http.StatusFound, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []PostWithCommentsCount

	findResult := app.db.Table("posts").
		Select("posts.*, COUNT(comments.id) as comments_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Preload("User").
		Group("posts.id").
		Find(&posts)

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

func (app *application) getPostListHandler(w http.ResponseWriter, r *http.Request) {
	var posts []PostWithCommentsCount

	payload, err := getFilter(w, r)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	findResult := app.db.Table("posts").
		Select("posts.*, COUNT(comments.id) as comments_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Preload("User").
		Group("posts.id").
		Limit(payload.PerPage).
		Offset(payload.Offset).
		Find(&posts)

	if findResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	err = writeJson(w, http.StatusFound, posts)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	var post model.Post
	findResult := app.db.Find(&post, postID)
	if findResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	if post.ID == 0 {
		app.notFoundError(w, r, errPostNotFound)
		return
	}

	var payload UpdatePostRequest
	err := readJson(w, r, &payload)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if payload.Title == post.Title && payload.Content == post.Content {
		app.badRequestError(w, r, errors.New("title and content are the same"))
		return
	}

	post.Title = payload.Title
	post.Content = payload.Content

	saveResult := app.db.Save(&post)
	if saveResult.Error != nil {
		app.internalServerError(w, r, findResult.Error)
		return
	}

	err = writeJson(w, http.StatusFound, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	var post model.Post
	result := app.db.Delete(&post, postID)
	if result.Error != nil {
		app.internalServerError(w, r, result.Error)
		return
	}

	if post.ID == 0 {
		app.notFoundError(w, r, errPostNotFound)
		return
	}

	err := writeJson(w, http.StatusFound, postID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
