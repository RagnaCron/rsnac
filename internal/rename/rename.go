// Package rename
package rename

import (
	"fmt"
	"os"
)

func Rename(oldPath, newPath string) error {
	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("rename %q -> %q: %w", oldPath, newPath, err)
	}
	return nil
}
