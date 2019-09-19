package day01_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day01"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	changes, err := day01.ReadFreqChanges(file)
	if err != nil {
		log.Fatalf("failed to read frequency changes: %v", err)
	}

	fmt.Printf("The resulting frequency is %d.\n", day01.PartOne(changes))
	fmt.Printf("The first frequency reached twice is %d.\n", day01.PartTwo(changes))
	// Output:
	// The resulting frequency is 510.
	// The first frequency reached twice is 69074.
}
