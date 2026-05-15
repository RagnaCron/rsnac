// Package config
package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Path   string
	DryRun bool
}

func Load() (*Config, error) {
	dryRun := flag.Bool("d", false, "dry run")
	flag.Parse()

	flagsCounter := 1

	if *dryRun {
		flagsCounter += 1
	}

	args := os.Args[flagsCounter:]
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 positional argument, got %d", len(args))
	}

	return &Config{
		Path:   args[0],
		DryRun: *dryRun,
	}, nil
}
