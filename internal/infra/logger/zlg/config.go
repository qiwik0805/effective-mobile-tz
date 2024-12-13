package zlg

import "os"

type Config struct {
	LogLevel string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() Config {
	c.LogLevel = os.Getenv("LOG_LEVEL")

	return *c
}
