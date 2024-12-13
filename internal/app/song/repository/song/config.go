package song

import "os"

type Config struct {
	DSN string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() Config {
	c.DSN = os.Getenv("SONG_REPOSITORY_DSN")
	return *c
}
