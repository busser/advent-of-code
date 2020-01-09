package day08

import (
	"bytes"
	"fmt"
	"io"

	"github.com/busser/adventofcode/util"
)

const width, height = 25, 6

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	layers, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	minLayer := layers[0]
	minCount := count(minLayer, '0')
	for _, l := range layers {
		c := count(l, '0')
		if c < minCount {
			minLayer, minCount = l, c
		}
	}

	_, err = fmt.Fprintf(w, "%d", count(minLayer, '1')*count(minLayer, '2'))
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	layers, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	image := make([][]byte, height)
	for y := range image {
		image[y] = make([]byte, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for l := 0; l < len(layers); l++ {
				if layers[l][y][x] == '2' {
					continue
				}
				image[y][x] = layers[l][y][x]
				break
			}
		}
	}

	_, err = fmt.Fprintf(w, "%s", bytes.Join(image, []byte{'\n'}))
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func count(layer [][]byte, digit byte) int {
	var c int

	for _, row := range layer {
		for _, d := range row {
			if d == digit {
				c++
			}
		}
	}

	return c
}

func readInput(r io.Reader) ([][][]byte, error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	digits := []byte(lines[0])

	layers := make([][][]byte, len(digits)/width/height)
	for l := range layers {
		layers[l] = make([][]byte, height)
		for y := range layers[l] {
			layers[l][y] = digits[:width]
			digits = digits[width:]
		}
	}

	return layers, nil
}
