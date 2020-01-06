package day05

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/2019/day05/intcode"
	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	program, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	diagnosticCode, err := runProgram(program, 1)
	if err != nil {
		return fmt.Errorf("failed to get diagnostic code: %v", err)
	}

	_, err = fmt.Fprintf(w, "%d", diagnosticCode)
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

	diagnosticCode, err := runProgram(program, 5)
	if err != nil {
		return fmt.Errorf("failed to get diagnostic code: %v", err)
	}

	_, err = fmt.Fprintf(w, "%d", diagnosticCode)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func runProgram(program []int, input int) (int, error) {
	var outputCount int
	var lastOutput int

	runner := intcode.NewRunner(
		program,
		func() int { return input },
		func(n int) { lastOutput = n; outputCount++ },
	)

	runner.Run()
	if runner.Err() != nil {
		return 0, fmt.Errorf("intcode runtime error: %w", runner.Err())
	}

	if outputCount == 0 {
		return 0, fmt.Errorf("program produced no output")
	}

	return lastOutput, nil
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
