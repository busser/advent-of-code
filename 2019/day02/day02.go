package day02

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	program, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	program[1], program[2] = 12, 2
	runProgram(program)

	_, err = fmt.Fprintf(w, "%d", program[0])
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

	noun, verb := findInputs(program, 19690720)

	_, err = fmt.Fprintf(w, "%d", 100*noun+verb)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func runProgram(program []int) {
	pos := 0
	for {
		opcode := program[pos]
		switch opcode {
		case 1:
			a, b, c := program[pos+1], program[pos+2], program[pos+3]
			program[c] = program[a] + program[b]
		case 2:
			a, b, c := program[pos+1], program[pos+2], program[pos+3]
			program[c] = program[a] * program[b]
		case 99:
			return
		}
		pos += 4
	}
}

func findInputs(program []int, output int) (noun, verb int) {
	memory := make([]int, len(program))
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(memory, program)
			memory[1], memory[2] = noun, verb
			runProgram(memory)
			if memory[0] == output {
				return noun, verb
			}
		}
	}

	panic(fmt.Sprintf("no inputs found that result in wanted output %d", output))
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
