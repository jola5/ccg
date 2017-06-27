package isdir

import "os"

// IsDirectory : Check if the given path string points to a directory
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}
