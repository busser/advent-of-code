package day06

import (
	"fmt"
	"io"
	"strings"

	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	orbitTree, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	direct, indirect := countOrbits(orbitTree)

	_, err = fmt.Fprintf(w, "%d", direct+indirect)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	orbitTree, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	distance := distanceToSanta(orbitTree)

	_, err = fmt.Fprintf(w, "%d", distance)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func countOrbits(orbitTree map[string][]string) (direct, indirect int) {
	seen := make(map[string]struct{})

	var helper func(string) (int, int)
	helper = func(object string) (direct, indirect int) {
		seen[object] = struct{}{}

		for _, child := range orbitTree[object] {
			if _, ok := seen[child]; ok {
				continue
			}

			childDirect, childIndirect := helper(child)
			direct += childDirect + 1
			indirect += childDirect + childIndirect
		}

		return
	}

	return helper("COM")
}

func distanceToSanta(orbitTree map[string][]string) int {
	distByObj := map[string]int{"YOU": 0}

	seen := make(map[string]struct{})
	seen["YOU"] = struct{}{}

	queue := []string{"YOU"}

	for len(queue) > 0 {
		object := queue[0]
		queue = queue[1:]

		dist := distByObj[object]

		for _, child := range orbitTree[object] {
			if _, ok := seen[child]; ok {
				continue
			}

			distByObj[child] = dist + 1
			seen[child] = struct{}{}
			queue = append(queue, child)
		}
	}

	return distByObj["SAN"] - 2
}

func readInput(r io.Reader) (map[string][]string, error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	orbitGraph := make(map[string][]string)

	for _, l := range lines {
		words := strings.Split(l, ")")
		parent, satellite := words[0], words[1]
		orbitGraph[parent] = append(orbitGraph[parent], satellite)
		orbitGraph[satellite] = append(orbitGraph[satellite], parent)
	}

	return orbitGraph, nil
}
