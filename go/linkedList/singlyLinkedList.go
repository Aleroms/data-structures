package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

//insert, remove, search, concatenate

type SinglyLinkedList struct {
	head *Node
	length int
}

type Node struct {
	data any
	next *Node
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{nil, 0}
}

func (sl *SinglyLinkedList) Insert(v any) {
	n := &Node{v, nil}
	
	if sl.head != nil {
		n.next = sl.head
	}
	sl.head = n
	sl.length++
}

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
}

func (sl *SinglyLinkedList) Clear() {
	// the garbage collector will notice there is no
	// reference to other nodes and free their memory.
	sl.head = nil
	sl.length = 0
}

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

func (sl *SinglyLinkedList) Print() string {
	if sl.head == nil {
		return "[]"
	}
	s := strings.Builder{}
	var t *Node

	s.WriteString("[")
	for t = sl.head; t.next != nil; t = t.next {
		s.WriteString(t.data.(string))
		s.WriteString(", ")
	}
	s.WriteString(fmt.Sprintf("%v]",t.data))
	return s.String()
}

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

func (sl *SinglyLinkedList) IsEmpty() bool{
	return sl.length == 0
}

func (sl *SinglyLinkedList) Size() int {
	return sl.length
}

func (sl *SinglyLinkedList) ToArray() []any {
	s := make([]any, sl.length)

	for t := sl.head; t != nil; t = t.next {
		s = append(s, t.data)
	}
	return s
}

func (sl *SinglyLinkedList) Clone() *SinglyLinkedList {
	l := New()

	for t := sl.head; t != nil; t = t.next {
		l.Insert(t.data)
	}
	return l
}

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

func (sl *SinglyLinkedList) GetFront() any {
	if sl.head != nil {
		return sl.head.data
	}
	return nil
}

func (sl *SinglyLinkedList) GetLast() any {
	
	for t := sl.head; t != nil; t = t.next {
		if t.next == nil {
			// last position
			return t.data
		}
	}
	return nil 
}