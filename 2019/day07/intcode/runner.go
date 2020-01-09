package intcode

import (
	"context"
	"fmt"
)

// Runner runs an IntCode program.
type Runner interface {
	// Run executes the program found in the Runner's memory.
	Run(context.Context)
	// Err returns any error that may have occured during program execution.
	Err() error
}

// NewRunner returns a new Runner ready to run program. It will call readFunc to
// read a value as input, writeFunc to write a value as output, and errFunc when
// an error occurs.
func NewRunner(program []int, readFunc func() int, writeFunc func(int)) Runner {
	return &runner{
		memory:    program,
		readFunc:  readFunc,
		writeFunc: writeFunc,
	}
}

type runner struct {
	memory []int

	readFunc  func() int
	writeFunc func(int)

	position int

	err error
}

func (r *runner) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			r.err = ctx.Err()
			return
		default:
		}

		opcode := r.memory[r.position] % 100

		switch opcode {
		case 1:
			r.add()
		case 2:
			r.multiply()
		case 3:
			r.read()
		case 4:
			r.write()
		case 5:
			r.jumpIfTrue()
		case 6:
			r.jumpIfFalse()
		case 7:
			r.lessThan()
		case 8:
			r.equals()
		case 99:
			return
		default:
			r.err = fmt.Errorf("unknwon opcode %d", opcode)
		}

		if r.err != nil {
			return
		}
	}
}

func (r *runner) Err() error {
	return r.err
}
