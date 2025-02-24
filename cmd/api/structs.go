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

type filter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Offset  int `json:"offset"`
}
