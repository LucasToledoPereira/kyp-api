package utils

import (
	"path/filepath"
	"runtime"
)

func Path() string {
	_, file, _, _ := runtime.Caller(0)
	abs, _ := filepath.Abs(file + "/../../..")

	return abs
}
