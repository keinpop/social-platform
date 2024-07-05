package api

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}

func NewH(db *gorm.DB) Handler {
	return Handler{db}
}

type App struct {
	config *Config

	Handler
}

func NewApp(cfg *Config) *App {
	return &App{config: cfg}
}
