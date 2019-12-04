package day03

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/util"
)

// PartOne solves the first part of the day's puzzle.
func PartOne(w io.Writer, r io.Reader) error {
	paths, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	positions := make(map[position]struct{})
	paths[0].walk(func(_ int, pos position) {
		positions[pos] = struct{}{}
	})

	intersections := make(map[position]struct{})
	paths[1].walk(func(_ int, pos position) {
		if _, intersects := positions[pos]; intersects {
			intersections[pos] = struct{}{}
		}
	})

	minDist := 1<<31 - 1
	for inter := range intersections {
		dist := distanceFromCenter(inter)
		if dist < minDist {
			minDist = dist
		}
	}

	_, err = fmt.Fprintf(w, "%d", minDist)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second part of the day's puzzle.
func PartTwo(w io.Writer, r io.Reader) error {
	paths, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	delaysByPosition := make(map[position]int)
	paths[0].walk(func(delay int, pos position) {
		if _, alreadyBeen := delaysByPosition[pos]; !alreadyBeen {
			delaysByPosition[pos] = delay
		}
	})

	delaySumsByIntersection := make(map[position]int)
	paths[1].walk(func(secondDelay int, pos position) {
		if firstDelay, intersects := delaysByPosition[pos]; intersects {
			if _, alreadyBeen := delaySumsByIntersection[pos]; !alreadyBeen {
				delaySumsByIntersection[pos] = firstDelay + secondDelay
			}
		}
	})

	minSum := 1<<31 - 1
	for _, sum := range delaySumsByIntersection {
		if sum < minSum {
			minSum = sum
		}
	}

	_, err = fmt.Fprintf(w, "%d", minSum)
	if err != nil {
		return fmt.Errorf("failed to write answer: %w", err)
	}

	return nil
}

func distanceFromCenter(pos position) int {
	return abs(pos.x) + abs(pos.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type position struct {
	x, y int
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

type step struct {
	dir  direction
	dist int
}

type path []step

// walk calls f on every position taken when following p from the center.
func (p path) walk(f func(step int, p position)) {
	pos := position{x: 0, y: 0}
	delay := 0
	for _, s := range p {
		switch s.dir {
		case up:
			for i := 0; i < s.dist; i++ {
				pos.y--
				delay++
				f(delay, pos)
			}
		case right:
			for i := 0; i < s.dist; i++ {
				pos.x++
				delay++
				f(delay, pos)
			}
		case down:
			for i := 0; i < s.dist; i++ {
				pos.y++
				delay++
				f(delay, pos)
			}
		case left:
			for i := 0; i < s.dist; i++ {
				pos.x--
				delay++
				f(delay, pos)
			}
		}
	}
}

func readInput(r io.Reader) ([]path, error) {
	lines, err := util.ReadLines(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	if len(lines) != 2 {
		return nil, fmt.Errorf("found %d lines of input, expected 2", len(lines))
	}

	var paths []path

	for _, l := range lines {
		var p path

		for _, word := range strings.Split(l, ",") {
			var s step

			switch word[0] {
			case 'U':
				s.dir = up
			case 'R':
				s.dir = right
			case 'D':
				s.dir = down
			case 'L':
				s.dir = left
			default:
				return nil, fmt.Errorf("unknown direction %q", word[0])
			}

			n, err := strconv.Atoi(word[1:])
			if err != nil {
				return nil, err
			}
			s.dist = n

			p = append(p, s)
		}

		paths = append(paths, p)
	}

	return paths, nil
}
