// Package files contains common file operations.
package files

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileExists return true if specified file exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

// ReadLines read file lines into a slice.
func ReadLines(path string) ([]string, error) {
	r := os.Stdin

	if path != "" {
		file, err := os.Open(filepath.Clean(path))

		if err != nil {
			return nil, err
		}

		defer file.Close()

		r = file
	}

	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadBytes reads bytes from specified file or from STDIN if file name is empty.
func ReadBytes(path string) ([]byte, error) {
	r := os.Stdin

	if path != "" {
		file, err := os.Open(filepath.Clean(path))

		if err != nil {
			return nil, err
		}

		defer file.Close()

		r = file
	}

	return ioutil.ReadAll(r)
}
