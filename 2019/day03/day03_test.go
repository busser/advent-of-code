package day03_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day03"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day03.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("3247"),
		},
		"PartTwo": {
			F:         day03.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("48054"),
		},
		"PartOneExample1": {
			F:         day03.PartOne,
			InputPath: "testdata/example1.txt",
			Expected:  []byte("6"),
		},
		"PartOneExample2": {
			F:         day03.PartOne,
			InputPath: "testdata/example2.txt",
			Expected:  []byte("159"),
		},
		"PartOneExample3": {
			F:         day03.PartOne,
			InputPath: "testdata/example3.txt",
			Expected:  []byte("135"),
		},
		"PartTwoExample1": {
			F:         day03.PartTwo,
			InputPath: "testdata/example1.txt",
			Expected:  []byte("30"),
		},
		"PartTwoExample2": {
			F:         day03.PartTwo,
			InputPath: "testdata/example2.txt",
			Expected:  []byte("610"),
		},
		"PartTwoExample3": {
			F:         day03.PartTwo,
			InputPath: "testdata/example3.txt",
			Expected:  []byte("410"),
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
			F:         day03.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day03.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
