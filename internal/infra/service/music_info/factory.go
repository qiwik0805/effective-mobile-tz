package music_info

import (
	songService "effective-mobile-tz/internal/app/song/service/song"
	"net/http"
)

func Factory(cfg Config) songService.MusicInfoClient {
	var musicInfoClient songService.MusicInfoClient
	if cfg.UseFake == "" {
		httpTransport := &http.Client{}
		baseURL := cfg.BaseURL
		musicInfoClient = NewClient(httpTransport, baseURL)
	} else {
		musicInfoClient = NewFakeClient()
	}

	return musicInfoClient
}
