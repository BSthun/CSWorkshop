package text

import (
	"path/filepath"
	"runtime"
)

func RelativePath(path string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "..", path)
}
