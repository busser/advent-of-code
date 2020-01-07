package day06_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day06"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day06.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("145250"),
		},
		"PartTwo": {
			F:         day06.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("274"),
		},
		"PartOneExample": {
			F:         day06.PartOne,
			InputPath: "testdata/example_1.txt",
			Expected:  []byte("42"),
		},
		"PartTwoExample": {
			F:         day06.PartTwo,
			InputPath: "testdata/example_2.txt",
			Expected:  []byte("4"),
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
			F:         day06.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day06.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
