package intcode

func (r *runner) add() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	c := r.paramPosition(3)
	if r.err != nil {
		return
	}

	r.memory[c] = a + b

	r.position += 4
}

func (r *runner) multiply() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	c := r.paramPosition(3)
	if r.err != nil {
		return
	}

	r.memory[c] = a * b

	r.position += 4
}

func (r *runner) read() {
	a := r.paramPosition(1)
	if r.err != nil {
		return
	}

	r.memory[a] = r.readFunc()

	r.position += 2
}

func (r *runner) write() {
	a := r.paramValue(1)
	if r.err != nil {
		return
	}

	r.writeFunc(a)

	r.position += 2
}

func (r *runner) jumpIfTrue() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	if r.err != nil {
		return
	}

	if a != 0 {
		r.position = b
		return
	}

	r.position += 3
}

func (r *runner) jumpIfFalse() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	if r.err != nil {
		return
	}

	if a == 0 {
		r.position = b
		return
	}

	r.position += 3
}

func (r *runner) lessThan() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	c := r.paramPosition(3)
	if r.err != nil {
		return
	}

	if a < b {
		r.memory[c] = 1
	} else {
		r.memory[c] = 0
	}

	r.position += 4
}

func (r *runner) equals() {
	a := r.paramValue(1)
	b := r.paramValue(2)
	c := r.paramPosition(3)
	if r.err != nil {
		return
	}

	if a == b {
		r.memory[c] = 1
	} else {
		r.memory[c] = 0
	}

	r.position += 4
}
