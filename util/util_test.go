package util_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/busser/adventofcode/util"
)

func TestReadLines(t *testing.T) {
	tests := map[string]struct {
		reader   io.Reader
		expected []string
	}{
		"test-empty-reader": {
			reader:   bytes.NewReader(nil),
			expected: nil,
		},
		"test-usual-case": {
			reader:   bytes.NewReader([]byte("this is\npretty\nstandard content.")),
			expected: []string{"this is", "pretty", "standard content."},
		},
		"test-trailing-newline": {
			reader:   bytes.NewReader([]byte("hello\nworld\n")),
			expected: []string{"hello", "world"},
		},
		"test-empty-lines": {
			reader:   bytes.NewReader([]byte("hello\n\nworld")),
			expected: []string{"hello", "", "world"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := util.ReadLines(tc.reader)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("did not get expected output:\n\texpected: %q\n\tactual: %q", tc.expected, actual)
			}
		})
	}
}

func BenchmarkReadLines(b *testing.B) {
	benchmarks := map[string]struct {
		inputPath string
	}{
		"5-paragraphs": {
			inputPath: "testdata/5-paragraphs.txt",
		},
		"50-paragraphs": {
			inputPath: "testdata/50-paragraphs.txt",
		},
		"150-paragraphs": {
			inputPath: "testdata/150-paragraphs.txt",
		},
	}

	for name, bc := range benchmarks {
		b.Run(name, func(b *testing.B) {
			file, err := os.Open(bc.inputPath)
			if err != nil {
				b.Fatalf("failed to open file at path %q: %v", bc.inputPath, err)
			}
			defer file.Close()

			input, err := ioutil.ReadAll(file)
			if err != nil {
				b.Fatalf("failed to read data from file at path %q: %v", bc.inputPath, err)
			}

			r := bytes.NewReader(nil)

			for n := 0; n < b.N; n++ {
				b.StopTimer()
				r.Reset(input)
				b.StartTimer()

				util.ReadLines(r)
			}
		})
	}
}
