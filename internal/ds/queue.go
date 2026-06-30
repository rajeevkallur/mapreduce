package ds

import (
	"errors"
)

// Queue is a first-in, first-out (FIFO) queue.
type Queue[T any] struct {
	queue []T
}

// Add appends value to the back of the queue.
func (q *Queue[T]) Add(value T) {
	if q.queue == nil {
		q.queue = make([]T, 0)
	}
	q.queue = append(q.queue, value)
}

// Remove removes and returns the value at the front of the queue. It returns an
// error if the queue is empty.
func (q *Queue[T]) Remove() (T, error) {
	if len(q.queue) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	value := q.queue[0]
	q.queue = q.queue[1:]
	return value, nil
}
