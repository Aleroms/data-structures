package linkedlist

import (
	"testing"
)

func TestInsertAtHead(t *testing.T) {
	l := New2()

	l.InsertAtHead(1)

	if expect, got := 1, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if head, tail := l.GetFront(), l.GetLast(); head != tail {
		t.Errorf("expected %v, got %v\n", head, tail)
	} 
	
	l.InsertAtHead(2)
	
	if expect, got := 2, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if head, tail := l.GetFront(), l.GetLast(); head == tail {
		t.Errorf("expected %v, got %v\n", head, tail)
	} 
	
	l.Clear()
	
	if got := l.GetFront(); nil != got {
		t.Errorf("expected nil but got %v\n", got)
	}
}

func TestInsertAtTail(t *testing.T) {
	l := New2()

	if got := l.GetLast(); got != nil {
		t.Errorf("expected nil but got %v\n", got)
	}
	
	l.InsertAtTail(1)
	
	expect := 1
	got := l.GetFront()
	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	got = l.GetLast()
	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	l.InsertAtTail(2)
	if head, tail := l.GetFront(), l.GetLast(); head == tail {
		t.Errorf("expected %v, got %v\n", head, tail)
		} 
		
	}
	
func TestInsertAt2(t *testing.T){
	l := New2()
	
	// insert at head via empty list
	l.InsertAt(5,"t")
	if expect, got := "t", l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	// insert at head via index 0
	l.InsertAt(0,"a")
	if expect, got := "a", l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	// insert at head via negative index
	l.InsertAt(-6,"b")
	if expect, got := "b", l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	// insert at tail
	l.InsertAt(2,"t")
	if expect, got := "t", l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	// insert at tail via negative index
	l.InsertAt(-1, "l")
	if expect, got := "l", l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	// insert at tail via index > length
	l.InsertAt(20, "e")
	l.Print()
	if expect, got := "e", l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	//insert in middle
	l.InsertAt(2,"s")
	l.Print()
	v, err := l.GetAtPosition(2)
	if err != nil {
		t.Error(err)
	}
	if expect, got := "s", v; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
}

func TestClear2(t *testing.T) {
	l := New2()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Clear()
	if expect, got := 0, l.length; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if l.head != nil && l.tail != nil {
		t.Error("Expected head to be nil")
	}

}

func TestIsEmpty2(t *testing.T){
	l := New2()

	if expect, got := true , l.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	l.Append(2)
	if expect, got := false , l.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	l.Clear()
	if expect, got := true , l.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestSize2(t *testing.T){
	l := New2()

	if expect, got := 0, l.Size(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	l.Append(2)
	l.Append(2)
	l.Append(2)
	l.Append(2)
	
	if expect, got := 4, l.Size(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	l.Clear()

	if expect, got := 0, l.Size(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}

func TestToArray2(t *testing.T){

	l := New2()
	

	l.InsertAt(5,"t")
	l.InsertAt(0,"a")
	l.InsertAt(-6,"b")
	l.InsertAt(2,"t")
	l.InsertAt(-1, "l")
	l.InsertAt(20, "e")
	l.InsertAt(2,"s")
	
	xs := l.ToArray()
	as := []string{"b","a","s","t","t","l","e"}
	for i := 0; i < len(as); i++ {
		if xs[i] != as[i] {
			t.Errorf("expected %v, got %v\n", as[i], xs[i])
		}
	}

	l.Clear()
	empty_x := l.ToArray()

	if expect, got := 0, len(empty_x); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	
}

func TestClone2(t *testing.T){
	l := New2()

	l.InsertAtTail(1)
	l.InsertAtTail(2)
	l.InsertAtTail(3)

	c := l.Clone()
	c.RemoveAtHead()

	if expect, got := c.GetFront(), l.GetFront(); expect == got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestAppend2(t *testing.T){
	l := New2()
	l.Append(1)
	l.Append(12)
	l.Append(123)

	if expect, got := 123, l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}

func TestFind2(t *testing.T){
	l := New2()
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)

	// find item not in linked list
	if _, err := l.Find(4); err == nil {
		t.Error(err)
	}

	if n, _ := l.Find(3); n == nil {
		t.Error("expected to find node")
	}

}

func TestRemoveAtHead(t *testing.T){
	l := New2()
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.InsertAtHead(2)

	l.RemoveAtHead()

	if expect, got := 3, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	l.RemoveAtHead()

	if expect, got := 2, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestRemoveAtTail(t *testing.T){
	l := New2()
	l.InsertAtTail(2)
	l.InsertAtTail(3)
	l.InsertAtTail(2)

	l.RemoveAtTail()

	if expect, got := 3, l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	l.RemoveAtTail()

	if expect, got := 2, l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestRemoveAt(t *testing.T){
	l := New2()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)

	l.RemoveAt(2)

	if expect, got := false, l.Contains(3); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if expect, got := 5, l.Size(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}


}

func TestGetFront2(t *testing.T){
	l := New2()
	l.InsertAt(0,1)

	if expect, got := 1, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestGetLast2(t *testing.T){
	l := New2()
	l.InsertAt(-1,1)

	if expect, got := 1, l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestGetAtPosition2(t *testing.T){
	l := New2()
	l.InsertAtTail(2)
	l.InsertAtTail(3)
	l.InsertAtTail(2)

	n, err := l.GetAtPosition(1)
	if err != nil {
		t.Error(err)
	}
	
	if expect, got := 3, n; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestContains2(t *testing.T){
	l := New2()
	l.Append(1)
	l.Append(1)
	l.Append(1)
	l.Append(1)

	if expect, got := true, l.Contains(1); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}
func TestConcatenate2(t *testing.T){
	l := New2()
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)

	n := New2()
	n.InsertAtTail(4)
	n.InsertAtTail(5)
	n.InsertAtTail(6)

	l.Concatenate(n)
	l.Print()
	
	v, err := l.GetAtPosition(3)
	if err != nil {
		t.Error(err)
	}

	if expect, got := 4, v; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}



