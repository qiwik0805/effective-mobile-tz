package music_info

import "os"

type Config struct {
	BaseURL string
	UseFake string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() Config {
	c.BaseURL = os.Getenv("MUSIC_INFO_BASE_URL")
	c.UseFake = os.Getenv("MUSIC_INFO_USE_FAKE")
	return *c
}
