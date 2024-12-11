package main

import (
	"context"
	"effective-mobile-tz/pkg/app"
	"fmt"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	if err := app.NewApp().Run(ctx); err != nil {
		log.Fatal().Msgf(fmt.Sprintf("app run: %s", err))
	}
}
