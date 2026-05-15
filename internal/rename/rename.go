// Package rename
package rename

import (
	"fmt"
	"os"
)

func Rename(oldPath, newPath string) error {
	oStat, err := os.Lstat(oldPath)
	if err != nil {
		return fmt.Errorf("error while reading lstat with: %s, with: %w", oldPath, err)
	}
	nStat, err := os.Lstat(newPath)
	if err != nil {
		return fmt.Errorf("error while reading lstat with: %s, with: %w", newPath, err)
	}

	if oStat.Name() == nStat.Name() {
		return fmt.Errorf("collision: %s allready exists", newPath)
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("error while renaming: %w", err)
	}

	return nil
}
