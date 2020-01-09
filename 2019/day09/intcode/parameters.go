package intcode

import "fmt"

func (r *runner) paramValue(offset int) int {
	if r.err != nil {
		return 0
	}

	instruction := r.get(r.position)
	param := r.get(r.position + offset)

	mode := instruction / pow10(offset+1) % 10
	switch mode {
	case 0: // position mode
		return r.get(param)
	case 1: // immediate mode
		return param
	case 2: // relative mode
		return r.get(param + r.relativeBase)
	default:
		r.err = fmt.Errorf("unknown parameter mode %d", mode)
		return 0
	}
}

func (r *runner) paramPosition(offset int) int {
	if r.err != nil {
		return 0
	}

	instruction := r.get(r.position)
	param := r.get(r.position + offset)

	mode := instruction / pow10(offset+1) % 10
	switch mode {
	case 0: // position mode
		return param
	case 1: // immediate mode
		r.err = fmt.Errorf("cannot get position of parameter in immediate mode")
		return 0
	case 2: // relative mode
		return param + r.relativeBase
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
