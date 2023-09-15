package pack_handler

import (
	"gimshark-test/ui/pkg/config"
	"time"
)

// Config for packs handler
type Config struct {
	Port          string
	Host          string
	Endpoint      string
	Timeout       time.Duration
	MinJitterWait time.Duration
	MaxJitterWait time.Duration
}

// Returns new instance of Config
func NewConfig(cnf *config.Config) *Config {
	return &Config{
		Port:          cnf.Packs.Port,
		Host:          cnf.Packs.Host,
		Endpoint:      cnf.Packs.Endpoint,
		Timeout:       cnf.Packs.Timeout,
		MinJitterWait: cnf.Packs.MinJitterWait,
		MaxJitterWait: cnf.Packs.MaxJitterWait,
	}
}
