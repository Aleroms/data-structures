package binarytree

import (
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

	n.Insert(6)
	n.Insert(3)
	n.Insert(0)
	n.Insert(9)
	n.Insert(4)
	
	// not in BST
	if expect, got := false, n.Search(20); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	// in BST
	if expect, got := true, n.Search(0); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if expect, got := true, n.Search(4); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}

func TestDelete(t *testing.T) {
	n := NewBST()
	d := []int{6,3,0,9,4}
	for _ , v := range d {
		n.Insert(v)
	}

	// delete leaf
	n.Delete(4)
	if expect, got := false, n.Search(4); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	

	// delete one child
	n.Delete(3)
	if expect, got := false, n.Search(3); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	// make sure 6 (root) points to 0
	if expect, got := 0, n.root.left.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	n.Insert(7)
	n.Insert(10)

	// delete two children
	n.Delete(9)
	if expect, got := false, n.Search(9); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestInorderTraversal(t *testing.T) {
	n := NewBST()
	n.Insert(2)
	n.Insert(3)
	n.Insert(1)
	n.InorderTraversal()
}

func TestPreorderTraversal(t *testing.T) {
	n := NewBST()
	n.Insert(2)
	n.Insert(3)
	n.Insert(1)
	n.PreorderTraversal()
}

func TestPostorderTraversal(t *testing.T) {
	n := NewBST()
	n.Insert(2)
	n.Insert(3)
	n.Insert(1)
	n.PostorderTraversal()
}

func TestLevelOrderTraversal(t *testing.T) {
	n := NewBST()
	n.Insert(10)
	n.Insert(5)
	n.Insert(15)
	n.Insert(12)
	n.Insert(20)
	n.LevelOrderTraversal()
}

func TestFindMin(t *testing.T) {
	n := NewBST()
	n.Insert(10)
	n.Insert(5)
	n.Insert(15)
	n.Insert(12)
	n.Insert(20)

	if expect, got:= 5, n.FindMin(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestFindMax(t *testing.T) {
	n := NewBST()
	n.Insert(10)
	n.Insert(5)
	n.Insert(15)
	n.Insert(12)
	n.Insert(20)

	if expect, got:= 20, n.FindMax(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestGetSuccessor(t *testing.T) {
	n := NewBST()
	n.Insert(10)
	n.Insert(5)
	n.Insert(15)
	n.Insert(12)
	n.Insert(20)

	s := n.GetSuccessor(n.root.right)
	
	if expect, got := 20, s.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestGetPredecessor(t *testing.T) {
	n := NewBST()
	n.Insert(10)
	n.Insert(5)
	n.Insert(15)
	n.Insert(12)
	n.Insert(20)

	s := n.GetPredecessor(n.root.right)
	
	if expect, got := 12, s.data; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}