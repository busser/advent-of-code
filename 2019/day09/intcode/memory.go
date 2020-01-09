package intcode

func (r *runner) get(pos int) int {
	if pos >= len(r.memory) {
		r.memory = append(r.memory, make([]int, pos-len(r.memory)+1)...)
	}
	return r.memory[pos]
}

func (r *runner) set(pos int, val int) {
	if pos >= len(r.memory) {
		r.memory = append(r.memory, make([]int, pos-len(r.memory)+1)...)
	}
	r.memory[pos] = val
}
