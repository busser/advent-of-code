package intcode

import "fmt"

func (r *runner) paramValue(offset int) int {
	if r.err != nil {
		return 0
	}

	instruction := r.memory[r.position]

	mode := instruction / pow10(offset+1) % 10
	switch mode {
	case 0: // position mode
		return r.memory[r.memory[r.position+offset]]
	case 1: // immediate mode
		return r.memory[r.position+offset]
	default:
		r.err = fmt.Errorf("unknown parameter mode %d", mode)
		return 0
	}
}

func (r *runner) paramPosition(offset int) int {
	if r.err != nil {
		return 0
	}

	instruction := r.memory[r.position]

	mode := instruction / pow10(offset+1) % 10
	switch mode {
	case 0: // position mode
		return r.memory[r.position+offset]
	case 1: // immediate mode
		r.err = fmt.Errorf("cannot get position of parameter in immediate mode")
		return 0
	default:
		r.err = fmt.Errorf("unknown parameter mode %d", mode)
		return 0
	}
}

func pow10(n int) int {
	val := 1
	for ; n > 0; n-- {
		val *= 10
	}
	return val
}
