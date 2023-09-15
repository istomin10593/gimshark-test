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
		return nil, fmt.Errorf("can't get config path: %s", err.Error())
	}

	yamlConf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read conf: %s", err.Error())
	}

	var config Config

	err = yaml.Unmarshal(yamlConf, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall conf: %s", err.Error())
	}

	return &config, nil
}
