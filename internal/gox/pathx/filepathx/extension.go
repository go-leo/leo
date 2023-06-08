package filepathx

import (
	"path/filepath"
	"strings"
)

func Extension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}
