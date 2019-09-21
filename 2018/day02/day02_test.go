package day02_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day02"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	ids, err := day02.ReadBoxIDs(file)
	if err != nil {
		log.Fatalf("failed to read box IDs: %v", err)
	}

	fmt.Printf("The checksum is %d.\n", day02.PartOne(ids))
	partTwo, err := day02.PartTwo(ids)
	if err != nil {
		log.Fatalf("failed to solve Part Two: %v", err)
	}
	fmt.Printf("The letters in common between the two IDs are %s.\n", string(partTwo))
	// Output:
	// The checksum is 5166.
	// The letters in common between the two IDs are cypueihajytordkgzxfqplbwn.
}
