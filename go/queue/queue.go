package queue

import "errors"

// Node represents a single element in the linked list-based queue.
type Node struct {
    data any
    next *Node
}

// Queue represents a linked list-based queue data structure.
type Queue struct {
    front, rear *Node
    length      int
}

// NewQueue creates and returns an empty queue.
func NewQueue() *Queue {
    return &Queue{
        front: nil,
        rear: nil,
        length: 0,
    }
}

// Enqueue adds an element to the rear
func (q *Queue) Enqueue(v any){
    n := &Node{v, nil}

    if q.rear == nil {
        q.rear = n
        q.front = q.rear
    }

    q.rear.next = n
    q.rear = n
    q.length++
}
// Dequeue removes and returns the front element
func (q *Queue) Dequeue() (any, error){
    if q.IsEmpty() {
        //empty Q
        return -1, errors.New("Dequeue from empty queue")
    }
    n := q.front
    q.front = q.front.next

    q.length--
    return n.data, nil
}

// Peek returns the front element without removing it. This method returns
// nil if the queue is empty
func (q *Queue) Peek() any {
    if !q.IsEmpty() {
        return q.front.data
    }
    return nil
}

// PeekFront returns the front element without removing it. This method returns
// nil if the queue is empty
func (q *Queue) PeekFront() any {
    return q.Peek()
}

// PeekRear returns the rear element without removing it. This method returns
// nil if the queue is empty
func (q *Queue) PeekRear() any {
    if q.rear != nil {
        return q.rear.data
    }
    return nil
}


// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
    return q.length == 0
}

// Size returns the number of elements in the
// queue.
func (q *Queue) Size() int {
    return q.length
}

// Clear clears the queue of any elements
func (q *Queue) Clear() {
    t := q.front

    for t != nil {
        next := t.next
        t = nil
        t = next
    }

    q.length = 0
}