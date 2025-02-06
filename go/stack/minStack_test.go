package stack

import (
	"math"
	"testing"
)

func TestNewMinStack(t *testing.T) {
	m := NewMinStack()

	// size should be empty
	if m.Size() != 0 {
		t.Errorf("expected size to be 0 but got %d\n", m.Size())
	}

	if m.minElement != math.MaxInt {
		t.Error("expected min element to be math.MaxInt")
	}

	if m.top != nil {
		t.Error("expected top to be nil")
	}
}

func TestIsEmpty2(t *testing.T) {
	m := NewMinStack()
	m.Push(1)

	if m.IsEmpty() {
		t.Error("expected IsEmpty to be false when an element is in stack")
	}

	m.Pop()

	if !m.IsEmpty() {
		t.Error("expected IsEmpty to be true on an empty stack")
	}
}

func TestSize2(t *testing.T) {
	m := NewMinStack()
	
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

func TestPeek2(t *testing.T) {
	m := NewMinStack()
	
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

func TestPush2(t *testing.T) {
	m := NewMinStack()
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

func TestGetMin(t *testing.T) {
	m := NewMinStack()
    m.Push(5);
	m.Push(3);
	
    v := m.GetMin();
	if expect, got := 3, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}

    m.Push(2);
    m.Push(1);
    v = m.GetMin();
	if expect, got := 1, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
    m.Pop();
    v = m.GetMin();
	if expect, got := 2, v; expect != got {
		t.Errorf("expected %d, got %d\n",expect,got)
	}
}

func TestPop2(t *testing.T) {
	m := NewMinStack()
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