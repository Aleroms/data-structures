package linkedlist

// Node struct has a data any type and
// next pointer
type Node struct {
	data any
	next *Node
}

type Node2 struct {
	data any
	next, prev *Node2
}