// Package config
package config

import "errors"

type Config struct {
	Path string
}

func Load() (Config, error) {
	return Config{}, errors.New("config Load not implemented")
}
