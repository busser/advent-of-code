package day07_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/adventofcode/2018/day07"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	instructions, err := day07.ReadInstructions(file)
	if err != nil {
		log.Fatalf("failed to read instructions: %v", err)
	}

	stepDurationFunc := func(s day07.Step) int {
		return int(rune(s) - 'A' + 61)
	}

	fmt.Printf("The steps should be completed in this order: %s.\n", day07.PartOne(instructions))
	fmt.Printf("It will take %d seconds to complete all the steps.\n", day07.PartTwo(instructions, 5, stepDurationFunc))
	// Output:
	// The steps should be completed in this order: ABGKCMVWYDEHFOPQUILSTNZRJX.
	// It will take 898 seconds to complete all the steps.
}
