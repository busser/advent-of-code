package day12_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day12"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	initialState, notes, err := day12.ReadInitialStateAndNotes(file)
	if err != nil {
		log.Fatalf("could not read initial state and notes: %v", err)
	}

	fmt.Printf("After 20 generations, the sum is %d.\n", day12.PartOne(initialState, notes))
	partTwo, err := day12.PartTwo(initialState, notes)
	if err != nil {
		log.Fatalf("could not solve Part Two: %v", err)
	}
	fmt.Printf("After fifty billion generations, the sum is %d.\n", partTwo)
	// Output:
	// After 20 generations, the sum is 3230.
	// After fifty billion generations, the sum is 4400000000304.
}
