package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func findInputFile(path string) (string, error) {
	if _, err := os.Stat(path); err == nil {
		return path, err
	}

	for skip := 1; skip < 5; skip++ {
		_, callerFile, _, ok := runtime.Caller(skip)
		if ok {
			callerDir := filepath.Dir(callerFile)
			dayDir := filepath.Base(callerDir)

			altPath := filepath.Join(dayDir, path)
			if _, err := os.Stat(altPath); err == nil {
				return altPath, nil
			}
		}
	}
	return path, nil
}

// ReadInput reads the entire input file as a string
func ReadInput(filepath string) (string, error) {
	resolvedPath, _ := findInputFile(filepath)
	data, err := os.ReadFile(resolvedPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads input file and returns slice of lines
func ReadLines(filepath string) ([]string, error) {
	content, err := ReadInput(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(content), "\n"), nil
}
