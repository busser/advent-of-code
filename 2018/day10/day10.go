package day10

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Point represents a point in the sky.
type Point struct {
	Pos Position
	Vel Velocity
}

// Position represents a point's position.
type Position struct {
	X, Y int
}

// Velocity represents a point's velocity.
type Velocity struct {
	X, Y int
}

// SecondOrderPolynomial represents A*x^2 + B*x + C.
type SecondOrderPolynomial struct {
	A, B, C int
}

// FirstOrderPolynomial represents A*x + C
type FirstOrderPolynomial struct {
	A, B int
}

// Add adds q to p.
func (p *SecondOrderPolynomial) Add(q SecondOrderPolynomial) {
	p.A += q.A
	p.B += q.B
	p.C += q.C
}

// Derive provides the derivative of p.
func (p *SecondOrderPolynomial) Derive() FirstOrderPolynomial {
	return FirstOrderPolynomial{A: 2 * p.A, B: p.B}
}

// PartOne solves Part One of the puzzle.
func PartOne(points []Point) string {
	costFunction := &SecondOrderPolynomial{}

	for i, p := range points {
		for _, q := range points[i+1:] {
			costFunction.A += (p.Vel.X-q.Vel.X)*(p.Vel.X-q.Vel.X) + (p.Vel.Y-q.Vel.Y)*(p.Vel.Y-q.Vel.Y)
			costFunction.B += 2 * ((p.Pos.X-q.Pos.X)*(p.Vel.X-q.Vel.X) + (p.Pos.Y-q.Pos.Y)*(p.Vel.Y-q.Vel.Y))
			costFunction.C += (p.Pos.X-q.Pos.X)*(p.Pos.X-q.Pos.X) + (p.Pos.Y-q.Pos.Y)*(p.Pos.Y-q.Pos.Y)
		}
	}

	costFunctionRate := costFunction.Derive()
	minCostTime := -costFunctionRate.B / costFunctionRate.A

	var convergedPoints []Point
	for _, p := range points {
		q := Point{Pos: Position{
			X: p.Pos.X + minCostTime*p.Vel.X,
			Y: p.Pos.Y + minCostTime*p.Vel.Y,
		}}
		convergedPoints = append(convergedPoints, q)
	}

	minX := convergedPoints[0].Pos.X
	maxX := minX
	minY := convergedPoints[0].Pos.Y
	maxY := minY
	for _, p := range convergedPoints[1:] {
		if p.Pos.X < minX {
			minX = p.Pos.X
		}
		if p.Pos.X > maxX {
			maxX = p.Pos.X
		}
		if p.Pos.Y < minY {
			minY = p.Pos.Y
		}
		if p.Pos.Y > maxY {
			maxY = p.Pos.Y
		}
	}

	pointMap := make([][]bool, maxX-minX+1)
	for x := 0; x <= maxX-minX; x++ {
		pointMap[x] = make([]bool, maxY-minY+1)
	}

	for _, p := range convergedPoints {
		pointMap[p.Pos.X-minX][p.Pos.Y-minY] = true
	}

	var msg strings.Builder

	for y := 0; y <= maxY-minY; y++ {
		var line strings.Builder
		for x := 0; x <= maxX-minX; x++ {
			if pointMap[x][y] {
				line.WriteRune('#')
			} else {
				line.WriteRune(' ')
			}
		}
		msg.WriteString(strings.TrimRight(line.String(), " "))
		msg.WriteRune('\n')
	}

	return msg.String()
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(points []Point) int {
	costFunction := &SecondOrderPolynomial{}

	for i, p := range points {
		for _, q := range points[i+1:] {
			costFunction.A += (p.Vel.X-q.Vel.X)*(p.Vel.X-q.Vel.X) + (p.Vel.Y-q.Vel.Y)*(p.Vel.Y-q.Vel.Y)
			costFunction.B += 2 * ((p.Pos.X-q.Pos.X)*(p.Vel.X-q.Vel.X) + (p.Pos.Y-q.Pos.Y)*(p.Vel.Y-q.Vel.Y))
			costFunction.C += (p.Pos.X-q.Pos.X)*(p.Pos.X-q.Pos.X) + (p.Pos.Y-q.Pos.Y)*(p.Pos.Y-q.Pos.Y)
		}
	}

	costFunctionRate := costFunction.Derive()
	minCostTime := -costFunctionRate.B / costFunctionRate.A

	return minCostTime
}

// ReadPoints reads points from r.
func ReadPoints(r io.Reader) ([]Point, error) {
	var points []Point

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var p Point

		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &p.Pos.X, &p.Pos.Y, &p.Vel.X, &p.Vel.Y)
		if err != nil {
			return nil, err
		}

		points = append(points, p)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return points, nil
}
