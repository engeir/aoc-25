package utils

import (
	"os"
	"strings"
)

// ReadInput reads the entire input file as a string
func ReadInput(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
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
