package binarytree

import (
	"fmt"
	"math"
)

type Node struct {
	data        int
	left, right *Node
}
type BinarySearchTree struct {
	root *Node
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{nil}
}

// search delete
// levelorder
// find min find max height isBalanced

func (b *BinarySearchTree) Insert(v int) {
	b.root = b.insertHelper(v, b.root)
}

func (b *BinarySearchTree) insertHelper(v int, n *Node) *Node {
	if n == nil {
		return &Node{v, nil, nil}
	}

	if v > n.data {
		n.right = b.insertHelper(v, n.right)
	} else if v < n.data {
		n.left = b.insertHelper(v, n.left)
	}
	return n
}

func (b *BinarySearchTree) Height() int {
	return int(b.heightHelper(b.root))
}
func (b *BinarySearchTree) heightHelper(n *Node) float64 {
	if n == nil {
		return -1
	}

	return math.Max(
		b.heightHelper(n.left),b.heightHelper(n.right)) + 1
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.root == nil
}

func (b *BinarySearchTree) Search(v int) bool {
	return b.searchHelper(v, b.root)
}

func (b *BinarySearchTree) searchHelper(v int, n *Node) bool {
	if n == nil {
		fmt.Println("| false |")
		return false
	}
	
	if v == n.data {
		fmt.Println("| true |")
		return true
	}

	if v > n.data {
		fmt.Printf("| %d | going right\n",n.data)
		return b.searchHelper(v, n.right)
		} else {
		fmt.Printf("| %d | going left\n",n.data)
		return b.searchHelper(v, n.left)
	}
}

func (b *BinarySearchTree) InorderTraversal() {
	fmt.Print("[ ")
	b.inorderHelper(b.root)
	fmt.Println("]")
}
func (b *BinarySearchTree) inorderHelper(n *Node) {
	if n == nil {
		return
	}

	b.inorderHelper(n.left)
	fmt.Printf("%v ", n.data)
	b.inorderHelper(n.right)
}

func (b *BinarySearchTree) PreorderTraversal(){
	fmt.Print("[")
	b.preorderHelper(b.root)
	fmt.Println("]")
}

func (b *BinarySearchTree) preorderHelper(n *Node){
	if n == nil {
		return
	}

	fmt.Printf("%v ", n.data)
	b.preorderHelper(n.left)
	b.preorderHelper(n.right)
}
func (b *BinarySearchTree) PostorderTraversal() {
	fmt.Print("[")
	b.postorderHelper(b.root)
	fmt.Println("]")
}

func (b *BinarySearchTree) postorderHelper(n *Node){
	if b == nil {
		return
	}
	b.preorderHelper(n.left)
	b.preorderHelper(n.right)	
	fmt.Printf("%v ", n.data)
}