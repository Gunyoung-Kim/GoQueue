package queue

import (
	"errors"
)

// ErrNoElement is returned when do something with empty queue which is impossible with empty queue
var ErrNoElement = errors.New("This queue has no element")

//Queue is first-in-first-out data structure.
type Queue struct {
	contents []interface{}
	start    int
	end      int
	length   int
}

//NewQueue returns new Queue with given type.
func NewQueue() *Queue {
	queue := &Queue{
		contents: make([]interface{}, 1),
		start:    -1,
		end:      -1,
		length:   0,
	}

	return queue
}

// Add add element to queue.
// Instead of append slice each time, expand size of slice double when there is no more space.
func (q *Queue) Add(e interface{}) error {
	q.expandSize()
	if q.start == -1 && q.end == -1 {
		q.start = 0
		q.end = 0
	} else {
		q.end++
	}
	q.contents[q.end] = e
	q.length++
	return nil
}

// Poll remove element which is head of queue.
// Return removed element.
// After poll resize slice in queue.
// If queue was empty, then return {@code nil}
func (q *Queue) Poll() interface{} {
	if q.length == 0 {
		return nil
	}

	target := q.contents[q.start]
	q.contents[q.start] = nil
	q.start++
	q.length--
	q.reduceSize()
	return target
}

// Remove remove element which is head of queue.
// Return removed element.
// After poll resize slice in queue.
// One difference with Poll() is this method return {@code ErrNoElement} if queue is empty
func (q *Queue) Remove() (interface{}, error) {
	if q.length == 0 {
		return nil, ErrNoElement
	}

	target := q.contents[q.start]
	q.start++
	q.length--
	q.reduceSize()
	return target, nil
}

// Peek return element which is head of queue.
// This does not remove element, just return element.
// If queue is empty, then return {@code nil}
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}

	return q.contents[q.start]
}

// Element return element which is head of queue.
// This does not remove element, just return element.
// One difference with Peek() is this method return {@code ErrNoElement} if queue was empty
func (q *Queue) Element() (interface{}, error) {
	if q.length == 0 {
		return nil, ErrNoElement
	}

	return q.contents[q.start], nil
}

// IsEmpty return whether queue is empty or not
// Return true if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// Size return length of queue
func (q *Queue) Size() int {
	return q.length
}

// reSize make new slice in queue for new size
func (q *Queue) reSize(size int) {
	newSlice := make([]interface{}, size)

	copy(newSlice, q.contents[q.start:q.end+1])

	q.contents = newSlice
	q.start = 0
	q.end = q.length - 1
}

// reduceSize reduce slice in queue when there is many empty block
func (q *Queue) reduceSize() {
	if q.length < len(q.contents)/4 {
		q.reSize(len(q.contents) / 2)
	}
}

// expandSize expand slice in queue when there is no more space in slice to add new element.
func (q *Queue) expandSize() {
	if q.end == len(q.contents)-1 {
		q.reSize(len(q.contents) * 2)
	}
}
