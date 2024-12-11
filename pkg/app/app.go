package app

import (
	"context"
	"effective-mobile-tz/pkg/app/song/server"
	"fmt"
	"github.com/joho/godotenv"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(ctx context.Context) error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("godotenv load: %w", err)
	}
	appConfig := NewConfig().Load()

	return server.NewServer(appConfig.songServerConfig).Run(ctx)
}
