package day09

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/2019/day09/intcode"
	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	program, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	keycode, err := runProgram(program, 1)
	if err != nil {
		return fmt.Errorf("failed to get keycode: %v", err)
	}

	_, err = fmt.Fprintf(w, "%d", keycode)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	program, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	coordinates, err := runProgram(program, 2)
	if err != nil {
		return fmt.Errorf("failed to get coordinates: %v", err)
	}

	_, err = fmt.Fprintf(w, "%d", coordinates)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func runProgram(program []int, input int) (int, error) {
	var output int

	runner := intcode.NewRunner(
		program,
		func() int { return input },
		func(n int) { output = n },
	)

	runner.Run(context.Background())
	if runner.Err() != nil {
		return 0, fmt.Errorf("intcode runtime error: %w", runner.Err())
	}

	return output, nil
}

func readInput(r io.Reader) ([]int, error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	numbers := strings.Split(lines[0], ",")
	program := make([]int, len(numbers))
	for i := range numbers {
		opcode, err := strconv.Atoi(numbers[i])
		if err != nil {
			return nil, err
		}
		program[i] = opcode
	}

	return program, nil
}
