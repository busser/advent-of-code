// Package util provides helper functions for writing Advent of Code solutions.
package util

import (
	"bufio"
	"fmt"
	"io"
)

// ReadLines returns a slice of all lines in r.
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	return lines, nil
}
