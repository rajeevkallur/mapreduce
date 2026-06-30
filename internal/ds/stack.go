package ds

import (
	"errors"
	"mrmodule/internal/mapreduce"
)

// StackInt is a last-in, first-out (LIFO) stack of integers.
type Stack[T ~int | float64] struct {
	stack []T
}

// Push adds value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	if s.stack == nil {
		s.stack = make([]T, 0)
	}
	s.stack = append(s.stack, value)
}

// Count returns the number of values currently on the stack.
func (s Stack[T]) Count() int {
	if s.stack == nil {
		return 0
	}
	return len(s.stack)
}

// Pop removes and returns the value at the top of the stack. It returns an
// error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.stack) == 0 {
		return zero, errors.New("stack is empty")
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value, nil
}

// Total returns the sum of all values on the stack.
func (s Stack[T]) Total() T {
	if s.stack == nil {
		var zero T
		return zero
	}

	return mapreduce.Reduce(
		s.stack,
		func(x, y T) T {
			return x + y
		},
		0)
}

// Average returns the arithmetic mean of the values on the stack, or 0 if the
// stack is empty.
func (s Stack[T]) Average() float64 {
	if len(s.stack) == 0 {
		return 0.0
	}
	return float64(s.Total()) / float64(len(s.stack))
}
