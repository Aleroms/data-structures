package linkedlist

import (
	"errors"
	"fmt"
)

// DoublyLinkedList represents a doubly linked list with a head and tail pointer.
type DoublyLinkedList struct {
    length      int
    head, tail *Node2
}


// New returns a newly created doubly linked list
// with default values.
func New2() *DoublyLinkedList {
	return &DoublyLinkedList{0, nil, nil}
}

// Insert inserts a newly created node at the
// head of a linked list.
func (dl *DoublyLinkedList) InsertAtHead(v any) {
	n := &Node2{v, nil, nil}

	// if dl.head != nil {
	// 	n.next = dl.head
	// }
	// dl.head = n
	// dl.length++

	if dl.head == nil {
		dl.tail = n
	} else {
		n.next = dl.head
		dl.head.prev = n
	}
		
	dl.head = n
	dl.length++
}

// Insert inserts a newly created node at the
// tail of a linked list.
func (dl *DoublyLinkedList) InsertAtTail(v any) {
	n := &Node2{v, nil, nil}

	if dl.tail == nil {
		dl.head = n
	} else {
		dl.tail.next = n
		n.prev = dl.tail
	}
	dl.tail = n
	dl.length++
}

// InsertAt inserts a newly created node at the specified
// index.
func (dl *DoublyLinkedList) InsertAt(index int, v any) {
	if dl.length == 0 || index == 0 || index < 0 - dl.length {
		dl.InsertAtHead(v)
		return
	}

	if index == dl.length - 1 || index >= dl.length {
		dl.InsertAtTail(v)
		return
	}


	// handles (+/-) indexs
	if index > 0 {
		t := dl.head
		for i := 1; i < index; i++ {
			t = t.next
		}
		n := &Node2{v, t.next, t}
		t.next = n
		n.next.prev = n
	} else {

		t := dl.tail
		for i := index; i < -1; i++ {
			t = t.prev
		}
		n := &Node2{v, nil, t }
		t.next = n

		dl.tail = n
	}
	
	dl.length++
}

// Print prints the linked list to the standard output.
func (dl *DoublyLinkedList) Print() {
	fmt.Print("[")
	for t := dl.head; t != nil; t = t.next {

		fmt.Print(t.data)

		if t.next != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}

// PrintReverse prints the linked list to the standard output 
// in reverse.
func (dl *DoublyLinkedList) PrintReverse(){
	fmt.Print("[")
	for t := dl.tail; t != nil; t = t.prev {
		fmt.Print(t.data)

		if t.prev != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}

// IsEmpty returns if the linked list is empty
func (dl *DoublyLinkedList) IsEmpty() bool{
	return dl.length == 0
}

// Size returns the current size of the linked list
func (dl *DoublyLinkedList) Size() int {
	return dl.length
}

// ToArray returns a slice of all elements in the linked list.
func (dl *DoublyLinkedList) ToArray() []any {
    s := make([]any, dl.length)
    i := 0
    for t := dl.head; t != nil; t = t.next {
        s[i] = t.data // Assign data directly to the pre-allocated slice
        i++
    }
    return s
}

// Clone creates a deep copy of all elements in the 
// linked list.
func (dl *DoublyLinkedList) Clone() *DoublyLinkedList {
	l := New2()

	for t := dl.head; t != nil; t = t.next {
		l.InsertAtTail(t.data)
	}
	return l
}

// Append appends a newly created node at the tail of
// the linked list.
func (dl *DoublyLinkedList) Append(v any) {
	dl.InsertAtTail(v)
}

// Clear empties the linked list
func (dl *DoublyLinkedList) Clear() {
	cur := dl.head
	for cur != nil {
		next := cur.next
		cur.prev = nil
		cur.next = nil
		cur = next
	}
	dl.head = nil
	dl.tail = nil
	dl.length = 0
}

// Find returns the first node that matches the value passed as argument.
// It returns an error if the node could not be found.
func (dl *DoublyLinkedList) Find(v any) (*Node2, error) {
	if dl.head == nil {
		return nil, errors.New("empty linked list")
	}

	for t := dl.head; t.next != nil; t = t.next {
		if t.data == v {
			return t, nil
		}
	}

	return nil, errors.New("did not find node")
}

// RemoveAtHead removes the first node in the linked list
func (dl *DoublyLinkedList) RemoveAtHead(){
	dl.head = dl.head.next
	dl.length--
}

// RemoveAtTail removes the last node in the linked list
func (dl *DoublyLinkedList) RemoveAtTail(){
	dl.tail = dl.tail.prev
	dl.length--
}

// RemoveAt removes the node at the specified index
func (dl *DoublyLinkedList) RemoveAt(index int) {
	if dl.head == nil {
		return
	}

	//check if more efficient to start at head || tail
	if index < dl.length / 2 {
		// start at head
		t := dl.head

		for i := 0; i < index; i++ {
			t = t.next
		}

		if t.next != nil {
			t.next.prev = t.prev
		}
		t.prev.next = t.next
		t = nil


	}else {
		// start at tail
		t := dl.tail

		for i := dl.length - 1; i > index; i-- {
			t = t.prev
		}

		if t.prev != nil {
			t.prev.next = t.next
		}
		t.next.prev = t.prev
		t = nil
	}

	dl.length--
}

// Remove removes a specified node from the linked list.
// it returns an error if list is empty or the node could
// not be found.
func (dl *DoublyLinkedList) Remove(node *Node2) error {
	if dl.head == nil {
		return errors.New("cannot remove empty list")
	}

	if dl.head == node {
		dl.RemoveAtHead()
		return nil
	}

	if dl.tail == node {
		dl.RemoveAtTail()
	}

	for t := dl.head; t != nil; t = t.next {
		if t == node {
			t.prev.next = t.next

			if t.next != nil {
				t.next.prev = t.prev
			}

			t = nil
			return nil
		}
	}
	return errors.New("node could not be found")
}

// GetFront returns the first element in the linked list
func (dl *DoublyLinkedList) GetFront() any {
	if dl.head != nil {
		return dl.head.data
	}

	return nil
}

// GetLast returns the last element in the linked list
func (dl *DoublyLinkedList) GetLast() any{
	if dl.tail != nil {
		return dl.tail.data
	}

	return nil
}

// GetAtPosition returns the element stored at the specified position.
// An error is returned if the list is empty or element could not be
// found at the position requested.
func (dl *DoublyLinkedList) GetAtPosition(index int) (any, error) {
	if dl.head == nil {
		return nil, errors.New("list is empty")
	}

	fmt.Println("cur length", dl.length)

	t := dl.head
	for i:=0; i < dl.length; i++ {
		if i == index {
			return t.data, nil
		}
		t = t.next
	}
	return nil, errors.New("did not find element")
}

// Contains returns a boolean if the element is found in the linked list
func (dl *DoublyLinkedList) Contains(v any) bool {

	for t := dl.head; t != nil; t = t.next {
		if (t.data == v){
			return true
		}
	}
	return false
}

// Concatenate concatenates the passed parameter list to the given 
// linked list
func (dl *DoublyLinkedList) Concatenate(ll *DoublyLinkedList) {
	for t := dl.head; t != nil; t = t.next {
		if t.next == nil {
			t.next = ll.head
			ll.head.prev = t
			break
		}
	}
	dl.length += ll.length
}
