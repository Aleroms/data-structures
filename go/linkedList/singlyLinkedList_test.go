package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	
}
func TestInsertAt(t *testing.T) {
	
}
func TestAppend(t *testing.T){

}
func TestClear(t *testing.T) {
	
}
func TestFind(t *testing.T){

}
func TestRemove(t *testing.T) {
	
}
func TestIsEmpty(t *testing.T) {
	
}
func TestSize(t *testing.T) {
	
}
func TestToArray(t *testing.T) {
	
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