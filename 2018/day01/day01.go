package day01

import (
	"bufio"
	"io"
	"strconv"
)

// Freq represents a frequency.
type Freq int

// FreqChange represents a frequency change.
type FreqChange int

// Apply applies a frequency change to f.
func (f *Freq) Apply(change FreqChange) {
	*f += Freq(change)
}

// PartOne solves Part One of the puzzle.
func PartOne(changes []FreqChange) Freq {
	var freq Freq

	for _, change := range changes {
		freq.Apply(change)
	}

	return freq
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(changes []FreqChange) Freq {
	var freq Freq

	seen := make(map[Freq]bool)
	seen[freq] = true

	for {
		for _, change := range changes {
			freq.Apply(change)

			if seen[freq] {
				return freq
			}

			seen[freq] = true
		}
	}
}

// ReadFreqChanges reads frequency changes from r.
func ReadFreqChanges(r io.Reader) ([]FreqChange, error) {
	var changes []FreqChange

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		changes = append(changes, FreqChange(x))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return changes, nil
}
