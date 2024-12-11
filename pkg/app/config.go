package app

import (
	songServer "effective-mobile-tz/pkg/app/song/server"
)

type Config struct {
	songServerConfig songServer.Config
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() Config {
	c.songServerConfig = songServer.NewConfig().Load()
	return *c
}
