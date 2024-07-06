package app

import (
	"mai-platform/internal/clients/db"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port uint `yaml:"port"`
	DB   db.Config
}

func NewConfig(configPath string) (*Config, error) {
	cfg, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal([]byte(cfg), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
