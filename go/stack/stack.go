package stack

import "errors"

type Stack struct {
	length int
	top    *Node
}
type Node struct {
	data any
	prev *Node
}

// New creates and returns a new Stack instance with default values
// initialized for its length and top pointer.
func New() *Stack {
	return &Stack{0, nil}
}

// IsEmpty returns a boolean value indicating whether 
// the current Stack instance is empty.
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// Size returns the current amount of elements on the Stack instance.
func (s *Stack) Size() int {
	return s.length
}


// Peek returns the element at the top of the Stack instance without
// removing it. It throws an error if the Stack is empty.
func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}

	return s.top.data, nil
}

// Push inserts an element at the top of the Stack instance
func (s *Stack) Push(v any) {
	n := &Node{v, nil}
	s.length++
	
	if s.IsEmpty() {
		s.top = n
	}
	
	n.prev = s.top
	s.top = n
}

// Pop removes the element at the top of the Stack instance. It
// throws an error if it tries to remove from an empty Stack.
func (s *Stack) Pop() (any, error) {

	if s.IsEmpty(){
		return -1, errors.New("cannot pop from an empty stack")
	}

	n := s.top.data
	s.top = s.top.prev
	s.length--
	return n, nil
}