package text

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func RelativePathOld(path string) string {
	_, filename, _, _ := runtime.Caller(0)
	if p, found := os.LookupEnv("RELATIVE_PATH"); found {
		return filepath.Join(p, path)
	}
	return filepath.Join(filepath.Dir(filename), "..", "..", path)
}

func RelativePath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(wd, path)
}
