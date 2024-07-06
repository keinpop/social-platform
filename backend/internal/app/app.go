package app

import (
	"mai-platform/internal/clients/db"
)

type App struct {
	config *Config
	DB     *db.DB
}

func NewApp(cfg *Config) *App {
	return &App{
		config: cfg,
		DB:     db.NewDB(&cfg.DB),
	}
}

func (app *App) Init() error {
	if err := app.DB.Init(); err != nil {
		return err
	}

	if err := app.DB.Migrate(); err != nil {
		return err
	}

	return nil
}
