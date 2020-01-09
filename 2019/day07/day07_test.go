package day07_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day07"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day07.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("273814"),
		},
		"PartTwo": {
			F:         day07.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("34579864"),
		},
		"PartOneExample1": {
			F:         day07.PartOne,
			InputPath: "testdata/example_1.txt",
			Expected:  []byte("43210"),
		},
		"PartOneExample2": {
			F:         day07.PartOne,
			InputPath: "testdata/example_2.txt",
			Expected:  []byte("54321"),
		},
		"PartOneExample3": {
			F:         day07.PartOne,
			InputPath: "testdata/example_3.txt",
			Expected:  []byte("65210"),
		},
		"PartTwoExample4": {
			F:         day07.PartTwo,
			InputPath: "testdata/example_4.txt",
			Expected:  []byte("139629729"),
		},
		"PartTwoExample5": {
			F:         day07.PartTwo,
			InputPath: "testdata/example_5.txt",
			Expected:  []byte("18216"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			util.RunTest(t, tc)
		})
	}
}

func BenchmarkDay(b *testing.B) {
	benchmarks := map[string]util.BenchmarkCase{
		"PartOne": {
			F:         day07.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day07.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
