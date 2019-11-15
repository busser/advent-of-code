package day06_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day06"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	coords, err := day06.ReadCoordinates(file)
	if err != nil {
		log.Fatalf("failed to read coordinates: %v", err)
	}

	fmt.Printf("The size of the largest finite area is %d.\n", day06.PartOne(coords))
	fmt.Printf("%d locations have a total distance to all given coordinates of less than 10000.\n", day06.PartTwo(coords))
	// Output:
	// The size of the largest finite area is 3260.
	// 42535 locations have a total distance to all given coordinates of less than 10000.
}
