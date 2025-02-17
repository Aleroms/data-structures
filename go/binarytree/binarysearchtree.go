package binarytree

import (
	"fmt"
	"math"

	"github.com/Aleroms/data-structures/go/queue"
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


// isBalanced

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
		return false
	}
	
	if v == n.data {
		return true
	}

	if v > n.data {
		return b.searchHelper(v, n.right)
		} else {
		return b.searchHelper(v, n.left)
	}
}

func (b *BinarySearchTree) Delete(v int) {
	b.root = b.deleteHelper(v, b.root)
}
func (b *BinarySearchTree) deleteHelper(v int, n *Node) *Node{
	if n == nil {
		return nil
	}
	
	if v > n.data {
		n.right = b.deleteHelper(v, n.right)
	} else if v < n.data {
		n.left = b.deleteHelper(v, n.left)
	} else {
		// found node to delete

		// case 1 - no children
		if n.left == nil && n.right == nil {
			return nil
		} else if n.right == nil {
			// case 2 - one child
			return n.left
		} else if n.left == nil {
			// case 2 - one child
			return n.right
		} else {
			// case 3 - two children
			// get in-order successor
			tmp := b.GetSuccessor(n)

			n.data = tmp.data
			n.right = b.deleteHelper(tmp.data, n.right)
		}
	}
	return n
}

func (b *BinarySearchTree) GetSuccessor(n *Node) *Node {
	tmp := n.right
	for tmp.left != nil {
		tmp = tmp.left
	}
	return tmp
}

func (b *BinarySearchTree) GetPredecessor(n *Node) *Node {
	tmp := n.left
	for tmp.right != nil {
		tmp = tmp.right
	}
	return tmp
}

func (b *BinarySearchTree) FindMin() int {
	if b.root == nil {
		return -1
	}
	t := b.root
	for t.left != nil {
		t = t.left
	} 
	return t.data
}

func (b *BinarySearchTree) FindMax() int {
	if b.root == nil {
		return -1
	}
	t := b.root
	for t.right != nil {
		t = t.right
	} 
	return t.data
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

func (b *BinarySearchTree) LevelOrderTraversal(){
	if b.root == nil {
		fmt.Println("[]")
		return
	}

	q := queue.NewQueue()
	q.Enqueue(b.root)

	fmt.Print("[")
	for !q.IsEmpty() {
		t := q.PeekFront()
		n := t.(*Node)

		fmt.Printf(" %v ",n.data)
		

		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
		q.Dequeue()
		
	}
	fmt.Println("]")
	
}