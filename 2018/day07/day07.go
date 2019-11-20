package day07

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strings"
)

// Step represents a step in the instructions.
type Step rune

// StepHeap is a min-heap of steps.
type StepHeap []Step

func (h StepHeap) Len() int           { return len(h) }
func (h StepHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StepHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push appends s to h.
func (h *StepHeap) Push(s interface{}) {
	*h = append(*h, s.(Step))
}

// Pop removes and returns the last element in h.
func (h *StepHeap) Pop() interface{} {
	n := len(*h)
	e := (*h)[n-1]
	*h = (*h)[:n-1]
	return e
}

// PartOne solves Part One of the puzzle.
func PartOne(instructions map[Step][]Step) string {
	indegreeByStep := getIndegreeByStep(instructions)

	availableSteps := &StepHeap{}
	heap.Init(availableSteps)
	for step, indegree := range indegreeByStep {
		if indegree == 0 {
			heap.Push(availableSteps, step)
		}
	}

	var orderedSteps []Step

	for availableSteps.Len() > 0 {
		nextStep := heap.Pop(availableSteps).(Step)

		for _, step := range instructions[nextStep] {
			indegreeByStep[step]--
			if indegreeByStep[step] == 0 {
				heap.Push(availableSteps, step)
			}
		}

		orderedSteps = append(orderedSteps, nextStep)
	}

	var sb strings.Builder
	for _, step := range orderedSteps {
		sb.WriteRune(rune(step))
	}

	return sb.String()
}

type worker struct {
	isWorking bool
	step      Step
	workLeft  int
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(instructions map[Step][]Step, numWorkers int, stepDurationFunc func(s Step) int) int {
	indegreeByStep := getIndegreeByStep(instructions)

	availableSteps := &StepHeap{}
	heap.Init(availableSteps)
	for step, indegree := range indegreeByStep {
		if indegree == 0 {
			heap.Push(availableSteps, step)
		}
	}

	var timeElapsed int

	var workers []*worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{})
	}

	var completedSteps []Step

	for len(completedSteps) < len(indegreeByStep) {
		// Assign available steps to available workers.
		for _, worker := range workers {
			if worker.isWorking || availableSteps.Len() == 0 {
				continue
			}
			worker.step = heap.Pop(availableSteps).(Step)
			worker.workLeft = stepDurationFunc(worker.step)
			worker.isWorking = true
		}

		// Workers make progress.
		for _, worker := range workers {
			if !worker.isWorking {
				continue
			}

			worker.workLeft--

			if worker.workLeft <= 0 { // If step complete:
				// Find newly available steps.
				for _, step := range instructions[worker.step] {
					indegreeByStep[step]--
					if indegreeByStep[step] == 0 {
						heap.Push(availableSteps, step)
					}
				}

				completedSteps = append(completedSteps, worker.step)
				worker.isWorking = false
			}
		}

		timeElapsed++
	}

	return timeElapsed
}

// ReadInstructions reads instructions from r and returns a map of steps to
// slices of steps. If step A must be finished before step B can begin, then
// the slice obtained for key A will contain B.
func ReadInstructions(r io.Reader) (map[Step][]Step, error) {
	instructions := make(map[Step][]Step)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var before, after Step

		_, err := fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.", &before, &after)
		if err != nil {
			return nil, fmt.Errorf("failed to parse text: %w", err)
		}

		instructions[before] = append(instructions[before], after)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return instructions, nil
}

func getIndegreeByStep(instructions map[Step][]Step) map[Step]int {
	indegreeByStep := make(map[Step]int)

	for step := range instructions {
		indegreeByStep[step] = 0
	}

	for _, steps := range instructions {
		for _, step := range steps {
			indegreeByStep[step]++
		}
	}

	return indegreeByStep
}
