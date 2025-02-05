package linkedlist

// Node represents a singly linked list
type Node struct {
	data any
	next *Node
}

// Node2 represents a doubly linked list
type Node2 struct {
	data any
	next, prev *Node2
}