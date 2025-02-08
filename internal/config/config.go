package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yaml:"port"`
}

func NewConfig(path string) (Config, error) {
	var cfg Config
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	if err = yaml.Unmarshal(file, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
