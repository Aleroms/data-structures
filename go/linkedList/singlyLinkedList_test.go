package linkedlist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	l := New()
	l.Insert(1)

	if expect, got := 1, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	l.Insert(2)
	l.Insert(3)
	if expect, got := 3, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestInsertAt(t *testing.T) {
	l := New()

	l.InsertAt(0, 3)
	l.InsertAt(0, 2)
	l.InsertAt(0, 1)

	if expect, got := 1, l.GetFront(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestAppend(t *testing.T){
	l := New()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Append(4)

	if expect, got := 4, l.GetLast(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestClear(t *testing.T) {
	l := New()
	l.Insert('c')
	l.Insert('c')
	l.Insert('c')
	l.Clear()

	if expect, got := 0, l.Size(); expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

}
func TestFind(t *testing.T){
	l := New()
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)

	n, err := l.Find(2)
	if err != nil {
		t.Error(err)
	}

	expect, got := 2, n.data
	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if _,  err = l.Find(15); err == nil {
		t.Error(err)
	}

}
func TestRemove(t *testing.T) {
	l := New()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	n, err := l.Find(4)
	l.Print()
	if err != nil {
		t.Error(err)
	}
	l.Remove(n)
	if l.Contains(4) {
		t.Error("Contains removed element")
	}

}
func TestIsEmpty(t *testing.T) {
	l := New()
	
	expect := true
	got := l.IsEmpty()

	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	l.Insert(1)

	expect = false
	got = l.IsEmpty()

	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
}
func TestSize(t *testing.T) {
	l := New()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	expect := 3
	got := l.Size()

	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}	
}
func TestToArray(t *testing.T) {
	l := New()
	l.Insert(1)
	l.Insert(2)
	l.Append(0)
	v := l.ToArray()

	if expect, got := 2, v[0]; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}

	if expect, got := 0, v[2]; expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
}
func TestClone(t *testing.T) {
	
}
func TestGetAtPosition(t *testing.T) {
	sc := "san clemente"
	sb := "san bernardino"
	sd := "san diego"
	l := New()
	l.Insert(sd)	
	l.Insert(sb)	
	l.Insert(sc)

	got, _ := l.GetAtPosition(0)
	expect := sc

	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}
	
	ll := New()
	ll.Print()
	_, err := ll.GetAtPosition(4)
	if err == nil {
		t.Error("expected an exception to be thrown on empty list")
	}
}
func TestGetFront(t *testing.T) {
	l := New()
	l.Insert(3)	
	l.Insert(2)	
	l.Insert(1)
	l.Print()
	expect := 1
	got := l.GetFront()
	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}	
}
func TestGetLast(t *testing.T) {
	l := New()
	l.Insert(3)	
	l.Insert(2)	
	l.Insert(1)
	l.Print()
	expect := 3
	got := l.GetLast()
	if expect != got {
		t.Errorf("expected %v, got %v\n", expect, got)
	}	
}