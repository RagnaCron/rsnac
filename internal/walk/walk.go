// Package walk
package walk

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RagnaCron/rsnac/internal/config"
	"github.com/RagnaCron/rsnac/internal/normalize"
)

type RenamePlan struct {
	Type    string
	OldPath string
	NewPath string
}

func ProcessDir(root string, cfg *config.Config) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("read dir %q: %w", root, err)
	}

	plans := buildRenamePlans(root, entries)

	valid, err := validatePlans(plans)
	if err != nil {
		return err
	}

	dirs, err := executePlans(valid, cfg)
	if err != nil {
		return err
	}

	for _, path := range dirs {
		err := ProcessDir(path, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildRenamePlans(root string, entries []os.DirEntry) []RenamePlan {
	plans := make([]RenamePlan, 0, len(entries))

	for _, entry := range entries {
		oldName := entry.Name()

		if strings.HasPrefix(oldName, ".") {
			continue // Skip hidden
		}

		var (
			newName string
			typ     string
		)

		if entry.IsDir() {
			newName = normalize.ToSnakeCaseDir(oldName)
			typ = "dir"
		} else {
			newName = normalize.ToSnakeCase(oldName)
			typ = "file"
		}

		if newName == oldName {
			continue
		}

		plans = append(plans, RenamePlan{
			Type:    typ,
			OldPath: filepath.Join(root, oldName),
			NewPath: filepath.Join(root, newName),
		})
	}

	return plans
}

func validatePlans(plans []RenamePlan) ([]RenamePlan, error) {
	valid := make([]RenamePlan, 0, len(plans))
	seen := map[string]struct{}{}
	for _, plan := range plans {
		_, err := os.Lstat(plan.OldPath)
		if err != nil {
			return nil, err
		}

		_, err = os.Lstat(plan.NewPath)
		if err == nil {
			log.Printf("collision: %q already exists\n", plan.NewPath)
			continue
		}

		if errors.Is(err, os.ErrNotExist) {
			if _, ok := seen[plan.NewPath]; ok {
				log.Printf("collision: %q already seen\n", plan.NewPath)
				continue
			}

			valid = append(valid, plan)
			seen[plan.NewPath] = struct{}{}
			continue
		}

		return nil, err
	}
	return valid, nil
}

func executePlans(plans []RenamePlan, cfg *config.Config) ([]string, error) {
	return nil, nil
}
