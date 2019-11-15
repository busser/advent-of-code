package day03_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day03"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	claims, err := day03.ReadClaims(file)
	if err != nil {
		log.Fatalf("failed to read claims: %v", err)
	}

	fmt.Printf("%d square inches of fabric are within two or more claims.\n", day03.PartOne(claims))
	partTwo, err := day03.PartTwo(claims)
	if err != nil {
		log.Fatalf("failed to solve Part Two: %v", err)
	}
	fmt.Printf("The ID of the only claim that doesn't overlap is %d.\n", partTwo)
	// Output:
	// 104712 square inches of fabric are within two or more claims.
	// The ID of the only claim that doesn't overlap is 840.
}
