package ds

import (
	"errors"
	"mrmodule/internal/mapreduce"
)

// StackInt is a last-in, first-out (LIFO) stack of integers.
type StackInt struct {
	stack []int
}

// Push adds value to the top of the stack.
func (s *StackInt) Push(value int) {
	if s.stack == nil {
		s.stack = make([]int, 0)
	}
	s.stack = append(s.stack, value)
}

// Count returns the number of values currently on the stack.
func (s StackInt) Count() int {
	if s.stack == nil {
		return 0
	}
	return len(s.stack)
}

// Pop removes and returns the value at the top of the stack. It returns an
// error if the stack is empty.
func (s *StackInt) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value, nil
}

// Total returns the sum of all values on the stack.
func (s StackInt) Total() int {
	if s.stack == nil {
		return 0
	}

	return mapreduce.ReduceInt(
		s.stack,
		func(x, y int) int {
			return x + y
		},
		0)
}

// Average returns the arithmetic mean of the values on the stack, or 0 if the
// stack is empty.
func (s StackInt) Average() float64 {
	if len(s.stack) == 0 {
		return 0.0
	}
	return float64(s.Total()) / float64(len(s.stack))
}
