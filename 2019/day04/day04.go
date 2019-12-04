package day04

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	min, max, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	count := 0
	for n := min; n <= max; n++ {
		dd := digits(n)
		if neverDecreases(dd) && hasDouble(dd) {
			count++
		}
	}

	_, err = fmt.Fprintf(w, "%d", count)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	min, max, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	count := 0
	for n := min; n <= max; n++ {
		dd := digits(n)
		if neverDecreases(dd) && hasStrictDouble(dd) {
			count++
		}
	}

	_, err = fmt.Fprintf(w, "%d", count)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// digits returns a slice containing to digits of n.
func digits(n int) []int {
	var dd []int

	div := 1
	for div <= n {
		div *= 10
	}
	div /= 10

	for div > 0 {
		d := (n / div) % 10
		dd = append(dd, d)
		div /= 10
	}

	return dd
}

// neverDecreases checks if values in num never decrease from left to right.
func neverDecreases(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

// hasDouble checks whether two consecutive values in nums are equal.
func hasDouble(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

// hasStrictDouble checks whether two consecutive values in nums are equal
// without those two values being consecutive with a third equal value.
func hasStrictDouble(nums []int) bool {
	var isDouble, isTriple bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			if isDouble {
				isTriple = true
			} else {
				isDouble = true
			}
		} else {
			if isDouble && !isTriple {
				return true
			}
			isDouble, isTriple = false, false
		}
	}

	return isDouble && !isTriple
}

func readInput(r io.Reader) (min, max int, err error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read lines: %w", err)
	}

	words := strings.Split(lines[0], "-")
	if len(words) != 2 {
		return 0, 0, fmt.Errorf("found %d words in first line of input, expected 2", len(words))
	}

	min, err = strconv.Atoi(words[0])
	if err != nil {
		return 0, 0, err
	}
	max, err = strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, err
	}

	return min, max, nil
}
