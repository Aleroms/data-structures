package queue

import (
	"errors"
	"fmt"
)

// CircularQueue represents a fixed-size circular queue data structure.
type CircularQueue struct {
	front, rear, length int
	data               []any
}

// NewCircularQueue creates a new CircularQueue with the given length.
func NewCircularQueue(len int) *CircularQueue {
	return &CircularQueue{
		front:  -1,
		rear:  -1,
		length: len,
		data:   make([]any, len),
	}
}

// Enqueue adds an element to the rear of the circular queue, overwriting the oldest value if full.
func (cq *CircularQueue) Enqueue(v any) {
	if cq.IsEmpty() {
		cq.front = 0
	} 

	cq.rear = (cq.rear + 1) % cq.length
	cq.data[cq.rear] = v

	// Allows overwriting the oldest values if the queue is full.
	if cq.front == cq.rear {
		cq.front = (cq.front + 1) % cq.length
	}
}

// Dequeue removes and returns the front element from the circular queue.
func (cq *CircularQueue) Dequeue() (any, error) {
	if cq.IsEmpty() {
		return nil, errors.New("dequeue on empty circular queue")
	}

	t := cq.data[cq.front]

	// If the queue had only one element, reset the pointers.
	if cq.front == cq.rear {
		cq.front, cq.rear = -1, -1
	} else {
		cq.front = (cq.front + 1) % cq.length
	}

	return t, nil
}

// Peek returns the front element without removing it from the circular queue.
func (cq *CircularQueue) Peek() (any, error) {
	if !cq.IsEmpty() {
		return cq.data[cq.front], nil
	}
	return nil, errors.New("no elements in circular queue")
}

// IsEmpty checks if the circular queue is empty.
func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == -1
}

// IsFull checks if the circular queue is full.
func (cq *CircularQueue) IsFull() bool {
	return cq.front == (cq.rear + 1) % cq.length
}

// ToArray returns a copy of the internal array representing the queue.
func (cq *CircularQueue) ToArray() []any {
	xa := make([]any, cq.length)
	copy(xa, cq.data)
	return xa
}

// Print displays the contents of the circular queue.
func (cq *CircularQueue) Print() {
	fmt.Println(cq.data...)
}

// Size returns the number of elements currently in the circular queue.
func (cq *CircularQueue) Size() int {
	if cq.IsEmpty() {
		return 0
	}

	if cq.front <= cq.rear {
		return cq.rear - cq.front
	}

	return cq.length - cq.front + cq.rear + 1
}
