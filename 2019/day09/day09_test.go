package day09_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day09"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day09.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("2752191671"),
		},
		"PartTwo": {
			F:         day09.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("87571"),
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
			F:         day09.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day09.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
