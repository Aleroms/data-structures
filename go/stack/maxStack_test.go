package stack

import (
	"math"
	"testing"
)

func TestNewMaxStack(t *testing.T) {
	m := NewMaxStack()

	// size should be empty
	if m.Size() != 0 {
		t.Errorf("expected size to be 0 but got %d\n", m.Size())
	}

	if m.maxElement != math.MinInt {
		t.Error("expected min element to be math.MinInt")
	}

	if m.top != nil {
		t.Error("expected top to be nil")
	}
}

func TestIsEmpty3(t *testing.T) {
	m := NewMaxStack()
	m.Push(1)

	if m.IsEmpty() {
		t.Error("expected IsEmpty to be false when an element is in stack")
	}

	m.Pop()

	if !m.IsEmpty() {
		t.Error("expected IsEmpty to be true on an empty stack")
	}
}

func TestSize3(t *testing.T) {
	m := NewMaxStack()
	
	if expect, got := 0, m.Size(); expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}

	for i:= 0; i < 10; i++ {
		m.Push(i)
	}

	if expect, got := 10, m.Size(); expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
}

func TestPeek3(t *testing.T) {
	m := NewMaxStack()
	
	if _, err := m.Peek(); err == nil {
		t.Error("expected nil on empty stack")
	}

	for i:= 0; i < 10; i++ {
		m.Push(i)
	}

	v, _ := m.Peek()
	if expect, got := 9, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
}

func TestPush3(t *testing.T) {
	m := NewMaxStack()
	m.Push(18)
	m.Push(19)
	m.Push(29)
	m.Push(15)
	m.Push(16)

	v, _ := m.Peek()

	if expect, got := 16, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
}

func TestGetMax(t *testing.T) {
	m := NewMaxStack()
    m.Push(5);
	m.Push(3);
	
    v := m.GetMax();
	if expect, got := 5, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}

    m.Push(4);
    m.Push(7);
    v = m.GetMax();
	if expect, got := 7, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
    m.Pop();
    v = m.GetMax();
	if expect, got := 5, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
}

func TestPop3(t *testing.T) {
	m := NewMaxStack()
	m.Push(18)
	m.Push(19)
	m.Push(29)
	m.Push(15)
	m.Push(16)

	v, _ := m.Pop()
	if expect, got := 16, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
	m.Pop()
	m.Pop()
	m.Pop()
	v, _ = m.Pop()
	if expect, got := 18, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
	m.Pop()
	if _, err := m.Pop(); err == nil {
		t.Error("expected error to be thrown")
	}
}