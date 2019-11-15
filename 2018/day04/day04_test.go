package day04_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day04"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	records, err := day04.ReadRecords(file)
	if err != nil {
		log.Fatalf("failed to read records: %v", err)
	}

	fmt.Printf("For strategy 1, the ID of the guard multiplied by the minute is %d.\n", day04.PartOne(records))
	fmt.Printf("For strategy 2, the ID of the guard multiplied by the minute is %d.\n", day04.PartTwo(records))
	// Output:
	// For strategy 1, the ID of the guard multiplied by the minute is 76357.
	// For strategy 2, the ID of the guard multiplied by the minute is 41668.
}
