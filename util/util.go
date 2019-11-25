// Package util provides helper functions for writing Advent of Code solutions.
package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

// ReadLines returns a slice of all lines in r.
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	return lines, nil
}

// TestCase represents test parameters for an Advent of Code solution.
type TestCase struct {
	F         func(w io.Writer, r io.Reader) error
	InputPath string
	Expected  []byte
}

// RunTest runs a test on f with the provided input. The test will fail if f
// returns an error or if f's output does not match expectedOutput.
func RunTest(t *testing.T, tc TestCase) {
	w := bytes.NewBuffer(nil)

	input, err := ioutil.ReadFile(tc.InputPath)
	if err != nil {
		t.Fatalf("failed to get input from file at path %q: %v", tc.InputPath, err)
	}
	r := bytes.NewReader(input)

	err = tc.F(w, r)
	if err != nil {
		t.Errorf("unexpected error when calling f: %v", err)
	}

	actual := w.Bytes()
	if !bytes.Equal(tc.Expected, actual) {
		t.Errorf("did not get expected output:\n\texpected: %q\n\tactual: %q", tc.Expected, actual)
	}
}

// A BenchmarkCase represents benchmark parameters for an Advent of Code solution.
type BenchmarkCase struct {
	F         func(w io.Writer, r io.Reader) error
	InputPath string
}

// RunBenchmark runs a benchmark on f with the provided input.
// RunBenchmark assumes f works as expected and will not check for correctness.
func RunBenchmark(b *testing.B, bc BenchmarkCase) {
	w := ioutil.Discard

	input, err := ioutil.ReadFile(bc.InputPath)
	if err != nil {
		b.Fatalf("failed to get input from file at path %q: %v", bc.InputPath, err)
	}

	r := bytes.NewReader(nil)

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		r.Reset(input)
		b.StartTimer()

		bc.F(w, r)
	}
}
