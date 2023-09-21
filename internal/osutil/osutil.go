// Package osutil provides utility functions for working with the operating system.
package osutil

import (
	"os"
	"path/filepath"
)

// UserDataDir returns the path to brane's data directory.
func UserDataDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ".local", "share", "brane")
}
