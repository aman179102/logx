package io

import (
	"bufio"
	"os"
)

// NewLineScanner returns a bufio.Scanner for the given input path.
// If path is empty, it reads from stdin.
// It returns the scanner and a close function that should be called to release resources.
// For stdin, the close function is a no-op.
func NewLineScanner(path string) (*bufio.Scanner, func(), error) {
	var f *os.File
	if path == "" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(path)
		if err != nil {
			return nil, nil, err
		}
	}
	scanner := bufio.NewScanner(f)
	// Increase buffer size for long lines (up to 1MB)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	return scanner, func() {
		if path != "" {
			f.Close()
		}
	}, nil
}