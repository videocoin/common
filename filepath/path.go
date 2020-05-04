package filepath

import (
	"path/filepath"
	"strings"
)

// MaybeSymlink ...
func MaybeSymlink(path string) string {
	path = strings.TrimSpace(path)
	symPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path
	}
	return symPath
}
