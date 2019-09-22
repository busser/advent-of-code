package day03

import (
	"bufio"
	"fmt"
	"io"
)

// Square represents a square inch of fabric.
type Square struct {
	X, Y int // Coordinates of the square on the fabric.
}

// Claim represents a claim to a rectangle of fabric.
type Claim struct {
	ID            int
	TopLeft       Square
	Width, Height int
}

// PartOne solves Part One of the puzzle.
func PartOne(claims []Claim) int {
	claimCountBySquare := make(map[Square]int)
	for _, claim := range claims {
		for x := claim.TopLeft.X; x < claim.TopLeft.X+claim.Width; x++ {
			for y := claim.TopLeft.Y; y < claim.TopLeft.Y+claim.Height; y++ {
				claimCountBySquare[Square{X: x, Y: y}]++
			}
		}
	}

	var overclaimedSquareCount int
	for _, claimCount := range claimCountBySquare {
		if claimCount >= 2 {
			overclaimedSquareCount++
		}
	}

	return overclaimedSquareCount
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(claims []Claim) (int, error) {
	overlappingClaimIDs := make(map[int]bool)

	claimIDsBySquare := make(map[Square][]int)
	for _, claim := range claims {
		for x := claim.TopLeft.X; x < claim.TopLeft.X+claim.Width; x++ {
			for y := claim.TopLeft.Y; y < claim.TopLeft.Y+claim.Height; y++ {
				sq := Square{X: x, Y: y}
				claimIDsBySquare[sq] = append(claimIDsBySquare[sq], claim.ID)
				if len(claimIDsBySquare[sq]) > 1 {
					for _, id := range claimIDsBySquare[sq] {
						overlappingClaimIDs[id] = true
					}
				}
			}
		}
	}

	for _, claim := range claims {
		if _, ok := overlappingClaimIDs[claim.ID]; !ok {
			// We assume there is only one claim that doesn't overlap.
			return claim.ID, nil
		}
	}

	return 0, fmt.Errorf("no non-overlapping claim found")
}

// ReadClaims reads claims to rectangles of fabric from r.
func ReadClaims(r io.Reader) ([]Claim, error) {
	var claims []Claim

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c := Claim{}

		_, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c.ID, &c.TopLeft.X, &c.TopLeft.Y, &c.Width, &c.Height)
		if err != nil {
			return nil, err
		}

		claims = append(claims, c)
	}

	return claims, nil
}
