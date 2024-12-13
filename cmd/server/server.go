package main

import (
	"context"
	_ "effective-mobile-tz/docs"
	"effective-mobile-tz/internal/app"
	"effective-mobile-tz/internal/infra/logger/zlg"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal().Msg(fmt.Sprintf("godotenv load: %s", err))
	}

	ctx := context.Background()
	zerologConfig := zlg.NewConfig().Load()
	zerologLevel := zlg.Level(zerologConfig)
	zerolog.SetGlobalLevel(zerologLevel)

	log.Info().Msgf("app starting")
	defer log.Info().Msgf("app closing")
	if err := app.NewApp().Run(ctx); err != nil {
		log.Fatal().Msgf(fmt.Sprintf("app run: %s", err))
	}

}
