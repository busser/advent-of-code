package day12

import (
	"bufio"
	"fmt"
	"io"
)

// State represents the state of the pots at a given generation.
type State struct {
	firstPotNum int
	pots        []bool
}

// NewState builds a new instance of State with space for pots numbered between
// min and max inclusive.
func NewState(min, max int) *State {
	return &State{
		firstPotNum: min,
		pots:        make([]bool, max-min+1),
	}
}

// Set sets pot number i as containing a plant.
func (s *State) Set(i int) {
	s.pots[i-s.firstPotNum] = true
}

// Get returns wether pot number i contains a plant.
func (s *State) Get(i int) bool {
	if i < s.firstPotNum || i >= s.firstPotNum+len(s.pots) {
		return false
	}
	return s.pots[i-s.firstPotNum]
}

// Sum returns the sum of the numbers of all pots containing a plant.
func (s *State) Sum() int {
	var sum int
	for i, v := range s.pots {
		if v {
			sum += s.firstPotNum + i
		}
	}
	return sum
}

// Next returns the next iteration of the state.
func (s *State) Next(notes *[32]bool) *State {
	var minPotNum, maxPotNum int
	for i, v := range s.pots {
		if v {
			minPotNum = i + s.firstPotNum
			break
		}
	}
	for i := len(s.pots) - 1; i >= 0; i-- {
		if s.pots[i] {
			maxPotNum = i + s.firstPotNum
			break
		}
	}

	nextState := NewState(minPotNum-2, maxPotNum+2)

	for i := nextState.firstPotNum; i < nextState.firstPotNum+len(nextState.pots); i++ {
		var context int
		for j := 0; j < 5; j++ {
			if s.Get(i + j - 2) {
				context += 1 << (4 - j)
			}
		}
		if notes[context] {
			nextState.Set(i)
		}
	}

	return nextState
}

// Similar checks whether s and t are similar, ie. they are identical only
// shifted by one or more pots.
func (s *State) Similar(t *State) bool {
	if len(s.pots) != len(t.pots) {
		return false
	}

	for i := range s.pots {
		if s.pots[i] != t.pots[i] {
			return false
		}
	}

	return true
}

// PartOne solves Part One of the puzzle.
func PartOne(initialState *State, notes *[32]bool) int {
	state := initialState

	for gen := 0; gen < 20; gen++ {
		state = state.Next(notes)
	}

	return state.Sum()
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(initialState *State, notes *[32]bool) (int, error) {
	currentState := initialState

	for gen := 0; gen < 1000; gen++ {
		nextState := currentState.Next(notes)

		if nextState.Similar(currentState) {
			currentSum := currentState.Sum()
			nextSum := nextState.Sum()
			return currentSum + (nextSum-currentSum)*(50000000000-gen), nil
		}

		currentState = nextState
	}

	return 0, fmt.Errorf("no pattern found")
}

// ReadInitialStateAndNotes reads the initial state and notes from r.
func ReadInitialStateAndNotes(r io.Reader) (*State, *[32]bool, error) {
	var stateStr string

	_, err := fmt.Fscanf(r, "initial state: %s\n\n", &stateStr)
	if err != nil {
		return nil, nil, err
	}

	state := NewState(0, len(stateStr))

	for i, c := range stateStr {
		if c == '#' {
			state.Set(i)
		}
	}

	var notes [32]bool

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var contextStr string
		var resultRune rune

		_, err := fmt.Sscanf(scanner.Text(), "%5s => %c", &contextStr, &resultRune)
		if err != nil {
			return nil, nil, err
		}

		if resultRune == '#' {
			var noteContext int
			for i, c := range contextStr {
				if c == '#' {
					noteContext += 1 << (4 - i)
				}
			}
			notes[noteContext] = true
		}

	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}

	return state, &notes, nil
}
