// Package test provides helper functions for testing and benchmarking Advent of
// Code solutions.
package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

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

	input, err := bytesFromFile(tc.InputPath)
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

// A BenchCase represents benchmark parameters for an Advent of Code solution.
type BenchCase struct {
	F         func(w io.Writer, r io.Reader) error
	InputPath string
}

// RunBenchmark runs a benchmark on f with the provided input.
// RunBenchmark assumes f works as expected and will not check for correctness.
func RunBenchmark(b *testing.B, bc BenchCase) {
	w := ioutil.Discard

	input, err := bytesFromFile(bc.InputPath)
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

func bytesFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file at path %q: %w", path, err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file at path %q: %w", path, err)
	}

	return b, nil
}
