package auth

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port uint   `yaml:"port"`
	Host string `yaml:"host"`
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
