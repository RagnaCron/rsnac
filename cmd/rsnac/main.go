// Package main
package main

import (
	"log"
	"path/filepath"

	"github.com/RagnaCron/rsnac/internal/app"
	"github.com/RagnaCron/rsnac/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	root, err := filepath.Abs(cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(root, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
