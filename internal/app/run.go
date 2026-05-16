// Package app
package app

import (
	"fmt"
	"os"

	"github.com/RagnaCron/rsnac/internal/config"
	"github.com/RagnaCron/rsnac/internal/walk"
)

func Run(root string, cfg *config.Config) error {
	info, err := os.Lstat(root)
	if err != nil {
		return fmt.Errorf("err reading lstat: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("root path is not a dir: %s, %w", root, err)
	}

	err = walk.ProcessDir(root, cfg)
	if err != nil {
		return err
	}

	return nil
}
