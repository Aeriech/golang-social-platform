package main

import (
	"github.com/aeriech/social/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	address string
	dns     string
}
