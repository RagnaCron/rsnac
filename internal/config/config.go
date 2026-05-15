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
	args := os.Args[1:]

	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	return &Config{
		Path: args[0],
	}, nil
}
