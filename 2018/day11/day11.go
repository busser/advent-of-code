package day11

import (
	"fmt"
	"io"
	"math"
)

// PartOne solves Part One of the puzzle.
func PartOne(gridSerialNumber int) (x, y int) {
	const (
		gridSize   = 300
		squareSize = 3
	)

	grid := newGrid(gridSize + 1) // Grid of partial sums.
	for x := 1; x <= gridSize; x++ {
		for y := 1; y <= gridSize; y++ {
			power := getCellPowerValue(x, y, gridSerialNumber)
			grid[x][y] = power + grid[x-1][y] + grid[x][y-1] - grid[x-1][y-1]
		}
	}

	var bestX, bestY int
	maxPower := math.MinInt32
	for x := squareSize; x <= gridSize; x++ {
		for y := squareSize; y <= gridSize; y++ {
			squarePower := getSquareTotalPower(grid, x, y, squareSize)
			if squarePower > maxPower {
				bestX, bestY, maxPower = x, y, squarePower
			}
		}
	}

	return bestX - squareSize + 1, bestY - squareSize + 1
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(gridSerialNumber int) (x, y, size int) {
	const (
		gridSize = 300
	)

	grid := newGrid(gridSize + 1) // Grid of partial sums.
	for x := 1; x <= gridSize; x++ {
		for y := 1; y <= gridSize; y++ {
			power := getCellPowerValue(x, y, gridSerialNumber)
			grid[x][y] = power + grid[x-1][y] + grid[x][y-1] - grid[x-1][y-1]
		}
	}

	var bestX, bestY, bestSquareSize int
	maxPower := math.MinInt32
	for squareSize := 1; squareSize < 300; squareSize++ {
		for x := squareSize; x <= gridSize; x++ {
			for y := squareSize; y <= gridSize; y++ {
				squarePower := getSquareTotalPower(grid, x, y, squareSize)
				if squarePower > maxPower {
					bestX, bestY, bestSquareSize, maxPower = x, y, squareSize, squarePower
				}
			}
		}
	}

	return bestX - bestSquareSize + 1, bestY - bestSquareSize + 1, bestSquareSize
}

// newGrid builds a new grid of size*size.
func newGrid(size int) [][]int {
	grid := make([][]int, size)

	for x := 0; x < size; x++ {
		grid[x] = make([]int, size)
	}

	return grid
}

// getCellPowerValue computes a cell's power value.
func getCellPowerValue(x, y, gridSerialNumber int) int {
	rackID := x + 10

	value := rackID*y + gridSerialNumber
	value *= rackID
	value = value / 100 % 10
	value -= 5

	return value
}

func getSquareTotalPower(grid [][]int, x, y, size int) int {
	return grid[x][y] - grid[x-size][y] - grid[x][y-size] + grid[x-size][y-size]
}

// ReadGridSerialNumber reads a grid ID from r.
func ReadGridSerialNumber(r io.Reader) (int, error) {
	var serialNumber int

	_, err := fmt.Fscanf(r, "%d", &serialNumber)
	if err != nil {
		return 0, err
	}

	return serialNumber, nil
}
