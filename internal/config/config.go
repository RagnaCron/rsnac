// Package config
package config

import (
	"fmt"
	"os"
)

type Config struct {
	Path string
}

func Load() (*Config, error) {
	if len(os.Args[1:]) != 1 {
		return &Config{}, fmt.Errorf("error: pass 1 argument, got: %d", len(os.Args[1:]))
	}

	cfg := &Config{
		Path: os.Args[1],
	}

	return cfg, nil
}
