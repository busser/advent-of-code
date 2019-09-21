package day02

import (
	"bufio"
	"fmt"
	"io"
)

// BoxID represents a box ID.
type BoxID []rune

// countLetters returns a map of the number of times each letter appears in id.
func (id BoxID) countLetters() map[rune]int {
	countByLetter := make(map[rune]int)
	for _, letter := range id {
		countByLetter[letter]++
	}
	return countByLetter
}

// lettersInCommon finds the letters that are in both in id and other, at the
// same index.
func (id BoxID) lettersInCommon(other BoxID) ([]rune, error) {
	if len(id) != len(other) {
		return nil, fmt.Errorf("IDs have different lengths")
	}

	var letters []rune
	for i, l := range id {
		if l == []rune(other)[i] {
			letters = append(letters, l)
		}
	}

	return letters, nil
}

// PartOne solves Part One of the puzzle.
func PartOne(ids []BoxID) int {
	var doubles, triples int

	for _, id := range ids {
		countByLetter := id.countLetters()

		var hasDouble, hasTriple bool
		for _, count := range countByLetter {
			switch count {
			case 2:
				hasDouble = true
			case 3:
				hasTriple = true
			}
		}

		if hasDouble {
			doubles++
		}
		if hasTriple {
			triples++
		}
	}

	return doubles * triples
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(ids []BoxID) ([]rune, error) {
	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			lettersInCommon, err := ids[i].lettersInCommon(ids[j])
			if err != nil {
				return nil, fmt.Errorf("could not find letters in common: %v", err)
			}

			if len(lettersInCommon) == len(ids[i])-1 {
				return lettersInCommon, nil
			}
		}
	}

	return nil, fmt.Errorf("no answer found")
}

// ReadBoxIDs reads box IDs from r.
func ReadBoxIDs(r io.Reader) ([]BoxID, error) {
	var ids []BoxID

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ids = append(ids, BoxID(scanner.Text()))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return ids, nil
}
