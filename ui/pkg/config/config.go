package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Config - config for app.
type Config struct {
	HTTP struct {
		Port              string        `yaml:"port"`
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	} `yaml:"http"`
	Packs struct {
		Port          string        `yaml:"port"`
		Host          string        `yaml:"host"`
		Endpoint      string        `yaml:"endpoint"`
		MinJitterWait time.Duration `yaml:"minJitterWait"`
		MaxJitterWait time.Duration `yaml:"maxJitterWait"`
		Timeout       time.Duration `yaml:"timeout"`
	} `yaml:"packs"`
}

// Parse - parse config from file.
func Parse(confPath string) (*Config, error) {
	filename, err := filepath.Abs(confPath)
	if err != nil {
		return nil, fmt.Errorf("can't get config path: %w", err)
	}

	yamlConf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read conf: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(yamlConf, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall conf: %w", err)
	}

	config.Packs.Host = getEnv("SERVER_HOST", config.Packs.Host)
	config.Packs.Port = getEnv("SERVER_PORT", config.Packs.Port)

	return &config, nil
}

// getEnv - get env value or default.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
