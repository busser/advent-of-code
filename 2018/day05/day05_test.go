package day05_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day05"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	poly, err := day05.ReadPolymer(file)
	if err != nil {
		log.Fatalf("failed to read polymer: %v", err)
	}

	fmt.Printf("%d units remain after fully reacting the polymer.\n", day05.PartOne(poly))
	fmt.Printf("%d units remain after optimizing and fully reacting the polymer.", day05.PartTwo(poly))
	// Output:
	// 9154 units remain after fully reacting the polymer.
	// 4556 units remain after optimizing and fully reacting the polymer.
}
