package api

type App struct {
	config *Config
}

func NewApp(cfg *Config) *App {
	return &App{config: cfg}
}
