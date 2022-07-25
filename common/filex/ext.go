package filex

import (
	"path/filepath"
	"strings"
)

func ExtName(file string) string {
	extName := filepath.Ext(file)
	if strings.HasPrefix(extName, ".") {
		return extName[1:]
	}
	return extName
}
