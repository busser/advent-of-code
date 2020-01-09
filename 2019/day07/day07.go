package day07

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/2019/day07/intcode"
	"github.com/busser/adventofcode/util"
	"golang.org/x/sync/errgroup"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	program, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	maxSignal, err := highestPossibleSignal(program, []int{0, 1, 2, 3, 4})
	if err != nil {
		return fmt.Errorf("failed to find highest possible signal: %w", err)
	}

	_, err = fmt.Fprintf(w, "%d", maxSignal)
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

	maxSignal, err := highestPossibleSignal(program, []int{5, 6, 7, 8, 9})
	if err != nil {
		return fmt.Errorf("failed to find highest possible signal: %w", err)
	}

	_, err = fmt.Fprintf(w, "%d", maxSignal)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func highestPossibleSignal(program []int, phaseSettings []int) (int, error) {
	g, ctx := errgroup.WithContext(context.Background())

	perms := make(chan []int)
	g.Go(func() error {
		defer close(perms)
		return permutations(ctx, phaseSettings, perms)
	})

	signals := make(chan int, factorial(len(phaseSettings)))

	const numWorkers = 8
	for i := 0; i < numWorkers; i++ {
		g.Go(func() error {
			for p := range perms {
				signal, err := computeSignal(ctx, program, p)
				if err != nil {
					return err
				}

				select {
				case <-ctx.Done():
					return ctx.Err()
				case signals <- signal:
				}
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}

	close(signals)

	maxSignal := <-signals
	for s := range signals {
		if s > maxSignal {
			maxSignal = s
		}
	}

	return maxSignal, nil
}

func computeSignal(ctx context.Context, program []int, phaseSettings []int) (int, error) {
	// Amplifier A reads from io[0] and writes to io[1], B reads from io[1] and
	// writes to io[2], etc.
	io := make([]chan int, len(phaseSettings))
	for i := range io {
		io[i] = make(chan int, 2)
	}

	for i := range phaseSettings {
		io[i] <- phaseSettings[i]
	}

	g, ctx := errgroup.WithContext(ctx)

	for i := range phaseSettings {
		i := i // https://golang.org/doc/faq#closures_and_goroutines

		g.Go(func() error {
			memory := make([]int, len(program))
			copy(memory, program)

			runner := intcode.NewRunner(
				memory,
				func() int {
					select {
					case <-ctx.Done():
						return 0
					case n := <-io[i]:
						return n
					}
				},
				func(n int) {
					select {
					case <-ctx.Done():
					case io[(i+1)%len(phaseSettings)] <- n:
					}
				},
			)

			runner.Run(ctx)
			return runner.Err()
		})
	}

	io[0] <- 0

	if err := g.Wait(); err != nil {
		return 0, err
	}

	signal := <-io[0]

	return signal, nil
}

func permutations(ctx context.Context, values []int, out chan<- []int) error {
	var helper func(int) error
	helper = func(size int) error {
		if size == 1 {
			tmp := make([]int, len(values))
			copy(tmp, values)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- tmp:
			}
		}

		for i := 0; i < size; i++ {
			err := helper(size - 1)
			if err != nil {
				return err
			}

			if size%2 == 1 {
				values[0], values[size-1] = values[size-1], values[0]
			} else {
				values[i], values[size-1] = values[size-1], values[i]
			}
		}

		return nil
	}

	return helper(len(values))
}

func factorial(n int) int {
	f := 1
	for n > 1 {
		f *= n
		n--
	}
	return f
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
