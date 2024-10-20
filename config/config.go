package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Service Service `yaml:"service"`
	// TODO: add configuration structs for any infrastructure
	// 	you need (DB, message broker, metrics etc.)
}

func NewConfig(path string) (*Config, error) {
	var config Config
	if err := config.readConfig(path); err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *Config) readConfig(path string) error {
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return err
	}
	return nil
}
