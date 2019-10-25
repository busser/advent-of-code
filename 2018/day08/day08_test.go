package day08_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day08"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	numbers, err := day08.ReadNumbers(file)
	if err != nil {
		log.Fatalf("failed to read numbers: %v", err)
	}

	fmt.Printf("The sum of all metadata entries is %d.\n", day08.PartOne(numbers))
	fmt.Printf("The value of the root node is %d.", day08.PartTwo(numbers))
	// Output:
	// The sum of all metadata entries is 42146.
	// The value of the root node is 26753.
}
