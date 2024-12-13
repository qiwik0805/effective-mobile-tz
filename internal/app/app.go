package app

import (
	"context"
	"effective-mobile-tz/internal/app/song/server"
	"github.com/rs/zerolog/log"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(ctx context.Context) error {
	appConfig := NewConfig().Load()

	log.Info().Msgf("server starting")
	defer log.Info().Msgf("server closing")
	return server.NewServer(appConfig.songServerConfig).Run(ctx)
}
