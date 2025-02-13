package stack

import (
	"errors"
	"fmt"
	"math"
)

// MaxStack represents a stack that supports retrieving the maximum element in constant time.
type MaxStack struct {
	maxElement int
	length     int
	top        *NodeInt
}

// NewMaxStack creates and returns an empty MaxStack.
func NewMaxStack() *MaxStack {
	return &MaxStack{
		maxElement: math.MinInt,
		length: 0,
		top: nil,
	}
}

// IsEmpty checks if the stack is empty.
func (mx *MaxStack) IsEmpty() bool {
	return mx.length == 0
}

// Size returns the number of elements in the stack.
func (mx *MaxStack) Size() int {
	return mx.length
}

// Peek returns the top element of the stack without removing it.
func (mx *MaxStack) GetMax() int {
	return mx.maxElement
}

// Push adds an element to the top of the stack and updates the maximum element.
func (mx *MaxStack) Push(x int) {
	if mx.IsEmpty() {
		mx.maxElement = x
	} else if x > mx.maxElement {
		t := 2*x - mx.maxElement
		mx.maxElement = x
		x = t
	}

	n := &NodeInt{x, mx.top}
	mx.top = n
	mx.length++
}

// Pop removes and returns the top element from the stack, updating the maximum element if necessary.
func (mx *MaxStack) Pop() (int, error) {
	if mx.IsEmpty() {
		return -1, errors.New("cannot pop from empty stack")
	}

	// pop the value from stack
	y := mx.top.data
	mx.top = mx.top.prev
	mx.length--

	if y > mx.maxElement {
		mx.maxElement = 2*mx.maxElement - y
	}

	return y, nil
}

// Print displays the stack elements from top to bottom.
func (mx *MaxStack) Print() {
	fmt.Print("top")
	for t := mx.top; t != nil; t = t.prev {
		fmt.Printf(" --- %d", t.data)
	}
	fmt.Print(" --- nil\n")
}

func (mx *MaxStack) Peek() (int, error) {
	if mx.IsEmpty() {
		return -1, errors.New("cannot peek from empty stack")
	}

	if mx.top.data > mx.maxElement {
		return mx.maxElement, nil
	}

	return mx.top.data, nil
}