package binarytree

import (
	"fmt"
	"testing"
)

func TestNewBST(t *testing.T) {
	n := NewBST()

	if n.root != nil {
		t.Errorf("expected root to be nil")
	}

	// if expect, got := nil, n.root; expect != got {
	// 	t.Errorf("expected %v, got %v\n", expect, got)
	// }
}

func TestInsert(t *testing.T) {
	n := NewBST()
	n.Insert(1)

	if expect, got := 1, n.root.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	n.Insert(2)
	
	if expect, got := 2, n.root.right.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	n.Insert(-1)
	
	if expect, got := -1, n.root.left.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	n.Insert(-1)
	n.InorderTraversal()
	

}

func TestHeight( t *testing.T){
	n := NewBST()

	// empty tree
	if expect, got := -1, n.Height(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}


	// tree with only root
	n.Insert(6)
	if expect, got := 0, n.Height(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	// tree with 1 level
	n.Insert(9)
	n.Insert(3)
	if expect, got := 1, n.Height(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	// tree with 3 level
	n.Insert(10)
	n.Insert(11)
	if expect, got := 1, n.Height(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}

func TestSearch(t *testing.T){
	n := NewBST()
	
	//empty 
	if expect, got := false, n.Search(6); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	fmt.Println("---")
	n.Insert(6)
	n.Insert(3)
	n.Insert(0)
	n.Insert(9)
	n.Insert(4)
	
	// not in BST
	if expect, got := false, n.Search(20); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	fmt.Println("---")
	// in BST
	if expect, got := true, n.Search(0); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	fmt.Println("---")
	if expect, got := true, n.Search(4); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}
// func TestInorderTraversal(t *testing.T) {
// 	n := NewBST()
// 	n.Insert(2)
// 	n.Insert(3)
// 	n.Insert(1)
// 	n.InorderTraversal()
// }

// func TestPreorderTraversal(t *testing.T) {
// 	n := NewBST()
// 	n.Insert(2)
// 	n.Insert(3)
// 	n.Insert(1)
// 	n.PreorderTraversal()
// }

// func TestPostorderTraversal(t *testing.T) {
// 	n := NewBST()
// 	n.Insert(2)
// 	n.Insert(3)
// 	n.Insert(1)
// 	n.PostorderTraversal()
// }