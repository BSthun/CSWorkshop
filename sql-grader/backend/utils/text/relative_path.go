package text

import (
	"os"
	"path/filepath"
	"runtime"
)

func RelativePath(path string) string {
	_, filename, _, _ := runtime.Caller(0)
	if p, found := os.LookupEnv("RELATIVE_PATH"); found {
		return filepath.Join(p, path)
	}
	return filepath.Join(filepath.Dir(filename), "..", "..", path)
}
