package day06

import (
	"bufio"
	"fmt"
	"io"
)

// Coordinate represents a
type Coordinate struct {
	X, Y int
}

// String implements the Stringer interface.
func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

// PartOne solves Part One of the puzzle.
func PartOne(coords []Coordinate) int {
	// Get range of X and Y values of coordinates.
	minX, maxX, minY, maxY := coords[0].X, coords[0].X, coords[0].Y, coords[0].Y
	for _, coord := range coords {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	// Find area of each coordinate.
	areaByCoordinate := make(map[Coordinate]int)
	areaIsInfiniteByCoordinate := make(map[Coordinate]bool)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			closest, tied := closestCoordinate(x, y, coords)
			if tied {
				continue
			}

			areaByCoordinate[closest]++
			if x == minX || x == maxX || y == minX || y == maxY {
				areaIsInfiniteByCoordinate[closest] = true
			}
		}
	}

	// Find largest area that isn't infinite.
	var largestFiniteArea int
	for _, coord := range coords {
		if !areaIsInfiniteByCoordinate[coord] && areaByCoordinate[coord] > largestFiniteArea {
			largestFiniteArea = areaByCoordinate[coord]
		}
	}

	return largestFiniteArea
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(coords []Coordinate) int {
	// Get range of X and Y values of coordinates.
	minX, maxX, minY, maxY := coords[0].X, coords[0].X, coords[0].Y, coords[0].Y
	for _, coord := range coords {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	const maxTotalDistance = 10000

	// Count locations which have a total distance to all coordinates of less
	// than maxTotalDistance.
	var count int
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if sumDistances(x, y, coords) < maxTotalDistance {
				count++
			}
		}
	}

	return count
}

// Distance returns the Manhattan distance between (x, y) and c.
func distance(x, y int, c Coordinate) int {
	return abs(x-c.X) + abs(y-c.Y)
}

// abs returns the absolute value of n.
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// closestCoordinate finds the coordinate closest to (x,y) and whether there
// is a tie. In case of a tie, any coordinate may be returned.
func closestCoordinate(x, y int, coords []Coordinate) (Coordinate, bool) {
	closestCoord := coords[0]
	minDistance := distance(x, y, closestCoord)
	var tied bool

	for _, coord := range coords[1:] {
		d := distance(x, y, coord)
		if d < minDistance {
			closestCoord = coord
			minDistance = d
			tied = false
		} else if d == minDistance {
			tied = true
		}
	}

	return closestCoord, tied
}

// sumDistances computes the sum of (x,y) to each coordinate.
func sumDistances(x, y int, coords []Coordinate) int {
	var sum int

	for _, coord := range coords {
		sum += distance(x, y, coord)
	}

	return sum
}

// ReadCoordinates reads a list of coordinates from r.
func ReadCoordinates(r io.Reader) ([]Coordinate, error) {
	var coords []Coordinate

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var c Coordinate

		_, err := fmt.Sscanf(scanner.Text(), "%d, %d", &c.X, &c.Y)
		if err != nil {
			return nil, fmt.Errorf("failed to parse text: %w", err)
		}

		coords = append(coords, c)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return coords, nil
}
