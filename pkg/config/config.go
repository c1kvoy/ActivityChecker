package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type URL struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Config struct {
	URLs []URL `yaml:"urls"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	return &config, nil
} 