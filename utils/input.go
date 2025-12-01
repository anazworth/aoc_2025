package utils

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput reads the input file for a given path
func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		return ""
	}
	return string(data)
}

func Lines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
