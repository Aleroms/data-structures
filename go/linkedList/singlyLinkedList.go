package linkedlist

import (
	"errors"
	"fmt"
)

// also make method concatenate

type SinglyLinkedList struct {
	head *Node
	length int
}

// Node struct has a data any type and
// next pointer
type Node struct {
	data any
	next *Node
}

// New returns a newly created singly linked list
// with default values.
func New() *SinglyLinkedList {
	return &SinglyLinkedList{nil, 0}
}

// Insert inserts a newly created node at the 
// head of a linked list.
func (sl *SinglyLinkedList) Insert(v any) {
	n := &Node{v, nil}
	
	if sl.head != nil {
		n.next = sl.head
	}
	sl.head = n
	sl.length++
}

// InsertAt inserts a newly created node at the specified
// index.
func (sl *SinglyLinkedList) InsertAt(index int, v any){
	if index == 0 {
		sl.Insert(v)
		return
	}

	t := sl.head
	for i := 1; i < index && i < sl.length; i++ {
		t = t.next
	}
	n := &Node{v, t.next}
	t.next = n
	sl.length++
}

// Append appends a newly created node at the tail of
// the linked list.
func (sl *SinglyLinkedList) Append(v any){
	if sl.head == nil {
		sl.Insert(v)
		return
	}
	t := sl.head
	for t.next != nil {
		t = t.next
	}
	n := &Node{v, nil}
	t.next = n
	sl.length++
}

// Clear empties the linked list by setting the head to
// nil and resetting the length. The garbage collector will
// notice there is no reference to other nodes and free their memory.
func (sl *SinglyLinkedList) Clear() {
	sl.head = nil
	sl.length = 0
}

// Find returns the first node that matches the value passed as argument.
// It returns an error if the node could not be found.
func (sl *SinglyLinkedList) Find(v any) (*Node, error) {
	if sl.head == nil {
		return nil, errors.New("empty linked list")
	}

	for t := sl.head; t.next != nil; t = t.next {
		if t.data == v {
			return t, nil
		}
	}

	return nil, errors.New("did not find node")
}

// Print prints the linked list to the standard output.
func (sl *SinglyLinkedList) Print() {
	fmt.Print("[")
	for t := sl.head; t != nil; t = t.next {
		
		fmt.Print(t.data)

		if t.next != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}

// Remove removes a specified node from the linked list.
// it returns an error if list is empty or the node could
// not be found.
func (sl *SinglyLinkedList) Remove(node *Node) error{
	if sl.head == nil {
		return errors.New("cannot remove empty list")
	}
	//[1]
	if sl.head == node {
		sl.head = sl.head.next
		sl.length--
		return nil
	}
	// [1,2,3]
	for t := sl.head; t != nil; t = t.next {
		if t.next == node {
			n := t.next
			t.next = n.next
			n = nil
			sl.length--
			return nil
		}
	}
	return errors.New("node could not be found")
}

// IsEmpty returns if the linked list is empty
func (sl *SinglyLinkedList) IsEmpty() bool{
	return sl.length == 0
}

// Size returns the current size of the linked list
func (sl *SinglyLinkedList) Size() int {
	return sl.length
}

// ToArray returns a slice of all elements in the linked list.
func (sl *SinglyLinkedList) ToArray() []any {
    s := make([]any, sl.length)
    i := 0
    for t := sl.head; t != nil; t = t.next {
        s[i] = t.data // Assign data directly to the pre-allocated slice
        i++
    }
    return s
}

// Clone creates a deep copy of all elements in the 
// linked list.
func (sl *SinglyLinkedList) Clone() *SinglyLinkedList {
	l := New()

	for t := sl.head; t != nil; t = t.next {
		l.Insert(t.data)
	}
	return l
}

// GetAtPosition returns the element stored at the specified position.
// An error is returned if the list is empty or element could not be
// found at the position requested.
func (sl *SinglyLinkedList) GetAtPosition(index int) (any, error) {
	//nil
	if sl.head == nil {
		return nil, errors.New("list is empty")
	}

	t := sl.head
	for i := 0; i < sl.length; i++ {
		if i == index {
			return t.data, nil
		}
		t = t.next
	}
	return nil, errors.New("did not find data at position requested")
}

// GetFront returns the element at the front of the linked list.
func (sl *SinglyLinkedList) GetFront() any {
	if sl.head != nil {
		return sl.head.data
	}
	return nil
}

// GetLast returns the last element in the linked list
func (sl *SinglyLinkedList) GetLast() any {
	
	for t := sl.head; t != nil; t = t.next {
		if t.next == nil {
			// last position
			return t.data
		}
	}
	return nil 
}

// Contains returns a boolean if the element is found in the linked list
func (sl *SinglyLinkedList) Contains(v any) bool {

	for t := sl.head; t != nil; t = t.next {
		if (t.data == v){
			return true
		}
	}
	return false
}