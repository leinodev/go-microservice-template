package config

import "time"

type Service struct {
	Address         string        `yaml:"address" env:"ADDRESS"`
	Name            string        `yaml:"name"`
	Development     bool          `yaml:"development" env:"DEVELOPMENT"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"SHUTDOWN_TIMEOUT"`
}
