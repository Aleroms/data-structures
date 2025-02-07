package stack

import (
	"errors"
	"fmt"
	"math"
)

// MinStack represents a stack that supports retrieving the minimum element in constant time.
type MinStack struct {
	minElement int
	length     int
	top        *NodeInt
}

// NodeInt represents a node in the MinStack.
type NodeInt struct {
	data int
	prev *NodeInt
}

// NewMinStack creates and returns an empty MinStack.
func NewMinStack() *MinStack {
	return &MinStack{
		minElement: math.MaxInt,
		length:     0,
		top:        nil,
	}
}

// IsEmpty checks if the stack is empty.
func (ms *MinStack) IsEmpty() bool {
	return ms.length == 0
}

// Size returns the number of elements in the stack.
func (ms *MinStack) Size() int {
	return ms.length
}

// Peek returns the top element of the stack without removing it.
func (ms *MinStack) Peek() (int, error) {
	if ms.IsEmpty() {
		return -1, errors.New("cannot peek on empty stack")
	}

	if ms.top.data < ms.minElement {
		return ms.minElement, nil
	}
	return ms.top.data, nil
}

// Push adds an element to the top of the stack and updates the minimum element.
func (ms *MinStack) Push(v int) {
	if ms.IsEmpty() {
		ms.minElement = v
	} else if v < ms.minElement {
		temp := 2*v - ms.minElement
		ms.minElement = v // Update minElement before storing the transformed value
		v = temp          // Store the transformed value in the new node
	}

	n := &NodeInt{v, ms.top} // Create new node
	ms.top = n
	ms.length++
}

// GetMin returns the minimum element in the stack.
func (ms *MinStack) GetMin() int {
	return ms.minElement
}

// Pop removes and returns the top element from the stack, updating the minimum element if necessary.
func (ms *MinStack) Pop() (int, error) {
	if ms.IsEmpty() {
		return -1, errors.New("cannot pop on empty stack")
	}
	y := ms.top
	ms.top = ms.top.prev

	if y.data < ms.minElement {
		ms.minElement = 2*ms.minElement - y.data
	}
	ms.length--
	return y.data, nil
}

// Print displays the stack elements from top to bottom.
func (ms *MinStack) Print() {
	fmt.Print("top")
	for t := ms.top; t != nil; t = t.prev {
		fmt.Printf(" --- %d", t.data)
	}
	fmt.Print(" --- nil\n")
}
