package api

import (
	"os"

	"gopkg.in/yaml.v3"
)

type dbConfig struct {
	login    string `yaml:"login"`
	password string `yaml:"password"`
	address  string `yaml:"address"`
}

type Config struct {
	port uint `yaml:"port"`

	dbConfig
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
