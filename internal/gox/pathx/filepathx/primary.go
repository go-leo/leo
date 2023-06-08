package filepathx

import (
	"path/filepath"
	"strings"
)

func Primary(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}
