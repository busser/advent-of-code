package day08

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// node represents a node in the license file tree.
type node struct {
	children []*node
	metadata []int
}

// PartOne solves Part One of the puzzle.
func PartOne(numbers []int) int {
	root, _ := readNode(numbers)
	return sumMetadata(root)
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(numbers []int) int {
	root, _ := readNode(numbers)
	return nodeValue(root)
}

// readNode reads a node from numbers and returns the number of numbers read.
// If the node has children, these are read recursivly.
func readNode(numbers []int) (*node, int) {
	n := &node{}

	numChildren := numbers[0]
	numMetadata := numbers[1]

	offset := 2

	for i := 0; i < numChildren; i++ {
		child, readCount := readNode(numbers[offset:])
		n.children = append(n.children, child)
		offset += readCount
	}

	n.metadata = numbers[offset : offset+numMetadata]
	offset += numMetadata

	return n, offset
}

// sum returns the sum of all values in nums.
func sum(nums []int) int {
	var s int
	for _, n := range nums {
		s += n
	}
	return s
}

// sumMetadata returns the sum of all metatdata values in n and its children.
func sumMetadata(n *node) int {
	if n == nil {
		return 0
	}

	s := sum(n.metadata)

	for _, c := range n.children {
		s += sumMetadata(c)
	}

	return s
}

// nodeValue computes the value of n.
func nodeValue(n *node) int {
	if n == nil {
		// This should never occur.
		return 0
	}

	if len(n.children) == 0 {
		return sum(n.metadata)
	}

	var value int

	for _, m := range n.metadata {
		if m > len(n.children) {
			continue
		}
		value += nodeValue(n.children[m-1])
	}

	return value
}

// ReadNumbers reads numbers from r and returns a slice of integers.
func ReadNumbers(r io.Reader) ([]int, error) {
	var nums []int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, w := range words {
			n, err := strconv.Atoi(w)
			if err != nil {
				return nil, err
			}
			nums = append(nums, n)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return nums, nil
}
