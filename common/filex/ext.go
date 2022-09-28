package filex

import (
	"path/filepath"
	"strings"
)

// Deprecated: Do not use. use github.com/go-leo/filex instead.
func ExtName(file string) string {
	extName := filepath.Ext(file)
	if strings.HasPrefix(extName, ".") {
		return extName[1:]
	}
	return extName
}
