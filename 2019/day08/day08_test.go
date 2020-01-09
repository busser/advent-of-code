package day08_test

import (
	"testing"

	"github.com/busser/adventofcode/util"

	"github.com/busser/adventofcode/2019/day08"
)

func TestDay(t *testing.T) {
	tests := map[string]util.TestCase{
		"PartOne": {
			F:         day08.PartOne,
			InputPath: "testdata/input.txt",
			Expected:  []byte("2520"),
		},
		"PartTwo": {
			F:         day08.PartTwo,
			InputPath: "testdata/input.txt",
			Expected:  []byte("1000011110011000011010001\n1000010000100100001010001\n1000011100100000001001010\n1000010000101100001000100\n1000010000100101001000100\n1111011110011100110000100"), // LEGJY
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
			F:         day08.PartOne,
			InputPath: "testdata/input.txt",
		},
		"PartTwo": {
			F:         day08.PartTwo,
			InputPath: "testdata/input.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			util.RunBenchmark(b, bc)
		})
	}
}
