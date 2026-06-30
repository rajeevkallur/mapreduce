package ds

import (
	"errors"
)

// QueueInt is a first-in, first-out (FIFO) queue of integers.
type QueueInt struct {
	queue []int
}

// Add appends value to the back of the queue.
func (q *QueueInt) Add(value int) {
	if q.queue == nil {
		q.queue = make([]int, 0)
	}
	q.queue = append(q.queue, value)
}

// Remove removes and returns the value at the front of the queue. It returns an
// error if the queue is empty.
func (q *QueueInt) Remove() (int, error) {
	if len(q.queue) == 0 {
		return 0, errors.New("queue is empty")
	}
	value := q.queue[0]
	q.queue = q.queue[1:]
	return value, nil
}
