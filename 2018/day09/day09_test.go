package day09_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day09"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	gameSetup, err := day09.ReadGameSetup(file)
	if err != nil {
		log.Fatalf("failed to read game setup: %v", err)
	}

	fmt.Printf("The winning elf's score is %d.\n", day09.PartOne(gameSetup))
	fmt.Printf("If the last marble's number was 100 times larger, the winning elf's score would be %d.\n", day09.PartTwo(gameSetup))
	// Output:
	// The winning elf's score is 424639.
	// If the last marble's number was 100 times larger, the winning elf's score would be 3516007333.
}
