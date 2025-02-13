package queue

import "testing"

func TestNew(t *testing.T) {
	q := NewQueue()

	//check length
	if expect, got := 0, q.length; expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	
	//check front and rear ptrs
	if f := q.front; f != nil {
		t.Error("expected nil ptrs for front")
	}
	
	if r := q.rear; r != nil {
		t.Error("expected nil ptrs for  rear")
	}

}

func TestSize(t *testing.T) {
	q := NewQueue()

	//check length
	if expect, got := 0, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	
	q.Enqueue(1)
	if expect, got := 1, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	
	q.Dequeue()
	if expect, got := 0, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue()

	if expect, got := true, q.IsEmpty(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	
	q.Enqueue(1)
	if expect, got := false, q.IsEmpty(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()

	// test nil
	if p := q.Peek(); p != nil {
		t.Error("expected nil on empty peek list")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if p := q.Peek(); p != 1 {
		t.Errorf("expected 1 but got %v\n", p)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	// test empty list
	q.Enqueue(1)

	if f, r := q.front, q.rear; f != r {
		t.Error("expected both front and rear to be equal")
	}

	// test on size 1
	q.Enqueue("table")

	if f, r := q.front, q.rear; f == r {
		t.Error("expected both front and rear to be different")
	}

	if q.Size() != 2 {
		t.Error("expected size to be two.")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)

	dq, err := q.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if expect, got := 2, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	if dq != 3 {
		t.Errorf("expected 3 but got %v\n", dq)
	}

	dq, _ = q.Dequeue()

	if dq != 2 {
		t.Errorf("expected 3 but got %v\n", dq)
	}

	// test empty list
	q.Dequeue()
	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestClear(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(1)
	q.Enqueue(1)
	q.Enqueue(1)
	q.Enqueue(1)

	if expect, got := 5, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	q.Clear()

	if expect, got := 0, q.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
}