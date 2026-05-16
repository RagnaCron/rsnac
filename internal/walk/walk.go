// Package walk
package walk

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RagnaCron/rsnac/internal/config"
	"github.com/RagnaCron/rsnac/internal/normalize"
	"github.com/RagnaCron/rsnac/internal/rename"
)

func ProcessDir(root string, cfg *config.Config) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("cound not read %s content: %w", root, err)
	}

	// Categorize Folder content
	var filenames []string
	var foldernames []string
	var symlinks []string

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue // Skip hidden
		} else if entry.Type()&fs.ModeSymlink != 0 {
			symlinks = append(symlinks, entry.Name())
		} else if entry.IsDir() {
			foldernames = append(foldernames, entry.Name())
		} else {
			filenames = append(filenames, entry.Name())
		}
	}

	// Rename files in current dir
	for _, oldName := range filenames {
		newName := normalize.ToSnakeCase(oldName)
		if newName == oldName {
			continue
		}

		oldPath := filepath.Join(root, oldName)
		newPath := filepath.Join(root, newName)

		if !cfg.DryRun {
			err := rename.Rename(oldPath, newPath)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}

		fmt.Printf("FILE: %s -> %s\n", oldName, newName)
	}

	// Rename symlinks in current dir
	for _, oldName := range symlinks {
		newName := normalize.ToSnakeCase(oldName)
		if newName == oldName {
			continue
		}

		oldPath := filepath.Join(root, oldName)
		newPath := filepath.Join(root, newName)

		if !cfg.DryRun {
			err := rename.Rename(oldPath, newPath)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}

		fmt.Printf("LINK: %s -> %s\n", oldName, newName)
	}

	newFolderpaths := make([]string, 0, len(foldernames))

	// Rename dirs in current dir
	for _, oldName := range foldernames {
		newName := normalize.ToSnakeCaseDir(oldName)
		if newName == oldName {
			continue
		}

		oldPath := filepath.Join(root, oldName)
		newPath := filepath.Join(root, newName)

		if !cfg.DryRun {
			err := rename.Rename(oldPath, newPath)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}

		newFolderpaths = append(newFolderpaths, newPath)

		fmt.Printf("DIR: %s -> %s\n", oldName, newName)
	}

	for _, path := range newFolderpaths {
		err := ProcessDir(path, cfg)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return nil
}
