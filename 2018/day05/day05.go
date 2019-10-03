package day05

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

// Unit represents a polymer unit.
type Unit rune

// sameType checks if u and v are of the same type.
// This method panics if u is neither uppercase or lowercase.
func (u Unit) sameType(v Unit) bool {
	return unicode.ToLower(rune(u)) == unicode.ToLower(rune(v))
}

// samePolarity checks if u and v are of the same polarity.
func (u Unit) samePolarity(v Unit) bool {
	return u == v ||
		unicode.IsLower(rune(u)) && unicode.IsLower(rune(v)) ||
		unicode.IsUpper(rune(u)) && unicode.IsUpper(rune(v))
}

// ReactsWith checks if u reacts with v.
func (u Unit) ReactsWith(v Unit) bool {
	return u.sameType(v) && !u.samePolarity(v)
}

// Polymer represents a polymer.
type Polymer []Unit

// String provides a printable representation of the polymer.
func (p Polymer) String() string {
	r := make([]rune, 0, len(p))
	for _, u := range p {
		r = append(r, rune(u))
	}
	return string(r)
}

// react reacts p.
func (p Polymer) react() Polymer {
	var unitsLeft []Unit

	for _, u := range p {
		switch {
		case len(unitsLeft) == 0:
			unitsLeft = append(unitsLeft, u)
		case u.ReactsWith(unitsLeft[len(unitsLeft)-1]):
			unitsLeft = unitsLeft[:len(unitsLeft)-1]
		default:
			unitsLeft = append(unitsLeft, u)
		}
	}

	return Polymer(unitsLeft)
}

// remove removes all units of the same type as u from p.
func (p Polymer) remove(u Unit) Polymer {
	var newPoly Polymer

	for _, v := range p {
		if !u.sameType(v) {
			newPoly = append(newPoly, v)
		}
	}

	return newPoly
}

// PartOne solves Part One of the puzzle.
func PartOne(poly Polymer) int {
	return len(poly.react())
}

// PartTwo solves PartTwo of the puzzle.
func PartTwo(poly Polymer) int {
	poly = poly.react() // This avoids doing the same work multiple times.

	reactedPolyByRemovedUnit := make(map[Unit]Polymer)

	for c := 'a'; c <= 'z'; c++ {
		u := Unit(c)
		reactedPolyByRemovedUnit[u] = poly.remove(u).react()
	}

	minPolyLength := len(poly)
	for _, p := range reactedPolyByRemovedUnit {
		if len(p) < minPolyLength {
			minPolyLength = len(p)
		}
	}

	return minPolyLength
}

// ReadPolymer reads a polymer from r.
func ReadPolymer(r io.Reader) (Polymer, error) {
	var poly Polymer

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if poly != nil {
			return nil, errors.New("more than one line in reader")
		}

		poly = Polymer(scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return poly, nil
}
