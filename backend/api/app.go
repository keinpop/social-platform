package api

import (
	"mai-platform/db"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewH(db *gorm.DB) Handler {
	return Handler{db}
}

type App struct {
	config *Config

	Handler
}

func NewApp(cfg *Config) *App {
	db := db.InitCompanyDB()

	return &App{
		config:  cfg,
		Handler: NewH(db),
	}
}
