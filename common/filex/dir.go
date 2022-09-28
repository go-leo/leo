package filex

import "os"

// IsDir reports whether the named file is a directory.
// Deprecated: Do not use. use github.com/go-leo/filex instead.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsDirectory reports whether the named file is a directory.
// Deprecated: Do not use. use github.com/go-leo/filex instead.
var IsDirectory = IsDir
