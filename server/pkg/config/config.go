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
	Usecases struct {
		PackSizes []uint64 `yaml:"packSizes"`
	} `yaml:"usecases"`
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

	return &config, nil
}
