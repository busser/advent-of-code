package day05_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day05"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day05.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("3122865"),
		},
		"PartTwo": {
			F:         day05.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("773660"),
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
			F:         day05.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day05.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
