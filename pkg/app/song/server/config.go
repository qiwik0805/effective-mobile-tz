package server

import (
	songRepository "effective-mobile-tz/pkg/app/song/repository/song"
	"effective-mobile-tz/pkg/infra/service/music_info"
)

type Config struct {
	MusicInfoConfig      music_info.Config
	SongRepositoryConfig songRepository.Config
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() Config {
	c.MusicInfoConfig = music_info.NewConfig().Load()
	c.SongRepositoryConfig = songRepository.NewConfig().Load()
	return *c
}
