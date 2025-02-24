package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Get("/", app.getPostsHandler)
			r.Post("/list", app.getPostListHandler)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", app.getPostByIdHandler)
				r.Delete("/", app.deletePostHandler)
			})
		})
	})

	return r

}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.address)

	return server.ListenAndServe()
}

func getFilter(w http.ResponseWriter, r *http.Request) (filter, error) {
	var payload filter

	err := readJson(w, r, &payload)
	if err != nil {
		return payload, err
	}

	if payload.Page == 0 {
		payload.Page = 1
	}

	if payload.PerPage == 0 {
		payload.PerPage = 20
	}

	payload.Offset = (payload.Page - 1) * payload.PerPage

	return payload, nil
}
