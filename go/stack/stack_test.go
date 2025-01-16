package stack

import "testing"

func TestPeek(t *testing.T) {
	expect := 3
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	got, err := s.Peek()
	if err != nil {
		t.Error(err)
	}

	if got != expect {
		t.Errorf("expected %v got %v\n", expect, got)
	}

	// Peek on empty stack
	s.Pop()
	s.Pop()
	s.Pop()

	if _, err := s.Peek(); err == nil {
		t.Error("expected an error to be thrown when Peek on empty stack")
	}
	
}
func TestPop(t *testing.T) {
	expect := 3
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	got, err := s.Pop()
	if err != nil {
		t.Error(err)
	}

	if got != expect {
		t.Errorf("expected %v got %v\n", expect, got)
	}

	// Pop on empty stack
	s.Pop()
	s.Pop()
	s.Pop()

	if _, err := s.Pop(); err == nil {
		t.Error("expected an error to be thrown when Pop on empty stack")
	}
	
}

func TestSize(t *testing.T) {
	s := Stack{}
	s.Push("Hello Test")

	if s.Size() != 1 {
		t.Error("Expected size to be 1 with one element in stack")
	}
}

func TestIsEmpty(t *testing.T) {
	s := Stack{}
	s.Push("Hello Test")

	if s.IsEmpty(){
		t.Error("Expected IsEmpty to be false with one element in stack")
	}
}

func TestPush(t *testing.T) {
	s := Stack{}

	s.Push(1)
	s.Push("2")
	s.Push(3.0)
}