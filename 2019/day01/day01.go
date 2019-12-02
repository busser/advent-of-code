package day01

import (
	"fmt"
	"io"
	"strconv"

	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	modules, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	var totalFuel int
	for _, mod := range modules {
		totalFuel += fuelRequired(mod.mass)
	}

	_, err = fmt.Fprintf(w, "%d", totalFuel)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	modules, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	var totalFuel int
	for _, mod := range modules {
		totalFuel += fuelRequiredIncludingFuel(mod.mass)
	}

	_, err = fmt.Fprintf(w, "%d", totalFuel)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

type module struct {
	mass int
}

func fuelRequired(mass int) int {
	return mass/3 - 2
}

func fuelRequiredIncludingFuel(mass int) int {
	fuel := fuelRequired(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + fuelRequiredIncludingFuel(fuel)
}

func readInput(r io.Reader) ([]module, error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	modules := make([]module, len(lines))
	for i := range lines {
		mass, err := strconv.Atoi(lines[i])
		if err != nil {
			return nil, err
		}
		modules[i].mass = mass
	}

	return modules, nil
}
