package main

import (
	"gorm.io/gorm"
)

type application struct {
	config config
	db     *gorm.DB
}

type config struct {
	address string
	dns     string
	env     string
}
