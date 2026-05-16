// Package rename
package rename

import (
	"fmt"
	"os"
)

func Rename(oldPath, newPath string) error {
	oStat, err := os.Lstat(oldPath)
	if err != nil {
		return fmt.Errorf("error while reading lstat from: %q, with: %w", oldPath, err)
	}
	nStat, err := os.Lstat(newPath)
	if err != nil {
		return fmt.Errorf("error while reading lstat from: %q, with: %w", newPath, err)
	}

	if oStat.Name() == nStat.Name() {
		return fmt.Errorf("collision: %q already exists", newPath)
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("rename %q -> %q: %w", oldPath, newPath, err)
	}

	return nil
}
