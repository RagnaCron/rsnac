// Package walk
package walk

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/RagnaCron/rsnac/internal/config"
	"github.com/RagnaCron/rsnac/internal/normalize"
	"github.com/RagnaCron/rsnac/internal/rename"
)

type Type string

const (
	FILE Type = "FILE"
	LINK Type = "LINK"
	DIR  Type = "DIR"
)

type RenamePlan struct {
	Type    Type
	OldPath string
	NewPath string
	Valid   bool
}

func ProcessDir(root string, cfg *config.Config) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("read dir %q: %w", root, err)
	}

	plans := buildRenamePlans(root, entries)

	err = validatePlans(plans)
	if err != nil {
		return err
	}

	err = executePlans(plans, cfg)
	if err != nil {
		return err
	}

	dirs := collectDirs(plans, cfg)

	for _, path := range dirs {
		err := ProcessDir(path, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildRenamePlans(root string, entries []os.DirEntry) (plans []RenamePlan) {
	for _, entry := range entries {
		oldName := entry.Name()

		if strings.HasPrefix(oldName, ".") {
			continue // Skip hidden
		}

		var (
			newName string
			typ     Type
		)

		if entry.IsDir() {
			newName = normalize.ToSnakeCaseDir(oldName)
			typ = DIR
		} else if entry.Type()&fs.ModeSymlink != 0 {
			newName = normalize.ToSnakeCase(oldName)
			typ = LINK
		} else {
			newName = normalize.ToSnakeCase(oldName)
			typ = FILE
		}

		plans = append(plans, RenamePlan{
			Type:    typ,
			OldPath: filepath.Join(root, oldName),
			NewPath: filepath.Join(root, newName),
		})
	}

	return plans
}

func validatePlans(plans []RenamePlan) error {
	seen := map[string]struct{}{}
	for i, plan := range plans {
		_, err := os.Lstat(plan.NewPath)
		if err == nil {
			return fmt.Errorf("collision: %s already exists", plan.NewPath)
		}

		if errors.Is(err, os.ErrNotExist) {
			if _, ok := seen[plan.NewPath]; ok {
				return fmt.Errorf("collision: %s already seen", plan.NewPath)
			}

			plans[i].Valid = plan.NewPath != plan.OldPath
			seen[plan.NewPath] = struct{}{}
			continue
		}

		return err
	}
	return nil
}

func executePlans(plans []RenamePlan, cfg *config.Config) error {
	for _, plan := range plans {
		if plan.Valid {
			if !cfg.DryRun {
				err := rename.Rename(plan.OldPath, plan.NewPath)
				if err != nil {
					return err
				}
			}

			fmt.Printf("%s: %q -> %q\n", plan.Type, plan.OldPath, plan.NewPath)
		}
	}

	return nil
}

func collectDirs(plans []RenamePlan, cfg *config.Config) []string {
	var dirs []string
	for _, plan := range plans {
		if plan.Type == DIR {
			path := plan.NewPath
			if cfg.DryRun {
				path = plan.OldPath
			}
			dirs = append(dirs, path)
		}
	}
	return dirs
}
