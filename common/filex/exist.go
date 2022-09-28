package filex

import (
	"fmt"
	"os"
)

// IsExist returns a boolean indicating whether a file or directory exist.
// Deprecated: Do not use. use github.com/go-leo/filex instead.
func IsExist(filepath string) bool {
	info, err := os.Stat(filepath)
	fmt.Println(info)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
