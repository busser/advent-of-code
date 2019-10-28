package day10_test

import (
	"fmt"
	"log"
	"os"

	"github.com/busser/advent-of-code/2018/day10"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	points, err := day10.ReadPoints(file)
	if err != nil {
		log.Fatalf("failed to read points: %v", err)
	}

	fmt.Printf("This message will eventually appear in the sky:\n%s", day10.PartOne(points))
	fmt.Printf("The message will appear after %d seconds.\n", day10.PartTwo(points))
	// Output:
	// This message will eventually appear in the sky:
	//   ##    #####    ####   #    #     ###  #####   #    #  ######
	//  #  #   #    #  #    #  #    #      #   #    #  #    #  #
	// #    #  #    #  #        #  #       #   #    #   #  #   #
	// #    #  #    #  #        #  #       #   #    #   #  #   #
	// #    #  #####   #         ##        #   #####     ##    #####
	// ######  #    #  #  ###    ##        #   #    #    ##    #
	// #    #  #    #  #    #   #  #       #   #    #   #  #   #
	// #    #  #    #  #    #   #  #   #   #   #    #   #  #   #
	// #    #  #    #  #   ##  #    #  #   #   #    #  #    #  #
	// #    #  #####    ### #  #    #   ###    #####   #    #  #
	// The message will appear after 10619 seconds.
}
