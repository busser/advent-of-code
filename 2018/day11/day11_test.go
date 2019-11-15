package day11_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day11"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	gridSerialNumber, err := day11.ReadGridSerialNumber(file)
	if err != nil {
		log.Fatalf("could not read grid ID: %v", err)
	}

	x, y := day11.PartOne(gridSerialNumber)
	fmt.Printf("The top-left coordinate of the 3x3 square with the largest total power is %d,%d.\n", x, y)
	x, y, size := day11.PartTwo(gridSerialNumber)
	fmt.Printf("The square with the largest total power is %d,%d,%d.\n", x, y, size)

	// Output:
	// The top-left coordinate of the 3x3 square with the largest total power is 235,35.
	// The square with the largest total power is 142,265,7.
}
