package filex

import "os"

// IsDir reports whether the named file is a directory.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsDirectory reports whether the named file is a directory.
var IsDirectory = IsDir
