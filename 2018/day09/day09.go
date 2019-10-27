package day09

import (
	"bufio"
	"fmt"
	"io"
)

// GameSetup represents the initial setup of the game: the number of players
// and the number of marbles.
type GameSetup struct {
	Players int
	Marbles int
}

// Marble represents a marble in the circle.
type Marble struct {
	Value      int
	Prev, Next *Marble
}

// Clockwise returns the marble n places clockwise of m.
func (m *Marble) Clockwise(n int) *Marble {
	for ; n > 0; n-- {
		m = m.Next
	}
	return m
}

// CounterClockwise returns the marble n places counter-clockwise of m.
func (m *Marble) CounterClockwise(n int) *Marble {
	for ; n > 0; n-- {
		m = m.Prev
	}
	return m
}

// InsertAfter inserts m after n in the circle.
func (m *Marble) InsertAfter(n *Marble) {
	m.Prev, m.Next = n, n.Next
	n.Next, n.Next.Prev = m, m
}

// Remove removes m from the circle.
func (m *Marble) Remove() {
	m.Prev.Next, m.Next.Prev = m.Next, m.Prev
	m.Prev, m.Next = nil, nil
}

// PartOne solves Part One of the puzzle.
func PartOne(gs GameSetup) int {
	scores := make([]int, gs.Players)

	currentMarble := &Marble{Value: 0}
	currentMarble.Prev, currentMarble.Next = currentMarble, currentMarble

	var currentPlayer int

	for value := 1; value <= gs.Marbles; value++ {
		switch {
		case value%23 == 0:
			scores[currentPlayer] += value
			removed := currentMarble.CounterClockwise(7)
			futureCurrent := removed.Clockwise(1)
			removed.Remove()
			scores[currentPlayer] += removed.Value
			currentMarble = futureCurrent
		default:
			m := &Marble{Value: value}
			m.InsertAfter(currentMarble.Clockwise(1))
			currentMarble = m
		}

		currentPlayer = (currentPlayer + 1) % gs.Players
	}

	return max(scores)
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(gs GameSetup) int {
	gs.Marbles *= 100
	return PartOne(gs)
}

// ReadGameSetup reads the game setup from r.
func ReadGameSetup(r io.Reader) (GameSetup, error) {
	var gs GameSetup
	var inputRead bool

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if inputRead {
			return gs, fmt.Errorf("more than one line in reader")
		}

		_, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &gs.Players, &gs.Marbles)
		if err != nil {
			return gs, err
		}
		inputRead = true
	}
	if scanner.Err() != nil {
		return gs, scanner.Err()
	}

	return gs, nil
}

func max(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
